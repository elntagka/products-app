variable "project_id" {
  type        = string
  description = "The project ID to host the cluster in"
}

variable "region" {
  type        = string
  description = "The region to host the cluster in"
  default     = "europe-west1"
}

variable "creds_file_path" {
  type        = string
  description = "gcp credentials file path"
  default     = "./creds/key.json"
}

variable "cluster_name" {
  description = "The name for the GKE cluster"
  default     = "products-app-cluster"
}
variable "env_name" {
  description = "The environment for the GKE cluster"
  default     = "prod"
}

variable "network" {
  description = "The VPC network created to host the cluster in"
  default     = "gke-network"
}
variable "subnetwork" {
  description = "The subnetwork created to host the cluster in"
  default     = "gke-subnet"
}
variable "ip_range_pods_name" {
  description = "The secondary ip range to use for pods"
  default     = "ip-range-pods"
}

variable "ip_range_services_name" {
  description = "The secondary ip range to use for services"
  default     = "ip-range-services"
}