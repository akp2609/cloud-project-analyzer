output "project_logs" {
  value = google_pubsub_topic.project_logs.name
}

output "metrics_pull_topic" {
  value = google_pubsub_topic.metrics_pull.name
}
