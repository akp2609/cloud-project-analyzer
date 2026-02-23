resource "google_pubsub_topic" "csv_uploads" {
  name = "csv-uploads"
  project = var.project_id
}

resource "google_pubsub_topic" "project_logs" {
  name = "project-logs"
  project = var.project_id
}

resource "google_pubsub_topic" "metrics_pull" {
  name = "metrics-pull"
  project = var.project_id
}

