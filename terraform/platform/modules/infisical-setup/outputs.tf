
#
# output "infisical_secret_yaml" {
#   value     = <<EOT
# apiVersion: v1
# kind: Secret
# metadata:
#   name: infisical-secrets
#   namespace: ${var.namespace}
# type: Opaque
# stringData:
#   AUTH_SECRET: ${local.auth_secret}
#   ENCRYPTION_KEY: ${local.encryption_key}
#   REDIS_URL: ${digitalocean_database_cluster.infisical_valkey.private_uri}
#   DB_CONNECTION_URI: ${digitalocean_database_cluster.infisical_postgres.private_uri}?ssl=true&sslmode=require&rejectUnauthorized=false
#   SITE_URL: ${var.site_url}
# EOT
#   sensitive = true
# }
