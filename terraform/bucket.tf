resource "google_storage_bucket" "bucket" {
  name                        = "loyaltypointagent"
  location                    = var.region
  uniform_bucket_level_access = true
}
