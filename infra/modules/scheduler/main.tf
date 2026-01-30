resource "google_cloud_scheduler_job" "metrics_pull" {
  name        = "metrics-pull-job"
  schedule    = "0 0 1,16 * *" 
  time_zone   = "UTC"

  pubsub_target {
    topic_name = var.metrics_pull_topic
    data       = base64encode("{}")
  }
}
