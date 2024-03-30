# List all site
data "cato_entities" "sites" {
    account_id = "1234"
    entity_type = "site"
}