
resource "random_bytes" "encryption_key" {
  length = 16
}

resource "random_password" "auth_secret" {
  length  = 32
  special = false
}

locals {

  encryption_key = random_bytes.encryption_key.hex
  auth_secret    = base64encode(random_password.auth_secret.result)
}

data "digitalocean_database_ca" "infisical_postgres" {
  cluster_id = digitalocean_database_cluster.infisical_postgres.id
}

data "digitalocean_database_ca" "infisical_valkey" {
  cluster_id = digitalocean_database_cluster.infisical_valkey.id
}

resource "random_password" "infisical_admin" {
  length  = 100
  special = false
}
locals {
  db_root_cert = sensitive(data.digitalocean_database_ca.infisical_postgres.certificate)
  infisical_creds = sensitive({
    AUTH_SECRET         = local.auth_secret
    ENCRYPTION_KEY      = local.encryption_key
    SITE_URL            = var.site_url
    DB_CONNECTION_URI   = "${digitalocean_database_cluster.infisical_postgres.private_uri}?sslmode=require"
    DB_ROOT_CERT        = local.db_root_cert
    REDIS_URL           = digitalocean_database_cluster.infisical_valkey.private_uri
    NODE_EXTRA_CA_CERTS = "/certs/do-ca.crt"
  })
  infisical_bootstrap_creds = sensitive({
    INFISICAL_ADMIN_EMAIL = var.bootstrap_email
    INFISICAL_ADMIN_PASSWORD = random_password.infisical_admin.result
  })
}

module "infisical_secret" {
  source      = "../../modules/sealed-secret"
  kubeconfig_path = var.kubeconfig_path
  namespace   = var.namespace
  secret_name = "infisical-secrets"
  string_data = local.infisical_creds
  # outfile     = "${path.root}/unsealed/infisical-secrets.yaml"
}

module "infisical_db_certs" {
  source      = "../../modules/sealed-secret"
  kubeconfig_path = var.kubeconfig_path
  namespace   = var.namespace
  secret_name = "infisical-db-certs"
  # outfile     = "${path.root}/unsealed/infisical-db-certs.yaml"

  string_data = {
    "do-ca.crt" = sensitive(data.digitalocean_database_ca.infisical_postgres.certificate)
  }
}

module "infisical_bootstrap_credentials" {
  source      = "../../modules/sealed-secret"
  kubeconfig_path = var.kubeconfig_path
  namespace   = var.namespace
  secret_name = "infisical-bootstrap-credentials"
  string_data = sensitive(local.infisical_bootstrap_creds)
}

output "infisical_admin_password" {
  value     = random_password.infisical_admin.result
  sensitive = true
}

# output "infisical_org_id" {
#   value     = jsondecode(base64decode(kubernetes_secret.infisical_bootstrap.data.INFISICAL_ORG_ID))
#   sensitive = true
# }
