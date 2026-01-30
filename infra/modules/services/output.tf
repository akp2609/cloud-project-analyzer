output "cloudrun_service" {
  value = google_project_service.cloudrun
}

output "cloudbuild_service" {
  value = google_project_service.cloudbuild
}

output "sqladmin_service" {
  value = google_project_service.sqladmin
}

output "servicenetworking_service" {
  value = google_project_service.servicenetworking
}

output "compute_service"{
    value = google_project_service.compute
}
