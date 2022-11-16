resource "google_sql_database_instance" "default" {
  project             = var.project
  name                = "mysql-f1"
  database_version    = "MYSQL_8_0"
  region              = var.region
  deletion_protection = false
  root_password       = var.mysql_root_password

  settings {
    tier              = "db-f1-micro"
    availability_type = "ZONAL"
    disk_size         = "10"

    ip_configuration {
      ipv4_enabled       = true
      allocated_ip_range = "default-ip-range"
      private_network    = "projects/loyaltypointagent/global/networks/default"

      authorized_networks {
        name  = "any"
        value = "0.0.0.0/0"
      }
    }

    backup_configuration {
      binary_log_enabled             = false
      enabled                        = false
      location                       = null
      start_time                     = null
      transaction_log_retention_days = null

      # backup_retention_settings {
      #   retained_backups = null
      #   retention_unit   = null
      # }
    }

    maintenance_window {
      day          = 1
      hour         = 23
      update_track = "canary"
    }
  }
}
