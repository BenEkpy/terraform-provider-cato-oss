---
subcategory: "Example Modules"
layout: "cato-oss"
page_title: "cato-oss_socketsite AWS VPC Resources and Modules"
description: |-
  Provides an combined example of creating a virtual socket site in Cato Management Application, and templates for creating a VPC and deploying a virtual socket instance in AWS.
---

# cato-oss_socketsite (Resource)

The `cato-oss_socketsite` resource contains the configuration parameters necessary to 
add a socket site to the Cato cloud 
([virtual socket in AWS/Azure, or physical socket](https://support.catonetworks.com/hc/en-us/articles/4413280502929-Working-with-X1500-X1600-and-X1700-Socket-Sites)).
Documentation for the underlying API used in this resource can be found at
[mutation.addSocketSite()](https://api.catonetworks.com/documentation/#mutation-site.addSocketSite).

## Example Usage

### cato-oss_socket Resource and Variables

```hcl
## cato-oss_socket Variables
variable cato_token {}

variable "account_id" {
    description = "Account ID"
    type        = number
    default	  = 12345
}

variable "connection_type" {
    type = string
    description = "Socket model connection type"
    validation {
        condition = contains(["SOCKET_AWS1500","SOCKET_AZ1500","SOCKET_ESX1500","SOCKET_X1500","SOCKET_X1600","SOCKET_X1600_LTE","SOCKET_X1700"])
        error_message = "The connection_type variable must be one of the following: 'SOCKET_AWS1500','SOCKET_AZ1500','SOCKET_ESX1500','SOCKET_X1500','SOCKET_X1600','SOCKET_X1600_LTE','SOCKET_X1700'"
    }
    description = "SOCKET_AWS1500"
}

variable "description" {
    type = string
    description = "Site description"
    default = "Your site description"
}

variable "name" {
    type = string
    description = "Site name"
    default = "Your site name"
}

variable "native_network_range" {
    type = string
    description = "Native network range"
    default = "172.16.1.0/24"
}

variable "country_code" {
    type = string
    description = "Country code"
    default = "US"
}

variable "state_code" {
    type = string
    description = "State code"
    default = "US-VA"
}

variable "timezone" {
    type = string
    description = "Timezone"
    default = "America/New_York"
}

variable "site_type" {
    type = string
    description = "Site type"
    validation {
        condition = contains(["BRANCH","CLOUD_DC","DATACENTER","HEADQUARTERS"])
        error_message = "The connection_type variable must be one of the following: 'BRANCH','CLOUD_DC','DATACENTER','HEADQUARTERS'"
    }
    default = "BRANCH"
}

## cato-oss_socket Provider and Resource

provider "cato-oss" {
    baseurl = "https://api.catonetworks.com/api/v1/graphql2"
    token = var.cato_token
    account_id = var.account_id
}

resource "cato-oss_socketsite" "aws-dev-site" {
    account_id = var.account_id
    connection_type  = "SOCKET_AWS1500"
    description = var.site_description
    name = var.project_name
    native_network_range = var.vpc_range
    site_location = {
        country_code = "US"
		state_code = "US-VA"
        timezone = "America/New_York"
    }
    site_type = "DATACENTER"
}
```

<details>
<summary>Create AWS VPC - Example Module</summary>

### Create AWS VPC - Example Module

```hcl
## AWS VPC Example Module Variables
variable "region" { 
  type = string
  default = "us-east-2" 
}

variable "project_name" { 
  type = string
  default = "Cato Lab" 
}

variable "ingress_cidr_blocks" { 
  type = list
  default = null
}

variable "subnet_range_mgmt" { 
  type = string
  default = null
}

variable "subnet_range_wan" { 
  type = string
  default = null
}

variable "subnet_range_lan" { 
  type = string
  default = null
}

variable "vpc_range" { 
  type = string
  default = null
}

## AWS VPC Example Module Resources

provider "aws" {
  region = var.region
}

resource "aws_vpc" "cato-lab" {
  cidr_block = var.vpc_range
  tags = {
    Name = "${var.project_name}-VPC"
  }
}

// Lookup data from region and VPC
data "aws_availability_zones" "available" {
  state = "available"
}

# Internet Gateway and Attachment
resource "aws_internet_gateway" "internet_gateway" {}

resource "aws_internet_gateway_attachment" "attach_gateway" {
  internet_gateway_id = aws_internet_gateway.internet_gateway.id
  vpc_id = aws_vpc.cato-lab.id
}

# Subnets
resource "aws_subnet" "mgmt_subnet" {
  vpc_id = aws_vpc.cato-lab.id
  cidr_block = var.subnet_range_mgmt
  availability_zone = data.aws_availability_zones.available.names[0]
  tags = {
    Name = "${var.project_name}-MGMT-Subnet"
  }
}

resource "aws_subnet" "wan_subnet" {
  vpc_id = aws_vpc.cato-lab.id
  cidr_block = var.subnet_range_wan
  availability_zone = data.aws_availability_zones.available.names[0]
  tags = {
    Name = "${var.project_name}-WAN-Subnet"
  }
}

resource "aws_subnet" "lan_subnet" {
  vpc_id = aws_vpc.cato-lab.id
  cidr_block = var.subnet_range_lan
  availability_zone = data.aws_availability_zones.available.names[0]
  tags = {
    Name = "${var.project_name}-LAN-Subnet"
  }
}

# Internal and External Security Groups
resource "aws_security_group" "internal_sg" {
  name = "${var.project_name}-Internal-SG"
  description = "CATO LAN Security Group - Allow all traffic Inbound"
  vpc_id = aws_vpc.cato-lab.id
  ingress = [
    {
      description 		= "Allow all traffic Inbound from Ingress CIDR Blocks"
	  protocol         	= -1
      from_port 		= 0
      to_port 			= 0
      cidr_blocks 	   	= var.ingress_cidr_blocks
      ipv6_cidr_blocks 	= []
      prefix_list_ids 	= []
      security_groups 	= []
      self 				= false
    }
  ]
  egress = [
	{
	  description 		= "Allow all traffic Outbound"
	  protocol 			= -1
	  from_port 		= 0
	  to_port 			= 0
	  cidr_blocks 	   	= ["0.0.0.0/0"]
	  ipv6_cidr_blocks 	= []
      prefix_list_ids 	= []
      security_groups 	= []
	  self 				= false
	}
  ]
  tags = {
    name = "${var.project_name}-Internal-SG"
  }
}

resource "aws_security_group" "external_sg" {
  name = "${var.project_name}-External-SG"
  description = "CATO WAN Security Group - Allow HTTPS In"
  vpc_id = aws_vpc.cato-lab.id
  ingress = [
    {
      description 		= "Allow HTTPS In"
	  protocol 			= "tcp"
      from_port 		= 443
      to_port 			= 443
      cidr_blocks 	   	= var.ingress_cidr_blocks
      ipv6_cidr_blocks 	= []
      prefix_list_ids 	= []
      security_groups 	= []
      self 				= false
    },
    {
      description 		= "Allow SSH In"
      protocol 			= "tcp"
      from_port 		= 22
      to_port 			= 22
      cidr_blocks 	   	= var.ingress_cidr_blocks
      ipv6_cidr_blocks 	= []
      prefix_list_ids 	= []
      security_groups 	= []
      self 				= false
    }
  ]
  egress = [
	{
	  description 		= "Allow all traffic Outbound"
	  protocol 			= -1
	  from_port 		= 0
	  to_port 			= 0
	  cidr_blocks 	   	= ["0.0.0.0/0"]
	  ipv6_cidr_blocks 	= []
      prefix_list_ids 	= []
      security_groups 	= []
	  self 				= false
	}
  ]
  tags = {
    name = "${var.project_name}-External-SG"
  }
}

##The following attributes are exported:
output "internet_gateway_id" { value = aws_internet_gateway.internet_gateway.id }
output "project_name" { value = var.project_name }
output "sg_internal" { value = aws_security_group.internal_sg.id }
output "sg_external" { value = aws_security_group.external_sg.id }
output "subnet_id_mgmt" { value = aws_subnet.mgmt_subnet.id }
output "subnet_id_wan" { value = aws_subnet.wan_subnet.id }
output "subnet_id_lan" { value = aws_subnet.lan_subnet.id }
output "vpc_id" { value = aws_vpc.cato-lab.id }
```
</details>

<details>
<summary>Create AWS vSocket - Example Module</summary>

### Create AWS vSocket - Example Module

```hcl
## AWS VPC Example Module Variables
variable "ebs_disk_size" {
  description = "Size of disk"
  type        = number
  default	  = 32
}

variable "ebs_disk_type" {
  description = "Size of disk"
  type        = string
  default	  = "gp2"
}

variable "instance_type" {
  description = "The instance type of the vSocket"
  type        = string
  default	 = "c5.xlarge"
  validation {
    condition = contains(["d2.xlarge","c3.xlarge","t3.large","t3.xlarge","c4.xlarge","c5.xlarge","c5d.xlarge","c5n.xlarge"], var.instance_type)
    error_message = "The instance_type variable must be one of 'd2.xlarge','c3.xlarge','t3.large','t3.xlarge','c4.xlarge','c5.xlarge','c5d.xlarge','c5n.xlarge'."
  }
}

variable "internet_gateway_id" {
  description = ""
  type        = string
  default	  = null
}

variable "key_pair" {
  description = "Name of an existing Key Pair for AWS encryption"
  type        = string
  default	  = null
}

variable "new_mgmt_eni" {
  description = "Choose an IP Address within the Management Subnet. You CANNOT use the first four assignable IP addresses within the subnet as it's reserved for the AWS virtual router interface. The accepted input format is X.X.X.X"
  type        = string
  default	  = null
}

variable "new_wan_eni" {
  description = "Choose an IP Address within the Public/WAN Subnet. You CANNOT use the first four assignable IP addresses within the subnet as it's reserved for the AWS virtual router interface. The accepted input format is X.X.X.X"
  type        = string
  default	  = null
}

variable "new_lan_eni" {
  description = "Choose an IP Address within the LAN Subnet. You CANNOT use the first four assignable IP addresses within the subnet as it's reserved for the AWS virtual router interface. The accepted input format is X.X.X.X"
  type        = string
  default	  = null
}

variable "project_name" {
  type = string
  description = "Your Cato Deployment Name Here"
  default = "Your Cato Deployment Name Here"
}

variable "region" { 
  type = string
  default = "us-east-2" 
}

variable "serial_number" {
  description = ""
  type        = string
  default	  = null
}

variable "subnet_id_mgmt" {
  description = ""
  type        = string
  default	  = null
}

variable "subnet_id_wan" {
  description = ""
  type        = string
  default	  = null
}

variable "subnet_id_lan" {
  description = ""
  type        = string
  default	  = null
}

variable "security_group_ingress" {
  description = ""
  type        = string
  default	  = null
}

variable "sg_external" {
  description = ""
  type        = string
  default	  = null
}

variable "sg_internal" {
  description = ""
  type        = string
  default	  = null
}

variable "vpc_id" {
  description = ""
  type        = string
  default	  = null
}

## Vsocket module Resources
provider "aws" {
  region = var.region
}

// Lookup data from region and VPC
data "aws_ami" "vSocket" {
  most_recent      = true
  name_regex       = "VSOCKET_AWS"
  owners           = ["679593333241"]
}

data "aws_availability_zones" "available" {
  state = "available"
}

// vSocket Network Interfaces
resource "aws_network_interface" "waneni" {
  source_dest_check = "true"
  subnet_id = var.subnet_id_wan
  private_ips = [var.new_wan_eni]
  tags = {
    Name = "${var.project_name}-WAN-INT"
  }
}

resource "aws_network_interface" "mgmteni" {
  source_dest_check = "true"
  subnet_id = var.subnet_id_mgmt
  private_ips = [var.new_mgmt_eni]
  tags = {
    Name = "${var.project_name}-MGMT-INT"
  }
}

resource "aws_network_interface" "laneni" {
  source_dest_check = "false"
  subnet_id = var.subnet_id_lan
  private_ips = [var.new_lan_eni]
  tags = {
    Name = "${var.project_name}-LAN-INT"
  }
}

// Elastic IP Addresses
resource "aws_eip" "wanip" {
  tags = {
    Name = "${var.project_name}-WAN-EIP"
  }
}

resource "aws_eip" "mgmteip" {
  tags = {
    Name = "${var.project_name}-MGMT-EIP"
  }
}

// Elastic IP Addresses Association - Required to properly destroy 
resource "aws_eip_association" "wanip_assoc" {
  network_interface_id = aws_network_interface.waneni.id
  allocation_id = aws_eip.wanip.id
}

resource "aws_eip_association" "mgmteip_assoc" {
  network_interface_id = aws_network_interface.mgmteni.id
  allocation_id = aws_eip.mgmteip.id
}

// Routing Tables
resource "aws_route_table" "wanrt" {
  vpc_id = var.vpc_id
  tags = {
    Name = "${var.project_name}-WAN-RT"
  }
}

resource "aws_route_table" "mgmtrt" {
  vpc_id = var.vpc_id
  tags = {
    Name = "${var.project_name}-MGMT-RT"
  }
}

resource "aws_route_table" "lanrt" {
  vpc_id = var.vpc_id
  tags = {
    Name = "${var.project_name}-LAN-RT"
  }
}

// Routes
resource "aws_route" "wan_route" {
  route_table_id = aws_route_table.wanrt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id = var.internet_gateway_id
}

resource "aws_route" "mgmt_route" {
  route_table_id = aws_route_table.mgmtrt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id = var.internet_gateway_id
}

resource "aws_route" "lan_route" {
  route_table_id = aws_route_table.lanrt.id
  destination_cidr_block = "0.0.0.0/0"
  network_interface_id = aws_network_interface.laneni.id
}

// Route Table Associations
resource "aws_route_table_association" "mgmt_subnet_route_table_association" {
  subnet_id = var.subnet_id_mgmt
  route_table_id = aws_route_table.mgmtrt.id
}

resource "aws_route_table_association" "wan_subnet_route_table_association" {
  subnet_id = var.subnet_id_wan
  route_table_id = aws_route_table.wanrt.id
}

resource "aws_route_table_association" "lan_subnet_route_table_association" {
  subnet_id = var.subnet_id_lan
  route_table_id = aws_route_table.lanrt.id
}

# vSocket Instance
resource "aws_instance" "vSocket" {
  depends_on = [
	aws_network_interface.waneni,
	aws_network_interface.mgmteni,
	aws_network_interface.laneni
  ]
  tenancy 				 = "default"
  ami 					 = data.aws_ami.vSocket.id
  availability_zone 	 = data.aws_availability_zones.available.names[0]
  key_name 				 = var.key_pair
  instance_type = var.instance_type
  user_data              = base64encode(var.serial_number)
  # Network Interfaces
  # WANENI
  network_interface {
    device_index         = 0
    network_interface_id = aws_network_interface.waneni.id
  }
  # MGMTENI
  network_interface {
    device_index         = 1
    network_interface_id = aws_network_interface.mgmteni.id
  }
  # LANENI
  network_interface {
    device_index         = 2
    network_interface_id = aws_network_interface.laneni.id
  }
  // CF Property(UserData) = base64encode(var.serial_number)
  ebs_block_device {
	device_name 		 = "/dev/sda1"
	volume_size 		 = var.ebs_disk_size
	volume_type 		 = var.ebs_disk_type
  }
  tags = {
    Name 				 = "${var.project_name}-vSocket"
  }
}

```
</details>

### cato-oss_socket VPC and Vsocket Module Usage

```hcl

module "vpc" {
	source = "./1-vpc"
	region = var.region 
	project_name = var.project_name
	vpc_range = var.vpc_range
	subnet_range_mgmt = var.subnet_range_mgmt
	subnet_range_wan = var.subnet_range_wan
	subnet_range_lan = var.subnet_range_lan
	ingress_cidr_blocks = var.ingress_cidr_blocks
}

module "vSocket" {
	source = "./2-vSocket"
	internet_gateway_id = module.vpc.internet_gateway_id
	instance_type = var.instance_type
	key_pair = var.key_pair
	new_mgmt_eni = var.new_mgmt_eni
	new_wan_eni = var.new_wan_eni
	new_lan_eni = var.new_lan_eni
	project_name = var.project_name
	region = var.region
	serial_number = data.cato-oss_accountSnapshot.aws-dev-site.sites[0].info.sockets[0].serial
	sg_external = module.vpc.sg_external
	sg_internal = module.vpc.sg_internal
	subnet_id_mgmt = module.vpc.subnet_id_mgmt
	subnet_id_wan = module.vpc.subnet_id_wan
	subnet_id_lan = module.vpc.subnet_id_lan
	vpc_id = module.vpc.vpc_id
}

```

### Required

- `account_id` (String) Cato Account ID (can be found into the URL on the CMA)
- `connection_type` (String) Connection type for the site (SOCKET_X1500, SOCKET_AWS1500, SOCKET_AZ1500, ...)
- `name` (String) Site name
- `native_network_range` (String) Site native IP range (CIDR)
- `site_location` (Attributes) Site location (see [below for nested schema](#nestedatt--site_location))
- `site_type` (String) Site type ('BRANCH','CLOUD_DC','DATACENTER','HEADQUARTERS')

### Optional

- `description` (String) Site description

### Read-Only

- `id` (String) Identifier for the site

<a id="nestedatt--site_location"></a>
### Nested Schema for `site_location`

Required:

- `country_code` (String) Site country code (can be retrieve from entityLookup)
- `timezone` (String) Site timezone (can be retrieve from entityLookup)

Optional:

- `state_code` (String) Optionnal site state code(can be retrieve from entityLookup)
