provider "google" {
  credentials = file(var.creds_file_path)
  project     = var.project_id
  region      = var.region
}

