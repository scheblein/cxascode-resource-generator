package speechandtextanalytics_dictionaryfeedback

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
   The data_source_genesyscloud_speechandtextanalytics_dictionaryfeedback.go contains the data source implementation
   for the resource.
*/

// dataSourceSpeechandtextanalyticsDictionaryfeedbackRead retrieves by name the id in question
func dataSourceSpeechandtextanalyticsDictionaryfeedbackRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := newSpeechandtextanalyticsDictionaryfeedbackProxy(sdkConfig)

	name := d.Get("name").(string)

	return util.WithRetries(ctx, 15*time.Second, func() *retry.RetryError {
		dictionaryFeedbackId, resp, retryable, err := proxy.getSpeechandtextanalyticsDictionaryfeedbackIdByName(ctx, name)

		if err != nil && !retryable {
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Error searching speechandtextanalytics dictionaryfeedback %s | error: %s", name, err), resp))
		}

		if retryable {
			return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("No speechandtextanalytics dictionaryfeedback found with name %s", name), resp))
		}

		d.SetId(dictionaryFeedbackId)
		return nil
	})
}
