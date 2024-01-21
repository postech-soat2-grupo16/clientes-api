variable "aws_region" {
  type    = string
  default = "us-east-1"
}

variable "s3_migration_bucket" {
  type = string
  default = "migrations-script"
}