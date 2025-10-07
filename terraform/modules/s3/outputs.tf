output "bucket_names" { value = [aws_s3_bucket.assets.id, aws_s3_bucket.backups.id] }
output "bucket_arns" { value = [aws_s3_bucket.assets.arn, aws_s3_bucket.backups.arn] }

