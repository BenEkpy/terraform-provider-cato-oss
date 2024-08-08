# Create a AWS Socket site
resource "cato-oss_socket_site" "site1" {
    name = "site1-tf"
    description = "site1 branch"
    site_type = "BRANCH"
    connection_type = "SOCKET_AWS1500"

    native_range = {
      native_network_range = "192.168.25.0/24"
      local_ip = "192.168.25.100"
      dhcp_settings ={
        dhcp_type = "DHCP_RANGE"
        ip_range = "192.168.25.10-192.168.25.22"
      }
    }

    site_location = {
        country_code = "FR",
        timezone = "Europe/Paris"
    }
}