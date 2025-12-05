resource "google_cloudbuild_trigger" "trigger" {
  project = var.project_id
  name = "${var.service_name}-trigger"

  github {
    owner = var.repo_owner
    name = var.repo_name

    push{
        branch = "${var.branch}"
    }
  }
   filename = "services/${var.service_name}/cloudbuild.yaml"

}