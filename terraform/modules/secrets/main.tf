resource "aws_secretsmanager_secret" "db" {
  name = "${var.name_prefix}-database-credentials"
  tags = var.tags
}

resource "aws_secretsmanager_secret_version" "db" {
  secret_id = aws_secretsmanager_secret.db.id
  secret_string = jsonencode({
    username = "gopro_admin"
    password = var.db_password
    endpoint = var.db_endpoint
  })
}

