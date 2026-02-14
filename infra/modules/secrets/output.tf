output "terraform_tfvars_secret_id" {
  value = google_secret_manager_secret.terraform_tfvars.secret_id
}

output "database_url_secret_id" {
  value = google_secret_manager_secret.database_url.secret_id
}
