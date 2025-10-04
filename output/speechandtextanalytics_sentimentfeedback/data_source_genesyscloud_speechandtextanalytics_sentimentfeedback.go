package speechandtextanalytics_sentimentfeedback

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
)

/*
   The data_source_genesyscloud_speechandtextanalytics_sentimentfeedback.go contains the data source implementation
   for the resource.
*/

// dataSourceSpeechandtextanalyticsSentimentfeedbackRead retrieves by name the id in question
func dataSourceSpeechandtextanalyticsSentimentfeedbackRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := newSpeechandtextanalyticsSentimentfeedbackProxy(sdkConfig)

	name := d.Get("name").(string)

	return util.WithRetries(ctx, 15*time.Second, func() *retry.RetryError {
		sentimentFeedbackId, resp, retryable, err := proxy.getSpeechandtextanalyticsSentimentfeedbackIdByName(ctx, name)

		if err != nil && !retryable {
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Error searching speechandtextanalytics sentimentfeedback %s | error: %s", name, err), resp))
		}

		if retryable {
			return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("No speechandtextanalytics sentimentfeedback found with name %s", name), resp))
		}

		d.SetId(sentimentFeedbackId)
		return nil
	})
}
