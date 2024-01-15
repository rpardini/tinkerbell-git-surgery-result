package reservation

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/go-logr/logr"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/tinkerbell/smee/internal/dhcp/data"
	oteldhcp "github.com/tinkerbell/smee/internal/dhcp/otel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/ipv4"
)

const tracerName = "github.com/tinkerbell/smee"

// setDefaults will update the Handler struct to have default values so as
// to avoid panic for nil pointers and such.
func (h *Handler) setDefaults() {
	if h.Backend == nil {
		h.Backend = noop{}
	}
	if h.Log.GetSink() == nil {
		h.Log = logr.Discard()
	}
}

// Handle responds to DHCP messages with DHCP server options.
func (h *Handler) Handle(ctx context.Context, conn *ipv4.PacketConn, p data.Packet) {
	h.setDefaults()
	if p.Pkt == nil {
		h.Log.Error(errors.New("incoming packet is nil"), "not able to respond when the incoming packet is nil")
		return
	}
	upeer, ok := p.Peer.(*net.UDPAddr)
	if !ok {
		h.Log.Error(errors.New("peer is not a UDP connection"), "not able to respond when the peer is not a UDP connection")
		return
	}
	if upeer == nil {
		h.Log.Error(errors.New("peer is nil"), "not able to respond when the peer is nil")
		return
	}
	if conn == nil {
		h.Log.Error(errors.New("connection is nil"), "not able to respond when the connection is nil")
		return
	}

	var ifName string
	if p.Md != nil {
		ifName = p.Md.IfName
	}
	log := h.Log.WithValues("mac", p.Pkt.ClientHWAddr.String(), "xid", p.Pkt.TransactionID.String(), "interface", ifName)
	tracer := otel.Tracer(tracerName)
	var span trace.Span
	ctx, span = tracer.Start(
		ctx,
		fmt.Sprintf("DHCP Packet Received: %v", p.Pkt.MessageType().String()),
		trace.WithAttributes(h.encodeToAttributes(p.Pkt, "request")...),
		trace.WithAttributes(attribute.String("DHCP.peer", p.Peer.String())),
		trace.WithAttributes(attribute.String("DHCP.server.ifname", ifName)),
	)

	defer span.End()

	var reply *dhcpv4.DHCPv4
	switch mt := p.Pkt.MessageType(); mt {
	case dhcpv4.MessageTypeDiscover:
		d, n, err := h.readBackend(ctx, p.Pkt.ClientHWAddr)
		if err != nil {
			if hardwareNotFound(err) {
				span.SetStatus(codes.Ok, "no reservation found")
				return
			}
			log.Info("error reading from backend", "error", err)
			span.SetStatus(codes.Error, err.Error())

			return
		}
		log.Info("received DHCP packet", "type", p.Pkt.MessageType().String())
		reply = h.updateMsg(ctx, p.Pkt, d, n, dhcpv4.MessageTypeOffer)
		log = log.WithValues("type", dhcpv4.MessageTypeOffer.String())
	case dhcpv4.MessageTypeRequest:
		d, n, err := h.readBackend(ctx, p.Pkt.ClientHWAddr)
		if err != nil {
			if hardwareNotFound(err) {
				span.SetStatus(codes.Ok, "no reservation found")
				return
			}
			log.Info("error reading from backend", "error", err)
			span.SetStatus(codes.Error, err.Error())

			return
		}
		log.Info("received DHCP packet", "type", p.Pkt.MessageType().String())
		reply = h.updateMsg(ctx, p.Pkt, d, n, dhcpv4.MessageTypeAck)
		log = log.WithValues("type", dhcpv4.MessageTypeAck.String())
	case dhcpv4.MessageTypeRelease:
		// Since the design of this DHCP server is that all IP addresses are
		// Host reservations, when a client releases an address, the server
		// doesn't have anything to do. This case is included for clarity of this
		// design decision.
		log.Info("received DHCP release packet, no response required, all IPs are host reservations", "type", p.Pkt.MessageType().String())
		span.SetStatus(codes.Ok, "received release, no response required")

		return
	default:
		log.Info("received unknown message type", "type", p.Pkt.MessageType().String())
		span.SetStatus(codes.Error, "received unknown message type")

		return
	}

	if bf := reply.BootFileName; bf != "" {
		log = log.WithValues("bootFileName", bf)
	}
	if ns := reply.ServerIPAddr; ns != nil {
		log = log.WithValues("nextServer", ns.String())
	}

	dst := replyDestination(p.Peer, p.Pkt.GatewayIPAddr)
	log = log.WithValues("ipAddress", reply.YourIPAddr.String(), "destination", dst.String())
	cm := &ipv4.ControlMessage{}
	if p.Md != nil {
		cm.IfIndex = p.Md.IfIndex
	}

	if _, err := conn.WriteTo(reply.ToBytes(), cm, dst); err != nil {
		log.Error(err, "failed to send DHCP")
		span.SetStatus(codes.Error, err.Error())

		return
	}

	log.Info("sent DHCP response")
	span.SetAttributes(h.encodeToAttributes(reply, "reply")...)
	span.SetStatus(codes.Ok, "sent DHCP response")
}

// replyDestination determines the destination address for the DHCP reply.
// If the giaddr is set, then the reply should be sent to the giaddr.
// Otherwise, the reply should be sent to the direct peer.
//
// From page 22 of https://www.ietf.org/rfc/rfc2131.txt:
// "If the 'giaddr' field in a DHCP message from a client is non-zero,
// the server sends any return messages to the 'DHCP server' port on
// the BOOTP relay agent whose address appears in 'giaddr'.".
func replyDestination(directPeer net.Addr, giaddr net.IP) net.Addr {
	if !giaddr.IsUnspecified() && giaddr != nil {
		return &net.UDPAddr{IP: giaddr, Port: dhcpv4.ServerPort}
	}

	return directPeer
}

// readBackend encapsulates the backend read and opentelemetry handling.
func (h *Handler) readBackend(ctx context.Context, mac net.HardwareAddr) (*data.DHCP, *data.Netboot, error) {
	h.setDefaults()

	tracer := otel.Tracer(tracerName)
	ctx, span := tracer.Start(ctx, "Hardware data get")
	defer span.End()

	d, n, err := h.Backend.GetByMac(ctx, mac)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return nil, nil, err
	}

	span.SetAttributes(d.EncodeToAttributes()...)
	span.SetAttributes(n.EncodeToAttributes()...)
	span.SetStatus(codes.Ok, "done reading from backend")

	return d, n, nil
}

// updateMsg handles updating DHCP packets with the data from the backend.
func (h *Handler) updateMsg(ctx context.Context, pkt *dhcpv4.DHCPv4, d *data.DHCP, n *data.Netboot, msgType dhcpv4.MessageType) *dhcpv4.DHCPv4 {
	h.setDefaults()
	mods := []dhcpv4.Modifier{
		dhcpv4.WithMessageType(msgType),
		dhcpv4.WithGeneric(dhcpv4.OptionServerIdentifier, h.IPAddr.AsSlice()),
		dhcpv4.WithServerIP(h.IPAddr.AsSlice()),
	}
	mods = append(mods, h.setDHCPOpts(ctx, pkt, d)...)

	if h.Netboot.Enabled && h.isNetbootClient(pkt) == nil {
		mods = append(mods, h.setNetworkBootOpts(ctx, pkt, n))
	}
	reply, err := dhcpv4.NewReplyFromRequest(pkt, mods...)
	if err != nil {
		return nil
	}

	return reply
}

// isNetbootClient returns true if the client is a valid netboot client.
//
// A valid netboot client will have the following in its DHCP request:
// 1. is a DHCP discovery/request message type.
// 2. option 93 is set.
// 3. option 94 is set.
// 4. option 97 is correct length.
// 5. option 60 is set with this format: "PXEClient:Arch:xxxxx:UNDI:yyyzzz" or "HTTPClient:Arch:xxxxx:UNDI:yyyzzz".
//
// See: http://www.pix.net/software/pxeboot/archive/pxespec.pdf
//
// See: https://www.rfc-editor.org/rfc/rfc4578.html
func (h *Handler) isNetbootClient(pkt *dhcpv4.DHCPv4) error {
	h.setDefaults()
	var err error
	// only response to DISCOVER and REQUEST packets
	if pkt.MessageType() != dhcpv4.MessageTypeDiscover && pkt.MessageType() != dhcpv4.MessageTypeRequest {
		// h.Log.Info("not a netboot client", "reason", "message type must be either Discover or Request", "mac", pkt.ClientHWAddr.String(), "message type", pkt.MessageType())
		err = errors.New("message type must be either Discover or Request")
	}
	// option 60 must be set
	if !pkt.Options.Has(dhcpv4.OptionClassIdentifier) {
		// h.Log.Info("not a netboot client", "reason", "option 60 not set", "mac", pkt.ClientHWAddr.String())
		err = fmt.Errorf("%w: option 60 not set", err)
	}
	// option 60 must start with PXEClient or HTTPClient
	opt60 := pkt.GetOneOption(dhcpv4.OptionClassIdentifier)
	if !strings.HasPrefix(string(opt60), string(pxeClient)) && !strings.HasPrefix(string(opt60), string(httpClient)) {
		// h.Log.Info("not a netboot client", "reason", "option 60 not PXEClient or HTTPClient", "mac", pkt.ClientHWAddr.String(), "option 60", string(opt60))
		err = fmt.Errorf("%w: option 60 not PXEClient or HTTPClient", err)
	}

	// option 93 must be set
	if !pkt.Options.Has(dhcpv4.OptionClientSystemArchitectureType) {
		// h.Log.Info("not a netboot client", "reason", "option 93 not set", "mac", pkt.ClientHWAddr.String())
		err = fmt.Errorf("%w: option 93 not set", err)
	}

	// option 94 must be set
	if !pkt.Options.Has(dhcpv4.OptionClientNetworkInterfaceIdentifier) {
		// h.Log.Info("not a netboot client", "reason", "option 94 not set", "mac", pkt.ClientHWAddr.String())
		err = fmt.Errorf("%w: option 94 not set", err)
	}

	// option 97 must be have correct length or not be set
	guid := pkt.GetOneOption(dhcpv4.OptionClientMachineIdentifier)
	switch len(guid) {
	case 0:
		// A missing GUID is invalid according to the spec, however
		// there are PXE ROMs in the wild that omit the GUID and still
		// expect to boot. The only thing we do with the GUID is
		// mirror it back to the client if it's there, so we might as
		// well accept these buggy ROMs.
	case 17:
		if guid[0] != 0 {
			h.Log.Info("not a netboot client", "reason", "option 97 does not start with 0", "mac", pkt.ClientHWAddr.String(), "option 97", string(guid))
			err = fmt.Errorf("%w: option 97 does not start with 0", err)
		}
	default:
		h.Log.Info("not a netboot client", "reason", "option 97 has invalid length (0 or 17)", "mac", pkt.ClientHWAddr.String(), "option 97", string(guid))
		err = fmt.Errorf("%w: option 97 has invalid length (0 or 17)", err)
	}

	return err
}

// encodeToAttributes takes a DHCP packet and returns opentelemetry key/value attributes.
func (h *Handler) encodeToAttributes(d *dhcpv4.DHCPv4, namespace string) []attribute.KeyValue {
	h.setDefaults()
	a := &oteldhcp.Encoder{Log: h.Log}

	return a.Encode(d, namespace, oteldhcp.AllEncoders()...)
}

// hardwareNotFound returns true if the error is from a hardware record not being found.
func hardwareNotFound(err error) bool {
	type hardwareNotFound interface {
		NotFound() bool
	}
	te, ok := err.(hardwareNotFound)
	return ok && te.NotFound()
}
