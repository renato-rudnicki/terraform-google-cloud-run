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

output "key_self_link" {
  description = "Key self link."
  value       = module.cloud_run_kms.keys[var.key_name]
}

output "keyring_self_link" {
  description = "Self link of the keyring."
  value       = module.cloud_run_kms.keyring
}

output "keyring_resource" {
  description = "Keyring resource."
  value       = module.cloud_run_kms.keyring_resource
}

# output "kms_project_id" {
#   value = module.cloud_run_kms.kms_project_id
#   description = "The project where KMS will be created."
# }

# output "keyring_name" {
#   value = module.cloud_run_kms.keyring_name
#   description = "Keyring name."
# }

# output "keys" {
#   value = module.cloud_run_kms.keys
#   description = "Key name."
# }
