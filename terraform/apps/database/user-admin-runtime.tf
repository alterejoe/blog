# # Tables that need read/write but NO delete
# resource "postgresql_grant" "admin_read_write_tables" {
#   role        = postgresql_role.admin_runtime.name
#   database    = var.postgres_database
#   schema      = postgresql_schema.clerk_schema.name
#   objects     = [] # alphabetically sorted
#   privileges  = ["INSERT", "SELECT", "UPDATE"]
#   object_type = "table"
#   depends_on = [
#     postgresql_role.admin_runtime,
#     postgresql_schema.clerk_schema
#   ]
# }
#
# # Tables that need full access including DELETE
# resource "postgresql_grant" "admin_full_access_tables" {
#   role        = postgresql_role.admin_runtime.name
#   database    = var.postgres_database
#   schema      = postgresql_schema.clerk_schema.name
#   objects     = [] # alphabetically sorted
#   privileges  = ["DELETE", "INSERT", "SELECT", "UPDATE"]
#   object_type = "table"
#   depends_on = [
#     postgresql_role.admin_runtime,
#     postgresql_schema.clerk_schema,
#     postgresql_grant.admin_read_write_tables # explicit dependency
#   ]
# }
#
# # Default privileges for future tables
# resource "postgresql_default_privileges" "admin_runtime_default" {
#   owner       = postgresql_role.migrator.name
#   role        = postgresql_role.admin_runtime.name
#   database    = var.postgres_database
#   schema      = postgresql_schema.clerk_schema.name
#   object_type = "table"
#   privileges  = ["INSERT", "SELECT", "UPDATE"]
# }
#
# # Schema usage
# resource "postgresql_grant" "admin_runtime_schema" {
#   role        = postgresql_role.admin_runtime.name
#   database    = var.postgres_database
#   schema      = postgresql_schema.clerk_schema.name
#   object_type = "schema"
#   privileges  = ["USAGE"]
#   depends_on = [
#     postgresql_role.admin_runtime
#   ]
# }
