package speechandtextanalytics_programs

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_speechandtextanalytics_programs_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the speechandtextanalytics_programs resource.
3.  The datasource schema definitions for the speechandtextanalytics_programs datasource.
4.  The resource exporter configuration for the speechandtextanalytics_programs exporter.
*/
const resourceName = "genesyscloud_speechandtextanalytics_programs"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(ResourceType, ResourceSpeechandtextanalyticsPrograms())
	regInstance.RegisterDataSource(ResourceType, DataSourceSpeechandtextanalyticsPrograms())
	regInstance.RegisterExporter(ResourceType, SpeechandtextanalyticsProgramsExporter())
}

// ResourceSpeechandtextanalyticsPrograms registers the genesyscloud_speechandtextanalytics_programs resource with Terraform
func ResourceSpeechandtextanalyticsPrograms() *schema.Resource {
	baseTopicEntitiyResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`name`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}

	addressableEntityRefResource := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics programs`,

		CreateContext: provider.CreateWithPooledClient(createSpeechandtextanalyticsPrograms),
		ReadContext:   provider.ReadWithPooledClient(readSpeechandtextanalyticsPrograms),
		UpdateContext: provider.UpdateWithPooledClient(updateSpeechandtextanalyticsPrograms),
		DeleteContext: provider.DeleteWithPooledClient(deleteSpeechandtextanalyticsPrograms),
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
			`topics`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        baseTopicEntitiyResource,
			},
			`tags`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
			`topic_links_job`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        addressableEntityRefResource,
			},
		},
	}
}

// SpeechandtextanalyticsProgramsExporter returns the resourceExporter object used to hold the genesyscloud_speechandtextanalytics_programs exporter's config
func SpeechandtextanalyticsProgramsExporter() *resourceExporter.ResourceExporter {
	return &resourceExporter.ResourceExporter{
		GetResourcesFunc: provider.GetAllWithPooledClient(getAllAuthSpeechandtextanalyticsProgramss),
		RefAttrs:         map[string]*resourceExporter.RefAttrSettings{
			// TODO: Add any reference attributes here
		},
	}
}

// DataSourceSpeechandtextanalyticsPrograms registers the genesyscloud_speechandtextanalytics_programs data source
func DataSourceSpeechandtextanalyticsPrograms() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics programs data source. Select an speechandtextanalytics programs by name`,
		ReadContext: provider.ReadWithPooledClient(dataSourceSpeechandtextanalyticsProgramsRead),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: `speechandtextanalytics programs name`,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
