resource "google_logging_project_sink" "project_logs_sink" {
    name = "export-project-logs"
    project = var.source_project_id
    destination = "pubsub.googleapis.com/projects/${var.dest_project_id}/topics/${var.topic_name}"

    filter = var.log_filter

    unique_writer_identity = true
}