resource "aws_s3_bucket" "assets" {
  bucket = "${var.name_prefix}-assets"
  tags   = var.tags
}

resource "aws_s3_bucket" "backups" {
  bucket = "${var.name_prefix}-backups"
  tags   = var.tags
}

