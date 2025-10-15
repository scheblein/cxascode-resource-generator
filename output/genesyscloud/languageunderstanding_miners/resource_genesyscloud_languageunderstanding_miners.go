package languageunderstanding_miners

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
	"log"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	"time"

	"terraform-provider-genesyscloud/genesyscloud/consistency_checker"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"terraform-provider-genesyscloud/genesyscloud/util/constants"
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

/*
The resource_genesyscloud_languageunderstanding_miners.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthLanguageunderstandingMiners retrieves all of the languageunderstanding miners via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthLanguageunderstandingMinerss(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	proxy := newLanguageunderstandingMinersProxy(clientConfig)
	resources := make(resourceExporter.ResourceIDMetaMap)

	miners, resp, err := proxy.getAllLanguageunderstandingMiners(ctx)
	if err != nil {
		return nil, util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to get languageunderstanding miners: %v", err), resp)
	}

	for _, miner := range *miners {
		resources[*miner.Id] = &resourceExporter.ResourceMeta{BlockLabel: *miner.Name}
	}

	return resources, nil
}

// createLanguageunderstandingMiners is used by the languageunderstanding_miners resource to create Genesys cloud languageunderstanding miners
func createLanguageunderstandingMiners(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getLanguageunderstandingMinersProxy(sdkConfig)

	languageunderstandingMiners := getLanguageunderstandingMinersFromResourceData(d)

	log.Printf("Creating languageunderstanding miners %s", *languageunderstandingMiners.Name)
	miner, resp, err := proxy.createLanguageunderstandingMiners(ctx, &languageunderstandingMiners)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to create languageunderstanding miners: %s", err), resp)
	}

	d.SetId(*miner.Id)
	log.Printf("Created languageunderstanding miners %s", *miner.Id)
	return readLanguageunderstandingMiners(ctx, d, meta)
}

// readLanguageunderstandingMiners is used by the languageunderstanding_miners resource to read an languageunderstanding miners from genesys cloud
func readLanguageunderstandingMiners(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getLanguageunderstandingMinersProxy(sdkConfig)
	cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceLanguageunderstandingMiners(), constants.ConsistencyChecks(), resourceName)

	log.Printf("Reading languageunderstanding miners %s", d.Id())

	return util.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		miner, resp, getErr := proxy.getLanguageunderstandingMinersById(ctx, d.Id())
		if getErr != nil {
			if util.IsStatus404(resp) {
				return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read languageunderstanding miners %s: %s", d.Id(), getErr), resp))
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read languageunderstanding miners %s: %s", d.Id(), getErr), resp))
		}

		resourcedata.SetNillableValue(d, "name", miner.Name)
		resourcedata.SetNillableValue(d, "language", miner.Language)
		resourcedata.SetNillableValue(d, "miner_type", miner.MinerType)
		resourcedata.SetNillableValue(d, "seeding", miner.Seeding)
		resourcedata.SetNillableValue(d, "status", miner.Status)
		resourcedata.SetNillableValue(d, "conversations_date_range_start", miner.ConversationsDateRangeStart)
		resourcedata.SetNillableValue(d, "conversations_date_range_end", miner.ConversationsDateRangeEnd)
		resourcedata.SetNillableValue(d, "date_completed", miner.DateCompleted)
		resourcedata.SetNillableValue(d, "message", miner.Message)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "error_info", miner.ErrorInfo, flattenMinerErrorInfo)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "warning_info", miner.WarningInfo, flattenMinerErrorInfo)
		resourcedata.SetNillableValue(d, "conversation_data_uploaded", miner.ConversationDataUploaded)
		resourcedata.SetNillableValue(d, "media_type", miner.MediaType)
		resourcedata.SetNillableValue(d, "participant_type", miner.ParticipantType)
		resourcedata.SetNillableValue(d, "queue_ids", miner.QueueIds)
		resourcedata.SetNillableValue(d, "date_triggered", miner.DateTriggered)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "latest_draft_version", miner.LatestDraftVersion, flattenDraft)
		resourcedata.SetNillableValue(d, "conversations_fetched_count", miner.ConversationsFetchedCount)
		resourcedata.SetNillableValue(d, "conversations_valid_count", miner.ConversationsValidCount)
		resourcedata.SetNillableValue(d, "getmined_item_count", miner.GetminedItemCount)

		log.Printf("Read languageunderstanding miners %s %s", d.Id(), *miner.Name)
		return cc.CheckState(d)
	})
}

// updateLanguageunderstandingMiners is used by the languageunderstanding_miners resource to update an languageunderstanding miners in Genesys Cloud
func updateLanguageunderstandingMiners(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// deleteLanguageunderstandingMiners is used by the languageunderstanding_miners resource to delete an languageunderstanding miners from Genesys cloud
func deleteLanguageunderstandingMiners(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getLanguageunderstandingMinersProxy(sdkConfig)

	resp, err := proxy.deleteLanguageunderstandingMiners(ctx, d.Id())
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to delete languageunderstanding miners %s: %s", d.Id(), err), resp)
	}

	return util.WithRetries(ctx, 180*time.Second, func() *retry.RetryError {
		_, resp, err := proxy.getLanguageunderstandingMinersById(ctx, d.Id())

		if err != nil {
			if util.IsStatus404(resp) {
				log.Printf("Deleted languageunderstanding miners %s", d.Id())
				return nil
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Error deleting languageunderstanding miners %s: %s", d.Id(), err), resp))
		}

		return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("languageunderstanding miners %s still exists", d.Id()), resp))
	})
}
