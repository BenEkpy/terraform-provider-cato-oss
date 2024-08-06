---
page_title: "cato-oss_socketsite Resource - terraform-provider-cato-oss"
subcategory: "Provider Reference"
description: |-
  Provides a cato-oss_socketsite resource  
---

# cato-oss_socketsite (Resource)

The `cato-oss_socketsite` resource contains the configuration parameters necessary to 
add a socket site to the Cato cloud 
([virtual socket in AWS/Azure, or physical socket](https://support.catonetworks.com/hc/en-us/articles/4413280502929-Working-with-X1500-X1600-and-X1700-Socket-Sites)).
Documentation for the underlying API used in this resource can be found at
[mutation.addSocketSite()](https://api.catonetworks.com/documentation/#mutation-site.addSocketSite).

## Example Usage

<details>
<summary>cato-oss_socket Resource Variables</summary>

### cato-oss_socket Resource Variables

```hcl
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
resource "cato-oss_socketsite" "aws-dev-site" {
    account_id = var.account_id
    connection_type  = var.connection_type
    description = var.description
    name = var.project_name
    native_network_range = var.vpc_range
    site_location = {
        country_code = var.country_code
		state_code = var.state_code
        timezone = var.timezone
    }
    site_type = "DATACENTER"
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
