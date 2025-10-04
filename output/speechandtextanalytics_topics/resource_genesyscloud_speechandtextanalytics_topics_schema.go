package speechandtextanalytics_topics

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_speechandtextanalytics_topics_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the speechandtextanalytics_topics resource.
3.  The datasource schema definitions for the speechandtextanalytics_topics datasource.
4.  The resource exporter configuration for the speechandtextanalytics_topics exporter.
*/
const resourceName = "genesyscloud_speechandtextanalytics_topics"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(resourceName, ResourceSpeechandtextanalyticsTopics())
	regInstance.RegisterDataSource(resourceName, DataSourceSpeechandtextanalyticsTopics())
	regInstance.RegisterExporter(resourceName, SpeechandtextanalyticsTopicsExporter())
}

// ResourceSpeechandtextanalyticsTopics registers the genesyscloud_speechandtextanalytics_topics resource with Terraform
func ResourceSpeechandtextanalyticsTopics() *schema.Resource {
	baseProgramEntityResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`name`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}

	phraseResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`text`: {
				Description: `The phrase text`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`strictness`: {
				Description: `The phrase strictness, default value is null`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`sentiment`: {
				Description: `The phrase sentiment, default value is Unspecified. Note: Sentiment value for phrases is currently not in use and has no impact to the system.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}

	addressableEntityRefResource := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics topics`,

		CreateContext: provider.CreateWithPooledClient(createSpeechandtextanalyticsTopics),
		ReadContext:   provider.ReadWithPooledClient(readSpeechandtextanalyticsTopics),
		UpdateContext: provider.UpdateWithPooledClient(updateSpeechandtextanalyticsTopics),
		DeleteContext: provider.DeleteWithPooledClient(deleteSpeechandtextanalyticsTopics),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			`name`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`description`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`published`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			`strictness`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`matching_type`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`programs`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        baseProgramEntityResource,
			},
			`tags`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			`dialect`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`participants`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`phrases`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        phraseResource,
			},
			`published_by`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        addressableEntityRefResource,
			},
			`date_published`: {
				Description: `Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z`,
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

// SpeechandtextanalyticsTopicsExporter returns the resourceExporter object used to hold the genesyscloud_speechandtextanalytics_topics exporter's config
func SpeechandtextanalyticsTopicsExporter() *resourceExporter.ResourceExporter {
	return &resourceExporter.ResourceExporter{
		GetResourcesFunc: provider.GetAllWithPooledClient(getAllAuthSpeechandtextanalyticsTopicss),
		RefAttrs:         map[string]*resourceExporter.RefAttrSettings{
			// TODO: Add any reference attributes here
		},
	}
}

// DataSourceSpeechandtextanalyticsTopics registers the genesyscloud_speechandtextanalytics_topics data source
func DataSourceSpeechandtextanalyticsTopics() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics topics data source. Select an speechandtextanalytics topics by name`,
		ReadContext: provider.ReadWithPooledClient(dataSourceSpeechandtextanalyticsTopicsRead),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: `speechandtextanalytics topics name`,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
