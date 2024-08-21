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

data "cato-oss_internet_fw_policy" "test" {}

output "test_internet_fw_policy" {
  value = data.cato-oss_internet_fw_policy.test
}
