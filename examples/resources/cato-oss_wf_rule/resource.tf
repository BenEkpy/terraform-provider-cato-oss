// wan firewall allowing all & logs
resource "cato-oss_wf_rule" "allow_all_and_log" {
  at = {
    position = "LAST_IN_POLICY"
  }
  rule = {
    name      = "Allow all & logs"
    enabled   = true
    action    = "ALLOW"
    direction = "BOTH"
    tracking = {
      event = {
        enabled = true
      }
    }
  }
}

// all SMBV3 for all domain users to the site named Datacenter
resource "cato-oss_wf_rule" "allow_smbv3_to_dc" {
  at = {
    position = "LAST_IN_POLICY"
  }
  rule = {
    name      = "Allow SMBv3 to DC"
    enabled   = true
    action    = "ALLOW"
    direction = "TO"
    source = {
      users_group = [
        {
          name = "Domain Users"
        }
      ]
    }
    destination = {
      site = [
        {
          name = "Datacenter"
        }
      ]
    }
    service = {
      standard = [
        {
          name = "SMBV3"
        }
      ]
    }
    tracking = {
      event = {
        enabled = true
      }
    }
  }
}


// block all remote users except "Marketing" using category domain "test.com"
resource "cato-oss_if_rule" "block_test_com_for_remote_users" {
  at = {
    position = "FIRST_IN_POLICY"
  }
  rule = {
    name              = "Block Test.com for Remote Users"
    enabled           = true
    action            = "BLOCK"
    connection_origin = "REMOTE"
    destination = {
      Domain = [
        "test.com"
      ]
    }
    tracking = {
      event = {
        enabled = true
      }
    }
    exceptions = [
      {
        name = "Exclude Marketing Teams"
        source = {
          users_group = [
            {
              name = "Marketing-Teams"
            }
          ]
        }
      }
    ]
  }
}

