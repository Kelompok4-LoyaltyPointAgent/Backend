resource "google_compute_network" "default" {
  project                  = var.project
  name                     = "default"
  description              = "Default network for the project"
  auto_create_subnetworks  = true
  enable_ula_internal_ipv6 = false
  routing_mode             = "REGIONAL"

  # mtu = 1460
}

# resource "google_project_service" "vpcaccess_api" {
#   project = var.project
#   service = "vpcaccess.googleapis.com"
# }

# module "serverless_network" {
#   source  = "terraform-google-modules/network/google"
#   version = "~> 6.0"

#   project_id   = var.project
#   network_name = "serverless-network"
#   mtu          = 1460

#   subnets = [
#     {
#       subnet_name   = "serverless-subnet"
#       subnet_ip     = "10.10.10.0/28"
#       subnet_region = "us-central1"
#     }
#   ]
# }

# module "serverless_connector" {
#   source     = "terraform-google-modules/network/google//modules/vpc-serverless-connector-beta"
#   project_id = var.project
#   vpc_connectors = [{
#     name            = "central-serverless"
#     region          = var.region
#     subnet_name     = module.serverless_network.subnets["${var.region}/serverless-subnet"].name
#     host_project_id = var.project
#     machine_type    = "f1-micro"
#     min_instances   = 2
#     max_instances   = 3
#     }
#   ]
#   depends_on = [
#     google_project_service.vpcaccess_api
#   ]
# }
