variable "project_id" {}
variable "service_name" {}
variable "region" {
  default = "us-central1"
}
variable "image" {}
variable "service_account" {}
variable "env_vars" {
  type = map(string)
  default = {}
}
