package speechandtextanalytics_settings

import (
	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

/*
The resource_genesyscloud_speechandtextanalytics_settings_test.go contains all of the test cases for running the resource
tests for speechandtextanalytics_settings.
*/

func TestAccResourceSpeechandtextanalyticsSettings(t *testing.T) {
	t.Parallel()
	var ()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { util.TestAccPreCheck(t) },
		ProviderFactories: provider.GetProviderFactories(providerResources, providerDataSources),
		Steps:             []resource.TestStep{},
		CheckDestroy:      testVerifySpeechandtextanalyticsSettingsDestroyed,
	})
}

func testVerifySpeechandtextanalyticsSettingsDestroyed(state *terraform.State) error {
	return nil
}
