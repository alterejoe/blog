variable "namespace" {
  description = "Namespace to deploy nginx ingress controller into"
  type        = string
  default     = "ingress-nginx"
}

variable "chart_version" {
  description = "Helm chart version for nginx-ingress"
  type        = string
  default     = "4.13.0"
}

variable "service_type" {
  description = "Type of service for controller (LoadBalancer, NodePort, etc.)"
  type        = string
  default     = "LoadBalancer"
}
