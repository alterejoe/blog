# terraform {
#   required_providers {
#     digitalocean = {
#       source  = "digitalocean/digitalocean"
#       version = "~> 2.0"
#     }
#   }
# }
#
# resource "digitalocean_kubernetes_cluster" "main" {
#   name     = var.cluster_name
#   region   = var.region
#   version  = var.cluster_version
#   ha       = var.ha_enabled
#   vpc_uuid = var.private_network_uuid
#   node_pool {
#     name       = var.node_pool_name
#     size       = var.node_size
#     node_count = var.node_count
#   }
# }
#
#
# output "endpoint" {
#   value = digitalocean_kubernetes_cluster.main.endpoint
# }
#
# output "cluster_ca_certificate" {
#   value     = digitalocean_kubernetes_cluster.main.kube_config[0].cluster_ca_certificate
#   sensitive = true
# }
#
# output "token" {
#   value     = digitalocean_kubernetes_cluster.main.kube_config[0].token
#   sensitive = true
# }
#
# output "kubeconfig" {
#   value     = digitalocean_kubernetes_cluster.main.kube_config[0].raw_config
#   sensitive = true
# }
