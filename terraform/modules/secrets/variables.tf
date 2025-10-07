variable "name_prefix" {
  type = string
}

variable "db_password" {
  type      = string
  sensitive = true
}

variable "redis_token" {
  type      = string
  sensitive = true
}

variable "db_endpoint" {
  type = string
}

variable "redis_endpoint" {
  type = string
}

variable "environment" {
  type = string
}

variable "tags" {
  type    = map(string)
  default = {}
}
