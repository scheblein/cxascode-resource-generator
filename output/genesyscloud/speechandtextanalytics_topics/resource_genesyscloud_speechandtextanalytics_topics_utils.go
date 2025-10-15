package speechandtextanalytics_topics

import (
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
)

/*
The resource_genesyscloud_speechandtextanalytics_topics_utils.go file contains various helper methods to marshal
and unmarshal data into formats consumable by Terraform and/or Genesys Cloud.
*/

// getSpeechandtextanalyticsTopicsFromResourceData maps data from schema ResourceData object to a platformclientv2.Topic
func getSpeechandtextanalyticsTopicsFromResourceData(d *schema.ResourceData) platformclientv2.Topic {
	return platformclientv2.Topic{
		Name:         platformclientv2.String(d.Get("name").(string)),
		Description:  platformclientv2.String(d.Get("description").(string)),
		Published:    platformclientv2.Bool(d.Get("published").(bool)),
		Strictness:   platformclientv2.String(d.Get("strictness").(string)),
		MatchingType: platformclientv2.String(d.Get("matching_type").(string)),
		Programs:     buildBaseProgramEntitys(d.Get("programs").([]interface{})),
		// TODO: Handle tags property
		Dialect:       platformclientv2.String(d.Get("dialect").(string)),
		Participants:  platformclientv2.String(d.Get("participants").(string)),
		Phrases:       buildPhrases(d.Get("phrases").([]interface{})),
		PublishedBy:   buildAddressableEntityRef(d.Get("published_by").([]interface{})),
		DatePublished: platformclientv2.String(d.Get("date_published").(string)),
	}
}

// buildBaseProgramEntitys maps an []interface{} into a Genesys Cloud *[]platformclientv2.Baseprogramentity
func buildBaseProgramEntitys(baseProgramEntitys []interface{}) *[]platformclientv2.Baseprogramentity {
	baseProgramEntitysSlice := make([]platformclientv2.Baseprogramentity, 0)
	for _, baseProgramEntity := range baseProgramEntitys {
		var sdkBaseProgramEntity platformclientv2.Baseprogramentity
		baseProgramEntitysMap, ok := baseProgramEntity.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkBaseProgramEntity.Name, baseProgramEntitysMap, "name")

		baseProgramEntitysSlice = append(baseProgramEntitysSlice, sdkBaseProgramEntity)
	}

	return &baseProgramEntitysSlice
}

// buildPhrases maps an []interface{} into a Genesys Cloud *[]platformclientv2.Phrase
func buildPhrases(phrases []interface{}) *[]platformclientv2.Phrase {
	phrasesSlice := make([]platformclientv2.Phrase, 0)
	for _, phrase := range phrases {
		var sdkPhrase platformclientv2.Phrase
		phrasesMap, ok := phrase.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkPhrase.Text, phrasesMap, "text")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkPhrase.Strictness, phrasesMap, "strictness")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkPhrase.Sentiment, phrasesMap, "sentiment")

		phrasesSlice = append(phrasesSlice, sdkPhrase)
	}

	return &phrasesSlice
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

// flattenBaseProgramEntitys maps a Genesys Cloud *[]platformclientv2.Baseprogramentity into a []interface{}
func flattenBaseProgramEntitys(baseProgramEntitys *[]platformclientv2.Baseprogramentity) []interface{} {
	if len(*baseProgramEntitys) == 0 {
		return nil
	}

	var baseProgramEntityList []interface{}
	for _, baseProgramEntity := range *baseProgramEntitys {
		baseProgramEntityMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(baseProgramEntityMap, "name", baseProgramEntity.Name)

		baseProgramEntityList = append(baseProgramEntityList, baseProgramEntityMap)
	}

	return baseProgramEntityList
}

// flattenPhrases maps a Genesys Cloud *[]platformclientv2.Phrase into a []interface{}
func flattenPhrases(phrases *[]platformclientv2.Phrase) []interface{} {
	if len(*phrases) == 0 {
		return nil
	}

	var phraseList []interface{}
	for _, phrase := range *phrases {
		phraseMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(phraseMap, "text", phrase.Text)
		resourcedata.SetMapValueIfNotNil(phraseMap, "strictness", phrase.Strictness)
		resourcedata.SetMapValueIfNotNil(phraseMap, "sentiment", phrase.Sentiment)

		phraseList = append(phraseList, phraseMap)
	}

	return phraseList
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
