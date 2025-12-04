resource "google_service_account" "sre_cost_sa" {
  account_id = "${var.prefix}-sre-sa"
  display_name = "SRE Cost Service Account"
}

resource "google_storage_bucket_iam_member" "raw_bucket_binding" {
  bucket = var.raw_bucket_name
  role = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_storage_bucket_iam_member" "processed_bucket_binding" {
  bucket = var.processed_bucket_name
  role = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}


resource "google_project_iam_member" "raw_runtime_user" {
  project = var.project_id
  role = "roles/run.invoker"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_project_iam_member" "artifact_registry_writer" {
  project = var.project_id
  role = "roles/artifactregistry.writer"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

