package speechandtextanalytics_sentimentfeedback

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
The resource_genesyscloud_speechandtextanalytics_sentimentfeedback.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthSpeechandtextanalyticsSentimentfeedback retrieves all of the speechandtextanalytics sentimentfeedback via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthSpeechandtextanalyticsSentimentfeedbacks(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	proxy := newSpeechandtextanalyticsSentimentfeedbackProxy(clientConfig)
	resources := make(resourceExporter.ResourceIDMetaMap)

	sentimentFeedbacks, resp, err := proxy.getAllSpeechandtextanalyticsSentimentfeedback(ctx)
	if err != nil {
		return nil, util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to get speechandtextanalytics sentimentfeedback: %v", err), resp)
	}

	for _, sentimentFeedback := range *sentimentFeedbacks {
		resources[*sentimentFeedback.Id] = &resourceExporter.ResourceMeta{BlockLabel: *sentimentFeedback.Name}
	}

	return resources, nil
}

// createSpeechandtextanalyticsSentimentfeedback is used by the speechandtextanalytics_sentimentfeedback resource to create Genesys cloud speechandtextanalytics sentimentfeedback
func createSpeechandtextanalyticsSentimentfeedback(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsSentimentfeedbackProxy(sdkConfig)

	speechandtextanalyticsSentimentfeedback := getSpeechandtextanalyticsSentimentfeedbackFromResourceData(d)

	log.Printf("Creating speechandtextanalytics sentimentfeedback %s", *speechandtextanalyticsSentimentfeedback.Name)
	sentimentFeedback, resp, err := proxy.createSpeechandtextanalyticsSentimentfeedback(ctx, &speechandtextanalyticsSentimentfeedback)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to create speechandtextanalytics sentimentfeedback: %s", err), resp)
	}

	d.SetId(*sentimentFeedback.Id)
	log.Printf("Created speechandtextanalytics sentimentfeedback %s", *sentimentFeedback.Id)
	return readSpeechandtextanalyticsSentimentfeedback(ctx, d, meta)
}

// readSpeechandtextanalyticsSentimentfeedback is used by the speechandtextanalytics_sentimentfeedback resource to read an speechandtextanalytics sentimentfeedback from genesys cloud
func readSpeechandtextanalyticsSentimentfeedback(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// updateSpeechandtextanalyticsSentimentfeedback is used by the speechandtextanalytics_sentimentfeedback resource to update an speechandtextanalytics sentimentfeedback in Genesys Cloud
func updateSpeechandtextanalyticsSentimentfeedback(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// deleteSpeechandtextanalyticsSentimentfeedback is used by the speechandtextanalytics_sentimentfeedback resource to delete an speechandtextanalytics sentimentfeedback from Genesys cloud
func deleteSpeechandtextanalyticsSentimentfeedback(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsSentimentfeedbackProxy(sdkConfig)

	resp, err := proxy.deleteSpeechandtextanalyticsSentimentfeedback(ctx, d.Id())
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to delete speechandtextanalytics sentimentfeedback %s: %s", d.Id(), err), resp)
	}

	return util.WithRetries(ctx, 180*time.Second, func() *retry.RetryError {
		_, resp, err := proxy.getSpeechandtextanalyticsSentimentfeedbackById(ctx, d.Id())

		if err != nil {
			if util.IsStatus404(resp) {
				log.Printf("Deleted speechandtextanalytics sentimentfeedback %s", d.Id())
				return nil
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Error deleting speechandtextanalytics sentimentfeedback %s: %s", d.Id(), err), resp))
		}

		return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("speechandtextanalytics sentimentfeedback %s still exists", d.Id()), resp))
	})
}

// getSpeechandtextanalyticsSentimentfeedbackFromResourceData maps data from schema ResourceData object to a platformclientv2.Sentimentfeedback
func getSpeechandtextanalyticsSentimentfeedbackFromResourceData(d *schema.ResourceData) platformclientv2.Sentimentfeedback {
	return platformclientv2.Sentimentfeedback{
		Phrase:        platformclientv2.String(d.Get("phrase").(string)),
		Dialect:       platformclientv2.String(d.Get("dialect").(string)),
		FeedbackValue: platformclientv2.String(d.Get("feedback_value").(string)),
	}
}
