variable "name_prefix" {
  type = string
}

variable "rds_arn" {
  type = string
}

variable "environment" {
  type = string
}

variable "tags" {
  type    = map(string)
  default = {}
}
