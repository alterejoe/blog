resource "postgresql_role" "admin" {
  name                      = "admin"
  login                     = false
  inherit                   = false
  bypass_row_level_security = false

  lifecycle {
    ignore_changes = [roles]
  }
}

resource "postgresql_role" "migrator" {
  name                      = "migrator"
  login                     = true
  inherit                   = true
  bypass_row_level_security = false
  password                  = var.migrator_password

  lifecycle {
    ignore_changes = [roles]
  }
}

resource "postgresql_role" "client_runtime" {
  name                      = "client_runtime"
  login                     = true
  inherit                   = true
  bypass_row_level_security = false
  password                  = var.client_runtime_password

}

resource "postgresql_role" "admin_runtime" {
  name                      = "admin_runtime"
  login                     = true
  inherit                   = true
  bypass_row_level_security = false
  password                  = var.admin_runtime_password

}

resource "postgresql_grant_role" "migrator_can_assume_admin" {
  role       = postgresql_role.admin.name
  grant_role = postgresql_role.migrator.name

}


