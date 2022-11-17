module "cloud_run" {
  source  = "GoogleCloudPlatform/cloud-run/google"
  version = "~> 0.2.0"

  service_name = "loyaltypointagent-staging"
  project_id   = var.project
  location     = var.region
  image        = "gcr.io/${var.project}/loyaltypointagent"

  ports = {
    "name" : "http1",
    "port" : var.env_vars[index(var.env_vars.*.name, "HTTP_PORT")].value
  }

  env_vars = var.env_vars

  members = ["allUsers"]

  template_annotations = {
    "autoscaling.knative.dev/minScale" : 0,
    "autoscaling.knative.dev/maxScale" : 1,
  }

  requests = {
    "cpu" : "100m",
    "memory" : "128Mi",
  }

  limits = {
    "cpu" : "1000m",
    "memory" : "1024Mi",
  }

  container_concurrency = 80
}
