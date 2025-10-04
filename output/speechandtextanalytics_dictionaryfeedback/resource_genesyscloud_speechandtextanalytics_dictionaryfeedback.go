package speechandtextanalytics_dictionaryfeedback

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
The resource_genesyscloud_speechandtextanalytics_dictionaryfeedback.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthSpeechandtextanalyticsDictionaryfeedback retrieves all of the speechandtextanalytics dictionaryfeedback via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthSpeechandtextanalyticsDictionaryfeedbacks(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	proxy := newSpeechandtextanalyticsDictionaryfeedbackProxy(clientConfig)
	resources := make(resourceExporter.ResourceIDMetaMap)

	dictionaryFeedbacks, resp, err := proxy.getAllSpeechandtextanalyticsDictionaryfeedback(ctx)
	if err != nil {
		return nil, util.BuildAPIDiagnosticError(resourceName, fmt.Sprintf("Failed to get speechandtextanalytics dictionaryfeedback: %v", err), resp)
	}

	for _, dictionaryFeedback := range *dictionaryFeedbacks {
		resources[*dictionaryFeedback.Id] = &resourceExporter.ResourceMeta{Name: *dictionaryFeedback.Name}
	}

	return resources, nil
}

// createSpeechandtextanalyticsDictionaryfeedback is used by the speechandtextanalytics_dictionaryfeedback resource to create Genesys cloud speechandtextanalytics dictionaryfeedback
func createSpeechandtextanalyticsDictionaryfeedback(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsDictionaryfeedbackProxy(sdkConfig)

	speechandtextanalyticsDictionaryfeedback := getSpeechandtextanalyticsDictionaryfeedbackFromResourceData(d)

	log.Printf("Creating speechandtextanalytics dictionaryfeedback %s", *speechandtextanalyticsDictionaryfeedback.Name)
	dictionaryFeedback, resp, err := proxy.createSpeechandtextanalyticsDictionaryfeedback(ctx, &speechandtextanalyticsDictionaryfeedback)
	if err != nil {
		return util.BuildAPIDiagnosticError(resourceName, fmt.Sprintf("Failed to create speechandtextanalytics dictionaryfeedback: %s", err), resp)
	}

	d.SetId(*dictionaryFeedback.Id)
	log.Printf("Created speechandtextanalytics dictionaryfeedback %s", *dictionaryFeedback.Id)
	return readSpeechandtextanalyticsDictionaryfeedback(ctx, d, meta)
}

// readSpeechandtextanalyticsDictionaryfeedback is used by the speechandtextanalytics_dictionaryfeedback resource to read an speechandtextanalytics dictionaryfeedback from genesys cloud
func readSpeechandtextanalyticsDictionaryfeedback(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsDictionaryfeedbackProxy(sdkConfig)
	cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceSpeechandtextanalyticsDictionaryfeedback(), constants.DefaultConsistencyChecks, resourceName)

	log.Printf("Reading speechandtextanalytics dictionaryfeedback %s", d.Id())

	return util.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		dictionaryFeedback, resp, getErr := proxy.getSpeechandtextanalyticsDictionaryfeedbackById(ctx, d.Id())
		if getErr != nil {
			if util.IsStatus404(resp) {
				return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Failed to read speechandtextanalytics dictionaryfeedback %s: %s", d.Id(), getErr), resp))
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Failed to read speechandtextanalytics dictionaryfeedback %s: %s", d.Id(), getErr), resp))
		}

		resourcedata.SetNillableValue(d, "term", dictionaryFeedback.Term)
		resourcedata.SetNillableValue(d, "dialect", dictionaryFeedback.Dialect)
		resourcedata.SetNillableValue(d, "boost_value", dictionaryFeedback.BoostValue)
		resourcedata.SetNillableValue(d, "source", dictionaryFeedback.Source)
		resourcedata.SetNillableValueWithInterfaceArrayWithFunc(d, "example_phrases", dictionaryFeedback.ExamplePhrases, flattenDictionaryFeedbackExamplePhrases)
		resourcedata.SetNillableValue(d, "sounds_like", dictionaryFeedback.SoundsLike)

		log.Printf("Read speechandtextanalytics dictionaryfeedback %s %s", d.Id(), *dictionaryFeedback.Name)
		return cc.CheckState(d)
	})
}

// updateSpeechandtextanalyticsDictionaryfeedback is used by the speechandtextanalytics_dictionaryfeedback resource to update an speechandtextanalytics dictionaryfeedback in Genesys Cloud
func updateSpeechandtextanalyticsDictionaryfeedback(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsDictionaryfeedbackProxy(sdkConfig)

	speechandtextanalyticsDictionaryfeedback := getSpeechandtextanalyticsDictionaryfeedbackFromResourceData(d)

	log.Printf("Updating speechandtextanalytics dictionaryfeedback %s", *speechandtextanalyticsDictionaryfeedback.Name)
	dictionaryFeedback, resp, err := proxy.updateSpeechandtextanalyticsDictionaryfeedback(ctx, d.Id(), &speechandtextanalyticsDictionaryfeedback)
	if err != nil {
		return util.BuildAPIDiagnosticError(resourceName, fmt.Sprintf("Failed to update speechandtextanalytics dictionaryfeedback %s: %s", d.Id(), err), resp)
	}

	log.Printf("Updated speechandtextanalytics dictionaryfeedback %s", *dictionaryFeedback.Id)
	return readSpeechandtextanalyticsDictionaryfeedback(ctx, d, meta)
}

// deleteSpeechandtextanalyticsDictionaryfeedback is used by the speechandtextanalytics_dictionaryfeedback resource to delete an speechandtextanalytics dictionaryfeedback from Genesys cloud
func deleteSpeechandtextanalyticsDictionaryfeedback(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*provider.ProviderMeta).ClientConfig
	proxy := getSpeechandtextanalyticsDictionaryfeedbackProxy(sdkConfig)

	resp, err := proxy.deleteSpeechandtextanalyticsDictionaryfeedback(ctx, d.Id())
	if err != nil {
		return util.BuildAPIDiagnosticError(resourceName, fmt.Sprintf("Failed to delete speechandtextanalytics dictionaryfeedback %s: %s", d.Id(), err), resp)
	}

	return util.WithRetries(ctx, 180*time.Second, func() *retry.RetryError {
		_, resp, err := proxy.getSpeechandtextanalyticsDictionaryfeedbackById(ctx, d.Id())

		if err != nil {
			if util.IsStatus404(resp) {
				log.Printf("Deleted speechandtextanalytics dictionaryfeedback %s", d.Id())
				return nil
			}
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("Error deleting speechandtextanalytics dictionaryfeedback %s: %s", d.Id(), err), resp))
		}

		return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError(resourceName, fmt.Sprintf("speechandtextanalytics dictionaryfeedback %s still exists", d.Id()), resp))
	})
}
