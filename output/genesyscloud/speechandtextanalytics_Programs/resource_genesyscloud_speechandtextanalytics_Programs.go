package speechandtextanalytics_programs

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
The resource_genesyscloud_speechandtextanalytics_programs.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthSpeechandtextanalyticsPrograms retrieves all of the speechandtextanalytics programs via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthSpeechandtextanalyticsProgramss(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	proxy := newSpeechandtextanalyticsProgramsProxy(clientConfig)
	resources := make(resourceExporter.ResourceIDMetaMap)

	programs, resp, err := proxy.getAllSpeechandtextanalyticsPrograms(ctx)
	if err != nil {
		return nil, util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to get speechandtextanalytics programs: %v", err), resp)
	}

	for _, program := range *programs {
		resources[*program.Id] = &resourceExporter.ResourceMeta{BlockLabel: *program.Name}
	}

	return resources, nil
}

// createSpeechandtextanalyticsPrograms is used by the speechandtextanalytics_programs resource to create Genesys cloud speechandtextanalytics programs
func createSpeechandtextanalyticsPrograms(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsProgramsProxy(sdkConfig)

	speechandtextanalyticsPrograms := getSpeechandtextanalyticsProgramsFromResourceData(d)

	log.Printf("Creating speechandtextanalytics programs %s", *speechandtextanalyticsPrograms.Name)
	program, resp, err := proxy.createSpeechandtextanalyticsPrograms(ctx, &speechandtextanalyticsPrograms)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to create speechandtextanalytics programs: %s", err), resp)
	}

	d.SetId(*program.Id)
	log.Printf("Created speechandtextanalytics programs %s", *program.Id)
	return readSpeechandtextanalyticsPrograms(ctx, d, meta)
}

// readSpeechandtextanalyticsPrograms is used by the speechandtextanalytics_programs resource to read an speechandtextanalytics programs from genesys cloud
func readSpeechandtextanalyticsPrograms(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsProgramsProxy(sdkConfig)
	cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceSpeechandtextanalyticsPrograms(), constants.ConsistencyChecks(), resourceName)

	log.Printf("Reading speechandtextanalytics programs %s", d.Id())

	return util.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		program, resp, getErr := proxy.getSpeechandtextanalyticsProgramsById(ctx, d.Id())
		if getErr != nil {
			if util.IsStatus404(resp) {
				return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read speechandtextanalytics programs %s: %s", d.Id(), getErr), resp))
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Failed to read speechandtextanalytics programs %s: %s", d.Id(), getErr), resp))
		}

		resourcedata.SetNillableValue(d, "name", program.Name)
		resourcedata.SetNillableValue(d, "description", program.Description)
		resourcedata.SetNillableValue(d, "published", program.Published)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "topics", program.Topics, flattenBaseTopicEntitiys)
		resourcedata.SetNillableValue(d, "tags", program.Tags)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "published_by", program.PublishedBy, flattenAddressableEntityRef)
		resourcedata.SetNillableValue(d, "date_published", program.DatePublished)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "topic_links_job", program.TopicLinksJob, flattenAddressableEntityRef)

		log.Printf("Read speechandtextanalytics programs %s %s", d.Id(), *program.Name)
		return cc.CheckState(d)
	})
}

// updateSpeechandtextanalyticsPrograms is used by the speechandtextanalytics_programs resource to update an speechandtextanalytics programs in Genesys Cloud
func updateSpeechandtextanalyticsPrograms(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsProgramsProxy(sdkConfig)

	speechandtextanalyticsPrograms := getSpeechandtextanalyticsProgramsFromResourceData(d)

	log.Printf("Updating speechandtextanalytics programs %s", *speechandtextanalyticsPrograms.Name)
	program, resp, err := proxy.updateSpeechandtextanalyticsPrograms(ctx, d.Id(), &speechandtextanalyticsPrograms)
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to update speechandtextanalytics programs %s: %s", d.Id(), err), resp)
	}

	log.Printf("Updated speechandtextanalytics programs %s", *program.Id)
	return readSpeechandtextanalyticsPrograms(ctx, d, meta)
}

// deleteSpeechandtextanalyticsPrograms is used by the speechandtextanalytics_programs resource to delete an speechandtextanalytics programs from Genesys cloud
func deleteSpeechandtextanalyticsPrograms(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsProgramsProxy(sdkConfig)

	resp, err := proxy.deleteSpeechandtextanalyticsPrograms(ctx, d.Id())
	if err != nil {
		return util.BuildAPIDiagnosticError(ResourceType, fmt.Sprintf("Failed to delete speechandtextanalytics programs %s: %s", d.Id(), err), resp)
	}

	return util.WithRetries(ctx, 180*time.Second, func() *retry.RetryError {
		_, resp, err := proxy.getSpeechandtextanalyticsProgramsById(ctx, d.Id())

		if err != nil {
			if util.IsStatus404(resp) {
				log.Printf("Deleted speechandtextanalytics programs %s", d.Id())
				return nil
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("Error deleting speechandtextanalytics programs %s: %s", d.Id(), err), resp))
		}

		return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(ResourceType, fmt.Sprintf("speechandtextanalytics programs %s still exists", d.Id()), resp))
	})
}
