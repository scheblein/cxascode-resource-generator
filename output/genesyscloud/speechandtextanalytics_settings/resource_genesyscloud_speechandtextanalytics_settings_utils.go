package speechandtextanalytics_settings

import (
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
)

/*
The resource_genesyscloud_speechandtextanalytics_settings_utils.go file contains various helper methods to marshal
and unmarshal data into formats consumable by Terraform and/or Genesys Cloud.
*/

// getSpeechandtextanalyticsSettingsFromResourceData maps data from schema ResourceData object to a platformclientv2.Speechtextanalyticssettingsresponse
func getSpeechandtextanalyticsSettingsFromResourceData(d *schema.ResourceData) platformclientv2.Speechtextanalyticssettingsresponse {
	return platformclientv2.Speechtextanalyticssettingsresponse{
		DefaultProgram: buildAddressableEntityRef(d.Get("default_program").([]interface{})),
		// TODO: Handle expected_dialects property
		TextAnalyticsEnabled: platformclientv2.Bool(d.Get("text_analytics_enabled").(bool)),
		AgentEmpathyEnabled:  platformclientv2.Bool(d.Get("agent_empathy_enabled").(bool)),
	}
}

// buildAddressableEntityRefs maps an []interface{} into a Genesys Cloud *[]platformclientv2.Addressableentityref
func buildAddressableEntityRefs(addressableEntityRefs []interface{}) *[]platformclientv2.Addressableentityref {
	addressableEntityRefsSlice := make([]platformclientv2.Addressableentityref, 0)
	for _, addressableEntityRef := range addressableEntityRefs {
		var sdkAddressableEntityRef platformclientv2.Addressableentityref
		addressableEntityRefsMap, ok := addressableEntityRef.(map[string]interface{})
		if !ok {
			continue
		}

		addressableEntityRefsSlice = append(addressableEntityRefsSlice, sdkAddressableEntityRef)
	}

	return &addressableEntityRefsSlice
}

// flattenAddressableEntityRefs maps a Genesys Cloud *[]platformclientv2.Addressableentityref into a []interface{}
func flattenAddressableEntityRefs(addressableEntityRefs *[]platformclientv2.Addressableentityref) []interface{} {
	if len(*addressableEntityRefs) == 0 {
		return nil
	}

	var addressableEntityRefList []interface{}
	for _, addressableEntityRef := range *addressableEntityRefs {
		addressableEntityRefMap := make(map[string]interface{})

		addressableEntityRefList = append(addressableEntityRefList, addressableEntityRefMap)
	}

	return addressableEntityRefList
}
