#••••••••••••••••••••••••••••••••••••••••••• PostgreSQL for Infisical 
resource "digitalocean_database_cluster" "infisical_postgres" {
  name                 = var.pg_name
  region               = var.region
  size                 = var.pg_size
  node_count           = var.pg_node_count
  engine               = "pg"
  version              = var.pg_version
  private_network_uuid = var.private_network_uuid
  lifecycle {
    prevent_destroy = true
  }
}

# Valkey (Redis-compatible) 
resource "digitalocean_database_cluster" "infisical_valkey" {
  name                 = var.valkey_name
  region               = var.region
  size                 = var.valkey_size
  node_count           = var.valkey_node_count
  version              = var.valkey_version
  engine               = "valkey"
  private_network_uuid = var.private_network_uuid
  lifecycle {
    prevent_destroy = true
  }
}


output "postgres_private_uri" {
  value     = digitalocean_database_cluster.infisical_postgres.private_uri
  sensitive = true
}

output "valkey_private_uri" {
  value     = digitalocean_database_cluster.infisical_valkey.private_uri
  sensitive = true
}

output "postgres_id" {
  value = digitalocean_database_cluster.infisical_postgres.id
}

output "valkey_id" {
  value = digitalocean_database_cluster.infisical_valkey.id
}
