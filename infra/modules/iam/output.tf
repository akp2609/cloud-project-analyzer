output "sre_sa_email" {
  value = google_service_account.sre_cost_sa.email
}

output "sre_sa_id" {
  value = google_service_account.sre_cost_sa.id
}

output "analysis_engine_email" {
  value = google_service_account.analysis_engine_sa.email
}

output "analysis_engine_id" {
  value = google_service_account.analysis_engine_sa.id
}