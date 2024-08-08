---
page_title: "cato-oss_wan_interface Resource - terraform-provider-cato-oss"
subcategory: "Provider Reference"
description: |-
  Provides a cato-oss_socketsite resource  
---

# cato-oss_wan_interface (Resource)

The `cato-oss_wan_interface` resource contains the configuration parameters necessary to 
add a wan interface to a socket. 
([virtual socket in AWS/Azure, or physical socket](https://support.catonetworks.com/hc/en-us/articles/4413280502929-Working-with-X1500-X1600-and-X1700-Socket-Sites)).
Documentation for the underlying API used in this resource can be found at
[mutation.updateSocketInterface()](https://api.catonetworks.com/documentation/#mutation-site.updateSocketInterface).

## Example Usage

<details>
<summary>cato-oss_wan_interface Resource Variables</summary>

### cato-oss_socket Resource Variables

```hcl
# Provider variables
variable cato_token {}

variable "account_id" {
    description = "Account ID"
    type        = number
    # default	  = 12345
}

# cato-oss_wan_interface variables
variable "wan1_interface_id" {
    type = string
    description = "WAN interface id"
    validation {
        condition = contains(["INT_1","INT_10","INT_11","INT_12","INT_2","INT_3","INT_4","INT_5","INT_6","INT_7","INT_8","INT_9","LAN1","LAN2","LTE","USB1","USB2","WAN1","WAN2","WLAN"],var.wan1_interface_id)
        error_message = "The connection_type variable must be one of the following: 'INT_1','INT_10','INT_11','INT_12','INT_2','INT_3','INT_4','INT_5','INT_6','INT_7','INT_8','INT_9','LAN1','LAN2','LTE','USB1','USB2','WAN1','WAN2','WLAN'"
    }
    default = "WAN1"
}

variable "wan1_interface_name" {
    type = string
    description = "WAN interface name"
    default = "WAN interface 1"
}

variable "wan1_upstream_bandwidth" {
    type = number
    description = "WAN interface upstream bandwidth"
    default = 1000
}

variable "wan1_downstream_bandwidth" {
    type = number
    description = "WAN interface downstream bandwidth"
    default = 1000
}

variable "wan1_role" {
    type = number
    description = "WAN interface upstream bandwidth"
    validation {
        condition = contains(["wan_1","wan_2","wan_3","wan_4"],var.wan1_role)
        error_message = "The wan1_role variable must be one of the following: 'wan_1','wan_2','wan_3','wan_4'"
    }
    default = "wan_1"
}

variable "wan1_precedence" {
    type = number
    description = "WAN interface upstream bandwidth"
    validation {
        condition = contains(["ACTIVE","LAST_RESORT","PASSIVE"],var.wan1_precedence)
        error_message = "The wan1_precedence variable must be one of the following: 'ACTIVE','LAST_RESORT','PASSIVE'"
    }
    default = "ACTIVE"
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
resource "cato-oss_wan_interface" "wan1" {
    site_id = cato-oss_socket_site.yoursite.id
    interface_id = var.wan1_interface_id
    name = var.wan1_interface_name
    upstream_bandwidth = var.wan1_upstream_bandwidth
    downstream_bandwidth = var.wan1_downstream_bandwidth
    role = var.wan1_role
    precedence = var.wan1_precedence
}
```

### Required

- `downstream_bandwidth` (Int) Maximum allowed bandwidth for traffic on this port, from the Cato Cloud to the site
- `interface_id` (Int) Socket interface id ('INT_1','INT_10','INT_11','INT_12','INT_2','INT_3','INT_4','INT_5','INT_6','INT_7','INT_8','INT_9','LAN1','LAN2','LTE','USB1','USB2','WAN1','WAN2','WLAN')
- `name` (String) Interface name
- `precedence` (String) WAN interface precedence ('ACTIVE','LAST_RESORT','PASSIVE')
- `role` (String) WAN interface role ('wan_1','wan_2','wan_3','wan_4')
- `site_id` (Int) Cato Account ID (can be found into the URL on the CMA)
- `upstream_bandwidth` (Int) Maximum allowed bandwidth on this port, for traffic from the site to the Cato Cloud

### Read-Only

- `id` (String) Identifier for the site
