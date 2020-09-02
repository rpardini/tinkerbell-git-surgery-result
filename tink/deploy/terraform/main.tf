# Configure the Packet Provider.
terraform {
  required_providers {
    packet = "~> 3.0.1"
  }
}

provider "packet" {
  auth_token = var.packet_api_token
}

# Create a new VLAN in datacenter "ewr1"
resource "packet_vlan" "provisioning-vlan" {
  description = "provisioning-vlan"
  facility    = var.facility
  project_id  = var.project_id
}

# Create a device and add it to tf_project_1
resource "packet_device" "tink-provisioner" {
  hostname         = "tink-provisioner"
  plan             = var.device_type
  facilities       = [var.facility]
  operating_system = "ubuntu_18_04"
  billing_cycle    = "hourly"
  project_id       = var.project_id
  user_data        = file("install_package.sh")

  provisioner "file" {
    source      = "./../../../tink"
    destination = "/root/"

    connection {
      type        = "ssh"
      user        = var.ssh_user
      host        = packet_device.tink-provisioner.network[0].address
      private_key = file(var.ssh_private_key)
    }
  }
}

resource "packet_device_network_type" "tink-provisioner-network-type" {
  device_id = packet_device.tink-provisioner.id
  type      = "hybrid"
}

# Create a device and add it to tf_project_1
resource "packet_device" "tink-worker" {
  hostname         = "tink-worker"
  plan             = var.device_type
  facilities       = [var.facility]
  operating_system = "custom_ipxe"
  ipxe_script_url  = "https://boot.netboot.xyz"
  always_pxe       = "true"
  billing_cycle    = "hourly"
  project_id       = var.project_id
}

resource "packet_device_network_type" "tink-worker-network-type" {
  device_id = packet_device.tink-worker.id
  type      = "layer2-individual"
}

# Attach VLAN to provisioner
resource "packet_port_vlan_attachment" "provisioner" {
  device_id = packet_device.tink-provisioner.id
  port_name = "eth1"
  vlan_vnid = packet_vlan.provisioning-vlan.vxlan
}

# Attach VLAN to worker
resource "packet_port_vlan_attachment" "worker" {
  device_id = packet_device.tink-worker.id
  port_name = "eth0"
  vlan_vnid = packet_vlan.provisioning-vlan.vxlan
}
