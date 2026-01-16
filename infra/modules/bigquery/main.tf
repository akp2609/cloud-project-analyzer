resource "google_bigquery_dataset" "cost_analytics" {
  dataset_id = "cost_analytics"
  project    = var.project_id
  location   = var.region

  delete_contents_on_destroy = false
}

resource "google_bigquery_table" "processed_costs" {
  dataset_id = google_bigquery_dataset.cost_analytics.dataset_id
  table_id   = "processed_costs"
  project    = var.project_id

  schema = jsonencode([
    {
      name = "tenant"
      type = "STRING"
      mode = "REQUIRED"
    },
    {
      name = "service"
      type = "STRING"
      mode = "REQUIRED"
    },
    {
      name = "date"
      type = "DATE"
      mode = "REQUIRED"
    },
    {
      name = "total_cost"
      type = "FLOAT"
      mode = "REQUIRED"
    }
  ])
}
