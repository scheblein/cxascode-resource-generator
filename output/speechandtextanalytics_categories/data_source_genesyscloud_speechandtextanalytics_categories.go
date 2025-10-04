package speechandtextanalytics_categories

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
   The data_source_genesyscloud_speechandtextanalytics_categories.go contains the data source implementation
   for the resource.
*/

// dataSourceSpeechandtextanalyticsCategoriesRead retrieves by name the id in question
func dataSourceSpeechandtextanalyticsCategoriesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := newSpeechandtextanalyticsCategoriesProxy(sdkConfig)

	name := d.Get("name").(string)

	return util.WithRetries(ctx, 15*time.Second, func() *retry.RetryError {
		staCategoryId, resp, retryable, err := proxy.getSpeechandtextanalyticsCategoriesIdByName(ctx, name)

		if err != nil && !retryable {
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Error searching speechandtextanalytics categories %s | error: %s", name, err), resp))
		}

		if retryable {
			return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("No speechandtextanalytics categories found with name %s", name), resp))
		}

		d.SetId(staCategoryId)
		return nil
	})
}
