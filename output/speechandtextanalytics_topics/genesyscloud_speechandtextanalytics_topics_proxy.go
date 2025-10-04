package speechandtextanalytics_topics

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
	"log"
)

/*
The genesyscloud_speechandtextanalytics_topics_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *speechandtextanalyticsTopicsProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createSpeechandtextanalyticsTopicsFunc func(ctx context.Context, p *speechandtextanalyticsTopicsProxy, topic *platformclientv2.Topic) (*platformclientv2.Topic, *platformclientv2.APIResponse, error)
type getAllSpeechandtextanalyticsTopicsFunc func(ctx context.Context, p *speechandtextanalyticsTopicsProxy) (*[]platformclientv2.Topic, *platformclientv2.APIResponse, error)
type getSpeechandtextanalyticsTopicsIdByNameFunc func(ctx context.Context, p *speechandtextanalyticsTopicsProxy, name string) (string, *platformclientv2.APIResponse, bool, error)
type getSpeechandtextanalyticsTopicsByIdFunc func(ctx context.Context, p *speechandtextanalyticsTopicsProxy, id string) (*platformclientv2.Topic, *platformclientv2.APIResponse, error)
type updateSpeechandtextanalyticsTopicsFunc func(ctx context.Context, p *speechandtextanalyticsTopicsProxy, id string, topic *platformclientv2.Topic) (*platformclientv2.Topic, *platformclientv2.APIResponse, error)
type deleteSpeechandtextanalyticsTopicsFunc func(ctx context.Context, p *speechandtextanalyticsTopicsProxy, id string) (*platformclientv2.APIResponse, error)

// speechandtextanalyticsTopicsProxy contains all of the methods that call genesys cloud APIs.
type speechandtextanalyticsTopicsProxy struct {
	clientConfig                                *platformclientv2.Configuration
	speechTextAnalyticsApi                      *platformclientv2.SpeechTextAnalyticsApi
	createSpeechandtextanalyticsTopicsAttr      createSpeechandtextanalyticsTopicsFunc
	getAllSpeechandtextanalyticsTopicsAttr      getAllSpeechandtextanalyticsTopicsFunc
	getSpeechandtextanalyticsTopicsIdByNameAttr getSpeechandtextanalyticsTopicsIdByNameFunc
	getSpeechandtextanalyticsTopicsByIdAttr     getSpeechandtextanalyticsTopicsByIdFunc
	updateSpeechandtextanalyticsTopicsAttr      updateSpeechandtextanalyticsTopicsFunc
	deleteSpeechandtextanalyticsTopicsAttr      deleteSpeechandtextanalyticsTopicsFunc
}

// newSpeechandtextanalyticsTopicsProxy initializes the speechandtextanalytics topics proxy with all of the data needed to communicate with Genesys Cloud
func newSpeechandtextanalyticsTopicsProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsTopicsProxy {
	api := platformclientv2.NewSpeechTextAnalyticsApiWithConfig(clientConfig)
	return &speechandtextanalyticsTopicsProxy{
		clientConfig:                                clientConfig,
		speechTextAnalyticsApi:                      api,
		createSpeechandtextanalyticsTopicsAttr:      createSpeechandtextanalyticsTopicsFn,
		getAllSpeechandtextanalyticsTopicsAttr:      getAllSpeechandtextanalyticsTopicsFn,
		getSpeechandtextanalyticsTopicsIdByNameAttr: getSpeechandtextanalyticsTopicsIdByNameFn,
		getSpeechandtextanalyticsTopicsByIdAttr:     getSpeechandtextanalyticsTopicsByIdFn,
		updateSpeechandtextanalyticsTopicsAttr:      updateSpeechandtextanalyticsTopicsFn,
		deleteSpeechandtextanalyticsTopicsAttr:      deleteSpeechandtextanalyticsTopicsFn,
	}
}

// getSpeechandtextanalyticsTopicsProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getSpeechandtextanalyticsTopicsProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsTopicsProxy {
	if internalProxy == nil {
		internalProxy = newSpeechandtextanalyticsTopicsProxy(clientConfig)
	}

	return internalProxy
}

// createSpeechandtextanalyticsTopics creates a Genesys Cloud speechandtextanalytics topics
func (p *speechandtextanalyticsTopicsProxy) createSpeechandtextanalyticsTopics(ctx context.Context, speechandtextanalyticsTopics *platformclientv2.Topic) (*platformclientv2.Topic, *platformclientv2.APIResponse, error) {
	return p.createSpeechandtextanalyticsTopicsAttr(ctx, p, speechandtextanalyticsTopics)
}

// getSpeechandtextanalyticsTopics retrieves all Genesys Cloud speechandtextanalytics topics
func (p *speechandtextanalyticsTopicsProxy) getAllSpeechandtextanalyticsTopics(ctx context.Context) (*[]platformclientv2.Topic, *platformclientv2.APIResponse, error) {
	return p.getAllSpeechandtextanalyticsTopicsAttr(ctx, p)
}

// getSpeechandtextanalyticsTopicsIdByName returns a single Genesys Cloud speechandtextanalytics topics by a name
func (p *speechandtextanalyticsTopicsProxy) getSpeechandtextanalyticsTopicsIdByName(ctx context.Context, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return p.getSpeechandtextanalyticsTopicsIdByNameAttr(ctx, p, name)
}

// getSpeechandtextanalyticsTopicsById returns a single Genesys Cloud speechandtextanalytics topics by Id
func (p *speechandtextanalyticsTopicsProxy) getSpeechandtextanalyticsTopicsById(ctx context.Context, id string) (*platformclientv2.Topic, *platformclientv2.APIResponse, error) {
	return p.getSpeechandtextanalyticsTopicsByIdAttr(ctx, p, id)
}

// updateSpeechandtextanalyticsTopics updates a Genesys Cloud speechandtextanalytics topics
func (p *speechandtextanalyticsTopicsProxy) updateSpeechandtextanalyticsTopics(ctx context.Context, id string, speechandtextanalyticsTopics *platformclientv2.Topic) (*platformclientv2.Topic, *platformclientv2.APIResponse, error) {
	return p.updateSpeechandtextanalyticsTopicsAttr(ctx, p, id, speechandtextanalyticsTopics)
}

// deleteSpeechandtextanalyticsTopics deletes a Genesys Cloud speechandtextanalytics topics by Id
func (p *speechandtextanalyticsTopicsProxy) deleteSpeechandtextanalyticsTopics(ctx context.Context, id string) (*platformclientv2.APIResponse, error) {
	return p.deleteSpeechandtextanalyticsTopicsAttr(ctx, p, id)
}

// createSpeechandtextanalyticsTopicsFn is an implementation function for creating a Genesys Cloud speechandtextanalytics topics
func createSpeechandtextanalyticsTopicsFn(ctx context.Context, p *speechandtextanalyticsTopicsProxy, speechandtextanalyticsTopics *platformclientv2.Topic) (*platformclientv2.Topic, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PostSpeechandtextanalyticsTopics(*speechandtextanalyticsTopics)
}

// getAllSpeechandtextanalyticsTopicsFn is the implementation for retrieving all speechandtextanalytics topics in Genesys Cloud
func getAllSpeechandtextanalyticsTopicsFn(ctx context.Context, p *speechandtextanalyticsTopicsProxy) (*[]platformclientv2.Topic, *platformclientv2.APIResponse, error) {
	var allTopics []platformclientv2.Topic
	const pageSize = 100

	topics, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsTopics()
	if err != nil {
		return nil, resp, err
	}
	if topics.Entities == nil || len(*topics.Entities) == 0 {
		return &allTopics, resp, nil
	}
	for _, topic := range *topics.Entities {
		allTopics = append(allTopics, topic)
	}

	for pageNum := 2; pageNum <= *topics.PageCount; pageNum++ {
		topics, _, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsTopics()
		if err != nil {
			return nil, resp, err
		}

		if topics.Entities == nil || len(*topics.Entities) == 0 {
			break
		}

		for _, topic := range *topics.Entities {
			allTopics = append(allTopics, topic)
		}
	}

	return &allTopics, resp, nil
}

// getSpeechandtextanalyticsTopicsIdByNameFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics topics by name
func getSpeechandtextanalyticsTopicsIdByNameFn(ctx context.Context, p *speechandtextanalyticsTopicsProxy, name string) (string, *platformclientv2.APIResponse, bool, error) {
	topics, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsTopics()
	if err != nil {
		return "", resp, false, err
	}

	if topics.Entities == nil || len(*topics.Entities) == 0 {
		return "", resp, true, err
	}

	for _, topic := range *topics.Entities {
		if *topic.Name == name {
			log.Printf("Retrieved the speechandtextanalytics topics id %s by name %s", *topic.Id, name)
			return *topic.Id, resp, false, nil
		}
	}

	return "", resp, true, fmt.Errorf("Unable to find speechandtextanalytics topics with name %s", name)
}

// getSpeechandtextanalyticsTopicsByIdFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics topics by Id
func getSpeechandtextanalyticsTopicsByIdFn(ctx context.Context, p *speechandtextanalyticsTopicsProxy, id string) (*platformclientv2.Topic, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.GetSpeechandtextanalyticsTopic(id)
}

// updateSpeechandtextanalyticsTopicsFn is an implementation of the function to update a Genesys Cloud speechandtextanalytics topics
func updateSpeechandtextanalyticsTopicsFn(ctx context.Context, p *speechandtextanalyticsTopicsProxy, id string, speechandtextanalyticsTopics *platformclientv2.Topic) (*platformclientv2.Topic, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PutSpeechandtextanalyticsTopic(id, *speechandtextanalyticsTopics)
}

// deleteSpeechandtextanalyticsTopicsFn is an implementation function for deleting a Genesys Cloud speechandtextanalytics topics
func deleteSpeechandtextanalyticsTopicsFn(ctx context.Context, p *speechandtextanalyticsTopicsProxy, id string) (*platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.DeleteSpeechandtextanalyticsTopic(id)
}
