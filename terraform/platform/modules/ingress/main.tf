terraform {
  required_providers {

    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.7.0"
    }
  }
}
resource "kubectl_manifest" "ingress" {
  yaml_body = templatefile("${path.module}/templates/ingress.yaml", {
    name               = var.name
    namespace          = var.namespace
    host               = var.host
    service_name       = var.service_name
    service_port       = var.service_port
    cluster_issuer     = var.cluster_issuer
    tls_secret_name    = "${var.name}-tls"
    ingress_class_name = var.ingress_class_name
  })
}

