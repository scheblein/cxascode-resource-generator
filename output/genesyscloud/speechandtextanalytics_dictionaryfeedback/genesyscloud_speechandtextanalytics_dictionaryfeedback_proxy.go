package speechandtextanalytics_dictionaryfeedback

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
	"log"
)

/*
The genesyscloud_speechandtextanalytics_dictionaryfeedback_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *speechandtextanalyticsDictionaryfeedbackProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createSpeechandtextanalyticsDictionaryfeedbackFunc func(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, dictionaryFeedback *platformclientv2.Dictionaryfeedback) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error)
type getAllSpeechandtextanalyticsDictionaryfeedbackFunc func(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy) (*[]platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error)
type getSpeechandtextanalyticsDictionaryfeedbackIdByNameFunc func(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, name string) (string, *platformclientv2.APIResponse, bool, error)
type getSpeechandtextanalyticsDictionaryfeedbackByIdFunc func(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, id string) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error)
type updateSpeechandtextanalyticsDictionaryfeedbackFunc func(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, id string, dictionaryFeedback *platformclientv2.Dictionaryfeedback) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error)
type deleteSpeechandtextanalyticsDictionaryfeedbackFunc func(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, id string) (*platformclientv2.APIResponse, error)

// speechandtextanalyticsDictionaryfeedbackProxy contains all of the methods that call genesys cloud APIs.
type speechandtextanalyticsDictionaryfeedbackProxy struct {
	clientConfig                                            *platformclientv2.Configuration
	speechTextAnalyticsApi                                  *platformclientv2.SpeechTextAnalyticsApi
	createSpeechandtextanalyticsDictionaryfeedbackAttr      createSpeechandtextanalyticsDictionaryfeedbackFunc
	getAllSpeechandtextanalyticsDictionaryfeedbackAttr      getAllSpeechandtextanalyticsDictionaryfeedbackFunc
	getSpeechandtextanalyticsDictionaryfeedbackIdByNameAttr getSpeechandtextanalyticsDictionaryfeedbackIdByNameFunc
	getSpeechandtextanalyticsDictionaryfeedbackByIdAttr     getSpeechandtextanalyticsDictionaryfeedbackByIdFunc
	updateSpeechandtextanalyticsDictionaryfeedbackAttr      updateSpeechandtextanalyticsDictionaryfeedbackFunc
	deleteSpeechandtextanalyticsDictionaryfeedbackAttr      deleteSpeechandtextanalyticsDictionaryfeedbackFunc
}

// newSpeechandtextanalyticsDictionaryfeedbackProxy initializes the speechandtextanalytics dictionaryfeedback proxy with all of the data needed to communicate with Genesys Cloud
func newSpeechandtextanalyticsDictionaryfeedbackProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsDictionaryfeedbackProxy {
	api := platformclientv2.NewSpeechTextAnalyticsApiWithConfig(clientConfig)
	return &speechandtextanalyticsDictionaryfeedbackProxy{
		clientConfig:           clientConfig,
		speechTextAnalyticsApi: api,
		createSpeechandtextanalyticsDictionaryfeedbackAttr:      createSpeechandtextanalyticsDictionaryfeedbackFn,
		getAllSpeechandtextanalyticsDictionaryfeedbackAttr:      getAllSpeechandtextanalyticsDictionaryfeedbackFn,
		getSpeechandtextanalyticsDictionaryfeedbackIdByNameAttr: getSpeechandtextanalyticsDictionaryfeedbackIdByNameFn,
		getSpeechandtextanalyticsDictionaryfeedbackByIdAttr:     getSpeechandtextanalyticsDictionaryfeedbackByIdFn,
		updateSpeechandtextanalyticsDictionaryfeedbackAttr:      updateSpeechandtextanalyticsDictionaryfeedbackFn,
		deleteSpeechandtextanalyticsDictionaryfeedbackAttr:      deleteSpeechandtextanalyticsDictionaryfeedbackFn,
	}
}

// getSpeechandtextanalyticsDictionaryfeedbackProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getSpeechandtextanalyticsDictionaryfeedbackProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsDictionaryfeedbackProxy {
	if internalProxy == nil {
		internalProxy = newSpeechandtextanalyticsDictionaryfeedbackProxy(clientConfig)
	}

	return internalProxy
}

// createSpeechandtextanalyticsDictionaryfeedback creates a Genesys Cloud speechandtextanalytics dictionaryfeedback
func (p *speechandtextanalyticsDictionaryfeedbackProxy) createSpeechandtextanalyticsDictionaryfeedback(ctx context.Context, speechandtextanalyticsDictionaryfeedback *platformclientv2.Dictionaryfeedback) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error) {
	return p.createSpeechandtextanalyticsDictionaryfeedbackAttr(ctx, p, speechandtextanalyticsDictionaryfeedback)
}

// getSpeechandtextanalyticsDictionaryfeedback retrieves all Genesys Cloud speechandtextanalytics dictionaryfeedback
func (p *speechandtextanalyticsDictionaryfeedbackProxy) getAllSpeechandtextanalyticsDictionaryfeedback(ctx context.Context) (*[]platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error) {
	return p.getAllSpeechandtextanalyticsDictionaryfeedbackAttr(ctx, p)
}

// getSpeechandtextanalyticsDictionaryfeedbackIdByName returns a single Genesys Cloud speechandtextanalytics dictionaryfeedback by a name
func (p *speechandtextanalyticsDictionaryfeedbackProxy) getSpeechandtextanalyticsDictionaryfeedbackIdByName(ctx context.Context, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return p.getSpeechandtextanalyticsDictionaryfeedbackIdByNameAttr(ctx, p, name)
}

// getSpeechandtextanalyticsDictionaryfeedbackById returns a single Genesys Cloud speechandtextanalytics dictionaryfeedback by Id
func (p *speechandtextanalyticsDictionaryfeedbackProxy) getSpeechandtextanalyticsDictionaryfeedbackById(ctx context.Context, id string) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error) {
	return p.getSpeechandtextanalyticsDictionaryfeedbackByIdAttr(ctx, p, id)
}

// updateSpeechandtextanalyticsDictionaryfeedback updates a Genesys Cloud speechandtextanalytics dictionaryfeedback
func (p *speechandtextanalyticsDictionaryfeedbackProxy) updateSpeechandtextanalyticsDictionaryfeedback(ctx context.Context, id string, speechandtextanalyticsDictionaryfeedback *platformclientv2.Dictionaryfeedback) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error) {
	return p.updateSpeechandtextanalyticsDictionaryfeedbackAttr(ctx, p, id, speechandtextanalyticsDictionaryfeedback)
}

// deleteSpeechandtextanalyticsDictionaryfeedback deletes a Genesys Cloud speechandtextanalytics dictionaryfeedback by Id
func (p *speechandtextanalyticsDictionaryfeedbackProxy) deleteSpeechandtextanalyticsDictionaryfeedback(ctx context.Context, id string) (*platformclientv2.APIResponse, error) {
	return p.deleteSpeechandtextanalyticsDictionaryfeedbackAttr(ctx, p, id)
}

// createSpeechandtextanalyticsDictionaryfeedbackFn is an implementation function for creating a Genesys Cloud speechandtextanalytics dictionaryfeedback
func createSpeechandtextanalyticsDictionaryfeedbackFn(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, speechandtextanalyticsDictionaryfeedback *platformclientv2.Dictionaryfeedback) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PostSpeechandtextanalyticsDictionaryfeedback(*speechandtextanalyticsDictionaryfeedback)
}

// getAllSpeechandtextanalyticsDictionaryfeedbackFn is the implementation for retrieving all speechandtextanalytics dictionaryfeedback in Genesys Cloud
func getAllSpeechandtextanalyticsDictionaryfeedbackFn(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy) (*[]platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error) {
	var allDictionaryFeedbacks []platformclientv2.Dictionaryfeedback
	const pageSize = 100

	dictionaryFeedbacks, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsDictionaryfeedback()
	if err != nil {
		return nil, resp, err
	}
	if dictionaryFeedbacks.Entities == nil || len(*dictionaryFeedbacks.Entities) == 0 {
		return &allDictionaryFeedbacks, resp, nil
	}
	for _, dictionaryFeedback := range *dictionaryFeedbacks.Entities {
		allDictionaryFeedbacks = append(allDictionaryFeedbacks, dictionaryFeedback)
	}

	for pageNum := 2; pageNum <= *dictionaryFeedbacks.PageCount; pageNum++ {
		dictionaryFeedbacks, _, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsDictionaryfeedback()
		if err != nil {
			return nil, resp, err
		}

		if dictionaryFeedbacks.Entities == nil || len(*dictionaryFeedbacks.Entities) == 0 {
			break
		}

		for _, dictionaryFeedback := range *dictionaryFeedbacks.Entities {
			allDictionaryFeedbacks = append(allDictionaryFeedbacks, dictionaryFeedback)
		}
	}

	return &allDictionaryFeedbacks, resp, nil
}

// getSpeechandtextanalyticsDictionaryfeedbackIdByNameFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics dictionaryfeedback by name
func getSpeechandtextanalyticsDictionaryfeedbackIdByNameFn(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, name string) (string, *platformclientv2.APIResponse, bool, error) {
	dictionaryFeedbacks, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsDictionaryfeedback()
	if err != nil {
		return "", resp, false, err
	}

	if dictionaryFeedbacks.Entities == nil || len(*dictionaryFeedbacks.Entities) == 0 {
		return "", resp, true, err
	}

	for _, dictionaryFeedback := range *dictionaryFeedbacks.Entities {
		if *dictionaryFeedback.Name == name {
			log.Printf("Retrieved the speechandtextanalytics dictionaryfeedback id %s by name %s", *dictionaryFeedback.Id, name)
			return *dictionaryFeedback.Id, resp, false, nil
		}
	}

	return "", resp, true, fmt.Errorf("Unable to find speechandtextanalytics dictionaryfeedback with name %s", name)
}

// getSpeechandtextanalyticsDictionaryfeedbackByIdFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics dictionaryfeedback by Id
func getSpeechandtextanalyticsDictionaryfeedbackByIdFn(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, id string) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.GetSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId(id)
}

// updateSpeechandtextanalyticsDictionaryfeedbackFn is an implementation of the function to update a Genesys Cloud speechandtextanalytics dictionaryfeedback
func updateSpeechandtextanalyticsDictionaryfeedbackFn(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, id string, speechandtextanalyticsDictionaryfeedback *platformclientv2.Dictionaryfeedback) (*platformclientv2.Dictionaryfeedback, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PutSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId(id, *speechandtextanalyticsDictionaryfeedback)
}

// deleteSpeechandtextanalyticsDictionaryfeedbackFn is an implementation function for deleting a Genesys Cloud speechandtextanalytics dictionaryfeedback
func deleteSpeechandtextanalyticsDictionaryfeedbackFn(ctx context.Context, p *speechandtextanalyticsDictionaryfeedbackProxy, id string) (*platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.DeleteSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId(id)
}
