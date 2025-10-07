# GO-PRO Learning Platform - Terraform Outputs

# VPC Outputs
output "vpc_id" {
  description = "ID of the VPC"
  value       = module.vpc.vpc_id
}

output "vpc_cidr" {
  description = "CIDR block of the VPC"
  value       = module.vpc.vpc_cidr
}

output "private_subnet_ids" {
  description = "IDs of private subnets"
  value       = module.vpc.private_subnet_ids
}

output "public_subnet_ids" {
  description = "IDs of public subnets"
  value       = module.vpc.public_subnet_ids
}

# RDS Outputs
output "rds_endpoint" {
  description = "RDS instance endpoint"
  value       = module.rds.endpoint
  sensitive   = true
}

output "rds_port" {
  description = "RDS instance port"
  value       = module.rds.port
}

output "rds_database_name" {
  description = "RDS database name"
  value       = module.rds.database_name
}

output "rds_arn" {
  description = "ARN of the RDS instance"
  value       = module.rds.arn
}

# Redis Outputs
output "redis_endpoint" {
  description = "Redis cluster endpoint"
  value       = module.redis.endpoint
  sensitive   = true
}

output "redis_port" {
  description = "Redis cluster port"
  value       = module.redis.port
}

output "redis_configuration_endpoint" {
  description = "Redis configuration endpoint"
  value       = module.redis.configuration_endpoint
  sensitive   = true
}

# EKS Outputs
output "eks_cluster_id" {
  description = "EKS cluster ID"
  value       = module.eks.cluster_id
}

output "eks_cluster_endpoint" {
  description = "EKS cluster endpoint"
  value       = module.eks.cluster_endpoint
}

output "eks_cluster_name" {
  description = "EKS cluster name"
  value       = module.eks.cluster_name
}

output "eks_cluster_certificate_authority_data" {
  description = "EKS cluster certificate authority data"
  value       = module.eks.cluster_certificate_authority_data
  sensitive   = true
}

output "eks_cluster_security_group_id" {
  description = "Security group ID attached to the EKS cluster"
  value       = module.eks.cluster_security_group_id
}

output "eks_node_group_ids" {
  description = "EKS node group IDs"
  value       = module.eks.node_group_ids
}

# MSK Outputs
output "msk_bootstrap_brokers" {
  description = "MSK bootstrap brokers"
  value       = module.msk.bootstrap_brokers
  sensitive   = true
}

output "msk_bootstrap_brokers_tls" {
  description = "MSK bootstrap brokers (TLS)"
  value       = module.msk.bootstrap_brokers_tls
  sensitive   = true
}

output "msk_zookeeper_connect_string" {
  description = "MSK Zookeeper connection string"
  value       = module.msk.zookeeper_connect_string
  sensitive   = true
}

output "msk_cluster_arn" {
  description = "ARN of the MSK cluster"
  value       = module.msk.cluster_arn
}

# S3 Outputs
output "s3_bucket_names" {
  description = "Names of S3 buckets"
  value       = module.s3.bucket_names
}

output "s3_bucket_arns" {
  description = "ARNs of S3 buckets"
  value       = module.s3.bucket_arns
}

# ALB Outputs
output "alb_dns_name" {
  description = "DNS name of the Application Load Balancer"
  value       = module.alb.dns_name
}

output "alb_arn" {
  description = "ARN of the Application Load Balancer"
  value       = module.alb.arn
}

output "alb_zone_id" {
  description = "Zone ID of the Application Load Balancer"
  value       = module.alb.zone_id
}

output "alb_target_group_arns" {
  description = "ARNs of ALB target groups"
  value       = module.alb.target_group_arns
}

# Route53 Outputs
output "route53_zone_id" {
  description = "Route53 hosted zone ID"
  value       = var.domain_name != "" ? module.route53[0].zone_id : null
}

output "route53_name_servers" {
  description = "Route53 name servers"
  value       = var.domain_name != "" ? module.route53[0].name_servers : null
}

# ACM Outputs
output "acm_certificate_arn" {
  description = "ARN of the ACM certificate"
  value       = var.domain_name != "" ? module.acm[0].certificate_arn : null
}

output "acm_certificate_status" {
  description = "Status of the ACM certificate"
  value       = var.domain_name != "" ? module.acm[0].certificate_status : null
}

# Secrets Manager Outputs
output "secrets_arn" {
  description = "ARN of the secrets in Secrets Manager"
  value       = module.secrets.secrets_arn
  sensitive   = true
}

# IAM Outputs
output "iam_role_arns" {
  description = "ARNs of IAM roles"
  value       = module.iam.role_arns
}

output "iam_policy_arns" {
  description = "ARNs of IAM policies"
  value       = module.iam.policy_arns
}

# CloudWatch Outputs
output "cloudwatch_log_group_names" {
  description = "Names of CloudWatch log groups"
  value       = module.cloudwatch.log_group_names
}

output "cloudwatch_dashboard_url" {
  description = "URL of the CloudWatch dashboard"
  value       = module.cloudwatch.dashboard_url
}

# Security Group Outputs
output "security_group_ids" {
  description = "Map of security group IDs"
  value = {
    alb   = module.security_groups.alb_security_group_id
    eks   = module.security_groups.eks_security_group_id
    rds   = module.security_groups.rds_security_group_id
    redis = module.security_groups.redis_security_group_id
    msk   = module.security_groups.msk_security_group_id
  }
}

# Connection Information
output "connection_info" {
  description = "Connection information for services"
  value = {
    database = {
      endpoint = module.rds.endpoint
      port     = module.rds.port
      name     = module.rds.database_name
    }
    redis = {
      endpoint = module.redis.endpoint
      port     = module.redis.port
    }
    kafka = {
      brokers = module.msk.bootstrap_brokers_tls
    }
    kubernetes = {
      cluster_name = module.eks.cluster_name
      endpoint     = module.eks.cluster_endpoint
    }
    load_balancer = {
      dns_name = module.alb.dns_name
    }
  }
  sensitive = true
}

# Kubectl Configuration Command
output "kubectl_config_command" {
  description = "Command to configure kubectl"
  value       = "aws eks update-kubeconfig --region ${var.aws_region} --name ${module.eks.cluster_name}"
}

# Application URL
output "application_url" {
  description = "Application URL"
  value       = var.domain_name != "" ? "https://${var.domain_name}" : "http://${module.alb.dns_name}"
}

# Environment Information
output "environment_info" {
  description = "Environment information"
  value = {
    environment = var.environment
    region      = var.aws_region
    account_id  = data.aws_caller_identity.current.account_id
  }
}

# Cost Estimation
output "estimated_monthly_cost" {
  description = "Estimated monthly cost (approximate) - varies by instance types and configuration"
  value = {
    note = "Costs vary significantly based on instance types, storage, and data transfer"
  }
}
