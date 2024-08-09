# Create Network Range
resource "cato-oss_network_range" "vlan1" {
    site_id = cato-oss_socket_site.site1.id
    name = "VLAN 100 network"
    range_type = "VLAN"
    subnet = "10.20.101.0/24"
    local_ip = "10.20.101.100"
    vlan = 100
}