variable "name" {}
variable "namespace" { default = "default" }
variable "host" {}
variable "service_name" {}
variable "service_port" { default = 80 }
variable "cluster_issuer" { default = "letsencrypt-staging" }
variable "ingress_class_name" { default = "nginx" }
