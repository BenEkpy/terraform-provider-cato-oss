---
subcategory: ""
page_title: "Azure VNET Socket Module"
description: |-
  Provides an combined example of creating a virtual socket site in Cato Management Application, and templates for creating an Resource Group with underlying network resources and deploying a virtual socket instance in Azure.
---

# Example Azure Module (cato-oss_socket_site)

The `cato-oss_socket_site` resource contains the configuration parameters necessary to 
add a socket site to the Cato cloud 
([virtual socket in AWS/Azure, or physical socket](https://support.catonetworks.com/hc/en-us/articles/4413280502929-Working-with-X1500-X1600-and-X1700-Socket-Sites)).
Documentation for the underlying API used in this resource can be found at
[mutation.addSocketSite()](https://api.catonetworks.com/documentation/#mutation-site.addSocketSite).

## Example Usage

### Create Azure VPC - Example Module

<details>
<summary>Socket Site Variables</summary>

### Socket Site Variables

```hcl
## Cato Provider Variables
variable cato_token {}

variable "account_id" {
  description = "Account ID"
  type        = number
  default	  = null
}

## Cato socket site variables
variable "site_description" {
  type = string
  description = "Site description"
}

variable "project_name" {
  type = string
  description = "Your Cato Deployment Name Here"
  default = "Your Cato Deployment Name Here"
}

variable "vnet_prefix" {
  type = string
  description = <<EOT
  	Choose a unique range for your new VPC that does not conflict with the rest of your Wide Area Network.
    The accepted input format is Standard CIDR Notation, e.g. X.X.X.X/X
	EOT
  default = "20.0.0.0/16"
}

variable "site_type" {
  description = "The type of the site"
  type        = string
  default	 = "DATACENTER"
  validation {
    condition = contains(["DATACENTER","BRANCH","CLOUD_DC","HEADQUARTERS"], var.site_type)
    error_message = "The site_type variable must be one of 'DATACENTER','BRANCH','CLOUD_DC','HEADQUARTERS'."
  }
}

variable "lan_eni_ip" {
   description = "Choose an IP Address within the LAN Subnet. You CANNOT use the first four assignable IP addresses within the subnet as it's reserved for the AWS virtual router interface. The accepted input format is X.X.X.X"
   type        = string
   default	  = null
}

```
</details>

### cato-oss_socket Resource

```
## cato-oss_socket Provider and Resource

provider "cato-oss" {
    baseurl = "https://api.catonetworks.com/api/v1/graphql2"
    token = var.cato_token
    account_id = var.account_id
}

resource "cato-oss_socket_site" "azure-site" {
    connection_type  = "SOCKET_AZ1500"
    description = var.site_description
    name = var.project_name
    native_range = {
      native_network_range = var.vnet_prefix
      local_ip = var.lan_ip
    }
    site_location = {
        country_code = "US"
		    state_code = "US-VA"
        timezone = "America/New_York"
    }
    site_type = var.site_type
}

data "cato-oss_accountSnapshotSite" "azure-site" {
	id = cato-oss_socket_site.azure-site.id
}

```

<details>
<summary>Create Azure VNET - Example Module</summary>

### Create Azure VNET - Example Module

```hcl
## Azure VNET Example Module Variables
variable "assetprefix" {
  type = string
  description = "Your asset prefix for resources created"
  default = "LAB1"
}

variable "location" { 
  type = string
  default = "East US"
}

variable "lan_ip" {
	type = string
	default = "20.0.3.4"
}

variable "project_name" { 
  type = string
  default = "Azure-Lab-Deployment"
}

variable "subnet_range_mgmt" {
  type = string
  description = <<EOT
    Choose a range within the VPC to use as the Management subnet. This subnet will be used initially to access the public internet and register your vSocket to the Cato Cloud.
    The minimum subnet length to support High Availability is /28.
    The accepted input format is Standard CIDR Notation, e.g. X.X.X.X/X
	EOT
  default = null
}

variable "subnet_range_wan" {
  type = string
  description = <<EOT
    Choose a range within the VPC to use as the Public/WAN subnet. This subnet will be used to access the public internet and securely tunnel to the Cato Cloud.
    The minimum subnet length to support High Availability is /28.
    The accepted input format is Standard CIDR Notation, e.g. X.X.X.X/X
	EOT
  default = null
}

variable "subnet_range_lan" {
  type = string
  description = <<EOT
    Choose a range within the VPC to use as the Private/LAN subnet. This subnet will host the target LAN interface of the vSocket so resources in the VPC (or AWS Region) can route to the Cato Cloud.
    The minimum subnet length to support High Availability is /29.
    The accepted input format is Standard CIDR Notation, e.g. X.X.X.X/X
	EOT
  default = null
}

variable "vnet_prefix" {
  type = string
  description = <<EOT
  	Choose a unique range for your new VPC that does not conflict with the rest of your Wide Area Network.
    The accepted input format is Standard CIDR Notation, e.g. X.X.X.X/X
	EOT
  default = null
}

## Azure VNET Example Module Resources
provider "azurerm" {
	features {}
}

resource "azurerm_resource_group" "azure-rg" {
  location = var.location
  name = var.project_name
}

resource "azurerm_availability_set" "availability-set" {
  location                     = var.location
  name                         = "${var.assetprefix}-availabilitySet"
  platform_fault_domain_count  = 2
  platform_update_domain_count = 2
  resource_group_name          = azurerm_resource_group.azure-rg.name
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}

## Create Network and Subnets
resource "azurerm_virtual_network" "vnet" {
  address_space       = [var.vnet_prefix]
  location            = var.location
  name                = "${var.assetprefix}-vsNet"
  resource_group_name = azurerm_resource_group.azure-rg.name
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}

resource "azurerm_subnet" "subnet-mgmt" {
  address_prefixes     = [var.subnet_range_mgmt]
  name                 = "subnetMGMT"
  resource_group_name  = azurerm_resource_group.azure-rg.name
  virtual_network_name = "${var.assetprefix}-vsNet"
  depends_on = [
    azurerm_virtual_network.vnet
  ]
}
resource "azurerm_subnet" "subnet-wan" {
  address_prefixes     = [var.subnet_range_wan]
  name                 = "subnetWAN"
  resource_group_name  = azurerm_resource_group.azure-rg.name
  virtual_network_name = "${var.assetprefix}-vsNet"
  depends_on = [
    azurerm_virtual_network.vnet
  ]
}

resource "azurerm_subnet" "subnet-lan" {
  address_prefixes     = [var.subnet_range_lan]
  name                 = "subnetLAN"
  resource_group_name  = azurerm_resource_group.azure-rg.name
  virtual_network_name = "${var.assetprefix}-vsNet"
  depends_on = [
    azurerm_virtual_network.vnet
  ]
}

# Allocate Public IPs
resource "azurerm_public_ip" "mgmt-public-ip" {
  allocation_method   = "Static"
  location            = var.location
  name                = "${var.assetprefix}-vs0nicMngPublicIP"
  resource_group_name = azurerm_resource_group.azure-rg.name
  sku                 = "Standard"
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}
resource "azurerm_public_ip" "wan-public-ip" {
  allocation_method   = "Static"
  location            = var.location
  name                = "${var.assetprefix}-vs0nicWanPublicIP"
  resource_group_name = azurerm_resource_group.azure-rg.name
  sku                 = "Standard"
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}

# Create Network Interfaces
resource "azurerm_network_interface" "mgmt-nic" {
  location            = var.location
  name                = "${var.assetprefix}-vs0nicMng"
  resource_group_name = azurerm_resource_group.azure-rg.name
  ip_configuration {
    name                          = "vs0nicMngIP"
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.mgmt-public-ip.id
    subnet_id                     = azurerm_subnet.subnet-mgmt.id
  }
  depends_on = [
    azurerm_public_ip.mgmt-public-ip,
    azurerm_subnet.subnet-mgmt
  ]
}

resource "azurerm_network_interface" "wan-nic" {
  enable_ip_forwarding = true
  location             = var.location
  name                 = "${var.assetprefix}-vs0nicWan"
  resource_group_name          = azurerm_resource_group.azure-rg.name
  ip_configuration {
    name                          = "vs0nicWanIP"
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.wan-public-ip.id
    subnet_id                     = azurerm_subnet.subnet-wan.id
  }
  depends_on = [
    azurerm_public_ip.wan-public-ip,
    azurerm_subnet.subnet-wan
  ]
}

resource "azurerm_network_interface" "lan-nic" {
  enable_ip_forwarding = true
  location             = var.location
  name                 = "${var.assetprefix}-vs0nicLan"
  resource_group_name          = azurerm_resource_group.azure-rg.name
  ip_configuration {
    name                          = "lanIPConfig"
    private_ip_address_allocation = "Dynamic"
    subnet_id                     = azurerm_subnet.subnet-lan.id
  }
  depends_on = [
    azurerm_subnet.subnet-lan
  ]
}

resource "azurerm_network_interface_security_group_association" "mgmt-nic-association" {
  network_interface_id      = azurerm_network_interface.mgmt-nic.id
  network_security_group_id = azurerm_network_security_group.mgmt-sg.id
  depends_on = [
    azurerm_network_interface.mgmt-nic,
    azurerm_network_security_group.mgmt-sg
  ]
}

resource "azurerm_network_interface_security_group_association" "wan-nic-association" {
  network_interface_id      = azurerm_network_interface.wan-nic.id
  network_security_group_id = azurerm_network_security_group.wan-sg.id
  depends_on = [
    azurerm_network_interface.wan-nic,
    azurerm_network_security_group.wan-sg
  ]
}

resource "azurerm_network_interface_security_group_association" "lan-nic-association" {
  network_interface_id      = azurerm_network_interface.lan-nic.id
  network_security_group_id = azurerm_network_security_group.lan-sg.id
  depends_on = [
    azurerm_network_interface.lan-nic,
    azurerm_network_security_group.lan-sg
  ]
}

# Create Security Groups
resource "azurerm_network_security_group" "mgmt-sg" {
  location            = var.location
  name                = "${var.assetprefix}-MGMTSecurityGroup"
  resource_group_name = azurerm_resource_group.azure-rg.name
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}
resource "azurerm_network_security_group" "wan-sg" {
  location            = var.location
  name                = "${var.assetprefix}-WANSecurityGroup"
  resource_group_name = azurerm_resource_group.azure-rg.name
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}
resource "azurerm_network_security_group" "lan-sg" {
  location            = var.location
  name                = "${var.assetprefix}-LANSecurityGroup"
  resource_group_name = azurerm_resource_group.azure-rg.name
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}

## Create Route Tables, Routes and Associations 
resource "azurerm_route_table" "private-rt" {
  disable_bgp_route_propagation = true
  location                      = var.location
  name                          = "${var.assetprefix}-viaCato"
  resource_group_name           = azurerm_resource_group.azure-rg.name
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}
resource "azurerm_route" "public-rt" {
  address_prefix      = "23.102.135.246/32"
  name                = "Microsoft-KMS"
  next_hop_type       = "Internet"
  resource_group_name = azurerm_resource_group.azure-rg.name
  route_table_name    = "${var.assetprefix}-viaCato"
  depends_on = [
    azurerm_route_table.private-rt
  ]
}
resource "azurerm_route" "lan-route" {
  address_prefix         = "0.0.0.0/0"
  name                   = "default"
  next_hop_in_ip_address = var.lan_ip
  next_hop_type          = "VirtualAppliance"
  resource_group_name    = azurerm_resource_group.azure-rg.name
  route_table_name       = "${var.assetprefix}-viaCato"
  depends_on = [
    azurerm_route_table.private-rt
  ]
}
resource "azurerm_route_table" "public-rt" {
  disable_bgp_route_propagation = true
  location                      = var.location
  name                          = "${var.assetprefix}-viaInternet"
  resource_group_name           = azurerm_resource_group.azure-rg.name
  depends_on = [
    azurerm_resource_group.azure-rg
  ]
}
resource "azurerm_route" "route-internet" {
  address_prefix      = "0.0.0.0/0"
  name                = "default"
  next_hop_type       = "Internet"
  resource_group_name = azurerm_resource_group.azure-rg.name
  route_table_name    = "${var.assetprefix}-viaInternet"
  depends_on = [
    azurerm_route_table.public-rt
  ]
}

resource "azurerm_subnet_route_table_association" "rt-table-association-mgmt" {
  route_table_id = azurerm_route_table.public-rt.id
  subnet_id      = azurerm_subnet.subnet-mgmt.id
  depends_on = [
    azurerm_route_table.public-rt,
    azurerm_subnet.subnet-mgmt
  ]
}

resource "azurerm_subnet_route_table_association" "rt-table-association-wan" {
  route_table_id = azurerm_route_table.public-rt.id
  subnet_id      = azurerm_subnet.subnet-wan.id
  depends_on = [
    azurerm_route_table.public-rt,
    azurerm_subnet.subnet-wan,
  ]
}

resource "azurerm_subnet_route_table_association" "rt-table-association-lan" {
  route_table_id = azurerm_route_table.private-rt.id
  subnet_id      = azurerm_subnet.subnet-lan.id
  depends_on = [
    azurerm_route_table.private-rt,
    azurerm_subnet.subnet-lan
  ]
}

## The following attributes are exported:
output "resource-group-name" { value = azurerm_resource_group.azure-rg.name }
output "mgmt-nic-id" { value = azurerm_network_interface.mgmt-nic.id }
output "wan-nic-id" { value = azurerm_network_interface.wan-nic.id }
output "lan-nic-id" { value = azurerm_network_interface.lan-nic.id }

```
</details>

<details>
<summary>Create Azure vSocket - Example Module</summary>

### Create Azure vSocket - Example Module

```hcl
## Azure vSocket Example Module Variables
variable "assetprefix" {
  type = string
  description = "Your asset prefix for resources created"
  default = null
}

variable "location" { 
  type = string
  description = "(Required) The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created."
  default = null
}

variable "resource-group-name" { 
  type = string
  description = "(Required) The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created."
  default = null
}

variable "socketsite-serial" { 
  type = string
  default = "AA-BB-CC-DD-EE-FF"
}

variable "mgmt-nic-id" {
  type = string
  default = null
}

variable "wan-nic-id" {
  type = string
  default = null
}

variable "lan-nic-id" {
  type = string
  default = null
}

variable "image_reference_id" {
	type = string
  description = "Path to image used to deploy specific version of the virutal socket"
	default = "/Subscriptions/38b5ec1d-b3b6-4f50-a34e-f04a67121955/Providers/Microsoft.Compute/Locations/eastus/Publishers/catonetworks/ArtifactTypes/VMImage/Offers/cato_socket/Skus/public-cato-socket/Versions/19.0.17805"
}

## Vsocket module Resources
provider "azurerm" {
	features {}
}

## Create Vsocket Virtual Machine
resource "azurerm_virtual_machine" "vsocket" {
  location                     = var.location
  name                         = "${var.assetprefix}-vSocket"
  network_interface_ids        = [var.mgmt-nic-id, var.wan-nic-id, var.lan-nic-id]
  primary_network_interface_id = var.mgmt-nic-id
  resource_group_name          = var.resource-group-name
  vm_size                      = "Standard_D8ls_v5"
  plan {
    name      = "public-cato-socket"
    product   = "cato_socket"
    publisher = "catonetworks"
  }
  boot_diagnostics {
    enabled     = true
    storage_uri = ""
  }
  storage_os_disk {
    create_option     = "Attach"
    name              = "${var.assetprefix}-vSocket-disk1"
    managed_disk_id   = azurerm_managed_disk.vSocket-disk1.id
    os_type = "Linux"
  }
  
  depends_on = [
    azurerm_managed_disk.vSocket-disk1
  ]
}

resource "azurerm_managed_disk" "vSocket-disk1" {
  name                 = "${var.assetprefix}-vSocket-disk1"
  location             = var.location
  resource_group_name  = var.resource-group-name
  storage_account_type = "Standard_LRS"
  create_option        = "FromImage"
  disk_size_gb         = 8
  os_type              = "Linux"
  image_reference_id   = var.image_reference_id
}

variable "commands" {
  type    = list(string)
  default = [
    "rm /cato/deviceid.txt",
    "rm /cato/socket/configuration/socket_registration.json",
    "nohup /cato/socket/run_socket_daemon.sh &"
   ]
}

resource "azurerm_virtual_machine_extension" "vsocket-custom-script" {
  auto_upgrade_minor_version = true
  name                       = "vsocket-custom-script"
  publisher                  = "Microsoft.Azure.Extensions"
  type                       = "CustomScript"
  type_handler_version       = "2.1"
  virtual_machine_id         = azurerm_virtual_machine.vsocket.id
  settings = <<SETTINGS
 {
  "commandToExecute": "${"echo '${var.socketsite-serial}' > /cato/serial.txt"};${join(";", var.commands)}"
 }
SETTINGS
  depends_on = [
    azurerm_virtual_machine.vsocket
  ]
}

```
</details>

### cato-oss_socket VPC and Vsocket Module Usage

```hcl

module "vnet" {
  source = "./1-vnet"
  location = var.location
  project_name = var.project_name
  assetprefix = var.assetprefix
  subnet_range_mgmt = var.subnet_range_mgmt
  subnet_range_wan = var.subnet_range_wan
  subnet_range_lan = var.subnet_range_lan
  lan_ip = var.lan_ip
  vnet_prefix = var.vnet_prefix
}

module "vSocket" {
  source = "./2-vSocket"
  location = var.location
  assetprefix = var.assetprefix
  resource-group-name = module.vnet.resource-group-name
  socketsite-serial = data.cato-oss_accountSnapshotSite.azure-site.info.sockets[0].serial
  mgmt-nic-id = module.vnet.mgmt-nic-id
  wan-nic-id = module.vnet.wan-nic-id
  lan-nic-id = module.vnet.lan-nic-id
}

```

