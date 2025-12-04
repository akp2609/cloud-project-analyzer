variable "prefix" {
  type = string
}

variable "region" {
  type = string
  default = "us-central1"

}

variable "project_id" {
  type = string
}


variable "raw_bucket_name" {
  type = string
}

variable "processed_bucket_name" {
    type = string
}