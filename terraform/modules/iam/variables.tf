variable "name_prefix" {
  type = string
}

variable "cluster_name" {
  type = string
}

variable "environment" {
  type = string
}

variable "tags" {
  type    = map(string)
  default = {}
}
