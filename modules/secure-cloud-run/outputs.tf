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
  value       = module.cloud_run_network.connector_id
}

output "load_balancer_ip" {
  value       = module.cloud_run_core.load-balancer-ip
  description = "IP Address used by Load Balancer."
}

output "revision" {
  value       = module.cloud_run_core.revision
  description = "Deployed revision for the service."
}

output "service_url" {
  value       = module.cloud_run_core.service_url
  description = "The URL on which the deployed service is available."
}

output "service_id" {
  value       = module.cloud_run_core.service_id
  description = "Unique Identifier for the created service."
}

output "service_status" {
  value       = module.cloud_run_core.service_status
  description = "Status of the created service."
}

output "domain_map_id" {
  value       = module.cloud_run_core.domain_map_id
  description = "Unique Identifier for the created domain map."
}

output "domain_map_status" {
  value       = module.cloud_run_core.domain_map_status
  description = "Status of Domain mapping."
}
