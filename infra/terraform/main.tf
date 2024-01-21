provider "aws" {
  region = var.aws_region
}

#Configuração do Terraform State
terraform {
  backend "s3" {
    bucket = "terraform-state-soat"
    key    = "infra-script-migration-clientes/terraform.tfstate"
    region = "us-east-1"

    dynamodb_table = "terraform-state-soat-locking"
    encrypt        = true
  }
}

resource "aws_s3_object" "clientes_migration_sql" {
  bucket = var.s3_migration_bucket
  source = "../../src/migrations/create_clientes_table.sql" #path to file from source
  key    = "scripts/clientes/create_clientes_table.sql"  #S3 Key
}
