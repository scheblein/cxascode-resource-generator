package speechandtextanalytics_sentimentfeedback

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
	"log"
)

/*
The genesyscloud_speechandtextanalytics_sentimentfeedback_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *speechandtextanalyticsSentimentfeedbackProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createSpeechandtextanalyticsSentimentfeedbackFunc func(ctx context.Context, p *speechandtextanalyticsSentimentfeedbackProxy, sentimentFeedback *platformclientv2.Sentimentfeedback) (*platformclientv2.Sentimentfeedback, *platformclientv2.APIResponse, error)
type getAllSpeechandtextanalyticsSentimentfeedbackFunc func(ctx context.Context, p *speechandtextanalyticsSentimentfeedbackProxy) (*[]platformclientv2.Sentimentfeedback, *platformclientv2.APIResponse, error)
type getSpeechandtextanalyticsSentimentfeedbackIdByNameFunc func(ctx context.Context, p *speechandtextanalyticsSentimentfeedbackProxy, name string) (string, *platformclientv2.APIResponse, bool, error)
type deleteSpeechandtextanalyticsSentimentfeedbackFunc func(ctx context.Context, p *speechandtextanalyticsSentimentfeedbackProxy, id string) (*platformclientv2.APIResponse, error)

// speechandtextanalyticsSentimentfeedbackProxy contains all of the methods that call genesys cloud APIs.
type speechandtextanalyticsSentimentfeedbackProxy struct {
	clientConfig                                           *platformclientv2.Configuration
	speechTextAnalyticsApi                                 *platformclientv2.SpeechTextAnalyticsApi
	createSpeechandtextanalyticsSentimentfeedbackAttr      createSpeechandtextanalyticsSentimentfeedbackFunc
	getAllSpeechandtextanalyticsSentimentfeedbackAttr      getAllSpeechandtextanalyticsSentimentfeedbackFunc
	getSpeechandtextanalyticsSentimentfeedbackIdByNameAttr getSpeechandtextanalyticsSentimentfeedbackIdByNameFunc
	deleteSpeechandtextanalyticsSentimentfeedbackAttr      deleteSpeechandtextanalyticsSentimentfeedbackFunc
}

// newSpeechandtextanalyticsSentimentfeedbackProxy initializes the speechandtextanalytics sentimentfeedback proxy with all of the data needed to communicate with Genesys Cloud
func newSpeechandtextanalyticsSentimentfeedbackProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsSentimentfeedbackProxy {
	api := platformclientv2.NewSpeechTextAnalyticsApiWithConfig(clientConfig)
	return &speechandtextanalyticsSentimentfeedbackProxy{
		clientConfig:           clientConfig,
		speechTextAnalyticsApi: api,
		createSpeechandtextanalyticsSentimentfeedbackAttr:      createSpeechandtextanalyticsSentimentfeedbackFn,
		getAllSpeechandtextanalyticsSentimentfeedbackAttr:      getAllSpeechandtextanalyticsSentimentfeedbackFn,
		getSpeechandtextanalyticsSentimentfeedbackIdByNameAttr: getSpeechandtextanalyticsSentimentfeedbackIdByNameFn,
		deleteSpeechandtextanalyticsSentimentfeedbackAttr:      deleteSpeechandtextanalyticsSentimentfeedbackFn,
	}
}

// getSpeechandtextanalyticsSentimentfeedbackProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getSpeechandtextanalyticsSentimentfeedbackProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsSentimentfeedbackProxy {
	if internalProxy == nil {
		internalProxy = newSpeechandtextanalyticsSentimentfeedbackProxy(clientConfig)
	}

	return internalProxy
}

// createSpeechandtextanalyticsSentimentfeedback creates a Genesys Cloud speechandtextanalytics sentimentfeedback
func (p *speechandtextanalyticsSentimentfeedbackProxy) createSpeechandtextanalyticsSentimentfeedback(ctx context.Context, speechandtextanalyticsSentimentfeedback *platformclientv2.Sentimentfeedback) (*platformclientv2.Sentimentfeedback, *platformclientv2.APIResponse, error) {
	return p.createSpeechandtextanalyticsSentimentfeedbackAttr(ctx, p, speechandtextanalyticsSentimentfeedback)
}

// getSpeechandtextanalyticsSentimentfeedback retrieves all Genesys Cloud speechandtextanalytics sentimentfeedback
func (p *speechandtextanalyticsSentimentfeedbackProxy) getAllSpeechandtextanalyticsSentimentfeedback(ctx context.Context) (*[]platformclientv2.Sentimentfeedback, *platformclientv2.APIResponse, error) {
	return p.getAllSpeechandtextanalyticsSentimentfeedbackAttr(ctx, p)
}

// getSpeechandtextanalyticsSentimentfeedbackIdByName returns a single Genesys Cloud speechandtextanalytics sentimentfeedback by a name
func (p *speechandtextanalyticsSentimentfeedbackProxy) getSpeechandtextanalyticsSentimentfeedbackIdByName(ctx context.Context, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return p.getSpeechandtextanalyticsSentimentfeedbackIdByNameAttr(ctx, p, name)
}

// deleteSpeechandtextanalyticsSentimentfeedback deletes a Genesys Cloud speechandtextanalytics sentimentfeedback by Id
func (p *speechandtextanalyticsSentimentfeedbackProxy) deleteSpeechandtextanalyticsSentimentfeedback(ctx context.Context, id string) (*platformclientv2.APIResponse, error) {
	return p.deleteSpeechandtextanalyticsSentimentfeedbackAttr(ctx, p, id)
}

// createSpeechandtextanalyticsSentimentfeedbackFn is an implementation function for creating a Genesys Cloud speechandtextanalytics sentimentfeedback
func createSpeechandtextanalyticsSentimentfeedbackFn(ctx context.Context, p *speechandtextanalyticsSentimentfeedbackProxy, speechandtextanalyticsSentimentfeedback *platformclientv2.Sentimentfeedback) (*platformclientv2.Sentimentfeedback, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PostSpeechandtextanalyticsSentimentfeedback(*speechandtextanalyticsSentimentfeedback)
}

// getAllSpeechandtextanalyticsSentimentfeedbackFn is the implementation for retrieving all speechandtextanalytics sentimentfeedback in Genesys Cloud
func getAllSpeechandtextanalyticsSentimentfeedbackFn(ctx context.Context, p *speechandtextanalyticsSentimentfeedbackProxy) (*[]platformclientv2.Sentimentfeedback, *platformclientv2.APIResponse, error) {
	var allSentimentFeedbacks []platformclientv2.Sentimentfeedback
	const pageSize = 100

	sentimentFeedbacks, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsSentimentfeedback()
	if err != nil {
		return nil, resp, err
	}
	if sentimentFeedbacks.Entities == nil || len(*sentimentFeedbacks.Entities) == 0 {
		return &allSentimentFeedbacks, resp, nil
	}
	for _, sentimentFeedback := range *sentimentFeedbacks.Entities {
		allSentimentFeedbacks = append(allSentimentFeedbacks, sentimentFeedback)
	}

	for pageNum := 2; pageNum <= *sentimentFeedbacks.PageCount; pageNum++ {
		sentimentFeedbacks, _, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsSentimentfeedback()
		if err != nil {
			return nil, resp, err
		}

		if sentimentFeedbacks.Entities == nil || len(*sentimentFeedbacks.Entities) == 0 {
			break
		}

		for _, sentimentFeedback := range *sentimentFeedbacks.Entities {
			allSentimentFeedbacks = append(allSentimentFeedbacks, sentimentFeedback)
		}
	}

	return &allSentimentFeedbacks, resp, nil
}

// getSpeechandtextanalyticsSentimentfeedbackIdByNameFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics sentimentfeedback by name
func getSpeechandtextanalyticsSentimentfeedbackIdByNameFn(ctx context.Context, p *speechandtextanalyticsSentimentfeedbackProxy, name string) (string, *platformclientv2.APIResponse, bool, error) {
	sentimentFeedbacks, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsSentimentfeedback()
	if err != nil {
		return "", resp, false, err
	}

	if sentimentFeedbacks.Entities == nil || len(*sentimentFeedbacks.Entities) == 0 {
		return "", resp, true, err
	}

	for _, sentimentFeedback := range *sentimentFeedbacks.Entities {
		if *sentimentFeedback.Name == name {
			log.Printf("Retrieved the speechandtextanalytics sentimentfeedback id %s by name %s", *sentimentFeedback.Id, name)
			return *sentimentFeedback.Id, resp, false, nil
		}
	}

	return "", resp, true, fmt.Errorf("Unable to find speechandtextanalytics sentimentfeedback with name %s", name)
}

// deleteSpeechandtextanalyticsSentimentfeedbackFn is an implementation function for deleting a Genesys Cloud speechandtextanalytics sentimentfeedback
func deleteSpeechandtextanalyticsSentimentfeedbackFn(ctx context.Context, p *speechandtextanalyticsSentimentfeedbackProxy, id string) (*platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.DeleteSpeechandtextanalyticsSentimentfeedbackSentimentFeedbackId(id)
}
