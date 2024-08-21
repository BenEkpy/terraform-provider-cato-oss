terraform {
  required_providers {
    cato-oss = {
      source = "BenEkpy/cato-oss"
    }
  }
}

provider "cato-oss" {
    baseurl = "https://api.catonetworks.com/api/v1/graphql2"
    token = "TOKEN_VALUE"
    account_id = "ACCOUNT_ID_VALUE"
}

resource "cato-oss_if_policy" "rname" {
  at = {
    position = "LAST_IN_POLICY"
  }
  // enabled = false
  publish = true
  sdk_key_name = "KEY_NANE_VALUE"
  rule = {
      name = "RULE_NAME_VALUE"
      description = "RULE_DESCRIPTION_VALUE"
      enabled =  false
      section = {}
      source = {
        ip = ["10.0.0.1"]
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
          }
        ]
        customcategory = []
        sanctionedappscategory = []
        country = []
        domain = []
        fqdn = []
        ip = ["10.0.0.99"]
        subnet = []
        iprange = []
        globaliprange = []
        remoteasn = []
      }
      service = {}
      action = "ALLOW"
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
