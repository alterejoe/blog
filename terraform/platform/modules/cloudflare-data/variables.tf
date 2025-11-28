variable "account_name" {
  type = string
}

# variable "subdomain" {
#   description = "Subdomain to protect (e.g. admin.example.com or *.example.com)"
#   type        = string
# }

# variable "account_id" {
#   description = "Cloudflare account ID"
#   type        = string
# }

# variable "email_domain" {
#   description = "Email domain allowed (e.g. domain.com)"
#   type        = string
# }
#
variable "root_domain" {
  description = "Base domain (e.g. example.com)"
  type        = string
}


# variable "session_duration" {
#   description = "Access session length"
#   type        = string
#   default     = "12h"
# }
