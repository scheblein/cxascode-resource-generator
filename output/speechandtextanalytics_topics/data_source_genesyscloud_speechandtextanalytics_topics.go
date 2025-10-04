package speechandtextanalytics_topics

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
   The data_source_genesyscloud_speechandtextanalytics_topics.go contains the data source implementation
   for the resource.
*/

// dataSourceSpeechandtextanalyticsTopicsRead retrieves by name the id in question
func dataSourceSpeechandtextanalyticsTopicsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := newSpeechandtextanalyticsTopicsProxy(sdkConfig)

	name := d.Get("name").(string)

	return util.WithRetries(ctx, 15*time.Second, func() *retry.RetryError {
		topicId, resp, retryable, err := proxy.getSpeechandtextanalyticsTopicsIdByName(ctx, name)

		if err != nil && !retryable {
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Error searching speechandtextanalytics topics %s | error: %s", name, err), resp))
		}

		if retryable {
			return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("No speechandtextanalytics topics found with name %s", name), resp))
		}

		d.SetId(topicId)
		return nil
	})
}
