---
page_title: "cato-oss_if_policy Resource - terraform-provider-cato-oss"
subcategory: "Provider Reference"
description: |-
  Provides a cato-oss_if_policy resource  
---

# cato-oss_if_policy (Resource)

The `cato-oss_if_policy` resource contains the configuration parameters necessary to 
add a network range to a cato site. 
([virtual socket in AWS/Azure, or physical socket](https://support.catonetworks.com/hc/en-us/articles/4413280502929-Working-with-X1500-X1600-and-X1700-Socket-Sites)).
Documentation for the underlying API used in this resource can be found at
[mutation.addNetworkRange()](https://api.catonetworks.com/documentation/#mutation-site.addNetworkRange).

## Example Usage

<details>
<summary>cato-oss_if_policy Resource Variables</summary>

### cato-oss_if_policy Resource Variables

```hcl
# Provider variables
variable cato_token {}

variable "account_id" {
    description = "Account ID"
    type        = number
    default	    = 12345
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
resource "cato-oss_if_policy" "example-block-rule" {
  at = {
    position = "LAST_IN_POLICY"
  }
  publish = true
  sdk_key_name = "your-api-key-name-here"
  rule = {
      name = "Example-Block-Rule"
      description = "Example-Block-Rule"
      enabled =  true
      source = {
        ip = []
        host = []
        site = []
        subnet = []
        iprange = []
        globaliprange = []
        networkinterface = []
        sitenetworksubnet = []
        floatingsubnet = []
        user = []
        usersgroup = []
        group = []
        systemgroup = []

      }
      connectionorigin = "ANY"
      country = []
      device = []
      deviceos = []
      destination = {
        application = []
        customapp = []
        appcategory = [
          {
            by = "NAME"
            input = "Beauty"
          },
          {
            by = "NAME"
            input = "SPAM"
          },
          {
            by = "NAME"
            input = "Hacking"
          }
        ]
        customcategory = []
        sanctionedappscategory = []
        country = []
        domain = []
        fqdn = []
        ip = []
        subnet = []
        iprange = []
        globaliprange = []
        remoteasn = []
      }
      service = {}
      action = "BLOCK"
      tracking = {
        event = {
          enabled = true
        }
        alert = {
          enabled = false
          frequency = "DAILY"
          subscriptiongroup = []
          webhook = []
          mailinglist = []
        }

      }
      schedule = {
        activeon = "ALWAYS"
      }
      exceptions = []
  }
}
```

### Required

- `at.position` (String) = Position relative to a policy, a section or another rule. Possible Values: AFTER_RULE,BEFORE_RULE,FIRST_IN_SECTION,LAST_IN_SECTION,FIRST_IN_POLICY,LAST_IN_POLICY
- `publish` (Boolean) = true
- `sdk_key_name` (String) = "your-api-key-name-here"
- `rule.name` (String) = "Example-Block-Rule"
- `rule.description` (String) = "Example-Block-Rule"
- `rule.enabled` (Boolean) =  true
- `rule.source` (Object) = Source traffic matching criteria. Logical ‘OR’ is applied within the criteria set. Logical ‘AND’ is applied between criteria sets.
- `rule.source.ip` [String] = IPv4 addresses
- `rule.source.host` (String) = Hosts and servers defined for your account.
- `rule.source.site` [String] = Site defined for the account.
- `rule.source.subnet` [String] = Subnets and network ranges defined for the LAN interfaces of a site.
- `rule.source.iprange` [String] = Multiple separate IP addresses or an IP range.
- `rule.source.globaliprange` [String] = Globally defined IP range, IP and subnet objects.
- `rule.source.networkinterface` [String] = Network range defined for a site.
- `rule.source.sitenetworksubnet` [String] = GlobalRange + InterfaceSubnet.
- `rule.source.floatingsubnet` [String] = Floating Subnets (ie. Floating Ranges) are used to identify traffic exactly matched to the route advertised by BGP. They are not associated with a specific site. This is useful in scenarios such as active-standby high availability routed via BGP.
- `rule.source.user` [String] = Individual users defined for the account
- `rule.source.usersgroup` [String] = Group of users.
- `rule.source.group` [String] = Groups defined for your account.
- `rule.source.systemgroup` [String] = Predefined Cato groups.
- `rule.connectionorigin` (String) = Connection origin of the traffic. ANY,REMOTE,SITE.
- `rule.country` [String] = Source country traffic matching criteria. Logical ‘OR’ is applied within the criteria set. Logical ‘AND’ is applied between criteria sets.
- `rule.device` [String] = Source Device Profile traffic matching criteria. Logical ‘OR’ is applied within the criteria set. Logical ‘AND’ is applied between criteria sets.
- `rule.deviceos` [String] = Source device Operating System traffic matching criteria. Logical ‘OR’ is applied within the criteria set. Logical ‘AND’ is applied between criteria sets.
- `rule.destination` (Object) = Destination traffic matching criteria. Logical ‘OR’ is applied within the criteria set. Logical ‘AND’ is applied between criteria sets.
- `rule.destination.application` [String] = Applications for the rule (pre-defined)
- `rule.destination.customapp` [String] = Custom (user-defined) applications
- `rule.destination.appcategory` [String] = Cato category of applications which are dynamically updated by Cato.
- `rule.destination.appcategory.rule. (Object) = Cato category of applications which are dynamically updated by Cato
- `rule.destination.appcategory.rule.by` (String) = Defines the object identification method by ID (default) or by name. Example: "NAME".
- `rule.destination.appcategory.rule.input` (String) = The object identification (ID or name) value. Example: "Beauty".
- `rule.destination.customcategory` [String] = Groups of objects such as predefined and custom applications, predefined and custom services, domains, FQDNs etc.
- `rule.destination.sanctionedappscategory` [String] = Predefined groups of objects such as predefined and custom applications, predefined and custom services, domains, FQDNs etc.
- sanctionedAppsCategory.
- `rule.destination.country` [String] = Country
- `rule.destination.domain` [String] = A Second-Level Domain (SLD). It matches all Top-Level Domains (TLD), and subdomains that include the Domain. Example: example.com.
- `rule.destination.fqdn` [String] = An exact match of the fully qualified domain (FQDN). Example: www.my.example.com.
- `rule.destination.ip` [String] = IPv4 addresses
- `rule.destination.subnet` [String] = Network subnets in CIDR notation
- `rule.destination.iprange` [String] = A range of IPs. Every IP within the range will be matched.
- `rule.destination.globaliprange` [String] = Globally defined IP range, IP and subnet objects.
- `rule.destination.remoteasn` [String] = 16 bit autonomous system number [0-65535].
- `rule.service` (String) = {}
- `rule.action` (String) = The action applied by the Internet Firewall if the rule is matched, example: "BLOCK". Possible Values: BLOCK,ALLOW,PROMPT,RBI
- `rule.tracking` (String) = {
- `rule.tracking.event` (Object) = Input of data if an alert is sent for a rule.
- `rule.tracking.event.enabled` (Boolean) = true
- `rule.tracking.alert` (Object) = Input of data for the alert settings for the rule
- `rule.tracking.alert.enabled` (String) = TRUE – send alerts when the rule is matched, FALSE – don’t send alerts when the rule is matched.
- `rule.tracking.alert.frequency` (String) = Returns data for the alert frequency. Possible Values: HOURLY,DAILY,WEEKLY,IMMEDIATE.
- `rule.tracking.alert.subscriptiongroup` [String] = Returns data for the Subscription Group that receives the alert.
- `rule.tracking.alert.subscriptiongroup.by` (String) = Defines the object identification method by ID (default) or by name. Example: "ID".
- `rule.tracking.alert.subscriptiongroup.input` (String) = The object identification (ID or name) value. Example: "xyz123".
- `rule.tracking.alert.webhook` [String] = Webhook
- `rule.tracking.alert.webhook.by` (String) = Defines the object identification method by ID (default) or by name. Example: "ID".
- `rule.tracking.alert.webhook.input` (String) = The object identification (ID or name) value. Example: "xyz123".
- `rule.tracking.alert.mailinglist` [String] = Returns data for the Mailing List that receives the alert.
- `rule.tracking.alert.mailinglist.by` (String) = Defines the object identification method by ID (default) or by name. Example: "ID".
- `rule.tracking.alert.mailinglist.input` (String) = The object identification (ID or name) value. Example: "xyz123".

### Read-Only

- `id` (String) Identifier for the site
