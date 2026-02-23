resource "google_service_account" "sre_cost_sa" {
  account_id = "${var.prefix}-sre-sa"
  display_name = "SRE Cost Service Account"
}

resource "google_service_account" "analysis_engine_sa" {
  account_id = "analysis-engine-sa"
  display_name = "Analysis Engine Service Account"
}

resource "google_storage_bucket_iam_member" "raw_bucket_binding" {
  bucket = var.raw_bucket_name
  role = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_storage_bucket_iam_member" "processed_bucket_binding" {
  bucket = var.processed_bucket_name
  role = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_project_iam_member" "raw_runtime_user" {
  project = var.project_id
  role = "roles/run.invoker"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_project_iam_member" "artifact_registry_writer" {
  project = var.project_id
  role = "roles/artifactregistry.writer"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_project_iam_member" "monitoring"{
  project = var.dest_project_id
  role = "roles/monitoring.viewer"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_project_iam_member" "sql_client"{
  project = var.project_id
  role = "roles/cloudsql.client"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_project_iam_member" "pubsub_subscriber" {
  project = var.dest_project_id
  role = "roles/pubsub.subscriber"
  member = "serviceAccount:${google_service_account.sre_cost_sa.email}"
}

resource "google_pubsub_topic_iam_member" "sink_publisher" {
  project = var.dest_project_id
  topic   = var.topic_name
  role    = "roles/pubsub.publisher"
  member  = "${var.project_log_sink_writer_identity}"
  depends_on = [var.project_log_sink_writer_identity]
}

resource "google_project_iam_member" "analysis_engine_sa" {
  project = var.project_id
  role = "roles/cloudsql.client"
  member = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

data "google_project" "project" {
  project_id = var.dest_project_id
}

resource "google_cloud_run_service_iam_member" "pubsub_invoker" {
  project  = var.dest_project_id
  location = "us-central1"
  service  = "logs-ingestor"

  role   = "roles/run.invoker"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "analysis_bq_viewer" {
  project = var.project_id
  role    = "roles/bigquery.dataViewer"
  member  = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_project_iam_member" "analysis_monitoring_viewer" {
  project = var.project_id
  role    = "roles/monitoring.viewer"
  member  = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_project_iam_member" "analysis_run_invoker" {
  project = var.project_id
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_project_iam_member" "analysis_bq_user" {
  project = var.project_id
  role    = "roles/bigquery.user"
  member  = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_project_iam_member" "analysis_engine_logs_writer" {
  project = var.project_id
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_project_iam_member" "analysis_engine_artifact_writer" {
  project = var.project_id
  role    = "roles/artifactregistry.writer"
  member  = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_storage_bucket_iam_member" "tf_state_bucket_access" {
  bucket =  var.state_bucket_name
  role   = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_secret_manager_secret_iam_member" "terraform_tfvars_access" {
  project   = var.project_id
  secret_id = var.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_project_iam_member" "analysis_engine_editor" {
  project = var.project_id
  role    = "roles/editor"
  member  = "serviceAccount:${google_service_account.analysis_engine_sa.email}"
}

resource "google_service_account" "extra_service_accounts" {
  for_each = var.extra_services

  account_id   = each.key
  display_name = "${each.key} service account"
}

locals {
  extra_service_role_map = merge([
    for svc, config in var.extra_services : {
      for role in config.roles :
      "${svc}-${role}" => {
        service = svc
        role    = role
      }
    }
  ]...)
}

resource "google_project_iam_member" "extra_service_roles" {
  for_each = local.extra_service_role_map

  project = var.project_id
  role    = each.value.role
  member  = "serviceAccount:${google_service_account.extra_service_accounts[each.value.service].email}"
}

resource "google_project_iam_member" "control_plane_external_access" {
  for_each = toset(var.external_projects)

  project = each.value
  role    = "roles/viewer"
  member  = "serviceAccount:${google_service_account.extra_service_accounts["project-hook"].email}"
}

resource "google_storage_bucket_iam_member" "cost_processor_raw_bucket" {
  bucket = var.raw_bucket_name
  role   = "roles/storage.objectViewer"
  member = "serviceAccount:${google_service_account.extra_service_accounts["cost-processor"].email}"
}

resource "google_storage_bucket_iam_member" "cost_processor_processed_bucket" {
  bucket = var.processed_bucket_name
  role   = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.extra_service_accounts["cost-processor"].email}"
}


resource "google_bigquery_dataset_iam_member" "cost_processor_dataset_access" {
  dataset_id = "cost_analytics"
  role       = "roles/bigquery.dataEditor"
  member     = "serviceAccount:${google_service_account.extra_service_accounts["cost-processor"].email}"
}


resource "google_project_iam_member" "cost_processor_run_invoker" {
  project = var.project_id
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_service_account.extra_service_accounts["cost-processor"].email}"
}
