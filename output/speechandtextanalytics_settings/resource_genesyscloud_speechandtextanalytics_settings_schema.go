package speechandtextanalytics_settings

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_speechandtextanalytics_settings_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the speechandtextanalytics_settings resource.
3.  The datasource schema definitions for the speechandtextanalytics_settings datasource.
4.  The resource exporter configuration for the speechandtextanalytics_settings exporter.
*/
const resourceName = "genesyscloud_speechandtextanalytics_settings"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(resourceName, ResourceSpeechandtextanalyticsSettings())
}

// ResourceSpeechandtextanalyticsSettings registers the genesyscloud_speechandtextanalytics_settings resource with Terraform
func ResourceSpeechandtextanalyticsSettings() *schema.Resource {
	addressableEntityRefResource := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics settings`,

		CreateContext: provider.CreateWithPooledClient(createSpeechandtextanalyticsSettings),
		ReadContext:   provider.ReadWithPooledClient(readSpeechandtextanalyticsSettings),
		UpdateContext: provider.UpdateWithPooledClient(updateSpeechandtextanalyticsSettings),
		DeleteContext: provider.DeleteWithPooledClient(deleteSpeechandtextanalyticsSettings),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			`default_program`: {
				Description: `Setting to choose name for the default program for topic detection`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        addressableEntityRefResource,
			},
			`expected_dialects`: {
				Description: `Setting to choose expected dialects`,
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			`text_analytics_enabled`: {
				Description: `Setting to enable/disable text analytics`,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			`agent_empathy_enabled`: {
				Description: `Setting to enable/disable Agent Empathy setting`,
				Optional:    true,
				Type:        schema.TypeBool,
			},
		},
	}
}
