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

variable "project_logs_topic"{
  type = string
}

variable "dest_project_id"{
  type = string
}

variable "log_filter"{
  type = string
}

variable "topic_name"{
  type = string
}

variable "project_log_sink_writer_identity" {
  type = string
}