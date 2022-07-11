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
	"github.com/tidwall/gjson"
)

func getResultFieldStrSlice(rs []gjson.Result, field string) []string {
	s := make([]string, 0)
	for _, r := range rs {
		s = append(s, r.Get(field).String())
	}
	return s
}

func TestCloudRun(t *testing.T) {
	secure_cloud_run := tft.NewTFBlueprintTest(t)
	secure_cloud_run.DefineVerify(func(assert *assert.Assertions) {
		//secure_cloud_run.DefaultVerify(assert)
		kmsProjectName := secure_cloud_run.GetStringOutput("kms_project_id")
		kmsKeyRingName := secure_cloud_run.GetStringOutput("keyring_name")
		kmsKey := secure_cloud_run.GetStringOutput("key_name")
		serviceId := secure_cloud_run.GetStringOutput("service_id")
		projectId := secure_cloud_run.GetStringOutput("project_id")
		vpcProjectId := secure_cloud_run.GetStringOutput("vpc_project_id")
		connectorId := secure_cloud_run.GetStringOutput("connector_id")
		run_identity_services_sa := secure_cloud_run.GetStringOutput("run_identity_services_sa")

		fmt.Println("------------------------------- KMS TEST -------------------------------")
		opKMS := gcloud.Runf(t, "kms keys list --keyring=%s --project=%s --location us-central1", kmsKeyRingName, kmsProjectName).Array()
		keyFullName := fmt.Sprintf("projects/%s/locations/us-central1/keyRings/%s/cryptoKeys/%s", kmsProjectName, kmsKeyRingName, kmsKey)
		assert.Equal(keyFullName, opKMS[0].Get("name").String(), fmt.Sprintf("should have key %s", keyFullName))

		fmt.Println("------------------------------- CLOUD RUN TEST -------------------------------")
		opCloudRun := gcloud.Runf(t, "run services list --project=%s", projectId).Array()
		cloudRunId := fmt.Sprintf("locations/us-central1/namespaces/%s/services/%s", projectId, opCloudRun[0].Get("metadata.name").String())
		assert.Equal(serviceId, cloudRunId, fmt.Sprintf("Should have same id: %s", serviceId))

		fmt.Println("------------------------------- VPC CONNECTOR TEST -------------------------------")
		opVPCConnector := gcloud.Runf(t, "compute networks vpc-access connectors list --region=us-central1 --project=%s", vpcProjectId).Array()
		vpcConnectorNames := getResultFieldStrSlice(opVPCConnector, "name")
		assert.Containsf(vpcConnectorNames, connectorId, fmt.Sprintf("Should have same id: %s", connectorId))

		fmt.Println("------------------------------- CLOUD ARMOR TEST -------------------------------")
		opCloudArmor := gcloud.Runf(t, "compute security-policies list --project=%s", projectId).Array()
		assert.Equal(fmt.Sprintf("cloud-armor-waf-policy"), opCloudArmor[0].Get("name").String(), "has expected name ")

		fmt.Println("------------------------------- FIREWALL TEST -------------------------------")
		expectFirewalls := []string{
			"serverless-to-vpc-connector", "vpc-connector-to-serverless",
			"vpc-connector-to-serverless-lb", "vpc-connector-health-checks",
			"vpc-connector-requests",
		}
		opFirewall := gcloud.Runf(t, "compute firewall-rules list --project=%s", vpcProjectId).Array()
		actualFirewalls := getResultFieldStrSlice(opFirewall, "name")
		assert.Subset(actualFirewalls, expectFirewalls, "Should have all 5 firewall rules")

		fmt.Println("------------------------------- IAM TEST -------------------------------")
		iamFilter := fmt.Sprintf("bindings.members:'serviceAccount:%s'", run_identity_services_sa)
		iamOpts := gcloud.WithCommonArgs([]string{"--flatten", "bindings", "--filter", iamFilter, "--format", "json"})
		orgIamPolicyRoles := gcloud.Run(t, fmt.Sprintf("projects get-iam-policy %s", vpcProjectId), iamOpts).Array()
		listRoles := getResultFieldStrSlice(orgIamPolicyRoles, "bindings.role")
		assert.Containsf(listRoles, "roles/vpcaccess.user", fmt.Sprintf("Service account %s should have VPC Access User role", run_identity_services_sa))

		fmt.Println("------------------------------- ORG POLICIES TEST -------------------------------")
		//orgArgs := gcloud.WithCommonArgs([]string{"--flatten", "listPolicy.allowedValues[]", "--format", "json"})

		opOrgPolicies := gcloud.Run(t, fmt.Sprintf("resource-manager org-policies describe constraints/run.allowedIngress -project=%s", projectId)).Array()
		fmt.Print(opOrgPolicies)
		//opOrgPolicies := gcloud.Run(t, fmt.Sprintf("resource-manager org-policies list --project=%s", projectId), orgArgs).Array()
		assert.Contains(opOrgPolicies[0].Get("listPolicy.allowedValues").String(), fmt.Sprintf("is:internal-and-cloud-load-balancing"), "has expected name")
		//assert.Contains(opOrgPolicies[1].Get("listPolicy.allowedValues").String(), fmt.Sprintf("private-ranges-only"), "has expected name")
	})
	secure_cloud_run.Test()
}
