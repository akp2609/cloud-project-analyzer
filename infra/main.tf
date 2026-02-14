
module "project_services" {
  source     = "./modules/services"
  project_id = var.project_id
}

module "project_secrets" {
  source = "./modules/secrets"
  project_id = var.project_id
}

module "storage" {
  source = "./modules/storage"

  prefix = var.project_id
  region = var.region
}

module "iam" {
  source = "./modules/iam"

  prefix = "cca"
  project_id = var.project_id
  region = var.region
  state_bucket_name = var.state_bucket
  raw_bucket_name = module.storage.cost_data_bucket
  processed_bucket_name = module.storage.processed_data_bucket
  secret_id = module.project_secrets.terraform_tfvars_secret_id
  project_logs_topic = module.pubsub.project_logs
  dest_project_id   = var.project_id
  topic_name        = "project-logs"
  log_filter = "severity>=ERROR"
  project_log_sink_writer_identity = module.hirebizz_log_sink.project_log_sink_writer_identity
}

module "repo" {
  source = "./modules/artifact-registry"
  project_id = var.project_id
  region = var.region
  repository_id = "repo"
}

module "upload_service" {
  source = "./modules/cloudrun"
  project_id = var.project_id
  region = "us-central1"
  service_name = "upload-service"
  image = "us-central1-docker.pkg.dev/${var.project_id}/repo/upload-service"
  env_vars = {
    RAW_BUCKET = module.storage.cost_data_bucket
    PROCESSED_BUCKET = module.storage.processed_data_bucket
  }
  service_account = module.iam.sre_sa_email
  depends_on = [module.project_services.cloudrun_service]
}

module "upload_service_trigger" {
  source = "./modules/cloudbuild"
  image = "us-central1-docker.pkg.dev/${var.project_id}/repo/upload-service"
  project_id = var.project_id
  branch = "^.*$"
  repo_owner = var.repo_owner
  repo_name = var.repo_name  
  service_name = "upload-service"
  depends_on = [
    module.project_services.cloudbuild_service
  ]
  service_account = module.iam.sre_sa_id

  included_files = ["services/upload-service/**"]
}

module "pubsub" {
  source = "./modules/pubsub"
  project_id = var.project_id
}


module "bigquery" {
  source = "./modules/bigquery"
  project_id = var.project_id
  region = var.region
}

module "sql" {
  source = "./modules/sql"
  region = var.region
  db_password = var.db_password
  depends_on = [module.project_services.sqladmin_service]
}

module "network" {
  source = "./modules/network"
  servicenetworking_dep = module.project_services.servicenetworking_service
  depends_on = [module.project_services.compute_service]
}

module "hirebizz_log_sink" {
  source = "./modules/log-sinks"

  source_project_id = "job-tracker-app-458110"
  dest_project_id   = var.project_id
  topic_name        = "projects-logs"

  log_filter = "severity>=ERROR"
}

module "analysis_engine_service" {
  source = "./modules/cloudrun"

  project_id   = var.project_id
  region       = var.region
  service_name = "analysis-engine"
  image        = "us-central1-docker.pkg.dev/${var.project_id}/repo/analysis-engine"

  env_vars = {
    DATABASE_URL       = "postgres://analyzer:${var.db_password}@/analyzer?host=/cloudsql/${var.cloudsql_instance_connection_name}&sslmode=disable"  
    GCP_PROJECT_ID     = var.project_id
    BQ_DATASET         = var.bq_dataset
    BQ_TABLE           = var.bq_table
    TARGET_PROJECT_ID  = "job-tracker-app-458110" 
  }

  cloud_sql_instances = [
    var.cloudsql_instance_connection_name
  ]

  service_account = module.iam.analysis_engine_email
  depends_on = [
    module.project_services.cloudrun_service,
    module.sql
  ]

}


module "analysis_engine_trigger" {
  source = "./modules/cloudbuild"

  project_id   = var.project_id
  service_name = "analysis-engine"
  image        = "us-central1-docker.pkg.dev/${var.project_id}/repo/analysis-engine"

  repo_owner = var.repo_owner
  repo_name  = var.repo_name
  branch     = "^main$"

  service_account = module.iam.analysis_engine_id

  depends_on = [
    module.project_services.cloudbuild_service
  ]
  included_files = ["services/analysis-engine/**"]
}
