# Create a AWS Socket site
resource "cato_socketsite" "aws_site_1" {
    account_id = "1234"
    name = "aws_site_1"
    site_type = "DATACENTER"
    connection_type = "SOCKET_AWS1500"
    native_network_range = "192.168.0.0/16"
    site_location = {
        country_code = "FR",
        timezone = "Europe/Paris"
    }
}