package speechandtextanalytics_dictionaryfeedback

import (
	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

/*
The resource_genesyscloud_speechandtextanalytics_dictionaryfeedback_test.go contains all of the test cases for running the resource
tests for speechandtextanalytics_dictionaryfeedback.
*/

func TestAccResourceSpeechandtextanalyticsDictionaryfeedback(t *testing.T) {
	t.Parallel()
	var ()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { util.TestAccPreCheck(t) },
		ProviderFactories: provider.GetProviderFactories(providerResources, providerDataSources),
		Steps:             []resource.TestStep{},
		CheckDestroy:      testVerifySpeechandtextanalyticsDictionaryfeedbackDestroyed,
	})
}

func testVerifySpeechandtextanalyticsDictionaryfeedbackDestroyed(state *terraform.State) error {
	return nil
}
