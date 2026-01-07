resource "google_cloudbuild_trigger" "trigger" {
  project = var.project_id
  name    = "${var.service_name}-trigger"
  filename = "services/${var.service_name}/cloudbuild.yaml"
  service_account = var.service_account

  github {
    owner = var.repo_owner
    name  = var.repo_name

    push {
      branch = "^main$"
    }
  }

  include_build_logs = "INCLUDE_BUILD_LOGS_WITH_STATUS"
}
