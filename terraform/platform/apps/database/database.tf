# data "postgresql_schemas" "clerk_db" {
#   database = var.postgres_database
# }


# resource "postgresql_database" "clerk_db" {
#   name  = "db"
#   owner = "postgres"
# }

resource "postgresql_schema" "clerk_schema" {
  name       = var.postgres_schema
  owner      = postgresql_role.migrator.name
  database   = var.postgres_database
  depends_on = [postgresql_role.migrator]
}
