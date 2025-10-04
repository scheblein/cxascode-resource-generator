package languageunderstanding_miners

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_languageunderstanding_miners_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the languageunderstanding_miners resource.
3.  The datasource schema definitions for the languageunderstanding_miners datasource.
4.  The resource exporter configuration for the languageunderstanding_miners exporter.
*/
const resourceName = "genesyscloud_languageunderstanding_miners"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(resourceName, ResourceLanguageunderstandingMiners())
	regInstance.RegisterDataSource(resourceName, DataSourceLanguageunderstandingMiners())
	regInstance.RegisterExporter(resourceName, LanguageunderstandingMinersExporter())
}

// ResourceLanguageunderstandingMiners registers the genesyscloud_languageunderstanding_miners resource with Terraform
func ResourceLanguageunderstandingMiners() *schema.Resource {
	minerErrorInfoResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`message`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`code`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`message_with_params`: {
				Description: `Error message with params included.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			// TODO: Handle message_params property
		},
	}

	draftResource := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	minerResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`name`: {
				Description: `Chat Corpus Name.`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`language`: {
				Description: `Language Localization code.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`miner_type`: {
				Description: `Type of the miner, intent or topic.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`seeding`: {
				Description: `Flag to indicate whether seeding is supported for this miner.`,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			`status`: {
				Description: `Status of the miner.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`conversations_date_range_start`: {
				Description: `Date from which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`conversations_date_range_end`: {
				Description: `Date till which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`date_completed`: {
				Description: `Date when the mining process was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`message`: {
				Description: `Mining message if present.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`error_info`: {
				Description: `Error Information`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        minerErrorInfoResource,
			},
			`warning_info`: {
				Description: `Warning Information`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        minerErrorInfoResource,
			},
			`conversation_data_uploaded`: {
				Description: `Flag to indicate whether data file to be mined was uploaded.`,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			`media_type`: {
				Description: `Media type for filtering conversations.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`participant_type`: {
				Description: `Type of the participant, either agent, customer or both.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`queue_ids`: {
				Description: `List of queue IDs for filtering conversations.`,
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			`date_triggered`: {
				Description: `Date when the miner started execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`latest_draft_version`: {
				Description: `Latest draft details of the miner.`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        draftResource,
			},
			`conversations_fetched_count`: {
				Description: `Number of conversations/transcripts fetched.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`conversations_valid_count`: {
				Description: `Number of conversations/recordings/transcripts that were found valid for mining purposes.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`getmined_item_count`: {
				Description: `Number of intents or topics based on the miner type.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
		},
	}

	draftIntentsResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`name`: {
				Description: `Name/Label for an intent.`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`utterances`: {
				Description: `The utterances that are extracted for an Intent.`,
				Required:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	draftTopicsResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			`name`: {
				Description: `Topic name.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`miner`: {
				Description: `The miner to which the topic belongs.`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        minerResource,
			},
			`conversation_count`: {
				Description: `Number of conversations where a topic has occurred.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`conversation_percent`: {
				Description: `Percentage of conversations where a topic has occurred.`,
				Optional:    true,
				Type:        schema.TypeFloat,
			},
			`utterance_count`: {
				Description: `Number of unique utterances where a topic has occurred.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`phrase_count`: {
				Description: `Number of unique phrases (sub-utterances) where a topic has occurred.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`phrases`: {
				Description: `The phrases that are extracted for a topic.`,
				Required:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	return &schema.Resource{
		Description: `Genesys Cloud languageunderstanding miners`,

		CreateContext: provider.CreateWithPooledClient(createLanguageunderstandingMiners),
		ReadContext:   provider.ReadWithPooledClient(readLanguageunderstandingMiners),
		UpdateContext: provider.UpdateWithPooledClient(updateLanguageunderstandingMiners),
		DeleteContext: provider.DeleteWithPooledClient(deleteLanguageunderstandingMiners),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			`name`: {
				Description: `Chat Corpus Name.`,
				Required:    true,
				Type:        schema.TypeString,
			},
			`language`: {
				Description: `Language Localization code.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`miner_type`: {
				Description: `Type of the miner, intent or topic.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`seeding`: {
				Description: `Flag to indicate whether seeding is supported for this miner.`,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			`status`: {
				Description: `Status of the miner.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`conversations_date_range_start`: {
				Description: `Date from which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`conversations_date_range_end`: {
				Description: `Date till which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`date_completed`: {
				Description: `Date when the mining process was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`message`: {
				Description: `Mining message if present.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`error_info`: {
				Description: `Error Information`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        minerErrorInfoResource,
			},
			`warning_info`: {
				Description: `Warning Information`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        minerErrorInfoResource,
			},
			`conversation_data_uploaded`: {
				Description: `Flag to indicate whether data file to be mined was uploaded.`,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			`media_type`: {
				Description: `Media type for filtering conversations.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`participant_type`: {
				Description: `Type of the participant, either agent, customer or both.`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`queue_ids`: {
				Description: `List of queue IDs for filtering conversations.`,
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			`date_triggered`: {
				Description: `Date when the miner started execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z`,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`latest_draft_version`: {
				Description: `Latest draft details of the miner.`,
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        draftResource,
			},
			`conversations_fetched_count`: {
				Description: `Number of conversations/transcripts fetched.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`conversations_valid_count`: {
				Description: `Number of conversations/recordings/transcripts that were found valid for mining purposes.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			`getmined_item_count`: {
				Description: `Number of intents or topics based on the miner type.`,
				Optional:    true,
				Type:        schema.TypeInt,
			},
		},
	}
}

// LanguageunderstandingMinersExporter returns the resourceExporter object used to hold the genesyscloud_languageunderstanding_miners exporter's config
func LanguageunderstandingMinersExporter() *resourceExporter.ResourceExporter {
	return &resourceExporter.ResourceExporter{
		GetResourcesFunc: provider.GetAllWithPooledClient(getAllAuthLanguageunderstandingMinerss),
		RefAttrs:         map[string]*resourceExporter.RefAttrSettings{
			// TODO: Add any reference attributes here
		},
	}
}

// DataSourceLanguageunderstandingMiners registers the genesyscloud_languageunderstanding_miners data source
func DataSourceLanguageunderstandingMiners() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud languageunderstanding miners data source. Select an languageunderstanding miners by name`,
		ReadContext: provider.ReadWithPooledClient(dataSourceLanguageunderstandingMinersRead),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: `languageunderstanding miners name`,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
