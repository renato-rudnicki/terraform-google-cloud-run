// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloud_run

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
)

func TestCloudRun(t *testing.T) {
	secure_cloud_run := tft.NewTFBlueprintTest(t)
	secure_cloud_run.DefineVerify(func(assert *assert.Assertions) {
		//secure_cloud_run.DefaultVerify(assert)

		fmt.Println("------------------------------- KMS TEST -------------------------------")
		kmsProjectName := secure_cloud_run.GetStringOutput("kms_project_id")
		kmsKeyRingName := secure_cloud_run.GetStringOutput("keyring_name")
		kmsKey := secure_cloud_run.GetStringOutput("key_name")
		fmt.Println(kmsProjectName)
		fmt.Println(kmsKeyRingName)
		fmt.Println(kmsKey)
		opKMS := gcloud.Runf(t, "kms keys list --keyring=%s --project=%s --location us-central1", kmsKeyRingName, kmsProjectName).Array()
		keyFullName := fmt.Sprintf("projects/%s/locations/us-central1/keyRings/%s/cryptoKeys/%s", kmsProjectName, kmsKeyRingName, kmsKey)
		assert.Equal(keyFullName, opKMS[0].Get("name").String(), fmt.Sprintf("should have key %s", keyFullName))

		fmt.Println("------------------------------- CLOUD RUN TEST -------------------------------")
		serviceId := secure_cloud_run.GetStringOutput("service_id")
		projectId := secure_cloud_run.GetStringOutput("project_id")
		fmt.Println(serviceId)
		fmt.Println(projectId)
		opCloudRun := gcloud.Runf(t, "run services list --project=%s", projectId).Array()
		cloudRunId := fmt.Sprintf("locations/us-central1/namespaces/%s/services/%s", projectId, opCloudRun[0].Get("metadata.name").String())
		assert.Equal(serviceId, cloudRunId, fmt.Sprintf("Should have same id: %s", serviceId))
	})
	secure_cloud_run.Test()
}
