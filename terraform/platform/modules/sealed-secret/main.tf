## part 1
terraform {
  required_providers {
    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.7.0"
    }
  }
}


locals {
  unsealed_secret_yaml = yamlencode({
    apiVersion = "v1"
    kind       = "Secret"
    metadata = {
      name      = var.secret_name
      namespace = var.namespace
    }
    type       = "Opaque"
    stringData = var.string_data
  })

  sealed_file_path = var.outfile != "" ? var.outfile : "${path.module}/tmp-${var.secret_name}-sealed.yaml"
}
resource "null_resource" "seal_secret" {
  triggers = {
    yaml_body = local.unsealed_secret_yaml
    outfile   = var.outfile
    cmd_hash  = sha1("v3")
  }

  provisioner "local-exec" {
    command = <<EOT
bash -c '
  mkdir -p "$(dirname ${local.sealed_file_path})"
  echo "${local.unsealed_secret_yaml}" | \
  kubeseal \
    --kubeconfig ${var.kubeconfig_path} \
    --controller-namespace ${var.controller_namespace} \
    --controller-name ${var.controller_name} \
    --format yaml > "${local.sealed_file_path}"
'
EOT
  }
}

# read sealed file only after sealing finishes
data "local_file" "sealed" {
  depends_on = [null_resource.seal_secret]
  filename   = local.sealed_file_path
}

# apply sealed manifest only after file exists
resource "kubectl_manifest" "sealed_secret" {
  depends_on = [data.local_file.sealed]
  yaml_body  = data.local_file.sealed.content
}

output "sealed_secret_path" {
  value = local.sealed_file_path
}
