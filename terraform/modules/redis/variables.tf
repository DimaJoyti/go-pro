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

variable "node_type" {
  type    = string
  default = "cache.t4g.micro"
}

variable "num_cache_nodes" {
  type    = number
  default = 1
}

variable "engine_version" {
  type    = string
  default = "7.0"
}

variable "parameter_group_family" {
  type    = string
  default = "redis7"
}

variable "port" {
  type    = number
  default = 6379
}

variable "at_rest_encryption_enabled" {
  type    = bool
  default = true
}

variable "transit_encryption_enabled" {
  type    = bool
  default = true
}

variable "auth_token" {
  type      = string
  sensitive = true
  default   = null
}

variable "snapshot_retention_limit" {
  type    = number
  default = 5
}

variable "snapshot_window" {
  type    = string
  default = "03:00-05:00"
}

variable "maintenance_window" {
  type    = string
  default = "sun:05:00-sun:07:00"
}

variable "notification_topic_arn" {
  type    = string
  default = null
}

variable "auto_minor_version_upgrade" {
  type    = bool
  default = true
}

variable "apply_immediately" {
  type    = bool
  default = false
}

variable "subnet_group_name" {
  type    = string
  default = ""
}

variable "environment" {
  type = string
}

variable "tags" {
  type    = map(string)
  default = {}
}

