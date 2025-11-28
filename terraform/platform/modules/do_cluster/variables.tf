variable "plan_space_name" {
  type        = string
  default     = "tfplan-space"
  description = "Remote store for terraform state"
}
variable "database_size" {
  default = "db-s-1vcpu-1gb"
}

variable "cluster_name" {
  type        = string
  default     = "development"
  description = "Name of the Kubernetes cluster"
}

variable "region" {
  type        = string
  description = "DigitalOcean region for the cluster"
}

variable "space_region" {
  type        = string
  default     = "nyc3"
  description = "DigitalOcean region for the cluster"
}

variable "cluster_version" {
  type        = string
  default     = "1.33.1-do.5"
  description = "Kubernetes version"
}

variable "ha_enabled" {
  type        = bool
  default     = true
  description = "Enable high availability for the cluster"
}

variable "primary_pool" {
  description = "Primary node pool (required)"
  type = object({
    name   = string
    size   = string
    count  = number
    labels = optional(map(string))
  })
}

variable "node_pools" {
  description = "Additional node pools"
  type = list(object({
    name   = string
    size   = string
    count  = number
    labels = optional(map(string))
  }))
  default = []
}

variable "private_network_uuid" {
  type        = string
  description = "ID of the private network to use for the cluster"
}
