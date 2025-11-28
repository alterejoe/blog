terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
    postgresql = {
      source  = "cyrilgdn/postgresql"
      version = "1.26.0"
    }
  }
}



provider "postgresql" {
  # Configuration options
  host     = var.postgres_host
  port     = var.postgres_port
  username = var.postgres_user
  password = var.postgres_password
  # database = var.postgres_database
  sslmode = var.postgres_sslmode
}


