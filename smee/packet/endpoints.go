package packet

import (
	"bytes"
	"context"
	"encoding/json"
	tink "github.com/tinkerbell/tink/protos/hardware"
	cacher "github.com/packethost/cacher/protos/cacher"
	"io"
	"net"
	"os"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tinkerbell/boots/metrics"
	//"github.com/tinkerbell/tink/protos/hardware"
)

const mimeJSON = "application/json"

type Component struct {
	Type            string      `json:"type"`
	Name            string      `json:"name"`
	Vendor          string      `json:"vendor"`
	Model           string      `json:"model"`
	Serial          string      `json:"serial"`
	FirmwareVersion string      `json:"firmware_version"`
	Data            interface{} `json:"data"`
}

type ComponentsResponse struct {
	Components []Component `json:"components"`
}

func (c *Client) DiscoverHardwareFromDHCP(mac net.HardwareAddr, giaddr net.IP, circuitID string) (*Discovery, error) {
	if mac == nil {
		return nil, errors.New("missing MAC address")
	}

	labels := prometheus.Labels{"from": "dhcp"}
	cacherTimer := prometheus.NewTimer(metrics.CacherDuration.With(labels))
	metrics.CacherRequestsInProgress.With(labels).Inc()
	metrics.CacherTotal.With(labels).Inc()

	var msg getRequest
	discoveryType := os.Getenv("DISCOVERY_TYPE")
	switch discoveryType {
	case discoveryTypeCacher:
		msg = &cacher.GetRequest{
			MAC: mac.String(),
		}
	case discoveryTypeTinkerbell:
		msg = &tink.GetRequest{
			Mac: mac.String(),
		}
	default:
		return nil, errors.New("invalid discovery type")
	}

	//msg := &hardware.GetRequest{
	//	Mac: mac.String(),
	//}
	resp, err := c.client.ByMAC(context.Background(), msg)

	cacherTimer.ObserveDuration()
	metrics.CacherRequestsInProgress.With(labels).Dec()

	if err != nil {
		return nil, errors.Wrap(err, "get hardware by mac from tink")
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("marshalling tink hardware")
	}

	// or maybe just check for error?
	if string(b) != "{}" {
		metrics.CacherCacheHits.With(labels).Inc()
		return NewDiscovery(b)
	}

	if giaddr == nil {
		return nil, errors.New("missing MAC address")
	}

	metrics.HardwareDiscovers.With(labels).Inc()
	metrics.DiscoversInProgress.With(labels).Inc()
	defer metrics.DiscoversInProgress.With(labels).Dec()
	discoverTimer := prometheus.NewTimer(metrics.DiscoverDuration.With(labels))
	defer discoverTimer.ObserveDuration()

	req := struct {
		MAC       string `json:"mac"`
		GIADDR    string `json:"giaddr,omitempty"`
		CIRCUITID string `json:"circuit_id,omitempty"`
	}{
		MAC:       mac.String(),
		GIADDR:    giaddr.String(),
		CIRCUITID: circuitID,
	}

	b, err = json.Marshal(&req)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling api discovery")
	}

	var res Discovery
	if err := c.Post("/staff/cacher/hardware-discovery", mimeJSON, bytes.NewReader(b), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DiscoverHardwareFromIP(ip net.IP) (*Discovery, error) {
	if ip.String() == net.IPv4zero.String() {
		return nil, errors.New("missing ip address")
	}

	labels := prometheus.Labels{"from": "ip"}
	cacherTimer := prometheus.NewTimer(metrics.CacherDuration.With(labels))
	defer cacherTimer.ObserveDuration()
	metrics.CacherRequestsInProgress.With(labels).Inc()
	defer metrics.CacherRequestsInProgress.With(labels).Dec()

	var msg getRequest
	discoveryType := os.Getenv("DISCOVERY_TYPE")
	switch discoveryType {
	case discoveryTypeCacher:
		msg = &cacher.GetRequest{
			IP: ip.String(),
		}
	case discoveryTypeTinkerbell:
		msg = &tink.GetRequest{
			Ip: ip.String(),
		}
	default:
		return nil, errors.New("invalid discovery type")
	}

	//msg := &hardware.GetRequest{
	//	Ip: ip.String(),
	//}
	resp, err := c.client.ByIP(context.Background(), msg)
	if err != nil {
		return nil, errors.Wrap(err, "get hardware by ip from tink")
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("marshalling tink hardware")
	}

	//if string(b) == "{}" {
	//	return nil, errors.New("empty response from cacher")
	//}
	metrics.CacherCacheHits.With(labels).Inc()
	return NewDiscovery(b)
}

// GetDeviceIDFromIP Looks up a device (instance) in cacher via ByIP
func (c *Client) GetInstanceIDFromIP(dip net.IP) (string, error) {
	d, err := c.DiscoverHardwareFromIP(dip)
	if err != nil {
		return "", err
	}
	if (*d).Instance() == nil {
		return "", nil
	}
	return (*d).Instance().ID, nil
}

// PostHardwareComponent - POSTs a HardwareComponent to the API
func (c *Client) PostHardwareComponent(hardwareID string, body io.Reader) (*ComponentsResponse, error) {
	var response ComponentsResponse

	if err := c.Post("/hardware/"+hardwareID+"/components", mimeJSON, body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) PostHardwareEvent(id string, body io.Reader) (string, error) {
	var res struct {
		ID string `json:"id"`
	}
	if err := c.Post("/hardware/"+id+"/events", mimeJSON, body, &res); err != nil {
		return "", err
	}
	return res.ID, nil
}
func (c *Client) PostHardwarePhoneHome(id string) error {
	return c.Post("/hardware/"+id+"/phone-home", "", nil, nil)
}
func (c *Client) PostHardwareFail(id string, body io.Reader) error {
	return c.Post("/hardware/"+id+"/fail", mimeJSON, body, nil)
}
func (c *Client) PostHardwareProblem(id string, body io.Reader) (string, error) {
	var res struct {
		ID string `json:"id"`
	}
	if err := c.Post("/hardware/"+id+"/problems", mimeJSON, body, &res); err != nil {
		return "", err
	}
	return res.ID, nil
}

func (c *Client) PostInstancePhoneHome(id string) error {
	return c.Post("/devices/"+id+"/phone-home", "", nil, nil)
}
func (c *Client) PostInstanceEvent(id string, body io.Reader) (string, error) {
	var res struct {
		ID string `json:"id"`
	}
	if err := c.Post("/devices/"+id+"/events", mimeJSON, body, &res); err != nil {
		return "", err
	}
	return res.ID, nil
}
func (c *Client) PostInstanceFail(id string, body io.Reader) error {
	return c.Post("/devices/"+id+"/fail", mimeJSON, body, nil)
}
func (c *Client) PostInstancePassword(id, pass string) error {
	var req = struct {
		Password string `json:"password"`
	}{
		Password: pass,
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return errors.Wrap(err, "marshalling instance password")
	}

	return c.Post("/devices/"+id+"/password", mimeJSON, bytes.NewReader(b), nil)
}
func (c *Client) UpdateInstance(id string, body io.Reader) error {
	return c.Patch("/devices/"+id, mimeJSON, body, nil)
}
