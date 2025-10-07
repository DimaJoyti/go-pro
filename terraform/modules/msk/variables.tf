variable "name_prefix" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "security_group_ids" {
  type = list(string)
}

variable "kafka_version" {
  type    = string
  default = "3.5.1"
}

variable "broker_node_count" {
  type    = number
  default = 2
}

variable "instance_type" {
  type    = string
  default = "kafka.t3.small"
}

variable "ebs_volume_size" {
  type    = number
  default = 100
}

variable "environment" {
  type = string
}

variable "tags" {
  type    = map(string)
  default = {}
}
