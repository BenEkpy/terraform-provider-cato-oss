terraform {
  required_providers {
    cato-oss = {
      source = "BenEkpy/cato-oss"
    }
  }
}

provider "cato-oss" {
    baseurl = "https://api.catonetworks.com/api/v1/graphql2"
    token = "TOKEN_VAUE"
    account_id = "ACCOUNT_ID_VALUE"
}

resource "cato-oss_ipsec_site" "ipsecsite1" {
  name = "TestIPSecSite001"
  site_type = "BRANCH"
  description = "TestIPSecSite001 description"
  nativenetworkrange = "192.168.99.0/24"
  sitelocation = {
      countrycode = ""
      statecode = ""
      timezone = ""
      address =  ""
      city = ""
  }
}
