resource "google_compute_instance" "redis" {
  name                = "redis-vm"
  machine_type        = "f1-micro"
  project             = var.project
  zone                = var.zone
  deletion_protection = false

  tags = ["redis-deployment"]

  labels = {
    "goog-dm" = "redis"
  }


  boot_disk {
    auto_delete = true
    device_name = "bitnami-vm-for-redis-vm-tmpl-boot-disk"
    mode        = "READ_WRITE"

    initialize_params {
      image  = "https://www.googleapis.com/compute/v1/projects/bitnami-launchpad/global/images/bitnami-redis-7-0-5-3-r04-linux-debian-11-x86-64-nami"
      labels = {}
      size   = 10
      type   = "pd-standard"
    }
  }

  network_interface {
    network = google_compute_network.default.self_link
    access_config {}
  }

  service_account {
    email = "1082596721499-compute@developer.gserviceaccount.com"
    scopes = [
      "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
      "https://www.googleapis.com/auth/cloudruntimeconfig",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring.write",
    ]
  }

  lifecycle {
    ignore_changes = [metadata]
  }
}
