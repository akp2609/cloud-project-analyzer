resource "google_cloudbuild_trigger" "trigger" {
  project = var.project_id
  name = "${var.service_name}-trigger"

  github {
    owner = var.repo_owner
    name = var.repo_name

    push{
        branch = "^${var.branch}$"
    }
  }
  build {
    step {
      name = "gcr.io/cloud-builders/docker" # Corrected typo from 'dockerl' to 'docker'
      args = ["build", "-t", var.image, "./services/upload-service"]
    }

    step {
      name = "gcr.io/cloud-builders/docker"
      args = ["push", var.image]
    }

    images = [var.image]
  }

}