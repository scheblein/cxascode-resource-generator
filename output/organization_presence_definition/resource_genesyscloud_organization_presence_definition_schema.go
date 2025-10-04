package organization_presence_definition

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_organization_presence_definition_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the organization_presence_definition resource.
3.  The datasource schema definitions for the organization_presence_definition datasource.
4.  The resource exporter configuration for the organization_presence_definition exporter.
*/
const resourceName = "genesyscloud_organization_presence_definition"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(ResourceType, ResourceOrganizationPresenceDefinition())
	regInstance.RegisterDataSource(ResourceType, DataSourceOrganizationPresenceDefinition())
	regInstance.RegisterExporter(ResourceType, OrganizationPresenceDefinitionExporter())
}

// ResourceOrganizationPresenceDefinition registers the genesyscloud_organization_presence_definition resource with Terraform
func ResourceOrganizationPresenceDefinition() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud organization presence definition`,

		CreateContext: provider.CreateWithPooledClient(createOrganizationPresenceDefinition),
		ReadContext:   provider.ReadWithPooledClient(readOrganizationPresenceDefinition),
		UpdateContext: provider.UpdateWithPooledClient(updateOrganizationPresenceDefinition),
		DeleteContext: provider.DeleteWithPooledClient(deleteOrganizationPresenceDefinition),
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
			// TODO: Handle language_labels property
			`system_presence`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`division_id`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeString,
			},
			`deactivated`: {
				Description: ``,
				Optional:    true,
				Type:        schema.TypeBool,
			},
		},
	}
}

// OrganizationPresenceDefinitionExporter returns the resourceExporter object used to hold the genesyscloud_organization_presence_definition exporter's config
func OrganizationPresenceDefinitionExporter() *resourceExporter.ResourceExporter {
	return &resourceExporter.ResourceExporter{
		GetResourcesFunc: provider.GetAllWithPooledClient(getAllAuthOrganizationPresenceDefinitions),
		RefAttrs:         map[string]*resourceExporter.RefAttrSettings{
			// TODO: Add any reference attributes here
		},
	}
}

// DataSourceOrganizationPresenceDefinition registers the genesyscloud_organization_presence_definition data source
func DataSourceOrganizationPresenceDefinition() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud organization presence definition data source. Select an organization presence definition by name`,
		ReadContext: provider.ReadWithPooledClient(dataSourceOrganizationPresenceDefinitionRead),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: `organization presence definition name`,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
