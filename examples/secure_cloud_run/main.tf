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

module "secure_cloud_run" {
  source = "../../modules/secure-cloud-run"

  connector_name                          = var.connector_name
  subnet_name                             = var.subnet_name
  vpc_project_id                          = var.vpc_project_id
  serverless_project_id                   = var.serverless_project_id
  shared_vpc_name                         = var.shared_vpc_name
  ip_cidr_range                           = var.ip_cidr_range
  kms_project_id                          = var.kms_project_id
  prevent_destroy                         = var.prevent_destroy
  key_name                                = var.key_name
  keyring_name                            = var.keyring_name
  key_rotation_period                     = var.key_rotation_period
  key_protection_level                    = var.key_protection_level
  service_name                            = var.service_name
  location                                = var.location
  region                                  = var.region
  image                                   = var.image
  cloud_run_sa                            = var.cloud_run_sa
  artifact_registry_repository_project_id = var.serverless_project_id
  artifact_registry_repository_name       = var.artifact_registry_repository_name
  artifact_registry_repository_location   = var.region
  connector_on_host_project               = var.connector_on_host_project
}
