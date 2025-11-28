# variable "plan_space_name" {
#   type        = string
#   default     = "tfplan-space"
#   description = "Remote store for terraform state"
# }
# variable "database_size" {
#   default = "db-s-1vcpu-1gb"
# }
#
# variable "cluster_name" {
#   type        = string
#   default     = "development"
#   description = "Name of the Kubernetes cluster"
# }
#
# variable "region" {
#   type        = string
#   description = "DigitalOcean region for the cluster"
# }
#
# variable "space_region" {
#   type        = string
#   default     = "nyc3"
#   description = "DigitalOcean region for the cluster"
# }
#
# variable "cluster_version" {
#   type        = string
#   default     = "1.33.1-do.5"
#   description = "Kubernetes version"
# }
#
# variable "ha_enabled" {
#   type        = bool
#   default     = false
#   description = "Enable high availability for the cluster"
# }
#
# variable "node_pool_name" {
#   type        = string
#   default     = "worker-pool"
#   description = "Name of the node pool"
# }
#
# variable "node_size" {
#   type        = string
#   default     = "s-1vcpu-2gb"
#   description = "Droplet size slug for worker nodes"
# }
#
# variable "node_count" {
#   type        = number
#   default     = 3
#   description = "Number of nodes in the pool"
# }
#
# variable "private_network_uuid" {
#   type        = string
#   description = "ID of the private network to use for the cluster"
# }
