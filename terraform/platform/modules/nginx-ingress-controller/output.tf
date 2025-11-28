data "kubernetes_service" "nginx_ingress" {
  metadata {
    name      = "nginx-ingress-ingress-nginx-controller"
    namespace = var.namespace
  }

  depends_on = [helm_release.nginx_ingress]
}

output "load_balancer_ip" {
  description = "External IP or hostname for the nginx ingress controller"
  value = coalesce(
    data.kubernetes_service.nginx_ingress.status[0].load_balancer[0].ingress[0].ip,
    data.kubernetes_service.nginx_ingress.status[0].load_balancer[0].ingress[0].hostname
  )
}
