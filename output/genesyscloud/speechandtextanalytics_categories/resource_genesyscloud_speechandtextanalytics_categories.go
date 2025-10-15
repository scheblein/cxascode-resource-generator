package speechandtextanalytics_categories

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
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
The resource_genesyscloud_speechandtextanalytics_categories.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthSpeechandtextanalyticsCategories retrieves all of the speechandtextanalytics categories via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthSpeechandtextanalyticsCategoriess(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	proxy := newSpeechandtextanalyticsCategoriesProxy(clientConfig)
	resources := make(resourceExporter.ResourceIDMetaMap)

	staCategorys, resp, err := proxy.getAllSpeechandtextanalyticsCategories(ctx)
	if err != nil {
		return nil, util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to get speechandtextanalytics categories: %v", err), resp)
	}

	for _, staCategory := range *staCategorys {
		resources[*staCategory.Id] = &resourceExporter.ResourceMeta{BlockLabel: *staCategory.Name}
	}

	return resources, nil
}

// createSpeechandtextanalyticsCategories is used by the speechandtextanalytics_categories resource to create Genesys cloud speechandtextanalytics categories
func createSpeechandtextanalyticsCategories(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsCategoriesProxy(sdkConfig)

	speechandtextanalyticsCategories := getSpeechandtextanalyticsCategoriesFromResourceData(d)

	log.Printf("Creating speechandtextanalytics categories %s", *speechandtextanalyticsCategories.Name)
	staCategory, resp, err := proxy.createSpeechandtextanalyticsCategories(ctx, &speechandtextanalyticsCategories)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to create speechandtextanalytics categories: %s", err), resp)
	}

	d.SetId(*staCategory.Id)
	log.Printf("Created speechandtextanalytics categories %s", *staCategory.Id)
	return readSpeechandtextanalyticsCategories(ctx, d, meta)
}

// readSpeechandtextanalyticsCategories is used by the speechandtextanalytics_categories resource to read an speechandtextanalytics categories from genesys cloud
func readSpeechandtextanalyticsCategories(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsCategoriesProxy(sdkConfig)
	cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceSpeechandtextanalyticsCategories(), constants.ConsistencyChecks(), resourceName)

	log.Printf("Reading speechandtextanalytics categories %s", d.Id())

	return util.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		staCategory, resp, getErr := proxy.getSpeechandtextanalyticsCategoriesById(ctx, d.Id())
		if getErr != nil {
			if util.IsStatus404(resp) {
				return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read speechandtextanalytics categories %s: %s", d.Id(), getErr), resp))
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read speechandtextanalytics categories %s: %s", d.Id(), getErr), resp))
		}

		resourcedata.SetNillableValue(d, "name", staCategory.Name)
		resourcedata.SetNillableValue(d, "description", staCategory.Description)
		resourcedata.SetNillableValue(d, "interaction_type", staCategory.InteractionType)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "criteria", staCategory.Criteria, flattenOperand)

		log.Printf("Read speechandtextanalytics categories %s %s", d.Id(), *staCategory.Name)
		return cc.CheckState(d)
	})
}

// updateSpeechandtextanalyticsCategories is used by the speechandtextanalytics_categories resource to update an speechandtextanalytics categories in Genesys Cloud
func updateSpeechandtextanalyticsCategories(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsCategoriesProxy(sdkConfig)

	speechandtextanalyticsCategories := getSpeechandtextanalyticsCategoriesFromResourceData(d)

	log.Printf("Updating speechandtextanalytics categories %s", *speechandtextanalyticsCategories.Name)
	staCategory, resp, err := proxy.updateSpeechandtextanalyticsCategories(ctx, d.Id(), &speechandtextanalyticsCategories)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to update speechandtextanalytics categories %s: %s", d.Id(), err), resp)
	}

	log.Printf("Updated speechandtextanalytics categories %s", *staCategory.Id)
	return readSpeechandtextanalyticsCategories(ctx, d, meta)
}

// deleteSpeechandtextanalyticsCategories is used by the speechandtextanalytics_categories resource to delete an speechandtextanalytics categories from Genesys cloud
func deleteSpeechandtextanalyticsCategories(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsCategoriesProxy(sdkConfig)

	resp, err := proxy.deleteSpeechandtextanalyticsCategories(ctx, d.Id())
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to delete speechandtextanalytics categories %s: %s", d.Id(), err), resp)
	}

	return util.WithRetries(ctx, 180*time.Second, func() *retry.RetryError {
		_, resp, err := proxy.getSpeechandtextanalyticsCategoriesById(ctx, d.Id())

		if err != nil {
			if util.IsStatus404(resp) {
				log.Printf("Deleted speechandtextanalytics categories %s", d.Id())
				return nil
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Error deleting speechandtextanalytics categories %s: %s", d.Id(), err), resp))
		}

		return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("speechandtextanalytics categories %s still exists", d.Id()), resp))
	})
}
