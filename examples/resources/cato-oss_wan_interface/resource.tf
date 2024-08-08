# Create a WAN Interface
resource "cato-oss_wan_interface" "wan1" {
    site_id = cato-oss_socket_site.site1.id
    interface_id = "WAN1"
    name = "WAN 01 test"
    upstream_bandwidth = 1000
    downstream_bandwidth = 1000
    role = "wan_1"
    precedence = "ACTIVE"
}