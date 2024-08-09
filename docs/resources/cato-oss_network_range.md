---
page_title: "cato-oss_wan_interface Resource - terraform-provider-cato-oss"
subcategory: "Provider Reference"
description: |-
  Provides a cato-oss_socketsite resource  
---

# cato-oss_wan_interface (Resource)

The `cato-oss_network_range` resource contains the configuration parameters necessary to 
add a network range to a cato site. 
([virtual socket in AWS/Azure, or physical socket](https://support.catonetworks.com/hc/en-us/articles/4413280502929-Working-with-X1500-X1600-and-X1700-Socket-Sites)).
Documentation for the underlying API used in this resource can be found at
[mutation.addNetworkRange()](https://api.catonetworks.com/documentation/#mutation-site.addNetworkRange).

## Example Usage

<details>
<summary>cato-oss_network_range Resource Variables</summary>

### cato-oss_network_range Resource Variables

```hcl
# Provider variables
variable cato_token {}

variable "account_id" {
    description = "Account ID"
    type        = number
    # default	  = 12345
}

# VLAN cato-oss_network_range variables
variable "vlan_network_range_name" {
    type = string
    description = "VLAN network range name"
    default = "VLAN 100 network"
}

variable "vlan_network_range_subnet" {
    type = string
    description = "VLAN network range subnet"
    default = "10.20.101.0/24"
}

variable "vlan_network_range_local_ip" {
    type = string
    description = "VLAN network range local ip"
    default = "10.20.101.100"
}

variable "vlan_network_range_vlan" {
    type = number
    description = "Network range VLAN"
    default = 100
}

# Routed cato-oss_network_range variables
variable "routed_network_range_name" {
    type = string
    description = "Routed network range name"
    default = "Routed network range"
}

variable "routed_network_range_subnet" {
    type = string
    description = "Routed network range subnet"
    default = "10.20.100.0/24"
}

variable "routed_network_range_local_ip" {
    type = string
    description = "Routed network range local ip"
    default = "10.20.101.0/24"
}

variable "routed_network_range_gateway" {
    type = string
    description = "Routed network range gateway"
    default = "192.168.25.254"
}

```
</details>

## Providers and Resources

```hcl
## Providers ###
provider "cato-oss" {
    baseurl = "https://api.catonetworks.com/api/v1/graphql2"
    token = var.cato_token
    account_id = var.account_id
}

### Resource ###
resource "cato-oss_network_range" "vlan1" {
    site_id = cato-oss_socket_site.site1.id
    name = var.vlan_network_range_name
    range_type = "VLAN"
    subnet = var.vlan_network_range_subnet
    local_ip = var.vlan_network_range_local_ip
    vlan = var.vlan_network_range_vlan
}

resource "cato-oss_network_range" "routed1" {
    site_id = cato-oss_socket_site.site1.id
    name = var.routed_network_range_name
    range_type = "Routed"
    subnet = var.routed_network_range_subnet
    gateway = var.routed_network_range_gateway
}
```

### Required

- `local_ip` (String) Network range local ip address
- `name` (String) Interface name
- `range_type` (String) Network range type ('Direct','Native','Routed','SecondaryNative','VLAN')
- `site_id` (Int) Cato Account ID (can be found into the URL on the CMA)
- `subnet` (String) Network range subnet

### Optional

- `gateway` (String) Network range gateway address
- `vlan` (Int) Network VLAN

### Read-Only

- `id` (String) Identifier for the site
