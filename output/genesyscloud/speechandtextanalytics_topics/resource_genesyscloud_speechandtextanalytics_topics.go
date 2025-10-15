package speechandtextanalytics_topics

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
The resource_genesyscloud_speechandtextanalytics_topics.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthSpeechandtextanalyticsTopics retrieves all of the speechandtextanalytics topics via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthSpeechandtextanalyticsTopicss(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	proxy := newSpeechandtextanalyticsTopicsProxy(clientConfig)
	resources := make(resourceExporter.ResourceIDMetaMap)

	topics, resp, err := proxy.getAllSpeechandtextanalyticsTopics(ctx)
	if err != nil {
		return nil, util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to get speechandtextanalytics topics: %v", err), resp)
	}

	for _, topic := range *topics {
		resources[*topic.Id] = &resourceExporter.ResourceMeta{BlockLabel: *topic.Name}
	}

	return resources, nil
}

// createSpeechandtextanalyticsTopics is used by the speechandtextanalytics_topics resource to create Genesys cloud speechandtextanalytics topics
func createSpeechandtextanalyticsTopics(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsTopicsProxy(sdkConfig)

	speechandtextanalyticsTopics := getSpeechandtextanalyticsTopicsFromResourceData(d)

	log.Printf("Creating speechandtextanalytics topics %s", *speechandtextanalyticsTopics.Name)
	topic, resp, err := proxy.createSpeechandtextanalyticsTopics(ctx, &speechandtextanalyticsTopics)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to create speechandtextanalytics topics: %s", err), resp)
	}

	d.SetId(*topic.Id)
	log.Printf("Created speechandtextanalytics topics %s", *topic.Id)
	return readSpeechandtextanalyticsTopics(ctx, d, meta)
}

// readSpeechandtextanalyticsTopics is used by the speechandtextanalytics_topics resource to read an speechandtextanalytics topics from genesys cloud
func readSpeechandtextanalyticsTopics(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsTopicsProxy(sdkConfig)
	cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceSpeechandtextanalyticsTopics(), constants.ConsistencyChecks(), resourceName)

	log.Printf("Reading speechandtextanalytics topics %s", d.Id())

	return util.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		topic, resp, getErr := proxy.getSpeechandtextanalyticsTopicsById(ctx, d.Id())
		if getErr != nil {
			if util.IsStatus404(resp) {
				return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read speechandtextanalytics topics %s: %s", d.Id(), getErr), resp))
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read speechandtextanalytics topics %s: %s", d.Id(), getErr), resp))
		}

		resourcedata.SetNillableValue(d, "name", topic.Name)
		resourcedata.SetNillableValue(d, "description", topic.Description)
		resourcedata.SetNillableValue(d, "published", topic.Published)
		resourcedata.SetNillableValue(d, "strictness", topic.Strictness)
		resourcedata.SetNillableValue(d, "matching_type", topic.MatchingType)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "programs", topic.Programs, flattenBaseProgramEntitys)
		resourcedata.SetNillableValue(d, "tags", topic.Tags)
		resourcedata.SetNillableValue(d, "dialect", topic.Dialect)
		resourcedata.SetNillableValue(d, "participants", topic.Participants)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "phrases", topic.Phrases, flattenPhrases)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "published_by", topic.PublishedBy, flattenAddressableEntityRef)
		resourcedata.SetNillableValue(d, "date_published", topic.DatePublished)

		log.Printf("Read speechandtextanalytics topics %s %s", d.Id(), *topic.Name)
		return cc.CheckState(d)
	})
}

// updateSpeechandtextanalyticsTopics is used by the speechandtextanalytics_topics resource to update an speechandtextanalytics topics in Genesys Cloud
func updateSpeechandtextanalyticsTopics(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsTopicsProxy(sdkConfig)

	speechandtextanalyticsTopics := getSpeechandtextanalyticsTopicsFromResourceData(d)

	log.Printf("Updating speechandtextanalytics topics %s", *speechandtextanalyticsTopics.Name)
	topic, resp, err := proxy.updateSpeechandtextanalyticsTopics(ctx, d.Id(), &speechandtextanalyticsTopics)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to update speechandtextanalytics topics %s: %s", d.Id(), err), resp)
	}

	log.Printf("Updated speechandtextanalytics topics %s", *topic.Id)
	return readSpeechandtextanalyticsTopics(ctx, d, meta)
}

// deleteSpeechandtextanalyticsTopics is used by the speechandtextanalytics_topics resource to delete an speechandtextanalytics topics from Genesys cloud
func deleteSpeechandtextanalyticsTopics(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsTopicsProxy(sdkConfig)

	resp, err := proxy.deleteSpeechandtextanalyticsTopics(ctx, d.Id())
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to delete speechandtextanalytics topics %s: %s", d.Id(), err), resp)
	}

	return util.WithRetries(ctx, 180*time.Second, func() *retry.RetryError {
		_, resp, err := proxy.getSpeechandtextanalyticsTopicsById(ctx, d.Id())

		if err != nil {
			if util.IsStatus404(resp) {
				log.Printf("Deleted speechandtextanalytics topics %s", d.Id())
				return nil
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Error deleting speechandtextanalytics topics %s: %s", d.Id(), err), resp))
		}

		return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("speechandtextanalytics topics %s still exists", d.Id()), resp))
	})
}
