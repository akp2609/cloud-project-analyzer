resource "google_storage_bucket" "cost_data_bucket" {
  name = "${var.prefix}-cost-data"
  location = var.region
  force_destroy = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket" "processed_data_bucket" {
  name = "${var.prefix}-processed-data"
  location = var.region
  force_destroy = true
  uniform_bucket_level_access = true
}