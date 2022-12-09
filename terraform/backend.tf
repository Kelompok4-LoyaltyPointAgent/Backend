terraform {
  backend "gcs" {
    bucket = "loyaltypointagent"
    prefix = "terraform/state"
  }
}
