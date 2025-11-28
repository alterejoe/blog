variable "zone_id" {
  description = "The Cloudflare Zone ID where the DNS record will be created."
  type        = string
}

variable "name" {
  description = "The name of the DNS record (e.g. 'app' or 'example.com')."
  type        = string
}

variable "type" {
  description = "The DNS record type (A, CNAME, TXT, etc)."
  type        = string
  default     = "A"
}

variable "content" {
  description = "The record content (e.g. IP address or target hostname)."
  type        = string
}

variable "ttl" {
  description = "TTL for the DNS record (1 = auto)."
  type        = number
  default     = 3600
}


variable "comment" {
  description = "Optional comment describing the record."
  type        = string
  default     = ""
}

variable "tags" {
  description = "Optional list of tags for the record."
  type        = list(string)
  default     = []
}
