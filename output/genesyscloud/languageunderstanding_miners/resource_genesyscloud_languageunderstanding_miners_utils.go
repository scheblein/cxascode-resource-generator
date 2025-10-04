package languageunderstanding_miners

import (
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
)

/*
The resource_genesyscloud_languageunderstanding_miners_utils.go file contains various helper methods to marshal
and unmarshal data into formats consumable by Terraform and/or Genesys Cloud.
*/

// getLanguageunderstandingMinersFromResourceData maps data from schema ResourceData object to a platformclientv2.Miner
func getLanguageunderstandingMinersFromResourceData(d *schema.ResourceData) platformclientv2.Miner {
	return platformclientv2.Miner{
		Name:                        platformclientv2.String(d.Get("name").(string)),
		Language:                    platformclientv2.String(d.Get("language").(string)),
		MinerType:                   platformclientv2.String(d.Get("miner_type").(string)),
		Seeding:                     platformclientv2.Bool(d.Get("seeding").(bool)),
		Status:                      platformclientv2.String(d.Get("status").(string)),
		ConversationsDateRangeStart: platformclientv2.String(d.Get("conversations_date_range_start").(string)),
		ConversationsDateRangeEnd:   platformclientv2.String(d.Get("conversations_date_range_end").(string)),
		DateCompleted:               platformclientv2.String(d.Get("date_completed").(string)),
		Message:                     platformclientv2.String(d.Get("message").(string)),
		ErrorInfo:                   buildMinerErrorInfo(d.Get("error_info").([]interface{})),
		WarningInfo:                 buildMinerErrorInfo(d.Get("warning_info").([]interface{})),
		ConversationDataUploaded:    platformclientv2.Bool(d.Get("conversation_data_uploaded").(bool)),
		MediaType:                   platformclientv2.String(d.Get("media_type").(string)),
		ParticipantType:             platformclientv2.String(d.Get("participant_type").(string)),
		// TODO: Handle queue_ids property
		DateTriggered:             platformclientv2.String(d.Get("date_triggered").(string)),
		LatestDraftVersion:        buildDraft(d.Get("latest_draft_version").([]interface{})),
		ConversationsFetchedCount: platformclientv2.Int(d.Get("conversations_fetched_count").(int)),
		ConversationsValidCount:   platformclientv2.Int(d.Get("conversations_valid_count").(int)),
		GetminedItemCount:         platformclientv2.Int(d.Get("getmined_item_count").(int)),
	}
}

// buildMinerErrorInfos maps an []interface{} into a Genesys Cloud *[]platformclientv2.Minererrorinfo
func buildMinerErrorInfos(minerErrorInfos []interface{}) *[]platformclientv2.Minererrorinfo {
	minerErrorInfosSlice := make([]platformclientv2.Minererrorinfo, 0)
	for _, minerErrorInfo := range minerErrorInfos {
		var sdkMinerErrorInfo platformclientv2.Minererrorinfo
		minerErrorInfosMap, ok := minerErrorInfo.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkMinerErrorInfo.Message, minerErrorInfosMap, "message")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMinerErrorInfo.Code, minerErrorInfosMap, "code")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMinerErrorInfo.MessageWithParams, minerErrorInfosMap, "message_with_params")
		// TODO: Handle message_params property

		minerErrorInfosSlice = append(minerErrorInfosSlice, sdkMinerErrorInfo)
	}

	return &minerErrorInfosSlice
}

// buildDrafts maps an []interface{} into a Genesys Cloud *[]platformclientv2.Draft
func buildDrafts(drafts []interface{}) *[]platformclientv2.Draft {
	draftsSlice := make([]platformclientv2.Draft, 0)
	for _, draft := range drafts {
		var sdkDraft platformclientv2.Draft
		draftsMap, ok := draft.(map[string]interface{})
		if !ok {
			continue
		}

		draftsSlice = append(draftsSlice, sdkDraft)
	}

	return &draftsSlice
}

// buildMiners maps an []interface{} into a Genesys Cloud *[]platformclientv2.Miner
func buildMiners(miners []interface{}) *[]platformclientv2.Miner {
	minersSlice := make([]platformclientv2.Miner, 0)
	for _, miner := range miners {
		var sdkMiner platformclientv2.Miner
		minersMap, ok := miner.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.Name, minersMap, "name")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.Language, minersMap, "language")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.MinerType, minersMap, "miner_type")
		sdkMiner.Seeding = platformclientv2.Bool(minersMap["seeding"].(bool))
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.Status, minersMap, "status")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.ConversationsDateRangeStart, minersMap, "conversations_date_range_start")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.ConversationsDateRangeEnd, minersMap, "conversations_date_range_end")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.DateCompleted, minersMap, "date_completed")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.Message, minersMap, "message")
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkMiner.ErrorInfo, minersMap, "error_info", buildMinerErrorInfo)
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkMiner.WarningInfo, minersMap, "warning_info", buildMinerErrorInfo)
		sdkMiner.ConversationDataUploaded = platformclientv2.Bool(minersMap["conversation_data_uploaded"].(bool))
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.MediaType, minersMap, "media_type")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.ParticipantType, minersMap, "participant_type")
		resourcedata.BuildSDKStringArrayValueIfNotNil(&sdkMiner.QueueIds, minersMap, "queue_ids")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkMiner.DateTriggered, minersMap, "date_triggered")
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkMiner.LatestDraftVersion, minersMap, "latest_draft_version", buildDraft)
		sdkMiner.ConversationsFetchedCount = platformclientv2.Int(minersMap["conversations_fetched_count"].(int))
		sdkMiner.ConversationsValidCount = platformclientv2.Int(minersMap["conversations_valid_count"].(int))
		sdkMiner.GetminedItemCount = platformclientv2.Int(minersMap["getmined_item_count"].(int))

		minersSlice = append(minersSlice, sdkMiner)
	}

	return &minersSlice
}

// buildDraftIntentss maps an []interface{} into a Genesys Cloud *[]platformclientv2.Draftintents
func buildDraftIntentss(draftIntentss []interface{}) *[]platformclientv2.Draftintents {
	draftIntentssSlice := make([]platformclientv2.Draftintents, 0)
	for _, draftIntents := range draftIntentss {
		var sdkDraftIntents platformclientv2.Draftintents
		draftIntentssMap, ok := draftIntents.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkDraftIntents.Name, draftIntentssMap, "name")
		resourcedata.BuildSDKStringArrayValueIfNotNil(&sdkDraftIntents.Utterances, draftIntentssMap, "utterances")

		draftIntentssSlice = append(draftIntentssSlice, sdkDraftIntents)
	}

	return &draftIntentssSlice
}

// buildDraftTopicss maps an []interface{} into a Genesys Cloud *[]platformclientv2.Drafttopics
func buildDraftTopicss(draftTopicss []interface{}) *[]platformclientv2.Drafttopics {
	draftTopicssSlice := make([]platformclientv2.Drafttopics, 0)
	for _, draftTopics := range draftTopicss {
		var sdkDraftTopics platformclientv2.Drafttopics
		draftTopicssMap, ok := draftTopics.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkDraftTopics.Name, draftTopicssMap, "name")
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkDraftTopics.Miner, draftTopicssMap, "miner", buildMiner)
		sdkDraftTopics.ConversationCount = platformclientv2.Int(draftTopicssMap["conversation_count"].(int))
		// TODO: Handle conversation_percent property
		sdkDraftTopics.UtteranceCount = platformclientv2.Int(draftTopicssMap["utterance_count"].(int))
		sdkDraftTopics.PhraseCount = platformclientv2.Int(draftTopicssMap["phrase_count"].(int))
		resourcedata.BuildSDKStringArrayValueIfNotNil(&sdkDraftTopics.Phrases, draftTopicssMap, "phrases")

		draftTopicssSlice = append(draftTopicssSlice, sdkDraftTopics)
	}

	return &draftTopicssSlice
}

// buildDrafts maps an []interface{} into a Genesys Cloud *[]platformclientv2.Draft
func buildDrafts(drafts []interface{}) *[]platformclientv2.Draft {
	draftsSlice := make([]platformclientv2.Draft, 0)
	for _, draft := range drafts {
		var sdkDraft platformclientv2.Draft
		draftsMap, ok := draft.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkDraft.Name, draftsMap, "name")
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkDraft.Miner, draftsMap, "miner", buildMiner)
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkDraft.Intents, draftsMap, "intents", buildDraftIntentss)
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkDraft.Topics, draftsMap, "topics", buildDraftTopicss)

		draftsSlice = append(draftsSlice, sdkDraft)
	}

	return &draftsSlice
}

// flattenMinerErrorInfos maps a Genesys Cloud *[]platformclientv2.Minererrorinfo into a []interface{}
func flattenMinerErrorInfos(minerErrorInfos *[]platformclientv2.Minererrorinfo) []interface{} {
	if len(*minerErrorInfos) == 0 {
		return nil
	}

	var minerErrorInfoList []interface{}
	for _, minerErrorInfo := range *minerErrorInfos {
		minerErrorInfoMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(minerErrorInfoMap, "message", minerErrorInfo.Message)
		resourcedata.SetMapValueIfNotNil(minerErrorInfoMap, "code", minerErrorInfo.Code)
		resourcedata.SetMapValueIfNotNil(minerErrorInfoMap, "message_with_params", minerErrorInfo.MessageWithParams)
		// TODO: Handle message_params property

		minerErrorInfoList = append(minerErrorInfoList, minerErrorInfoMap)
	}

	return minerErrorInfoList
}

// flattenDrafts maps a Genesys Cloud *[]platformclientv2.Draft into a []interface{}
func flattenDrafts(drafts *[]platformclientv2.Draft) []interface{} {
	if len(*drafts) == 0 {
		return nil
	}

	var draftList []interface{}
	for _, draft := range *drafts {
		draftMap := make(map[string]interface{})

		draftList = append(draftList, draftMap)
	}

	return draftList
}

// flattenMiners maps a Genesys Cloud *[]platformclientv2.Miner into a []interface{}
func flattenMiners(miners *[]platformclientv2.Miner) []interface{} {
	if len(*miners) == 0 {
		return nil
	}

	var minerList []interface{}
	for _, miner := range *miners {
		minerMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(minerMap, "name", miner.Name)
		resourcedata.SetMapValueIfNotNil(minerMap, "language", miner.Language)
		resourcedata.SetMapValueIfNotNil(minerMap, "miner_type", miner.MinerType)
		resourcedata.SetMapValueIfNotNil(minerMap, "seeding", miner.Seeding)
		resourcedata.SetMapValueIfNotNil(minerMap, "status", miner.Status)
		resourcedata.SetMapValueIfNotNil(minerMap, "conversations_date_range_start", miner.ConversationsDateRangeStart)
		resourcedata.SetMapValueIfNotNil(minerMap, "conversations_date_range_end", miner.ConversationsDateRangeEnd)
		resourcedata.SetMapValueIfNotNil(minerMap, "date_completed", miner.DateCompleted)
		resourcedata.SetMapValueIfNotNil(minerMap, "message", miner.Message)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(minerMap, "error_info", miner.ErrorInfo, flattenMinerErrorInfo)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(minerMap, "warning_info", miner.WarningInfo, flattenMinerErrorInfo)
		resourcedata.SetMapValueIfNotNil(minerMap, "conversation_data_uploaded", miner.ConversationDataUploaded)
		resourcedata.SetMapValueIfNotNil(minerMap, "media_type", miner.MediaType)
		resourcedata.SetMapValueIfNotNil(minerMap, "participant_type", miner.ParticipantType)
		resourcedata.SetMapStringArrayValueIfNotNil(minerMap, "queue_ids", miner.QueueIds)
		resourcedata.SetMapValueIfNotNil(minerMap, "date_triggered", miner.DateTriggered)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(minerMap, "latest_draft_version", miner.LatestDraftVersion, flattenDraft)
		resourcedata.SetMapValueIfNotNil(minerMap, "conversations_fetched_count", miner.ConversationsFetchedCount)
		resourcedata.SetMapValueIfNotNil(minerMap, "conversations_valid_count", miner.ConversationsValidCount)
		resourcedata.SetMapValueIfNotNil(minerMap, "getmined_item_count", miner.GetminedItemCount)

		minerList = append(minerList, minerMap)
	}

	return minerList
}

// flattenDraftIntentss maps a Genesys Cloud *[]platformclientv2.Draftintents into a []interface{}
func flattenDraftIntentss(draftIntentss *[]platformclientv2.Draftintents) []interface{} {
	if len(*draftIntentss) == 0 {
		return nil
	}

	var draftIntentsList []interface{}
	for _, draftIntents := range *draftIntentss {
		draftIntentsMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(draftIntentsMap, "name", draftIntents.Name)
		resourcedata.SetMapStringArrayValueIfNotNil(draftIntentsMap, "utterances", draftIntents.Utterances)

		draftIntentsList = append(draftIntentsList, draftIntentsMap)
	}

	return draftIntentsList
}

// flattenDraftTopicss maps a Genesys Cloud *[]platformclientv2.Drafttopics into a []interface{}
func flattenDraftTopicss(draftTopicss *[]platformclientv2.Drafttopics) []interface{} {
	if len(*draftTopicss) == 0 {
		return nil
	}

	var draftTopicsList []interface{}
	for _, draftTopics := range *draftTopicss {
		draftTopicsMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(draftTopicsMap, "name", draftTopics.Name)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(draftTopicsMap, "miner", draftTopics.Miner, flattenMiner)
		resourcedata.SetMapValueIfNotNil(draftTopicsMap, "conversation_count", draftTopics.ConversationCount)
		resourcedata.SetMapValueIfNotNil(draftTopicsMap, "conversation_percent", draftTopics.ConversationPercent)
		resourcedata.SetMapValueIfNotNil(draftTopicsMap, "utterance_count", draftTopics.UtteranceCount)
		resourcedata.SetMapValueIfNotNil(draftTopicsMap, "phrase_count", draftTopics.PhraseCount)
		resourcedata.SetMapStringArrayValueIfNotNil(draftTopicsMap, "phrases", draftTopics.Phrases)

		draftTopicsList = append(draftTopicsList, draftTopicsMap)
	}

	return draftTopicsList
}

// flattenDrafts maps a Genesys Cloud *[]platformclientv2.Draft into a []interface{}
func flattenDrafts(drafts *[]platformclientv2.Draft) []interface{} {
	if len(*drafts) == 0 {
		return nil
	}

	var draftList []interface{}
	for _, draft := range *drafts {
		draftMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(draftMap, "name", draft.Name)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(draftMap, "miner", draft.Miner, flattenMiner)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(draftMap, "intents", draft.Intents, flattenDraftIntentss)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(draftMap, "topics", draft.Topics, flattenDraftTopicss)

		draftList = append(draftList, draftMap)
	}

	return draftList
}
