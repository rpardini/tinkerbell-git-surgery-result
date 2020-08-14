package main

import (
	"context"
	"encoding/json"
	"os"

	tink "github.com/tinkerbell/tink/protos/hardware"

	"github.com/packethost/cacher/protos/cacher"
	"google.golang.org/grpc"
)

// hardwareGetterMock is a mock implentation of the hardwareGetter interface
// hardwareResp represents the hardware data stored inside tink db
type hardwareGetterMock struct {
	hardwareResp string
}

func (hg hardwareGetterMock) ByIP(ctx context.Context, in getRequest, opts ...grpc.CallOption) (hardware, error) {
	var hw hardware
	dataModelVersion := os.Getenv("DATA_MODEL_VERSION")
	switch dataModelVersion {
	case "1":
		hw = &tink.Hardware{}
		err := json.Unmarshal([]byte(hg.hardwareResp), hw)
		if err != nil {
			return nil, err
		}
	default:
		hw = &cacher.Hardware{JSON: hg.hardwareResp}
	}

	return hw, nil
}

func (hg hardwareGetterMock) Watch(ctx context.Context, in getRequest, opts ...grpc.CallOption) (watchClient, error) {
	// TODO (kdeng3849)
	return nil, nil
}

const (
	mockUserIP      = "192.168.1.5" // value is completely arbitrary, as long as it's an IP to be parsed by getIPFromRequest (could even be 0.0.0.0)
	cacherDataModel = `
	{
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "arch": "x86_64",
	  "name": "node-name",
	  "state": "provisioning",
	  "allow_pxe": true,
	  "allow_workflow": true,
	  "plan_slug": "t1.small.x86",
	  "facility_code": "onprem",
      "efi_boot": false,
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "device": "/dev/sda",
			  "wipeTable": true,
			  "partitions": [
				{
				  "size": 4096,
				  "label": "BIOS",
				  "number": 1
				},
				{
				  "size": "3993600",
				  "label": "SWAP",
				  "number": 2
				},
				{
				  "size": 0,
				  "label": "ROOT",
				  "number": 3
				}
			  ]
			}
		  ],
		  "filesystems": [
			{
			  "mount": {
				"point": "/",
				"create": {
				  "options": ["-L", "ROOT"]
				},
				"device": "/dev/sda3",
				"format": "ext4"
			  }
			},
			{
			  "mount": {
				"point": "none",
				"create": {
				  "options": ["-L", "SWAP"]
				},
				"device": "/dev/sda2",
				"format": "swap"
			  }
			}
		  ]
		},
		"crypted_root_password": "$6$qViImWbWFfH/a4pq$s1bpFFXMpQj1eQbHWsruLy6/",
		"operating_system_version": {
		  "distro": "ubuntu",
		  "version": "16.04",
		  "os_slug": "ubuntu_16_04"
		}
	  },
	  "ip_addresses": [
		{
		  "cidr": 29,
		  "public": false,
		  "address": "192.168.1.5",
		  "enabled": true,
		  "gateway": "192.168.1.1",
		  "netmask": "255.255.255.248",
		  "network": "192.168.1.0",
		  "address_family": 4
		}
	  ],
	  "network_ports": [
		{
		  "data": {
			"mac": "98:03:9b:48:de:bc"
		  },
		  "name": "eth0",
		  "type": "data"
		}
	  ]
	}
`
	cacherPartitionSizeInt = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": 4096,
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeString = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "3333",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeStringLeadingZeros = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "007",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeWhitespace = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": " \t 1234\n   ",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeInterceptingWhitespace = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "12\tmb",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeBLower = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "1000000b",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeBUpper = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "1000000B",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeK = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "24K",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeKBLower = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "24kb",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeKBUpper = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "24KB",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeKBMixed = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "24Kb",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeM = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "3m",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeTB = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "2TB",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeInvalidSuffix = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "3kmgtb",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeInvalidIntertwined = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "12kb3",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeInvalidIntertwined2 = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "k123b",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeEmpty = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	cacherPartitionSizeReversedPlacement = `
    {
	  "id": "8978e7d4-1a55-4845-8a66-a5259236b104",
	  "instance": {
		"storage": {
		  "disks": [
			{
			  "partitions": [
				{
				  "size": "mb10",
				  "label": "BIOS",
				  "number": 1
				}
			  ]
			}
		  ]
        }
	  }
	}
`
	tinkerbellDataModel = `
	{
	   "network":{
		  "interfaces":[
			 {
				"dhcp":{
				   "mac":"ec:0d:9a:c0:01:0c",
				   "hostname":"server001",
				   "lease_time":86400,
				   "arch":"x86_64",
				   "ip":{
					  "address":"192.168.1.5",
					  "netmask":"255.255.255.248",
					  "gateway":"192.168.1.1"
				   }
				},
				"netboot":{
				   "allow_pxe":true,
				   "allow_workflow":true,
				   "ipxe":{
					  "url":"http://url/menu.ipxe",
					  "contents":"#!ipxe"
				   },
				   "osie":{
					  "kernel":"vmlinuz-x86_64"
				   }
				}
			 }
		  ]
	   },
	   "id":"fde7c87c-d154-447e-9fce-7eb7bdec90c0",
	   "metadata":"{\"components\":{\"id\":\"bc9ce39b-7f18-425b-bc7b-067914fa9786\",\"type\":\"DiskComponent\"},\"userdata\":\"#!/bin/bash\\necho \\\"Hello world!\\\"\",\"bonding_mode\":5,\"custom\":{\"preinstalled_operating_system_version\":{},\"private_subnets\":[]},\"facility\":{\"facility_code\":\"ewr1\",\"plan_slug\":\"c2.medium.x86\",\"plan_version_slug\":\"\"},\"instance\":{\"crypted_root_password\":\"redacted/\",\"operating_system_version\":{\"distro\":\"ubuntu\",\"os_slug\":\"ubuntu_18_04\",\"version\":\"18.04\"},\"storage\":{\"disks\":[{\"device\":\"/dev/sda\",\"partitions\":[{\"label\":\"BIOS\",\"number\":1,\"size\":4096},{\"label\":\"SWAP\",\"number\":2,\"size\":3993600},{\"label\":\"ROOT\",\"number\":3,\"size\":0}],\"wipe_table\":true}],\"filesystems\":[{\"mount\":{\"create\":{\"options\":[\"-L\",\"ROOT\"]},\"device\":\"/dev/sda3\",\"format\":\"ext4\",\"point\":\"/\"}},{\"mount\":{\"create\":{\"options\":[\"-L\",\"SWAP\"]},\"device\":\"/dev/sda2\",\"format\":\"swap\",\"point\":\"none\"}}]}},\"manufacturer\":{\"id\":\"\",\"slug\":\"\"},\"state\":\"\"}"
	}
`
	tinkerbellNoMetadata = `
	{
	   "network":{
		  "interfaces":[
			 {
				"dhcp":{
				   "mac":"ec:0d:9a:c0:01:0c",
				   "hostname":"server001",
				   "lease_time":86400,
				   "arch":"x86_64",
				   "ip":{
					  "address":"192.168.1.5",
					  "netmask":"255.255.255.248",
					  "gateway":"192.168.1.1"
				   }
				},
				"netboot":{
				   "allow_pxe":true,
				   "allow_workflow":true,
				   "ipxe":{
					  "url":"http://url/menu.ipxe",
					  "contents":"#!ipxe"
				   },
				   "osie":{
					  "kernel":"vmlinuz-x86_64"
				   }
				}
			 }
		  ]
	   },
	   "id":"363115b0-f03d-4ce5-9a15-5514193d131a"
	}
`
	tinkerbellKant = `
	{
	   "network":{
		  "interfaces":[
			 {
				"dhcp":{
				   "mac":"ec:0d:9a:c0:01:0c",
				   "hostname":"server001",
				   "lease_time":86400,
				   "arch":"x86_64",
				   "ip":{
					  "address":"192.168.1.5",
					  "netmask":"255.255.255.248",
					  "gateway":"192.168.1.1"
				   }
				},
				"netboot":{
				   "allow_pxe":true,
				   "allow_workflow":true,
				   "ipxe":{
					  "url":"http://url/menu.ipxe",
					  "contents":"#!ipxe"
				   },
				   "osie":{
					  "kernel":"vmlinuz-x86_64"
				   }
				}
			 }
		  ]
	   },
	   "id":"fde7c87c-d154-447e-9fce-7eb7bdec90c0",
       "metadata": "{\"components\":{\"id\":\"bc9ce39b-7f18-425b-bc7b-067914fa9786\",\"type\":\"DiskComponent\"},\"instance\":{\"facility\":\"sjc1\",\"hostname\":\"tink-provisioner\",\"id\":\"f955e31a-cab6-44d6-872c-9614c2024bb4\"},\"userdata\":\"#!/bin/bash\\n\\necho \\\"Hello world!\\\"\"}"
	}
`
)
