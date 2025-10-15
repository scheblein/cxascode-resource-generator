package recording_settings

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
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
The resource_genesyscloud_recording_settings.go contains all of the methods that perform the core logic for a resource.
*/

// createRecordingSettings is used by the recording_settings resource to create Genesys cloud recording settings
func createRecordingSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// readRecordingSettings is used by the recording_settings resource to read an recording settings from genesys cloud
func readRecordingSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getRecordingSettingsProxy(sdkConfig)
	cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceRecordingSettings(), constants.ConsistencyChecks(), resourceName)

	log.Printf("Reading recording settings %s", d.Id())

	return util.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		recordingSettings, resp, getErr := proxy.getRecordingSettingsById(ctx, d.Id())
		if getErr != nil {
			if util.IsStatus404(resp) {
				return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read recording settings %s: %s", d.Id(), getErr), resp))
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read recording settings %s: %s", d.Id(), getErr), resp))
		}

		resourcedata.SetNillableValue(d, "max_simultaneous_streams", recordingSettings.MaxSimultaneousStreams)
		resourcedata.SetNillableValue(d, "max_configurable_screen_recording_streams", recordingSettings.MaxConfigurableScreenRecordingStreams)
		resourcedata.SetNillableValue(d, "regional_recording_storage_enabled", recordingSettings.RegionalRecordingStorageEnabled)
		resourcedata.SetNillableValue(d, "recording_playback_url_ttl", recordingSettings.RecordingPlaybackUrlTtl)
		resourcedata.SetNillableValue(d, "recording_batch_download_url_ttl", recordingSettings.RecordingBatchDownloadUrlTtl)

		log.Printf("Read recording settings %s %s", d.Id(), *recordingSettings.Name)
		return cc.CheckState(d)
	})
}

// updateRecordingSettings is used by the recording_settings resource to update an recording settings in Genesys Cloud
func updateRecordingSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getRecordingSettingsProxy(sdkConfig)

	recordingSettings := getRecordingSettingsFromResourceData(d)

	log.Printf("Updating recording settings %s", *recordingSettings.Name)
	recordingSettings, resp, err := proxy.updateRecordingSettings(ctx, d.Id(), &recordingSettings)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to update recording settings %s: %s", d.Id(), err), resp)
	}

	log.Printf("Updated recording settings %s", *recordingSettings.Id)
	return readRecordingSettings(ctx, d, meta)
}

// deleteRecordingSettings is used by the recording_settings resource to delete an recording settings from Genesys cloud
func deleteRecordingSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// getRecordingSettingsFromResourceData maps data from schema ResourceData object to a platformclientv2.Recordingsettings
func getRecordingSettingsFromResourceData(d *schema.ResourceData) platformclientv2.Recordingsettings {
	return platformclientv2.Recordingsettings{
		MaxSimultaneousStreams:                platformclientv2.Int(d.Get("max_simultaneous_streams").(int)),
		MaxConfigurableScreenRecordingStreams: platformclientv2.Int(d.Get("max_configurable_screen_recording_streams").(int)),
		RegionalRecordingStorageEnabled:       platformclientv2.Bool(d.Get("regional_recording_storage_enabled").(bool)),
		RecordingPlaybackUrlTtl:               platformclientv2.Int(d.Get("recording_playback_url_ttl").(int)),
		RecordingBatchDownloadUrlTtl:          platformclientv2.Int(d.Get("recording_batch_download_url_ttl").(int)),
	}
}
