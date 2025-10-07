resource "aws_backup_vault" "main" {
  name = "${var.name_prefix}-backup-vault"
  tags = var.tags
}

