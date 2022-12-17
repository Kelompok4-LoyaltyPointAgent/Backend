module "cloudrun_production" {
  source  = "GoogleCloudPlatform/cloud-run/google"
  version = "~> 0.2.0"

  service_name = "loyaltypointagent"
  project_id   = var.project
  location     = var.region
  image        = "gcr.io/${var.project}/loyaltypointagent"

  ports = {
    "name" : "http1",
    "port" : var.env_vars_production[index(var.env_vars_production.*.name, "HTTP_PORT")].value
  }

  env_vars = var.env_vars_production

  members = ["allUsers"]

  template_annotations = {
    "autoscaling.knative.dev/minScale" : 0,
    "autoscaling.knative.dev/maxScale" : 5,
    "run.googleapis.com/cloudsql-instances" : google_sql_database_instance.mysql.connection_name
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
