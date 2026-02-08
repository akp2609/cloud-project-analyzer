variable "project_id" {}
variable "region" {
  default = "us-central1"
}
variable "service_account" {}
variable "env_vars" {
  type = map(string)
  default = {}
}
variable "image" {
  type = string
}
variable "service_name" {
  type = string
}
