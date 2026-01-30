resource "google_compute_network" "vpc" {
    name = "analyzer-vpc"
    auto_create_subnetworks = true
}