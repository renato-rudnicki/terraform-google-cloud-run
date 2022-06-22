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

  #cloud_run_network
  # vpc connector
  # firewall rules
  # iam roles
  #
  # connector_name        = var.connector_name
  # subnet_name           = var.subnet_name
  # location              = var.location
  # vpc_project_id        = var.vpc_project_id
  # serverless_project_id = var.serverless_project_id
  # shared_vpc_name       = var.shared_vpc_name
  # connector_on_host_project = true
  # ip_cidr_range         = var.ip_cidr_range

  #cloud_run_security
  # kms
  # org policies
  #
  # kms_project_id        = var.kms_project_id
  # location              = var.location
  # serverless_project_id = var.serverless_project_id
  # prevent_destroy       = var.prevent_destroy
  # key_name              = var.key_name
  # keyring_name          = var.keyring_name
  # key_rotation_period   = var.key_rotation_period
  # key_protection_level  = var.key_protection_level

  # cloud_run_core
  # cloud run
  # load balancer
  # ssl
  # cloud armor
  service_name          = var.service_name
  location              = var.location
  serverless_project_id = var.serverless_project_id
  image                 = var.image
  cloud_run_sa          = var.cloud_run_sa
  vpc_connector_id      = module.cloud_run_network.connector_id
  encryption_key        = module.cloud_run_security.key
  env_vars              = var.env_vars
  members               = var.members
  name                            = var.lb_name
  project                         = var.serverless_project_id
  ssl                             = var.ssl
  managed_ssl_certificate_domains = [var.domain]
}
