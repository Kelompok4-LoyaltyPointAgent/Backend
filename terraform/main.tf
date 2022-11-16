terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.43.0"
    }
  }

  backend "gcs" {
    bucket = "loyaltypointagent"
    prefix = "terraform/state"
  }
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}

resource "google_storage_bucket" "bucket" {
  name                        = "loyaltypointagent"
  location                    = var.region
  uniform_bucket_level_access = true
}

module "mysql" {
  source = "./modules/mysql"

  project             = var.project
  region              = var.region
  mysql_root_password = var.mysql_root_password
}
