output "cost_data_bucket" {
  value = module.storage.cost_data_bucket
}

output "processed_data_bucket" {
  value = module.storage.processed_data_bucket
}

output "cca_sa_email" {
  value = module.iam.sre_sa_email
}