package organization_presence_definition

import (
	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

/*
The resource_genesyscloud_organization_presence_definition_test.go contains all of the test cases for running the resource
tests for organization_presence_definition.
*/

func TestAccResourceOrganizationPresenceDefinition(t *testing.T) {
	t.Parallel()
	var ()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { util.TestAccPreCheck(t) },
		ProviderFactories: provider.GetProviderFactories(providerResources, providerDataSources),
		Steps:             []resource.TestStep{},
		CheckDestroy:      testVerifyOrganizationPresenceDefinitionDestroyed,
	})
}

func testVerifyOrganizationPresenceDefinitionDestroyed(state *terraform.State) error {
	return nil
}
