variable "domain_name" {
  type = string
}

variable "zone_id" {
  type = string
}

variable "environment" {
  type = string
}

variable "tags" {
  type    = map(string)
  default = {}
}
