package main

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/tinkerbell/tink/protos/packet"
)

func TestGetByIPCacher(t *testing.T) {
	t.Log("DATA_MODEL_VERSION (empty to use cacher):", os.Getenv("DATA_MODEL_VERSION"))

	for name, test := range cacherGrpcTests {
		t.Log(name)

		hegelTestServer := &server{
			log:            logger,
			hardwareClient: hardwareGetterMock{test.json},
		}
		ehw, err := getByIP(context.Background(), hegelTestServer, test.remote)
		if err != nil {
			if err.Error() != test.error {
				t.Fatalf("unexpected error in getByIP, want: %v, got: %v\n", test.error, err.Error())
			}
			continue
		}
		hw := exportedHardwareCacher{}
		err = json.Unmarshal(ehw, &hw)
		if err != nil {
			t.Fatal("Error in unmarshalling hardware", err)
		}

		if hw.State != test.state {
			t.Fatalf("unexpected state, want: %v, got: %v\n", test.state, hw.State)
		}
		if hw.Facility != test.facility {
			t.Fatalf("unexpected facility, want: %v, got: %v\n", test.facility, hw.Facility)
		}
		if len(hw.NetworkPorts) > 0 && hw.NetworkPorts[0]["data"].(map[string]interface{})["mac"] != test.mac {
			t.Fatalf("unexpected mac, want: %v, got: %v\n", test.mac, hw.NetworkPorts[0]["data"].(map[string]interface{})["mac"])
		}
		if len(hw.Instance.Storage.Disks) > 0 {
			if hw.Instance.Storage.Disks[0].Device != test.diskDevice {
				t.Fatalf("unexpected storage disk device, want: %v, got: %v\n", test.diskDevice, hw.Instance.Storage.Disks[0].Device)
			}
			if hw.Instance.Storage.Disks[0].WipeTable != test.wipeTable {
				t.Fatalf("unexpected storage disk wipe table, want: %v, got: %v\n", test.wipeTable, hw.Instance.Storage.Disks[0].WipeTable)
			}
			t.Log("want:", test.partitionSize, " got:", hw.Instance.Storage.Disks[0].Paritions[0].Size)
			if int(hw.Instance.Storage.Disks[0].Paritions[0].Size) != test.partitionSize {
				t.Fatalf("unexpected storage disk partition size, want: %v, got: %v\n", test.partitionSize, hw.Instance.Storage.Disks[0].Paritions[0].Size)
			}
		}
		if len(hw.Instance.Storage.Filesystems) > 0 {
			if hw.Instance.Storage.Filesystems[0].Mount.Device != test.filesystemDevice {
				t.Fatalf("unexpected storage filesystem mount device, want: %v, got: %v\n", test.filesystemDevice, hw.Instance.Storage.Filesystems[0].Mount.Device)
			}
			if hw.Instance.Storage.Filesystems[0].Mount.Format != test.filesystemFormat {
				t.Fatalf("unexpected storage filesystem mount format, want: %v, got: %v\n", test.filesystemFormat, hw.Instance.Storage.Filesystems[0].Mount.Format)
			}
		}
		if hw.Instance.OS.OsSlug != test.osSlug {
			t.Fatalf("unexpected os slug, want: %v, got: %v\n", test.osSlug, hw.Instance.OS.OsSlug)
		}
		if hw.PlanSlug != test.planSlug {
			t.Fatalf("unexpected plan slug, want: %v, got: %v\n", test.planSlug, hw.PlanSlug)
		}
	}
}

func TestGetByIPTinkerbell(t *testing.T) {
	os.Setenv("DATA_MODEL_VERSION", "1")
	t.Log("DATA_MODEL_VERSION:", os.Getenv("DATA_MODEL_VERSION"))

	for name, test := range tinkerbellGrpcTests {
		t.Log(name)

		hegelTestServer := &server{
			log:            logger,
			hardwareClient: hardwareGetterMock{test.json},
		}
		ehw, err := getByIP(context.Background(), hegelTestServer, test.remote)
		if err != nil {
			if err.Error() != test.error {
				t.Fatalf("unexpected error in getByIP, want: %v, got: %v\n", test.error, err.Error())
			}
			continue
		}
		hw := exportedHardwareTinkerbell{}
		err = json.Unmarshal(ehw, &hw)
		if err != nil {
			t.Error("Error in unmarshalling hardware:", err)
		}

		if hw.Metadata == nil {
			return
		}

		var metadata packet.Metadata
		md, err := json.Marshal(hw.Metadata)
		if err != nil {
			t.Error("Error in unmarshalling hardware", err)
		}
		err = json.Unmarshal(md, &metadata)
		if err != nil {
			t.Error("Error in unmarshalling hardware metadata", err)
		}

		if metadata.State != test.state {
			t.Fatalf("unexpected state, want: %v, got: %v\n", test.state, metadata.State)
		}
		if metadata.BondingMode != test.bondingMode {
			t.Fatalf("unexpected bonding mode, want: %v, got: %v\n", test.bondingMode, metadata.BondingMode)
		}
		if len(metadata.Instance.Storage.Disks) > 0 {
			if metadata.Instance.Storage.Disks[0].Device != test.diskDevice {
				t.Fatalf("unexpected disk device, want: %v, got: %v\n", test.diskDevice, metadata.Instance.Storage.Disks[0].Device)
			}
			if metadata.Instance.Storage.Disks[0].WipeTable != test.wipeTable {
				t.Fatalf("unexpected wipe table, want: %v, got: %v\n", test.wipeTable, metadata.Instance.Storage.Disks[0].WipeTable)
			}
			if metadata.Instance.Storage.Disks[0].Partitions[0].Size != test.partitionSize {
				t.Fatalf("unexpected partition size, want: %v, got: %v\n", test.partitionSize, metadata.Instance.Storage.Disks[0].Partitions[0].Size)
			}
		}
		if len(metadata.Instance.Storage.Filesystems) > 0 {
			if metadata.Instance.Storage.Filesystems[0].Mount.Device != test.filesystemDevice {
				t.Fatalf("unexpected filesystem mount device, want: %v, got: %v\n", test.filesystemDevice, metadata.Instance.Storage.Filesystems[0].Mount.Device)
			}
			if metadata.Instance.Storage.Filesystems[0].Mount.Format != test.filesystemFormat {
				t.Fatalf("unexpected filesystem mount format, want: %v, got: %v\n", test.filesystemFormat, metadata.Instance.Storage.Filesystems[0].Mount.Format)
			}
		}
		if metadata.Instance.OperatingSystemVersion.OsSlug != test.osSlug {
			t.Fatalf("unexpected os slug, want: %v, got: %v\n", test.osSlug, metadata.Instance.OperatingSystemVersion.OsSlug)
		}
		if metadata.Facility.PlanSlug != test.planSlug {
			t.Fatalf("unexpected os slug, want: %v, got: %v\n", test.planSlug, metadata.Facility.PlanSlug)
		}
	}
}

var cacherGrpcTests = map[string]struct {
	remote           string
	state            string
	facility         string
	mac              string
	diskDevice       string
	wipeTable        bool
	partitionSize    int
	filesystemDevice string
	filesystemFormat string
	osSlug           string
	planSlug         string
	json             string
	error            string
}{
	"cacher": {
		remote:           "192.168.1.5",
		state:            "provisioning",
		facility:         "onprem",
		mac:              "98:03:9b:48:de:bc",
		diskDevice:       "/dev/sda",
		wipeTable:        true,
		partitionSize:    4096,
		filesystemDevice: "/dev/sda3",
		filesystemFormat: "ext4",
		osSlug:           "ubuntu_16_04",
		planSlug:         "t1.small.x86",
		json:             cacherDataModel,
	},
	"cacher_partition_size_int": { // 4096
		partitionSize: 4096,
		json:          cacherPartitionSizeInt,
	},
	"cacher_partition_size_string": { // "3333"
		partitionSize: 3333,
		json:          cacherPartitionSizeString,
	},
	"cacher_partition_size_string_leading_zeros": { // "007"
		partitionSize: 7,
		json:          cacherPartitionSizeStringLeadingZeros,
	},
	"cacher_partition_size_whitespace": { // " \t 1234\n  "
		partitionSize: 1234,
		json:          cacherPartitionSizeWhitespace,
	},
	"cacher_partition_size_intercepting_whitespace": { // "12\tmb"
		partitionSize: 12582912,
		json:          cacherPartitionSizeInterceptingWhitespace,
	},
	"cacher_partition_size_b_lower": { // "1000000b"
		partitionSize: 1000000,
		json:          cacherPartitionSizeBLower,
	},
	"cacher_partition_size_b_upper": { // "1000000B"
		partitionSize: 1000000,
		json:          cacherPartitionSizeBUpper,
	},
	"cacher_partition_size_k": { // "24K"
		partitionSize: 24576,
		json:          cacherPartitionSizeK,
	},
	"cacher_partition_size_kb_lower": { // "24kb"
		partitionSize: 24576,
		json:          cacherPartitionSizeKBLower,
	},
	"cacher_partition_size_kb_upper": { // "24KB"
		partitionSize: 24576,
		json:          cacherPartitionSizeKBUpper,
	},
	"cacher_partition_size_kb_mixed": { // "24Kb"
		partitionSize: 24576,
		json:          cacherPartitionSizeKBMixed,
	},
	"cacher_partition_size_m": { // "3m"
		partitionSize: 3145728,
		json:          cacherPartitionSizeM,
	},
	"cacher_partition_size_t": { // "2TB"
		partitionSize: 2199023255552,
		json:          cacherPartitionSizeTB,
	},
	"cacher_partition_size_invalid_suffix": { // "3kmgtb"
		partitionSize: -1,
		json:          cacherPartitionSizeInvalidSuffix,
		error:         "invalid suffix",
	},
	"cacher_partition_size_invalid_intertwined": { // "12kb3"
		partitionSize: -1,
		json:          cacherPartitionSizeInvalidIntertwined,
		error:         `strconv.Atoi: parsing "12kb3": invalid syntax`,
	},
	"cacher_partition_size_invalid_intertwined_2": { // "k123b"
		partitionSize: -1,
		json:          cacherPartitionSizeInvalidIntertwined2,
		error:         "invalid suffix",
	},
	"cacher_partition_size_empty": { // ""
		partitionSize: 0,
		json:          cacherPartitionSizeEmpty,
	},
	"cacher_partition_size_reverse_placement": { // "mb10" // (kdeng3849) should we allow this?
		partitionSize: 10485760,
		json:          cacherPartitionSizeReversedPlacement,
	},
}

var tinkerbellGrpcTests = map[string]struct {
	remote           string
	state            string
	bondingMode      int64
	diskDevice       string
	wipeTable        bool
	partitionSize    int64
	filesystemDevice string
	filesystemFormat string
	osSlug           string
	planSlug         string
	json             string
	error            string
}{
	"tinkerbell": {
		remote:           "192.168.1.5",
		bondingMode:      5,
		diskDevice:       "/dev/sda",
		wipeTable:        true,
		partitionSize:    4096,
		filesystemDevice: "/dev/sda3",
		filesystemFormat: "ext4",
		osSlug:           "ubuntu_18_04",
		planSlug:         "c2.medium.x86",
		json:             tinkerbellDataModel,
	},
	"tinkerbell no metadata": {
		remote: "192.168.1.5",
		json:   tinkerbellNoMetadata,
	},
}
