resource "cato-oss_static_host" "host1" {
    site_id = cato-oss_socket_site.site1.id
    name = "your-hostname-here"
    ip = "192.168.25.24"
}