resource "google_compute_network" "default" {
  project                  = var.project
  name                     = "default"
  description              = "Default network for the project"
  auto_create_subnetworks  = true
  enable_ula_internal_ipv6 = false
  routing_mode             = "REGIONAL"

  # mtu = 1460
}
