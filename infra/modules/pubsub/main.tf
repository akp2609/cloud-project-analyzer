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

# resource "google_pubsub_subscription" "logs_ingestor_sub" {
#   name  = "logs-ingestor-sub"
#   topic = google_pubsub_topic.project_logs.name

#   push_config {
#     push_endpoint = "https://YOUR_CLOUD_RUN_URL/"
#     # optional: set an OIDC token if your ingestor requires auth
#     # oidc_token {
#     #   service_account_email = google_service_account.ingestor.email
#     # }
#   }

#   ack_deadline_seconds = 10
# }
