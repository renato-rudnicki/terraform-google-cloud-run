/**
 * Copyright 2025 Google LLC
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

variable "regions" {
  type        = list(string)
  description = "Regions where serverless vpc access connectors will be created."
  default     = ["us-west1", "europe-west1"]
}

variable "project_id" {
  type        = string
  description = "Project where the Cloud Run v2 will be deployed."
}

variable "cloud_run_deletion_protection" {
  type        = bool
  description = "Prevents Terraform from destroying/recreating Cloud Run jobs/services."
  default     = true
}

variable "vpc_connectors" {
  description = "Configuration for Serverless VPC Access connectors by regions."
  type = map(object({
    name        = string
    region      = string
    subnet_name = string
  }))
}
