
# --- Common ---
variable "region" {
  type        = string
  description = "DigitalOcean region for all resources"
  default     = "nyc1"
}

variable "namespace" {
  type        = string
  description = "Kubernetes namespace for Infisical"
  default     = "secrets"
}

# --- Database: PostgreSQL ---
variable "pg_name" {
  type    = string
  default = "infisical-postgres"
}

variable "pg_size" {
  type    = string
  default = "db-s-1vcpu-1gb"
}

variable "pg_node_count" {
  type    = number
  default = 1
}

variable "pg_version" {
  type    = string
  default = "16"
}

# --- Database: Valkey ---
variable "valkey_name" {
  type    = string
  default = "infisical-valkey"
}

variable "valkey_size" {
  type    = string
  default = "db-s-1vcpu-1gb"
}

variable "valkey_node_count" {
  type    = number
  default = 1
}

variable "valkey_version" {
  type    = string
  default = "8"
}


variable "site_url" {
  type = string
  # default = "https://secrets.replaceme.com"
}

variable "private_network_uuid" {
  type        = string
  description = "DigitalOcean private network UUID"
}

variable "bootstrap_email" {
  type = string
  description = "Bootstrap email"
}

variable "reset_infisical" {
  type    = bool
  default = false
}

variable "kubeconfig_path" {
  type    = string
}
