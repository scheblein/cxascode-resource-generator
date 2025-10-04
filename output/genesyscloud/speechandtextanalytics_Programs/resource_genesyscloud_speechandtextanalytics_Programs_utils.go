package speechandtextanalytics_programs

import (
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
)

/*
The resource_genesyscloud_speechandtextanalytics_programs_utils.go file contains various helper methods to marshal
and unmarshal data into formats consumable by Terraform and/or Genesys Cloud.
*/

// getSpeechandtextanalyticsProgramsFromResourceData maps data from schema ResourceData object to a platformclientv2.Program
func getSpeechandtextanalyticsProgramsFromResourceData(d *schema.ResourceData) platformclientv2.Program {
	return platformclientv2.Program{
		Name:        platformclientv2.String(d.Get("name").(string)),
		Description: platformclientv2.String(d.Get("description").(string)),
		Published:   platformclientv2.Bool(d.Get("published").(bool)),
		Topics:      buildBaseTopicEntitiys(d.Get("topics").([]interface{})),
		// TODO: Handle tags property
		PublishedBy:   buildAddressableEntityRef(d.Get("published_by").([]interface{})),
		DatePublished: platformclientv2.String(d.Get("date_published").(string)),
		TopicLinksJob: buildAddressableEntityRef(d.Get("topic_links_job").([]interface{})),
	}
}

// buildBaseTopicEntitiys maps an []interface{} into a Genesys Cloud *[]platformclientv2.Basetopicentitiy
func buildBaseTopicEntitiys(baseTopicEntitiys []interface{}) *[]platformclientv2.Basetopicentitiy {
	baseTopicEntitiysSlice := make([]platformclientv2.Basetopicentitiy, 0)
	for _, baseTopicEntitiy := range baseTopicEntitiys {
		var sdkBaseTopicEntitiy platformclientv2.Basetopicentitiy
		baseTopicEntitiysMap, ok := baseTopicEntitiy.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkBaseTopicEntitiy.Name, baseTopicEntitiysMap, "name")

		baseTopicEntitiysSlice = append(baseTopicEntitiysSlice, sdkBaseTopicEntitiy)
	}

	return &baseTopicEntitiysSlice
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

// flattenBaseTopicEntitiys maps a Genesys Cloud *[]platformclientv2.Basetopicentitiy into a []interface{}
func flattenBaseTopicEntitiys(baseTopicEntitiys *[]platformclientv2.Basetopicentitiy) []interface{} {
	if len(*baseTopicEntitiys) == 0 {
		return nil
	}

	var baseTopicEntitiyList []interface{}
	for _, baseTopicEntitiy := range *baseTopicEntitiys {
		baseTopicEntitiyMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(baseTopicEntitiyMap, "name", baseTopicEntitiy.Name)

		baseTopicEntitiyList = append(baseTopicEntitiyList, baseTopicEntitiyMap)
	}

	return baseTopicEntitiyList
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
