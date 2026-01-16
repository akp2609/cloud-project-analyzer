
resource "google_project_service" "iam_api" {
  project = var.project_id
  service = "iam.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "pubsub" {
  project = var.project_id
  service = "pubsub.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "cloudbuild" {
  project             = var.project_id
  service             = "cloudbuild.googleapis.com"
  disable_on_destroy  = false
}

resource "google_project_service" "cloudrun" {
  project             = var.project_id
  service             = "run.googleapis.com"
  disable_on_destroy  = false
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
  raw_bucket_name = module.storage.cost_data_bucket
  processed_bucket_name = module.storage.processed_data_bucket
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
  depends_on = [google_project_service.cloudrun]
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
    google_project_service.cloudbuild
  ]
  service_account = module.iam.sre_sa_id
}

module "bigquery" {
  source = "./modules/bigquery"
  project_id = var.project_id
  region = var.region
}