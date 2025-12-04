output "cost_data_bucket" {
  value = google_storage_bucket.cost_data_bucket.name
}

output "processed_data_bucket" {
  value = google_storage_bucket.processed_data_bucket.name
}