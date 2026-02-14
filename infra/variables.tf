variable "project_id" {
  type        = string
  description = "The GCP project ID."
}
variable "region" {
  type        = string
  description = "The GCP region for resources."
}
variable "zone" { 
  type        = string
  description = "The GCP zone for resources."
}
variable "repo_owner" {
  type        = string
  description = "The owner of the GitHub repository."
}
variable "repo_name" {
  type        = string
  description = "The name of the GitHub repository."
}
variable "branch" {
  type        = string
  description = "The branch name to trigger the build."
}
variable "service_name" {
  type        = string
  description = "The name of the service."
}

variable "db_password"{
  type = string
  description = "The name of the practice"
}

variable "bq_dataset" {
  description = "BigQuery dataset containing processed cost data"
  type        = string
}

variable "bq_table" {
  description = "BigQuery table containing processed cost data"
  type        = string
}

variable "target_project_id" {
  type = string
}

variable "cloudsql_instance_connection_name" {
  description = "Cloud SQL instance connection name for Analysis Engine"
  type        = string
}

variable "state_bucket" {
  description = "Terraform state bucket name"
  type        = string
}

