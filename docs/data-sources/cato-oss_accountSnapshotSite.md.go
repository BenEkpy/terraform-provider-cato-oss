---
page_title: "cato-oss__accountSnapshotSite Data Source - terraform-provider-cato-oss"
subcategory: "Provider Reference"
description: |-
Provides a cato-oss_accountSnapshotSite data source
---

# cato-oss_accountSnapshotSite (Resource)

The `cato-oss_accountSnapshotSite` data source contains the configuration parameters necessary to
retrieve a site serial number.
Documentation for the underlying API used in this resource can be found at
[query.accountSnapshot()](https://api.catonetworks.com/documentation/#query-accountSnapshot).

## Example Usage

<details>
<summary>cato-oss_accountSnapshotSite Data Source Variables</summary>

### cato-oss_accountSnapshotSite Data Source Variables

```hcl
# Provider variables
variable cato_token {}

variable "account_id" {
    description = "Account ID"
    type        = number
    # default	  = 12345
}

# VLAN cato-oss_network_range variables
variable "site_id" {
    type = number
    description = "Site ID"
    default = 12345
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

### Data Source ###
data "cato-oss_accountSnapshotSite" "aws-dev-site" {
	id = var.site_id
}

```

### Required

- `id` (Int) Cato Site ID

### Optional

- `gateway` (String) Network range gateway address
- `vlan` (Int) Network VLAN

### Read-Only

- `id` (String) Identifier for the site
