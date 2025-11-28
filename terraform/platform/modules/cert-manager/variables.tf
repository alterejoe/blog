
variable "namespace" {
  type    = string
  default = "cert-manager"
}

variable "chart_version" {
  type    = string
  default = "v1.19.1" # latest according to docs :contentReference[oaicite:1]{index=1}
}

variable "install_crds" {
  type    = bool
  default = false
}
