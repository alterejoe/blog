variable "namespace" {
  description = "Namespace where the secret will be created"
  type        = string
}

variable "secret_name" {
  description = "Name of the secret to seal"
  type        = string
}

variable "string_data" {
  description = "Map of key-value pairs for the secret's stringData"
  type        = map(string)
}

variable "controller_namespace" {
  description = "Namespace of the sealed-secrets controller"
  type        = string
  default     = "kube-system"
}

variable "controller_name" {
  description = "Name of the sealed-secrets controller"
  type        = string
  default     = "sealed-secrets"
}

variable "outfile" {
  description = "Optional path to save sealed YAML to disk"
  type        = string
  default     = ""
}

# variable "kubeconfig_b64" {
#   type = string
# }
variable "kubeconfig_path" {
  type = string
}
