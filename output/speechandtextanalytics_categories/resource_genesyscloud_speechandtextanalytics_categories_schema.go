package speechandtextanalytics_categories

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_speechandtextanalytics_categories_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the speechandtextanalytics_categories resource.
3.  The datasource schema definitions for the speechandtextanalytics_categories datasource.
4.  The resource exporter configuration for the speechandtextanalytics_categories exporter.
*/
const resourceName = "genesyscloud_speechandtextanalytics_categories"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(ResourceType, ResourceSpeechandtextanalyticsCategories())
	regInstance.RegisterDataSource(ResourceType, DataSourceSpeechandtextanalyticsCategories())
	regInstance.RegisterExporter(ResourceType, SpeechandtextanalyticsCategoriesExporter())
}

// ResourceSpeechandtextanalyticsCategories registers the genesyscloud_speechandtextanalytics_categories resource with Terraform
func ResourceSpeechandtextanalyticsCategories() *schema.Resource {
	termResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`word`: {
				Description: `Find term in interaction`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`participant_type`: {
				Description: `Dictates if term operand must come from the internal, external or both participants`,
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}

	operandPositionResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`starting_position_value`: {
				Description: `Defines starting point of a position range - number of seconds or words from the start or from the end of the interaction`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`starting_position_direction`: {
				Description: `Dictates starting position directionality`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`ending_position_value`: {
				Description: `Defines ending point of a position range - number of seconds or words from the start or from the end of the interaction`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`ending_position_direction`: {
				Description: `Dictates ending position directionality`,
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}

	operatorPositionResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`voice_seconds_position`: {
				Description: `Number of seconds (for voice interactions) from operand match`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`digital_words_position`: {
				Description: `Number of words (for digital interactions) from operand match`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
		},
	}

	infixOperatorResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`operator_type`: {
				Description: `The logical operation that is applied on the operand against the following operand`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`operator_position`: {
				Description: `Dictates when the following operand should occur relative to current operand`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        operatorPositionResource,
			},
		},
	}

	operandResource := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics categories`,

		CreateContext: provider.CreateWithPooledClient(createSpeechandtextanalyticsCategories),
		ReadContext:   provider.ReadWithPooledClient(readSpeechandtextanalyticsCategories),
		UpdateContext: provider.UpdateWithPooledClient(updateSpeechandtextanalyticsCategories),
		DeleteContext: provider.DeleteWithPooledClient(deleteSpeechandtextanalyticsCategories),
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
				Description: `The description of the category`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`interaction_type`: {
				Description: `The type of interaction the category will apply to`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`criteria`: {
				Description: `A collection of conditions joined together by logical operation to provide more refined filtering of conversations`,
				Required:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        operandResource,
			},
		},
	}
}

// SpeechandtextanalyticsCategoriesExporter returns the resourceExporter object used to hold the genesyscloud_speechandtextanalytics_categories exporter's config
func SpeechandtextanalyticsCategoriesExporter() *resourceExporter.ResourceExporter {
	return &resourceExporter.ResourceExporter{
		GetResourcesFunc: provider.GetAllWithPooledClient(getAllAuthSpeechandtextanalyticsCategoriess),
		RefAttrs:         map[string]*resourceExporter.RefAttrSettings{
			// TODO: Add any reference attributes here
		},
	}
}

// DataSourceSpeechandtextanalyticsCategories registers the genesyscloud_speechandtextanalytics_categories data source
func DataSourceSpeechandtextanalyticsCategories() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud speechandtextanalytics categories data source. Select an speechandtextanalytics categories by name`,
		ReadContext: provider.ReadWithPooledClient(dataSourceSpeechandtextanalyticsCategoriesRead),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: `speechandtextanalytics categories name`,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
