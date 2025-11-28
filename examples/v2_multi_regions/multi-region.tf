terraform {
  required_version = ">= 1.3"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 6, < 7.2"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = ">= 6, < 7.2"
    }
  }
}

provider "google" {
  project = "prj-secure-cloud-run-1fbd"
}

provider "google-beta" {
  project = "prj-secure-cloud-run-1fbd"
}

variable "regions" {
  type    = list(string)
  default = ["us-west1", "europe-west1"]
}

locals {
  vpc_connectors = {
    "us-west1"     = "projects/prj-secure-cloud-run-1fbd/locations/us-west1/connectors/con-secure-cloud-run"          # Remove connectors hardcode
    "europe-west1" = "projects/prj-secure-cloud-run-1fbd/locations/europe-west1/connectors/con-secure-cloud-run-mult" # Remove connectors hardcode
  }
}

resource "google_cloud_run_v2_service" "multi" {
  provider = google-beta

  for_each = toset(var.regions)

  name     = "my-multiregion-service-${each.key}"
  location = each.key

  ingress = "INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER"

  template {
    vpc_access {
      connector = local.vpc_connectors[each.key]
      egress    = "PRIVATE_RANGES_ONLY"
    }

    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello:latest"
    }
  }

  traffic {
    percent = 100
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
  }
}

