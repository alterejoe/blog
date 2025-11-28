
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5"
    }
  }
}

# Lookup account by name
data "cloudflare_accounts" "account" {
  direction = "desc"
  name      = var.account_name
  max_items = 1
}

# Lookup zone by name
data "cloudflare_zones" "zone" {
  account = {
    id = data.cloudflare_accounts.account.result[0].id
  }
  direction = "desc"
  name      = var.root_domain
  status    = "active"
}

locals {
  account_id = data.cloudflare_accounts.account.result[0].id
  zone_id    = data.cloudflare_zones.zone.result[0].id
}



