# main.tf
# core infrastructure for net-scope
# provisions the VPC, subnets, routing, and other networking

terraform {
  required_versions = ">= 1.6.0"
  
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  # using local state for now, but I will
  # uncomment the following once VPC provisions correctly
  # backend "s3" {
  #  bucket = "tfstate-bucket"
  #  key    = "net-scope/terraform.tfstate"
  #  region = "us-east-1
  # }

  module "vpc" {
    source              = "terraform-aws-modules/vpc/aws"
    name                = "net-scope-vpc"
    cidr                = "10.0.0.0/16"
    enable_nat_gateway  = true
  }
}
