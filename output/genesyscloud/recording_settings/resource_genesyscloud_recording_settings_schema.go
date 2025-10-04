package recording_settings

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_recording_settings_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the recording_settings resource.
3.  The datasource schema definitions for the recording_settings datasource.
4.  The resource exporter configuration for the recording_settings exporter.
*/
const resourceName = "genesyscloud_recording_settings"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(ResourceType, ResourceRecordingSettings())
}

// ResourceRecordingSettings registers the genesyscloud_recording_settings resource with Terraform
func ResourceRecordingSettings() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud recording settings`,

		CreateContext: provider.CreateWithPooledClient(createRecordingSettings),
		ReadContext:   provider.ReadWithPooledClient(readRecordingSettings),
		UpdateContext: provider.UpdateWithPooledClient(updateRecordingSettings),
		DeleteContext: provider.DeleteWithPooledClient(deleteRecordingSettings),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			`max_simultaneous_streams`: {
				Description: `Maximum number of simultaneous screen recording streams`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`max_configurable_screen_recording_streams`: {
				Description: `Upper limit that maxSimultaneousStreams can be configured`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`regional_recording_storage_enabled`: {
				Description: `Store call recordings in the region where they are intended to be recorded, otherwise in the organization's home region`,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			`recording_playback_url_ttl`: {
				Description: `The duration in minutes for which the generated URL for recording playback remains valid.The default duration is set to 60 minutes, with a minimum allowable duration of 2 minutes and a maximum of 60 minutes.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`recording_batch_download_url_ttl`: {
				Description: `The duration in minutes for which the generated URL for recording batch download remains valid.The default duration is set to 60 minutes, with a minimum allowable duration of 2 minutes and a maximum of 60 minutes.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
		},
	}
}
