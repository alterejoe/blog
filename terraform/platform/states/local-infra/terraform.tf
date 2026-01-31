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

  # https://tfstate-space.nyc3.digitaloceanspaces.com/terraform/state.tfstate
  backend "local" {
    path = "terraform.tfstate"
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

provider "aws" {
  access_key                  = var.aws_access_key
  secret_key                  = var.aws_secret_key
  region                      = "us-east-1"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  s3_use_path_style           = true


  endpoints {
    s3  = "http://localhost:4566"
    iam = "http://localhost:4566"
    sts = "http://localhost:4566"
  }
}

