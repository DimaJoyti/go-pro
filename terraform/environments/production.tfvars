# GO-PRO Learning Platform - Production Environment Configuration

# General Configuration
environment = "production"
owner       = "DevOps Team"
cost_center = "Engineering"

# AWS Configuration
aws_region = "us-east-1"
vpc_cidr   = "10.1.0.0/16"

# GCP Configuration
gcp_project_id = "gopro-production"
gcp_region     = "us-central1"

# RDS Configuration
rds_instance_class        = "db.r6g.xlarge"
rds_allocated_storage     = 500
rds_max_allocated_storage = 2000
rds_backup_retention_days = 30
db_engine_version         = "15.4"
db_parameter_group_family = "postgres15"

# ElastiCache Configuration
redis_node_type       = "cache.r6g.large"
redis_num_cache_nodes = 3

# EKS Configuration
eks_cluster_version = "1.28"
eks_node_groups = {
  general = {
    desired_size   = 3
    min_size       = 3
    max_size       = 10
    instance_types = ["t3.large"]
    capacity_type  = "ON_DEMAND"
    disk_size      = 100
  }
  compute = {
    desired_size   = 2
    min_size       = 2
    max_size       = 8
    instance_types = ["c6i.xlarge"]
    capacity_type  = "ON_DEMAND"
    disk_size      = 100
  }
}

# MSK Configuration
msk_kafka_version     = "3.5.1"
msk_instance_type     = "kafka.m5.large"
msk_ebs_volume_size   = 500
msk_broker_node_count = 3

# Domain Configuration
domain_name = "gopro.example.com" # Replace with actual domain

# Monitoring Configuration
enable_enhanced_monitoring    = true
cloudwatch_log_retention_days = 90

# Security Configuration
enable_waf          = true
enable_guardduty    = true
enable_security_hub = true

# Backup Configuration
enable_automated_backups = true
backup_retention_days    = 30

# Feature Flags
enable_multi_az            = true
enable_encryption          = true
enable_deletion_protection = true
enable_spot_instances      = false
enable_autoscaling         = true
enable_nat_gateway         = true
single_nat_gateway         = false # High availability

# Scaling Configuration
min_capacity              = 3
max_capacity              = 20
target_cpu_utilization    = 60
target_memory_utilization = 70

# Application Configuration
app_port          = 8080
health_check_path = "/health"

# Additional Tags
additional_tags = {
  Team           = "Backend"
  Purpose        = "Production"
  Compliance     = "SOC2"
  DataClass      = "Confidential"
  BackupRequired = "true"
}
