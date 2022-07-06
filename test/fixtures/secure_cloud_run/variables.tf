/**
 * Copyright 2021 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

variable "ssl" {
  description = "Run load balancer on HTTPS and provision managed certificate with provided `domain`."
  type        = bool
  default     = true
}

variable "domain" {
  description = "Domain name to run the load balancer on. Used if `ssl` is `true`. Modify the default value below for your `domain` name."
  type        = string
  default     = "my-domain.com"
}

variable "lb_name" {
  description = "Name for load balancer and associated resources."
}

variable "location" {
  description = "The location where resources are going to be deployed."
  type        = string
}

variable "region" {
  description = "Location for load balancer and Cloud Run resources."
  type        = string
}

variable "serverless_project_id" {
  description = "The project where cloud run is going to be deployed."
  type        = string
}

variable "service_name" {
  description = "The name of the Cloud Run service to create."
  type        = string
}

variable "image" {
  description = "GAR hosted image URL to deploy."
  type        = string
}

variable "cloud_run_sa" {
  description = "Service account to be used on Cloud Run."
  type        = string
}

variable "vpc_connector_id" {
  description = "VPC Connector id in the format projects/PROJECT/locations/LOCATION/connectors/NAME."
  type        = string
}

variable "encryption_key" {
  description = "CMEK encryption key self-link expected in the format projects/PROJECT/locations/LOCATION/keyRings/KEY-RING/cryptoKeys/CRYPTO-KEY."
  type        = string
}

variable "env_vars" {
  type = list(object({
    value = string
    name  = string
  }))
  description = "Environment variables."
  default     = []
}

variable "members" {
  type        = list(string)
  description = "Users/SAs to be given invoker access to the service with the prefix `serviceAccount:' for SAs and `user:` for users."
  default     = []
}

variable "connector_name" {
  description = "The name of the serverless connector which is going to be created."
  type        = string
}

variable "subnet_name" {
  description = "Subnet name to be re-used to create Serverless Connector."
  type        = string
}

variable "shared_vpc_name" {
  description = "Shared VPC name which is going to be used to create Serverless Connector."
  type        = string
}

variable "ip_cidr_range" {
  description = "The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported"
  type        = string
}

variable "vpc_project_id" {
  description = "The project where shared vpc is."
  type        = string
}

variable "kms_project_id" {
  description = "The project where KMS will be created."
  type        = string
}

variable "prevent_destroy" {
  description = "Set the prevent_destroy lifecycle attribute on keys."
  type        = bool
  default     = false
}

variable "keyring_name" {
  description = "Keyring name."
  type        = string
}

variable "key_rotation_period" {
  description = "Period of key rotation in seconds."
  type        = string
  default     = "2592000s"
}

variable "key_name" {
  description = "Key name."
  type        = string
}

variable "key_protection_level" {
  description = "The protection level to use when creating a version based on this template. Possible values: [\"SOFTWARE\", \"HSM\"]"
  type        = string
  default     = "HSM"
}

variable "artifact_registry_repository_project_id" {
  description = "Artifact Registry Repository Project ID to grant serverless identity viewer role."
  type        = string
}

variable "artifact_registry_repository_location" {
  description = "Artifact Registry Repository location to grant serverless identity viewer role."
  type        = string
}

variable "artifact_registry_repository_name" {
  description = "Artifact Registry Repository name to grant serverless identity viewer role"
  type        = string
}

variable "connector_on_host_project" {
  description = "Connector is going to be created on the host project if true. When false, connector is going to be created on service project. For more information, access [documentation](https://cloud.google.com/run/docs/configuring/connecting-shared-vpc)."
  type        = bool
  default     = true
}
