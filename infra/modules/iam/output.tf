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

output "extra_service_account_emails" {
  value = {
    for k, v in google_service_account.extra_service_accounts :
    k => v.email
  }
}


output "extra_service_account_ids" {
  value = {
    for k, v in google_service_account.extra_service_accounts :
    k => v.id
  }
}