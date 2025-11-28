
terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

resource "digitalocean_kubernetes_cluster" "main" {
  name     = var.cluster_name
  region   = var.region
  version  = var.cluster_version
  ha       = false
  vpc_uuid = var.private_network_uuid

  node_pool {
    name       = var.primary_pool.name
    size       = var.primary_pool.size
    node_count = var.primary_pool.count
    labels     = lookup(var.primary_pool, "labels", null)
  }
}

# Additional node pools
resource "digitalocean_kubernetes_node_pool" "extra" {
  for_each   = { for np in var.node_pools : np.name => np }
  cluster_id = digitalocean_kubernetes_cluster.main.id
  name       = each.value.name
  size       = each.value.size
  node_count = each.value.count
  labels     = lookup(each.value, "labels", null)
}

output "endpoint" {
  value = digitalocean_kubernetes_cluster.main.endpoint
}

output "cluster_ca_certificate" {
  value     = digitalocean_kubernetes_cluster.main.kube_config[0].cluster_ca_certificate
  sensitive = true
}

output "token" {
  value     = digitalocean_kubernetes_cluster.main.kube_config[0].token
  sensitive = true
}

output "kubeconfig" {
  value     = digitalocean_kubernetes_cluster.main.kube_config[0].raw_config
  sensitive = true
}
