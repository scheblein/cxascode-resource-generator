package speechandtextanalytics_sentimentfeedback

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_speechandtextanalytics_sentimentfeedback_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the speechandtextanalytics_sentimentfeedback resource.
3.  The datasource schema definitions for the speechandtextanalytics_sentimentfeedback datasource.
4.  The resource exporter configuration for the speechandtextanalytics_sentimentfeedback exporter.
*/
const resourceName = "genesyscloud_speechandtextanalytics_sentimentfeedback"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(ResourceType, ResourceSpeechandtextanalyticsSentimentfeedback())
	regInstance.RegisterDataSource(ResourceType, DataSourceSpeechandtextanalyticsSentimentfeedback())
	regInstance.RegisterExporter(ResourceType, SpeechandtextanalyticsSentimentfeedbackExporter())
}

// ResourceSpeechandtextanalyticsSentimentfeedback registers the genesyscloud_speechandtextanalytics_sentimentfeedback resource with Terraform
func ResourceSpeechandtextanalyticsSentimentfeedback() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics sentimentfeedback`,

		CreateContext: provider.CreateWithPooledClient(createSpeechandtextanalyticsSentimentfeedback),
		ReadContext:   provider.ReadWithPooledClient(readSpeechandtextanalyticsSentimentfeedback),
		UpdateContext: provider.UpdateWithPooledClient(updateSpeechandtextanalyticsSentimentfeedback),
		DeleteContext: provider.DeleteWithPooledClient(deleteSpeechandtextanalyticsSentimentfeedback),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			`phrase`: {
				Description: `The phrase for which sentiment feedback is provided`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`dialect`: {
				Description: `The dialect for the given phrase, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`feedback_value`: {
				Description: `The sentiment feedback value for the given phrase`,
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

// SpeechandtextanalyticsSentimentfeedbackExporter returns the resourceExporter object used to hold the genesyscloud_speechandtextanalytics_sentimentfeedback exporter's config
func SpeechandtextanalyticsSentimentfeedbackExporter() *resourceExporter.ResourceExporter {
	return &resourceExporter.ResourceExporter{
		GetResourcesFunc: provider.GetAllWithPooledClient(getAllAuthSpeechandtextanalyticsSentimentfeedbacks),
		RefAttrs:         map[string]*resourceExporter.RefAttrSettings{
			// TODO: Add any reference attributes here
		},
	}
}

// DataSourceSpeechandtextanalyticsSentimentfeedback registers the genesyscloud_speechandtextanalytics_sentimentfeedback data source
func DataSourceSpeechandtextanalyticsSentimentfeedback() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics sentimentfeedback data source. Select an speechandtextanalytics sentimentfeedback by name`,
		ReadContext: provider.ReadWithPooledClient(dataSourceSpeechandtextanalyticsSentimentfeedbackRead),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: `speechandtextanalytics sentimentfeedback name`,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
