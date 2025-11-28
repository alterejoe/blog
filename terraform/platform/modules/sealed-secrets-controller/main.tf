terraform {
  required_providers {
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.17"
    }
  }
}

resource "helm_release" "sealed_secrets" {
  name             = "sealed-secrets"
  repository       = "https://bitnami-labs.github.io/sealed-secrets"
  chart            = "sealed-secrets"
  namespace        = "kube-system"
  create_namespace = true
  version          = "2.17.7"

  set {
    name  = "image.pullPolicy"
    value = "IfNotPresent"
  }
}

output "name" {
  value = helm_release.sealed_secrets.name
}

