variable "cli_image_tag" {
  type = string
  description = "Infisical CLI container image tag"
  default     = "0.41.97"
}
variable "site_url" {
  type        = string
  description = "Infisical site URL"
  # default     = "https://infisical.com"
}
variable "app_image_tag" {
  type        = string
  description = "Infisical container image tag"
  default     = "v0.146.0-postgres"
}

variable "replica_count" {
  type        = number
  description = "Number of Infisical replicas"
  # default     = 2
}


variable "namespace" {
  type        = string
  description = "Kubernetes namespace for Infisical"
}

variable "auto_bootstrap_enabled" {
  type        = bool
  description = "Enable auto-bootstrap of Infisical"
  default     = false
}

variable "ingress_enabled" {
  type        = bool
  description = "Enable ingress for Infisical"
}

variable "ingress_host" {
  type        = string
  description = "Ingress hostname for Infisical"
  # default     = ""
}

variable "service_type" {
  type        = string
  description = "Kubernetes service type (ClusterIP/LoadBalancer)"
  default     = "ClusterIP"
}

variable "service_node_port" {
  type        = number
  description = "Kubernetes service node port"
  default     = 8080
}

# --- Helm ---
variable "helm_name" {
  type = string
  # default = "infisical"
}

variable "helm_repo" {
  type = string
  # default = "https://dl.cloudsmith.io/public/infisical/helm-charts/helm/charts/"
}

variable "helm_chart" {
  type = string
  # default = "infisical-standalone"
}

variable "helm_chart_version" {
  type        = string
  description = "Helm chart version of Infisical to deploy"
  # default     = "1.17.0"
}

#organization
variable "organization_name" {
  type        = string
  description = "Organization name"
  # default     = "Infisical"
}

# bootstrap-secret-name
variable "bootstrap_secret_credentials" {
  type        = string
  description = "Bootstrap secret name"
  default     = "infisical-bootstrap-credentials"
}

# bootstrap-secret-name
variable "bootstrap_secret_name" {
  type        = string
  description = "Bootstrap secret name"
  default     = "infisical-bootstrap-secret"
}
