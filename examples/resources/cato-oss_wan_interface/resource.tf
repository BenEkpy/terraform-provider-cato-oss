// wan interface active
resource "cato-oss_wan_interface" "wan1" {
  site_id              = cato-oss_socket_site.site1.id
  interface_id         = "WAN1"
  name                 = "Interface WAN 1"
  upstream_bandwidth   = "100"
  downstream_bandwidth = "100"
  role                 = "wan_1"
  precedence           = "ACTIVE"
}