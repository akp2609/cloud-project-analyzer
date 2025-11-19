terraform {
  backend "gcs" {
    bucket = "cloud-cost-resource-analyzer-tf-state"
    prefix = "terraform/state"
  }
}