module "notes_database" {
  source = "../../apps/database/"

  postgres_host           = var.postgres_host
  postgres_port           = var.postgres_port
  postgres_user           = var.postgres_user
  postgres_password       = var.postgres_password
  postgres_database       = var.postgres_database
  postgres_sslmode        = var.postgres_sslmode
  postgres_schema         = var.postgres_schema
  migrator_password       = "migrator"
  admin_runtime_password  = "admin_runtime"
  client_runtime_password = "client_runtime"
}
