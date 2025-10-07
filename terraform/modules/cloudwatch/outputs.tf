output "log_group_names" { value = [aws_cloudwatch_log_group.app.name] }
output "dashboard_url" { value = "https://console.aws.amazon.com/cloudwatch" }

