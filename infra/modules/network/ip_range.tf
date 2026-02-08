resource "google_compute_global_address" "sql_range" {
  name = "sql-private-range"
  purpose = "VPC_PEERING"
  address_type = "INTERNAL"
  prefix_length = 16
  network = google_compute_network.vpc.self_link
}