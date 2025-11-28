resource "helm_release" "nginx_ingress" {
  name       = "nginx-ingress"
  namespace  = var.namespace
  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"
  version    = var.chart_version

  values = [yamlencode({
    controller = {
      ingressClassResource = {
        enabled = true
        default = true
      }

      admissionWebhooks = {
        enabled = true
        patch = {
          enabled = true
        }
        certManager = {
          enabled = true
        }
      }

      config = {
        "strict-validate-path-type" = "false"
        "proxy-body-size"           = "20m"
      }
    }
  })]
}
