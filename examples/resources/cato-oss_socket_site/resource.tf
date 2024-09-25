// socket site for AWS
resource "cato-oss_socket_site" "aws_site" {
  name            = "aws_site"
  description     = "site description"
  site_type       = "DATACENTER"
  connection_type = "SOCKET_AWS1500"

  native_range = {
    native_network_range = "192.168.25.0/24"
    local_ip             = "192.168.25.5"
  }

  site_location = {
    country_code = "FR"
    timezone     = "Europe/Paris"
  }
}

// socket site x1500 with DHCP settings
resource "cato-oss_socket_site" "branch_site" {
  name            = "branch_site"
  description     = "site description"
  site_type       = "BRANCH"
  connection_type = "SOCKET_X1500"

  native_range = {
    native_network_range = "192.168.20.0/24"
    local_ip             = "192.168.20.1"
    dhcp_settings = {
      dhcp_type = "DHCP_RANGE"
      ip_range  = "192.168.20.10-192.168.20.22"
    }
  }

  site_location = {
    country_code = "FR"
    timezone     = "Europe/Paris"
  }
}