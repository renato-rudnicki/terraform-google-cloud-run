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

  connector_name                          = "serverless-connector"
  subnet_name                             = "vpc-subnet"
  vpc_project_id                          = var.vpc_project_id
  serverless_project_id                   = var.serverless_project_id
  kms_project_id                          = var.kms_project_id
  shared_vpc_name                         = "vpc-p-shared-restricted"
  ip_cidr_range                           = var.ip_cidr_range
  key_name                                = "cloud-run"
  keyring_name                            = "cloud-run-keyring"
  service_name                            = "hello-world"
  location                                = "us-central1"
  region                                  = "us-central1"
  image                                   = "us-docker.pkg.dev/cloudrun/container/hello"
  cloud_run_sa                            = var.cloud_run_sa
  connector_on_host_project               = true
}
