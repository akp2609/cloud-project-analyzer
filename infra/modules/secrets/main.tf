resource "google_secret_manager_secret" "database_url" {
  project   = var.project_id
  secret_id = "database-url"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "terraform_tfvars" {
  project   = var.project_id
  secret_id = "terraform-tfvars"

  replication {
    auto {}
  }
}
