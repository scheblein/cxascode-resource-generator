package organization_presence_definition

import (
	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

/*
Test Class for the organization presence definition Data Source
*/

func TestAccDataSourceOrganizationPresenceDefinition(t *testing.T) {
	t.Parallel()
	var ()
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { util.TestAccPreCheck(t) },
		ProviderFactories: provider.GetProviderFactories(providerResources, providerDataSources),
		Steps:             []resource.TestStep{},
	})
}
