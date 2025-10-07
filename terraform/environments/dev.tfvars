# GO-PRO Learning Platform - Development Environment Configuration

# General Configuration
environment = "dev"
owner       = "DevOps Team"
cost_center = "Engineering"

# AWS Configuration
aws_region = "us-east-1"
vpc_cidr   = "10.0.0.0/16"

# GCP Configuration
gcp_project_id = "gopro-dev"
gcp_region     = "us-central1"

# RDS Configuration
rds_instance_class        = "db.t3.micro"
rds_allocated_storage     = 20
rds_max_allocated_storage = 100
rds_backup_retention_days = 3
db_engine_version         = "15.4"
db_parameter_group_family = "postgres15"

# ElastiCache Configuration
redis_node_type       = "cache.t4g.micro"
redis_num_cache_nodes = 1

# EKS Configuration
eks_cluster_version = "1.28"
eks_node_groups = {
  general = {
    desired_size   = 1
    min_size       = 1
    max_size       = 3
    instance_types = ["t3.small"]
    capacity_type  = "ON_DEMAND"
    disk_size      = 30
  }
}

# MSK Configuration
msk_kafka_version     = "3.5.1"
msk_instance_type     = "kafka.t3.small"
msk_ebs_volume_size   = 50
msk_broker_node_count = 2

# Domain Configuration
domain_name = "" # No custom domain for dev

# Monitoring Configuration
enable_enhanced_monitoring    = false
cloudwatch_log_retention_days = 7

# Security Configuration
enable_waf          = false
enable_guardduty    = false
enable_security_hub = false

# Backup Configuration
enable_automated_backups = true
backup_retention_days    = 7

# Feature Flags
enable_multi_az            = false
enable_encryption          = true
enable_deletion_protection = false
enable_spot_instances      = false
enable_autoscaling         = true
enable_nat_gateway         = true
single_nat_gateway         = true # Cost optimization for dev

# Scaling Configuration
min_capacity              = 1
max_capacity              = 3
target_cpu_utilization    = 70
target_memory_utilization = 80

# Application Configuration
app_port          = 8080
health_check_path = "/health"

# Additional Tags
additional_tags = {
  Team         = "Backend"
  Purpose      = "Development"
  AutoShutdown = "true"
}
