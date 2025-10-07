resource "aws_cloudwatch_log_group" "app" {
  name              = "/aws/eks/${var.cluster_name}/application"
  retention_in_days = 30
  tags              = var.tags
}

