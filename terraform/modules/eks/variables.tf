variable "cluster_name" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "cluster_version" {
  type    = string
  default = "1.28"
}

variable "node_groups" {
  type    = any
  default = {}
}

variable "enable_irsa" {
  type    = bool
  default = true
}

variable "environment" {
  type = string
}

variable "tags" {
  type    = map(string)
  default = {}
}
