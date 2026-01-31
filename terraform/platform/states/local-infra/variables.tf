variable "aws_access_key" {
  default = "test"
}
variable "aws_secret_key" {
  default = "test"
}

variable "postgres_host" {
  default = "localhost"
}

variable "postgres_port" {
  default = 5000
}

variable "postgres_user" {
  default = "postgres"
}

variable "postgres_password" {
  default = "postgres"
}

variable "postgres_database" {
  default = "db"
}

variable "postgres_sslmode" {
  default = "disable"
}

variable "postgres_schema" {
  default = "blog"
}

variable "migrator_password" {
  default = "migrator"
}

variable "client_runtime_password" {
  default = "client_runtime"
}

variable "admin_runtime_password" {
  default = "admin_runtime"
}
