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
