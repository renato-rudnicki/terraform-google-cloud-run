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
		secure_cloud_run.DefaultVerify(assert)

		// gcloudArgsCloudRun := gcloud.WithCommonArgs([]string{"--format", "json"})
		// opCloudRun := gcloud.Run(t, "run services list", gcloudArgsCloudRun).Array()
		// assert.Equal(fmt.Sprintf("hello-world-with-apis-test07"), opCloudRun[0].Get("metadata.name").String(), "has expected name")

		kmsProjectName := secure_cloud_run.GetStringOutput("kms_project_id")
		kmsKeyRingName := secure_cloud_run.GetStringOutput("data_ingestion_topic_name")
		//gcloudArgsCloudRun := gcloud.WithCommonArgs([]string{"--format", "json"})
		opCloudRun := gcloud.Run(t, fmt.Sprintf("alpha kms keys list --keyring=%s --project=%s --location us-central1", kmsKeyRingName, kmsProjectName))
		assert.Equal(fmt.Sprintf("projects/%s/locations/us-central1/keyRings/%s/cryptoKeys/cloud-run07", kmsProjectName, kmsKeyRingName), opCloudRun.Get("name").String(), "has expected name")
	})
	secure_cloud_run.Test()
}
