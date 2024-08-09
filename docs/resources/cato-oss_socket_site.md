---
page_title: "cato-oss_socket_site Resource - terraform-provider-cato-oss"
subcategory: "Provider Reference"
description: |-
  Provides a cato-oss_socket_site resource  
---

# cato-oss_socket_site (Resource)

The `cato-oss_socket_site` resource contains the configuration parameters necessary to 
add a socket site to the Cato cloud 
([virtual socket in AWS/Azure, or physical socket](https://support.catonetworks.com/hc/en-us/articles/4413280502929-Working-with-X1500-X1600-and-X1700-Socket-Sites)).
Documentation for the underlying API used in this resource can be found at
[mutation.addSocketSite()](https://api.catonetworks.com/documentation/#mutation-site.addSocketSite).

## Example Usage

<details>
<summary>cato-oss_socket_site Resource Variables</summary>

### cato-oss_socket_site Resource Variables

```hcl
# Provider variables
variable cato_token {}

variable "account_id" {
    description = "Account ID"
    type        = number
    # default	  = 12345
}

# cato-oss_socket_site variables
variable "connection_type" {
    type = string
    description = "Socket model connection type"
    validation {
        condition = contains(["SOCKET_AWS1500","SOCKET_AZ1500","SOCKET_ESX1500","SOCKET_X1500","SOCKET_X1600","SOCKET_X1600_LTE","SOCKET_X1700"],var.connection_type)
        error_message = "The connection_type variable must be one of the following: 'SOCKET_AWS1500','SOCKET_AZ1500','SOCKET_ESX1500','SOCKET_X1500','SOCKET_X1600','SOCKET_X1600_LTE','SOCKET_X1700'"
    }
    default = "SOCKET_AWS1500"
}

variable "native_network_range" {
    type = string
    description = "Native network range"
    default = "192.168.25.0/24"
}

variable "local_ip" {
    type = string
    description = "Socket local ip"
    default = "192.168.25.100"
}

variable "dhcp_type" {
    type = string
    description = "Network range dhcp type"
    validation {
        condition = contains(["ACCOUNT_DEFAULT","DHCP_RANGE","DHCP_DISABLED","DHCP_RELAY"],var.dhcp_type)
        error_message = "The dhcp_type variable must be one of the following: 'ACCOUNT_DEFAULT','DHCP_RANGE','DHCP_DISABLED','DHCP_RELAY'"
    }
    default = "DHCP_RANGE"
}

variable "dchp_range" {
    type = string
    description = "Socket local ip"
    default = "192.168.25.10-192.168.25.22"
}

variable "site_description" {
    type = string
    description = "Site description"
    default = "Your site description"
}

variable "site_name" {
    type = string
    description = "Site name"
    default = "Your site name"
}

variable "site_type" {
    type = string
    description = "Site type"
    validation {
        condition = contains(["BRANCH","CLOUD_DC","DATACENTER","HEADQUARTERS"],var.site_type)
        error_message = "The connection_type variable must be one of the following: 'BRANCH','CLOUD_DC','DATACENTER','HEADQUARTERS'"
    }
    default = "BRANCH"
}

## siteLocation variables
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
resource "cato-oss_socket_site" "site1" {
    name = var.site_name
    description = var.site_description
    site_type = var.site_type
    connection_type = var.connection_type

    native_range = {
      native_network_range = var.native_network_range
      local_ip = var.local_ip
      dhcp_settings ={
        dhcp_type = var.dhcp_type
        ip_range = var.dchp_range
      }
    }

    site_location = {
        country_code = var.country_code,
        state_code = var.state_code,
        timezone = var.timezone
    }
}
```

### Required

- `account_id` (String) Cato Account ID (can be found into the URL on the CMA)
- `connection_type` (String) Connection type for the site (SOCKET_X1500, SOCKET_AWS1500, SOCKET_AZ1500, ...)
- `description` (String) Site description 
- `name` (String) Site name
- `native_range.native_network_range` (String) Site native IP range (CIDR)
- `native_range.local_ip` (String) Local socket local lan ip address
- `native_range.dhcp_settings.dhcp_type` (String) Network range dhcp type (ACCOUNT_DEFAULT,DHCP_RANGE,DHCP_DISABLED,DHCP_RELAY)
- `native_range.dhcp_settings.ip_range` (String) Network dhcp IP range, example: "192.168.1.10-192.168.1.100"
- `site_location.country_code` (String) Site country code (can be retrieved from entityLookup)
- `site_location.timezone` (String) Timezone
- `site_type` (String) Site type ('BRANCH','CLOUD_DC','DATACENTER','HEADQUARTERS')

### Optional

- `description` (String) Site description
- `site_location.city` (String) Optional site city name (can be retrieved from entityLookup)
- `site_location.state_code` (String) Optionnal site state code(can be retrieve from entityLookup)

### Read-Only

- `id` (String) Identifier for the site

