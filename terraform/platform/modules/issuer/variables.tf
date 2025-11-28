// variables email, server, private_key_secret, ingress_class

variable "environment" {
  type    = string
  default = "staging"
}
variable "email" {}
variable "server" {
  type    = string
  default = "https://acme-staging-v02.api.letsencrypt.org/directory"
}
variable "private_key_secret" {
  type    = string
  default = "letsencrypt-staging"
}
variable "ingress_class_name" {
  type    = string
  default = "nginx"
}
