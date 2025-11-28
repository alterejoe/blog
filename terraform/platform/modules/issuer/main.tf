terraform {
  required_providers {

    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.7.0"
    }
  }
}

resource "kubectl_manifest" "issuer" {
  yaml_body = templatefile("${path.module}/templates/letsencrypt.yaml", {
    environment        = var.environment
    email              = var.email
    server             = var.server
    private_key_secret = var.private_key_secret
    ingress_class      = var.ingress_class_name
  })
}
