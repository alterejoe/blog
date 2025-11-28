## part 1
terraform {
  required_providers {
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.17"
    }
    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.7.0"
    }
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5"
    }
  }
}
locals {
  infisical_values = {
    postgresql = { enabled = false }
    redis      = { enabled = false }

    infisical = {
      image = {
        tag = var.app_image_tag != "" ? var.app_image_tag : null
      }

      replicaCount  = var.replica_count > 0 ? var.replica_count : null
      kubeSecretRef = "infisical-secrets"
      autoBootstrap = {
        enabled = var.auto_bootstrap_enabled
        organization = var.organization_name
        secretDestination = {
          name      = var.bootstrap_secret_name
          namespace = var.namespace
        }
        image = {
            tag = var.cli_image_tag
        }
        credentialSecret = {
          name = var.bootstrap_secret_credentials
        }
      }

      # # environment variables
      # extraEnv = [
      #   { name = "JWT_AUTH_LIFETIME", value = "15m" },
      #   { name = "JWT_REFRESH_LIFETIME", value = "24h" },
      #   { name = "JWT_SERVICE_LIFETIME", value = "1h" },
      #   # { name = "NODE_EXTRA_CA_CERTS", value = "/certs/do-ca.crt" }
      # ]

      # in-cluster service definition
      service = {
        type     = var.service_type
        nodePort = var.service_type == "NodePort" ? var.service_node_port : null
      }
    }

    # ingress is top-level in this chart
    ingress = {
      enabled          = var.ingress_enabled
      hostName         = var.ingress_host != "" ? var.ingress_host : null
      ingressClassName = "nginx" # explicitly use your existing ingress class
      # createIngressClassResource = false   # <- if the chart supports it
      nginx = {
        enabled = false
      }
    }
  }

  base_values = yamldecode(file("${path.root}/values/infisical.yaml"))
  merged_values = merge(
    local.base_values,
    {
      infisical  = merge(local.base_values.infisical, local.infisical_values.infisical)
      ingress    = merge(local.base_values.ingress, local.infisical_values.ingress)
      postgresql = local.infisical_values.postgresql
      redis      = local.infisical_values.redis
    }
  )
}

resource "helm_release" "infisical" {
  name             = var.helm_name
  repository       = var.helm_repo
  chart            = var.helm_chart
  namespace        = var.namespace
  create_namespace = false
  version          = var.helm_chart_version
  values           = [yamlencode(local.merged_values)]
  timeout          = 900 # 15 minutes
}

locals {
  service_name = "${var.helm_name}-${var.helm_chart}-infisical"
}

output "service_name" {
  value = local.service_name
}
