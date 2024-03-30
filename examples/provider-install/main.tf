terraform {
  required_providers {
    cato = {
      source = "registry.terraform.io/benekpy/cato-oss"
    }
  }
}

provider "cato" {}

resource "cato_socketsite" "site1" {
    account_id = "5242"
    name = "test22"
    site_type = "BRANCH"
    connection_type = "SOCKET_AWS1500"
    native_network_range = "192.168.122.0/24"
    site_location = {
        country_code = "FR",
        timezone = "Europe/Paris"
    }
}

data "cato_accountSnapshot" "site1" {
    account_id = "5242"
    site_id = cato_socketsite.site1.id
}