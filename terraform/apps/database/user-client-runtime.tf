resource "postgresql_grant" "client_read_write_permission_tables" {
  role        = postgresql_role.client_runtime.name
  database    = var.postgres_database
  schema      = postgresql_schema.clerk_schema.name
  objects     = ["users"]
  privileges  = ["INSERT", "SELECT", "UPDATE"]
  object_type = "table"

  depends_on = [
    postgresql_role.client_runtime,
    postgresql_schema.clerk_schema
  ]
}
#
# resource "postgresql_grant" "client_read_only_tables" {
#   role        = postgresql_role.client_runtime.name
#   database    = var.postgres_database
#   schema      = postgresql_schema.clerk_schema.name
#   objects     = []
#   privileges  = ["SELECT"]
#   object_type = "table"
#
#   depends_on = [
#     postgresql_role.client_runtime,
#     postgresql_schema.clerk_schema,
#     postgresql_grant.client_read_write_permission_tables
#   ]
# }
#
resource "postgresql_grant" "client_full_access_tables" {
  role        = postgresql_role.client_runtime.name
  database    = var.postgres_database
  schema      = postgresql_schema.clerk_schema.name
  objects     = ["casbin_rule", "sessions", "user_sessions"]
  privileges  = ["DELETE", "INSERT", "SELECT", "UPDATE"]
  object_type = "table"

  depends_on = [
    postgresql_role.client_runtime,
    postgresql_schema.clerk_schema,
    postgresql_grant.client_runtime_schema
  ]
}

resource "postgresql_grant" "client_runtime_schema" {
  role        = postgresql_role.client_runtime.name
  database    = var.postgres_database
  schema      = postgresql_schema.clerk_schema.name
  object_type = "schema"
  privileges  = ["USAGE"]
  depends_on = [
    postgresql_role.client_runtime
  ]
}
