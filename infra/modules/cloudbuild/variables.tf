variable "project_id" {
  type        = string
  description = "The GCP project ID."
}

variable "service_name" {
  type        = string
  description = "The name of the service."
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

variable "image" {
  type        = string
  description = "The full name of the container image to build and push."
}