module "storage" {
  source = "./modules/storage"

  prefix = var.project_id
  region = var.region
}