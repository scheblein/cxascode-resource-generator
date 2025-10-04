package organization_presence_definition

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
	"log"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	"time"

	"terraform-provider-genesyscloud/genesyscloud/consistency_checker"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"terraform-provider-genesyscloud/genesyscloud/util/constants"
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

/*
The resource_genesyscloud_organization_presence_definition.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthOrganizationPresenceDefinition retrieves all of the organization presence definition via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthOrganizationPresenceDefinitions(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	proxy := newOrganizationPresenceDefinitionProxy(clientConfig)
	resources := make(resourceExporter.ResourceIDMetaMap)

	organizationPresenceDefinitions, resp, err := proxy.getAllOrganizationPresenceDefinition(ctx)
	if err != nil {
		return nil, util.BuildAPIDiagnosticError(resourceName, fmt.Sprintf("Failed to get organization presence definition: %v", err), resp)
	}

	for _, organizationPresenceDefinition := range *organizationPresenceDefinitions {
		resources[*organizationPresenceDefinition.Id] = &resourceExporter.ResourceMeta{Name: *organizationPresenceDefinition.Name}
	}

	return resources, nil
}

// createOrganizationPresenceDefinition is used by the organization_presence_definition resource to create Genesys cloud organization presence definition
func createOrganizationPresenceDefinition(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getOrganizationPresenceDefinitionProxy(sdkConfig)

	organizationPresenceDefinition := getOrganizationPresenceDefinitionFromResourceData(d)

	log.Printf("Creating organization presence definition %s", *organizationPresenceDefinition.Name)
	organizationPresenceDefinition, resp, err := proxy.createOrganizationPresenceDefinition(ctx, &organizationPresenceDefinition)
	if err != nil {
		return util.BuildAPIDiagnosticError(resourceName, fmt.Sprintf("Failed to create organization presence definition: %s", err), resp)
	}

	d.SetId(*organizationPresenceDefinition.Id)
	log.Printf("Created organization presence definition %s", *organizationPresenceDefinition.Id)
	return readOrganizationPresenceDefinition(ctx, d, meta)
}

// readOrganizationPresenceDefinition is used by the organization_presence_definition resource to read an organization presence definition from genesys cloud
func readOrganizationPresenceDefinition(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getOrganizationPresenceDefinitionProxy(sdkConfig)
	cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceOrganizationPresenceDefinition(), constants.DefaultConsistencyChecks, resourceName)

	log.Printf("Reading organization presence definition %s", d.Id())

	return util.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		organizationPresenceDefinition, resp, getErr := proxy.getOrganizationPresenceDefinitionById(ctx, d.Id())
		if getErr != nil {
			if util.IsStatus404(resp) {
				return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Failed to read organization presence definition %s: %s", d.Id(), getErr), resp))
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Failed to read organization presence definition %s: %s", d.Id(), getErr), resp))
		}

		resourcedata.SetNillableValue(d, "name", organizationPresenceDefinition.Name)
		// TODO: Handle language_labels property
		resourcedata.SetNillableValue(d, "system_presence", organizationPresenceDefinition.SystemPresence)
		resourcedata.SetNillableReferenceWritableDivision(d, "division_id", organizationPresenceDefinition.Division)
		resourcedata.SetNillableValue(d, "deactivated", organizationPresenceDefinition.Deactivated)

		log.Printf("Read organization presence definition %s %s", d.Id(), *organizationPresenceDefinition.Name)
		return cc.CheckState(d)
	})
}

// updateOrganizationPresenceDefinition is used by the organization_presence_definition resource to update an organization presence definition in Genesys Cloud
func updateOrganizationPresenceDefinition(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getOrganizationPresenceDefinitionProxy(sdkConfig)

	organizationPresenceDefinition := getOrganizationPresenceDefinitionFromResourceData(d)

	log.Printf("Updating organization presence definition %s", *organizationPresenceDefinition.Name)
	organizationPresenceDefinition, resp, err := proxy.updateOrganizationPresenceDefinition(ctx, d.Id(), &organizationPresenceDefinition)
	if err != nil {
		return util.BuildAPIDiagnosticError(resourceName, fmt.Sprintf("Failed to update organization presence definition %s: %s", d.Id(), err), resp)
	}

	log.Printf("Updated organization presence definition %s", *organizationPresenceDefinition.Id)
	return readOrganizationPresenceDefinition(ctx, d, meta)
}

// deleteOrganizationPresenceDefinition is used by the organization_presence_definition resource to delete an organization presence definition from Genesys cloud
func deleteOrganizationPresenceDefinition(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getOrganizationPresenceDefinitionProxy(sdkConfig)

	resp, err := proxy.deleteOrganizationPresenceDefinition(ctx, d.Id())
	if err != nil {
		return util.BuildAPIDiagnosticError(resourceName, fmt.Sprintf("Failed to delete organization presence definition %s: %s", d.Id(), err), resp)
	}

	return util.WithRetries(ctx, 180*time.Second, func() *retry.RetryError {
		_, resp, err := proxy.getOrganizationPresenceDefinitionById(ctx, d.Id())

		if err != nil {
			if util.IsStatus404(resp) {
				log.Printf("Deleted organization presence definition %s", d.Id())
				return nil
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Error deleting organization presence definition %s: %s", d.Id(), err), resp))
		}

		return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("organization presence definition %s still exists", d.Id()), resp))
	})
}

// getOrganizationPresenceDefinitionFromResourceData maps data from schema ResourceData object to a platformclientv2.Organizationpresencedefinition
func getOrganizationPresenceDefinitionFromResourceData(d *schema.ResourceData) platformclientv2.Organizationpresencedefinition {
	return platformclientv2.Organizationpresencedefinition{
		Name: platformclientv2.String(d.Get("name").(string)),
		// TODO: Handle language_labels property
		SystemPresence: platformclientv2.String(d.Get("system_presence").(string)),
		Division:       &platformclientv2.Writabledivision{Id: platformclientv2.String(d.Get("division_id").(string))},
		Deactivated:    platformclientv2.Bool(d.Get("deactivated").(bool)),
	}
}
