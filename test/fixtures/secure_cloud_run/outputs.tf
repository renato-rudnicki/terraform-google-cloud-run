/**
 * Copyright 2022 Google LLC
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

output "connector_id" {
  description = "VPC serverless connector ID."
  value       = module.example_cloud_run.connector_id
}

output "load_balancer_ip" {
  value       = module.example_cloud_run.load_balancer_ip
  description = "IP Address used by Load Balancer."
}

output "revision" {
  value       = module.example_cloud_run.revision
  description = "Deployed revision for the service."
}

output "service_url" {
  value       = module.example_cloud_run.service_url
  description = "The URL on which the deployed service is available."
}

output "service_id" {
  value       = module.example_cloud_run.service_id
  description = "Unique Identifier for the created service."
}

output "service_status" {
  value       = module.example_cloud_run.service_status
  description = "Status of the created service."
}

output "domain_map_id" {
  value       = module.example_cloud_run.domain_map_id
  description = "Unique Identifier for the created domain map."
}

output "domain_map_status" {
  value       = module.example_cloud_run.domain_map_status
  description = "Status of Domain mapping."
}

output "kms_project_id" {
  value = local.kms_project_id #module.example_cloud_run.kms_project_id
  description = "The project where KMS will be created."
}

output "keyring_name" {
  value = local.keyring_name #module.example_cloud_run.keyring_name
  description = "Keyring name."
}

output "key_name" {
  value = local.key_name #module.example_cloud_run.key_name
  description = "Key name."
}

output "serverless_project_id" {
  value = local.serverless_project_id #module.example_cloud_run.key_name
  description = "The project where cloud run is going to be deployed."
}

output "project_id" {
  value = module.example_cloud_run.project_id
  description = "The project where cloud run is going to be deployed."
}

output "vpc_project_id" {
  value = module.example_cloud_run.vpc_project_id
  description = "The project where VPC Connector is going to be deployed."
}