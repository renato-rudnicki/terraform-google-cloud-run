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
		// kmsProjectName := secure_cloud_run.GetStringOutput("kms_project_id")
		// kmsKeyRingName := secure_cloud_run.GetStringOutput("keyring_name")
		// kmsKey := secure_cloud_run.GetStringOutput("key_name")
		// serviceId := secure_cloud_run.GetStringOutput("service_id")
		projectId := secure_cloud_run.GetStringOutput("project_id")
		vpcProjectId := secure_cloud_run.GetStringOutput("vpc_project_id")
		// connectorId := secure_cloud_run.GetStringOutput("connector_id")

		// fmt.Println("------------------------------- KMS TEST -------------------------------")
		// fmt.Println(kmsProjectName)
		// fmt.Println(kmsKeyRingName)
		// fmt.Println(kmsKey)
		// opKMS := gcloud.Runf(t, "kms keys list --keyring=%s --project=%s --location us-central1", kmsKeyRingName, kmsProjectName).Array()
		// keyFullName := fmt.Sprintf("projects/%s/locations/us-central1/keyRings/%s/cryptoKeys/%s", kmsProjectName, kmsKeyRingName, kmsKey)
		// assert.Equal(keyFullName, opKMS[0].Get("name").String(), fmt.Sprintf("should have key %s", keyFullName))

		// fmt.Println("------------------------------- CLOUD RUN TEST -------------------------------")
		// fmt.Println(serviceId)
		// fmt.Println(projectId)
		// opCloudRun := gcloud.Runf(t, "run services list --project=%s", projectId).Array()
		// cloudRunId := fmt.Sprintf("locations/us-central1/namespaces/%s/services/%s", projectId, opCloudRun[0].Get("metadata.name").String())
		// assert.Equal(serviceId, cloudRunId, fmt.Sprintf("Should have same id: %s", serviceId))

		// fmt.Println("------------------------------- VPC CONNECTOR TEST -------------------------------")
		// fmt.Println(vpcProjectId)
		// fmt.Println(connectorId)
		// opVPCConnector := gcloud.Runf(t, "compute networks vpc-access connectors list --region=us-central1 --project=%s", vpcProjectId).Array()
		// containsName := false
		// for _, connectorJson := range opVPCConnector {
		// 	if connectorJson.Get("name").String() == connectorId {
		// 		containsName = true
		// 		break
		// 	}
		// }
		// assert.True(containsName, fmt.Sprintf("Should have same id: %s", connectorId))

		// fmt.Println("------------------------------- CLOUD ARMOR TEST -------------------------------")
		// fmt.Println(projectId)
		// opCloudArmor := gcloud.Runf(t, "compute security-policies list --project=%s", projectId).Array()
		// assert.Equal(fmt.Sprintf("cloud-armor-waf-policy0a22"), opCloudArmor[0].Get("name").String(), "has expected name ")

		fmt.Println("------------------------------- FIREWALL TEST -------------------------------")
		//gcloud compute firewall-rules list --project=cloud-run-test-355015 --format=json
		opFirewall := gcloud.Runf(t, "compute firewall-rules list --project=%s", vpcProjectId).Array()
		serverlessToVpcConnector := false
		vpcConnectorToServerless := false
		vpcConnectorToLoadbalancer := false
		vpcConnectorHealthCheck := false
		vpcConnectorRequests := false
		for _, firewallJson := range opFirewall {
			if firewallJson.Get("name").String() == "serverless-to-vpc-connector" {
				serverlessToVpcConnector = true
			}
			if firewallJson.Get("name").String() == "vpc-connector-to-serverless" {
				vpcConnectorToServerless = true
			}
			if firewallJson.Get("name").String() == "vpc-connector-to-serverless-lb" {
				vpcConnectorToLoadbalancer = true
			}
			if firewallJson.Get("name").String() == "vpc-connector-health-checks" {
				vpcConnectorHealthCheck = true
			}
			if firewallJson.Get("name").String() == "vpc-connector-requests" {
				vpcConnectorRequests = true
			}
		}
		assert.True(serverlessToVpcConnector, "Should have serverless-to-vpc-connector firewall rule")
		assert.True(vpcConnectorToServerless, "Should have vpc-connector-to-serverless firewall rule")
		assert.True(vpcConnectorToLoadbalancer, "Should have vpc-connector-to-serverless-lb firewall rule")
		assert.True(vpcConnectorHealthCheck, "Should have vpc-connector-health-checks firewall rule")
		assert.True(vpcConnectorRequests, "Should have vpc-connector-requests firewall rule")

		fmt.Println("------------------------------- LOADBALANCER TEST -------------------------------")
		fmt.Println(projectId)
		opLoadBalancer := gcloud.Runf(t, "compute addresses list --filter=\"addressType=('EXTERNAL')\" --project=%s", projectId).Array()
		assert.Equal(fmt.Sprintf("tf-cr-lb-address"), opLoadBalancer[0].Get("name").String(), "has expected name ")
		//opLoadBalancer := gcloud.Runf(t, "compute addresses list --filter=\"addressType=('EXTERNAL')\" --project=%s", projectId)
		//assert.Equal(fmt.Sprintf("tf-cr-lb-address"), opLoadBalancer.Get("name").String(), "has expected name ")
	})
	secure_cloud_run.Test()
}
