output "secrets_arn" {
  value     = aws_secretsmanager_secret.db.arn
  sensitive = true
}
