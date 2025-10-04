package speechandtextanalytics_settings

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
	"log"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"

	"terraform-provider-genesyscloud/genesyscloud/consistency_checker"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"terraform-provider-genesyscloud/genesyscloud/util/constants"
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

/*
The resource_genesyscloud_speechandtextanalytics_settings.go contains all of the methods that perform the core logic for a resource.
*/

// createSpeechandtextanalyticsSettings is used by the speechandtextanalytics_settings resource to create Genesys cloud speechandtextanalytics settings
func createSpeechandtextanalyticsSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// readSpeechandtextanalyticsSettings is used by the speechandtextanalytics_settings resource to read an speechandtextanalytics settings from genesys cloud
func readSpeechandtextanalyticsSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsSettingsProxy(sdkConfig)
	cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceSpeechandtextanalyticsSettings(), constants.ConsistencyChecks(), resourceName)

	log.Printf("Reading speechandtextanalytics settings %s", d.Id())

	return util.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		speechTextAnalyticsSettingsResponse, resp, getErr := proxy.getSpeechandtextanalyticsSettingsById(ctx, d.Id())
		if getErr != nil {
			if util.IsStatus404(resp) {
				return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read speechandtextanalytics settings %s: %s", d.Id(), getErr), resp))
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read speechandtextanalytics settings %s: %s", d.Id(), getErr), resp))
		}

		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "default_program", speechTextAnalyticsSettingsResponse.DefaultProgram, flattenAddressableEntityRef)
		resourcedata.SetNillableValue(d, "expected_dialects", speechTextAnalyticsSettingsResponse.ExpectedDialects)
		resourcedata.SetNillableValue(d, "text_analytics_enabled", speechTextAnalyticsSettingsResponse.TextAnalyticsEnabled)
		resourcedata.SetNillableValue(d, "agent_empathy_enabled", speechTextAnalyticsSettingsResponse.AgentEmpathyEnabled)

		log.Printf("Read speechandtextanalytics settings %s %s", d.Id(), *speechTextAnalyticsSettingsResponse.Name)
		return cc.CheckState(d)
	})
}

// updateSpeechandtextanalyticsSettings is used by the speechandtextanalytics_settings resource to update an speechandtextanalytics settings in Genesys Cloud
func updateSpeechandtextanalyticsSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsSettingsProxy(sdkConfig)

	speechandtextanalyticsSettings := getSpeechandtextanalyticsSettingsFromResourceData(d)

	log.Printf("Updating speechandtextanalytics settings %s", *speechandtextanalyticsSettings.Name)
	speechTextAnalyticsSettingsResponse, resp, err := proxy.updateSpeechandtextanalyticsSettings(ctx, d.Id(), &speechandtextanalyticsSettings)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to update speechandtextanalytics settings %s: %s", d.Id(), err), resp)
	}

	log.Printf("Updated speechandtextanalytics settings %s", *speechTextAnalyticsSettingsResponse.Id)
	return readSpeechandtextanalyticsSettings(ctx, d, meta)
}

// deleteSpeechandtextanalyticsSettings is used by the speechandtextanalytics_settings resource to delete an speechandtextanalytics settings from Genesys cloud
func deleteSpeechandtextanalyticsSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
