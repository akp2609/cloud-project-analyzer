resource "google_artifact_registry_repository" "repo" {
  project       = "${var.project_id}"
  location = "${var.region}"
  repository_id = "${var.repository_id}"
  description   = "Docker repo for microservices"
  format        = "DOCKER"


   cleanup_policies {
    id     = "delete-old-artifacts"
    action = "DELETE"
    condition {
      tag_state = "ANY"
    }
  }

  cleanup_policies {
    id     = "keep-latest-5"
    action = "KEEP"
    most_recent_versions {
      keep_count = 5
    }
  }
}