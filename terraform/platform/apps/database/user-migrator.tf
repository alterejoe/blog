resource "postgresql_grant" "migrator_schema" {
  role        = postgresql_role.migrator.name
  database    = var.postgres_database
  schema      = postgresql_schema.clerk_schema.name
  object_type = "schema"
  privileges  = ["USAGE", "CREATE"]
}


