package speechandtextanalytics_categories

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
	"log"
)

/*
The genesyscloud_speechandtextanalytics_categories_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *speechandtextanalyticsCategoriesProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createSpeechandtextanalyticsCategoriesFunc func(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, staCategory *platformclientv2.Stacategory) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error)
type getAllSpeechandtextanalyticsCategoriesFunc func(ctx context.Context, p *speechandtextanalyticsCategoriesProxy) (*[]platformclientv2.Stacategory, *platformclientv2.APIResponse, error)
type getSpeechandtextanalyticsCategoriesIdByNameFunc func(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, name string) (string, *platformclientv2.APIResponse, bool, error)
type getSpeechandtextanalyticsCategoriesByIdFunc func(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, id string) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error)
type updateSpeechandtextanalyticsCategoriesFunc func(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, id string, staCategory *platformclientv2.Stacategory) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error)
type deleteSpeechandtextanalyticsCategoriesFunc func(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, id string) (*platformclientv2.APIResponse, error)

// speechandtextanalyticsCategoriesProxy contains all of the methods that call genesys cloud APIs.
type speechandtextanalyticsCategoriesProxy struct {
	clientConfig                                    *platformclientv2.Configuration
	speechTextAnalyticsApi                          *platformclientv2.SpeechTextAnalyticsApi
	createSpeechandtextanalyticsCategoriesAttr      createSpeechandtextanalyticsCategoriesFunc
	getAllSpeechandtextanalyticsCategoriesAttr      getAllSpeechandtextanalyticsCategoriesFunc
	getSpeechandtextanalyticsCategoriesIdByNameAttr getSpeechandtextanalyticsCategoriesIdByNameFunc
	getSpeechandtextanalyticsCategoriesByIdAttr     getSpeechandtextanalyticsCategoriesByIdFunc
	updateSpeechandtextanalyticsCategoriesAttr      updateSpeechandtextanalyticsCategoriesFunc
	deleteSpeechandtextanalyticsCategoriesAttr      deleteSpeechandtextanalyticsCategoriesFunc
}

// newSpeechandtextanalyticsCategoriesProxy initializes the speechandtextanalytics categories proxy with all of the data needed to communicate with Genesys Cloud
func newSpeechandtextanalyticsCategoriesProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsCategoriesProxy {
	api := platformclientv2.NewSpeechTextAnalyticsApiWithConfig(clientConfig)
	return &speechandtextanalyticsCategoriesProxy{
		clientConfig:           clientConfig,
		speechTextAnalyticsApi: api,
		createSpeechandtextanalyticsCategoriesAttr:      createSpeechandtextanalyticsCategoriesFn,
		getAllSpeechandtextanalyticsCategoriesAttr:      getAllSpeechandtextanalyticsCategoriesFn,
		getSpeechandtextanalyticsCategoriesIdByNameAttr: getSpeechandtextanalyticsCategoriesIdByNameFn,
		getSpeechandtextanalyticsCategoriesByIdAttr:     getSpeechandtextanalyticsCategoriesByIdFn,
		updateSpeechandtextanalyticsCategoriesAttr:      updateSpeechandtextanalyticsCategoriesFn,
		deleteSpeechandtextanalyticsCategoriesAttr:      deleteSpeechandtextanalyticsCategoriesFn,
	}
}

// getSpeechandtextanalyticsCategoriesProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getSpeechandtextanalyticsCategoriesProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsCategoriesProxy {
	if internalProxy == nil {
		internalProxy = newSpeechandtextanalyticsCategoriesProxy(clientConfig)
	}

	return internalProxy
}

// createSpeechandtextanalyticsCategories creates a Genesys Cloud speechandtextanalytics categories
func (p *speechandtextanalyticsCategoriesProxy) createSpeechandtextanalyticsCategories(ctx context.Context, speechandtextanalyticsCategories *platformclientv2.Stacategory) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error) {
	return p.createSpeechandtextanalyticsCategoriesAttr(ctx, p, speechandtextanalyticsCategories)
}

// getSpeechandtextanalyticsCategories retrieves all Genesys Cloud speechandtextanalytics categories
func (p *speechandtextanalyticsCategoriesProxy) getAllSpeechandtextanalyticsCategories(ctx context.Context) (*[]platformclientv2.Stacategory, *platformclientv2.APIResponse, error) {
	return p.getAllSpeechandtextanalyticsCategoriesAttr(ctx, p)
}

// getSpeechandtextanalyticsCategoriesIdByName returns a single Genesys Cloud speechandtextanalytics categories by a name
func (p *speechandtextanalyticsCategoriesProxy) getSpeechandtextanalyticsCategoriesIdByName(ctx context.Context, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return p.getSpeechandtextanalyticsCategoriesIdByNameAttr(ctx, p, name)
}

// getSpeechandtextanalyticsCategoriesById returns a single Genesys Cloud speechandtextanalytics categories by Id
func (p *speechandtextanalyticsCategoriesProxy) getSpeechandtextanalyticsCategoriesById(ctx context.Context, id string) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error) {
	return p.getSpeechandtextanalyticsCategoriesByIdAttr(ctx, p, id)
}

// updateSpeechandtextanalyticsCategories updates a Genesys Cloud speechandtextanalytics categories
func (p *speechandtextanalyticsCategoriesProxy) updateSpeechandtextanalyticsCategories(ctx context.Context, id string, speechandtextanalyticsCategories *platformclientv2.Stacategory) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error) {
	return p.updateSpeechandtextanalyticsCategoriesAttr(ctx, p, id, speechandtextanalyticsCategories)
}

// deleteSpeechandtextanalyticsCategories deletes a Genesys Cloud speechandtextanalytics categories by Id
func (p *speechandtextanalyticsCategoriesProxy) deleteSpeechandtextanalyticsCategories(ctx context.Context, id string) (*platformclientv2.APIResponse, error) {
	return p.deleteSpeechandtextanalyticsCategoriesAttr(ctx, p, id)
}

// createSpeechandtextanalyticsCategoriesFn is an implementation function for creating a Genesys Cloud speechandtextanalytics categories
func createSpeechandtextanalyticsCategoriesFn(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, speechandtextanalyticsCategories *platformclientv2.Stacategory) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PostSpeechandtextanalyticsCategories(*speechandtextanalyticsCategories)
}

// getAllSpeechandtextanalyticsCategoriesFn is the implementation for retrieving all speechandtextanalytics categories in Genesys Cloud
func getAllSpeechandtextanalyticsCategoriesFn(ctx context.Context, p *speechandtextanalyticsCategoriesProxy) (*[]platformclientv2.Stacategory, *platformclientv2.APIResponse, error) {
	var allStaCategorys []platformclientv2.Stacategory
	const pageSize = 100

	staCategorys, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsCategories()
	if err != nil {
		return nil, resp, err
	}
	if staCategorys.Entities == nil || len(*staCategorys.Entities) == 0 {
		return &allStaCategorys, resp, nil
	}
	for _, staCategory := range *staCategorys.Entities {
		allStaCategorys = append(allStaCategorys, staCategory)
	}

	for pageNum := 2; pageNum <= *staCategorys.PageCount; pageNum++ {
		staCategorys, _, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsCategories()
		if err != nil {
			return nil, resp, err
		}

		if staCategorys.Entities == nil || len(*staCategorys.Entities) == 0 {
			break
		}

		for _, staCategory := range *staCategorys.Entities {
			allStaCategorys = append(allStaCategorys, staCategory)
		}
	}

	return &allStaCategorys, resp, nil
}

// getSpeechandtextanalyticsCategoriesIdByNameFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics categories by name
func getSpeechandtextanalyticsCategoriesIdByNameFn(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, name string) (string, *platformclientv2.APIResponse, bool, error) {
	staCategorys, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsCategories()
	if err != nil {
		return "", resp, false, err
	}

	if staCategorys.Entities == nil || len(*staCategorys.Entities) == 0 {
		return "", resp, true, err
	}

	for _, staCategory := range *staCategorys.Entities {
		if *staCategory.Name == name {
			log.Printf("Retrieved the speechandtextanalytics categories id %s by name %s", *staCategory.Id, name)
			return *staCategory.Id, resp, false, nil
		}
	}

	return "", resp, true, fmt.Errorf("Unable to find speechandtextanalytics categories with name %s", name)
}

// getSpeechandtextanalyticsCategoriesByIdFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics categories by Id
func getSpeechandtextanalyticsCategoriesByIdFn(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, id string) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.GetSpeechandtextanalyticsCategory(id)
}

// updateSpeechandtextanalyticsCategoriesFn is an implementation of the function to update a Genesys Cloud speechandtextanalytics categories
func updateSpeechandtextanalyticsCategoriesFn(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, id string, speechandtextanalyticsCategories *platformclientv2.Stacategory) (*platformclientv2.Stacategory, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PutSpeechandtextanalyticsCategory(id, *speechandtextanalyticsCategories)
}

// deleteSpeechandtextanalyticsCategoriesFn is an implementation function for deleting a Genesys Cloud speechandtextanalytics categories
func deleteSpeechandtextanalyticsCategoriesFn(ctx context.Context, p *speechandtextanalyticsCategoriesProxy, id string) (*platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.DeleteSpeechandtextanalyticsCategory(id)
}
