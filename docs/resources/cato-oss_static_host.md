---
page_title: "cato-oss_static_host Resource - terraform-provider-cato-oss"
subcategory: "Provider Reference"
description: |-
  Provides a cato-oss_socketsite resource  
---

# cato-oss_static_host (Resource)

The `cato-oss_static_host` resource contains the configuration parameters necessary to 
add a static host. 
Documentation for the underlying API used in this resource can be found at
[mutation.addStaticHost()](https://api.catonetworks.com/documentation/#mutation-site.addStaticHost).

## Example Usage

<details>
<summary>cato-oss_static_host Resource Variables</summary>

### cato-oss_static_host Resource Variables

```hcl
# Provider variables
variable cato_token {}

variable "account_id" {
    description = "Account ID"
    type        = number
    # default	  = 12345
}

# VLAN cato-oss_static_host variables
variable "static_host_name" {
    type = string
    description = "Static hostmame"
    default = "static-hostname-here"
}

variable "static_host_ip" {
    type = string
    description = "Static host ip address"
    default = "192.168.25.24"
}

variable "static_host_mac" {
    type = string
    description = "Static host mac address"
    default = "00:00:00:00:00:00"
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
resource "cato-oss_static_host" "host1" {
    site_id = cato-oss_socket_site.site1.id
    name = var.static_host_name
    ip = var.static_host_ip
    mac = var.static_host_mac
}
```

### Required

- `name` (Int) Static host hostname
- `ip` (String) Static host ip address
- `site_id` (Int) Site ID to add static host to

### Optional

- `mac` (String) Static host mac address

### Read-Only

- `id` (String) Identifier for the site
