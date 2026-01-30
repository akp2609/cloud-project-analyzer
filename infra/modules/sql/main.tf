resource "google_sql_database_instance" "control_plane" {
  name = "control-plane-db"
  database_version = "POSTGRES_14"
  region = var.region

  settings {
    tier = "db-f1-micro"

    disk_size = 10
    disk_type = "PD_SSD"

    availability_type = "ZONAL"

    ip_configuration {
        ipv4_enabled = true
    }

    backup_configuration {
      enabled = false
    }
  }

  deletion_protection = false
}

resource "google_sql_database" "analyzer" {
    name = "analyzer"
    instance = google_sql_database_instance.control_plane.name
}

resource "google_sql_user" "analyzer" {
    name = "analyzer"
    instance = google_sql_database_instance.control_plane.name
    password = var.db_password
}