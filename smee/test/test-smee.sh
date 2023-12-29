#!/bin/sh
# shellcheck shell=dash disable=SC1091,SC2154

# useful for debugging sometimes
# tcpdump -ni eth0 &
# alternatively, only show DHCP and pretty print the packets
# tcpdump -nvvei eth0 port 67 or port 68 &

sleep_at_start=3
echo "starting DHCP in $sleep_at_start seconds"
sleep $sleep_at_start

# busybox udhcpc will happily set arbitrary DHCP options and is easy
# to configure with a custom setup script to call on DHCPOFFER
#
# dummy setup script for -s is copied in by Dockerfile
# -q tells udhcpc to exit after getting a lease, otherwise it will keep generating new traces
# opt60 (-V PXEClient) pretend to be an Intel PXE client. required to be noticed by smee
# opt93 (-x 0x5d) set to 0 for "Intel x86PC" platform, required by smee
# opt94 (-x 0x5e) set to 0 for "UNDI" firmware type, required by smee
# opt97 (-x 0x61) sets the client guid (https://datatracker.ietf.org/doc/html/rfc4578#section-2.3)
#              first 8 octets should be zeroes to make smee happy (Intel PXE does this)
#              ID: 4a525bd43517df7f8b4799c18d (randomly generated and hard-coded here)
busybox udhcpc \
	-q \
	-s /busybox-udhcpc-script.sh \
	-V PXEClient \
	-x 0x5d:0000 \
	-x 0x5e:0000 \
	-x 0x61:000000004a525bd43517df7f8b4799c18d

# set boot_file variable ahead of sourcing dhcpoffer-vars.sh to please the linter
boot_file=""

# the busybox script writes the DHCP variables to /tmp/dhcpoffer-vars.sh
# shellcheck disable=SC1091
. /tmp/dhcpoffer-vars.sh

# smee sets 2 values in option 43, check out dhcp/pxe.go
# these can come in out of order so we have to look for the traceparent's
# id and length which is always 0x451a
# busybox udhcpc helpfully returns options in hex
# option43 ordering is not guaranteed, at least not in this implementation
. extract-traceparent-from-opt43.sh     # load a function to do the parsing
extract_traceparent_from_opt43 "$opt43" # parse the value, exports TRACEPARENT
echo "got traceparent $TRACEPARENT from opt43 value $opt43"
# write it to the shell profile.d for easy loading
echo "export TRACEPARENT=$TRACEPARENT" >/etc/profile.d/smee-traceparent.sh

# fetch / from the server with the traceparent set
tp_header="Traceparent: $TRACEPARENT"
curl -H "$tp_header" http://192.168.99.42/auto.ipxe
# TODO: test opportunity here: validate the returned traceparent matches the one in boot_file

# boot_file is set by the DHCP envvars
# OTEL in Smee is enabled by default.
tftp 192.168.99.42 -c get "${boot_file}"

# sleep a long time so you can enter the container with
# docker exec -ti smee_client_1 /bin/sh
sleep 30000
