package languageunderstanding_miners

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
	"log"
)

/*
The genesyscloud_languageunderstanding_miners_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *languageunderstandingMinersProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createLanguageunderstandingMinersFunc func(ctx context.Context, p *languageunderstandingMinersProxy, miner *platformclientv2.Miner) (*platformclientv2.Miner, *platformclientv2.APIResponse, error)
type getAllLanguageunderstandingMinersFunc func(ctx context.Context, p *languageunderstandingMinersProxy) (*[]platformclientv2.Miner, *platformclientv2.APIResponse, error)
type getLanguageunderstandingMinersIdByNameFunc func(ctx context.Context, p *languageunderstandingMinersProxy, name string) (string, *platformclientv2.APIResponse, bool, error)
type getLanguageunderstandingMinersByIdFunc func(ctx context.Context, p *languageunderstandingMinersProxy, id string) (*platformclientv2.Miner, *platformclientv2.APIResponse, error)
type deleteLanguageunderstandingMinersFunc func(ctx context.Context, p *languageunderstandingMinersProxy, id string) (*platformclientv2.APIResponse, error)

// languageunderstandingMinersProxy contains all of the methods that call genesys cloud APIs.
type languageunderstandingMinersProxy struct {
	clientConfig                               *platformclientv2.Configuration
	languageUnderstandingApi                   *platformclientv2.LanguageUnderstandingApi
	createLanguageunderstandingMinersAttr      createLanguageunderstandingMinersFunc
	getAllLanguageunderstandingMinersAttr      getAllLanguageunderstandingMinersFunc
	getLanguageunderstandingMinersIdByNameAttr getLanguageunderstandingMinersIdByNameFunc
	getLanguageunderstandingMinersByIdAttr     getLanguageunderstandingMinersByIdFunc
	deleteLanguageunderstandingMinersAttr      deleteLanguageunderstandingMinersFunc
}

// newLanguageunderstandingMinersProxy initializes the languageunderstanding miners proxy with all of the data needed to communicate with Genesys Cloud
func newLanguageunderstandingMinersProxy(clientConfig *platformclientv2.Configuration) *languageunderstandingMinersProxy {
	api := platformclientv2.NewLanguageUnderstandingApiWithConfig(clientConfig)
	return &languageunderstandingMinersProxy{
		clientConfig:                               clientConfig,
		languageUnderstandingApi:                   api,
		createLanguageunderstandingMinersAttr:      createLanguageunderstandingMinersFn,
		getAllLanguageunderstandingMinersAttr:      getAllLanguageunderstandingMinersFn,
		getLanguageunderstandingMinersIdByNameAttr: getLanguageunderstandingMinersIdByNameFn,
		getLanguageunderstandingMinersByIdAttr:     getLanguageunderstandingMinersByIdFn,
		deleteLanguageunderstandingMinersAttr:      deleteLanguageunderstandingMinersFn,
	}
}

// getLanguageunderstandingMinersProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getLanguageunderstandingMinersProxy(clientConfig *platformclientv2.Configuration) *languageunderstandingMinersProxy {
	if internalProxy == nil {
		internalProxy = newLanguageunderstandingMinersProxy(clientConfig)
	}

	return internalProxy
}

// createLanguageunderstandingMiners creates a Genesys Cloud languageunderstanding miners
func (p *languageunderstandingMinersProxy) createLanguageunderstandingMiners(ctx context.Context, languageunderstandingMiners *platformclientv2.Miner) (*platformclientv2.Miner, *platformclientv2.APIResponse, error) {
	return p.createLanguageunderstandingMinersAttr(ctx, p, languageunderstandingMiners)
}

// getLanguageunderstandingMiners retrieves all Genesys Cloud languageunderstanding miners
func (p *languageunderstandingMinersProxy) getAllLanguageunderstandingMiners(ctx context.Context) (*[]platformclientv2.Miner, *platformclientv2.APIResponse, error) {
	return p.getAllLanguageunderstandingMinersAttr(ctx, p)
}

// getLanguageunderstandingMinersIdByName returns a single Genesys Cloud languageunderstanding miners by a name
func (p *languageunderstandingMinersProxy) getLanguageunderstandingMinersIdByName(ctx context.Context, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return p.getLanguageunderstandingMinersIdByNameAttr(ctx, p, name)
}

// getLanguageunderstandingMinersById returns a single Genesys Cloud languageunderstanding miners by Id
func (p *languageunderstandingMinersProxy) getLanguageunderstandingMinersById(ctx context.Context, id string) (*platformclientv2.Miner, *platformclientv2.APIResponse, error) {
	return p.getLanguageunderstandingMinersByIdAttr(ctx, p, id)
}

// deleteLanguageunderstandingMiners deletes a Genesys Cloud languageunderstanding miners by Id
func (p *languageunderstandingMinersProxy) deleteLanguageunderstandingMiners(ctx context.Context, id string) (*platformclientv2.APIResponse, error) {
	return p.deleteLanguageunderstandingMinersAttr(ctx, p, id)
}

// createLanguageunderstandingMinersFn is an implementation function for creating a Genesys Cloud languageunderstanding miners
func createLanguageunderstandingMinersFn(ctx context.Context, p *languageunderstandingMinersProxy, languageunderstandingMiners *platformclientv2.Miner) (*platformclientv2.Miner, *platformclientv2.APIResponse, error) {
	return p.languageUnderstandingApi.PostLanguageunderstandingMiners(*languageunderstandingMiners)
}

// getAllLanguageunderstandingMinersFn is the implementation for retrieving all languageunderstanding miners in Genesys Cloud
func getAllLanguageunderstandingMinersFn(ctx context.Context, p *languageunderstandingMinersProxy) (*[]platformclientv2.Miner, *platformclientv2.APIResponse, error) {
	var allMiners []platformclientv2.Miner
	const pageSize = 100

	miners, resp, err := p.languageUnderstandingApi.GetLanguageunderstandingMiners()
	if err != nil {
		return nil, resp, err
	}
	if miners.Entities == nil || len(*miners.Entities) == 0 {
		return &allMiners, resp, nil
	}
	for _, miner := range *miners.Entities {
		allMiners = append(allMiners, miner)
	}

	for pageNum := 2; pageNum <= *miners.PageCount; pageNum++ {
		miners, _, err := p.languageUnderstandingApi.GetLanguageunderstandingMiners()
		if err != nil {
			return nil, resp, err
		}

		if miners.Entities == nil || len(*miners.Entities) == 0 {
			break
		}

		for _, miner := range *miners.Entities {
			allMiners = append(allMiners, miner)
		}
	}

	return &allMiners, resp, nil
}

// getLanguageunderstandingMinersIdByNameFn is an implementation of the function to get a Genesys Cloud languageunderstanding miners by name
func getLanguageunderstandingMinersIdByNameFn(ctx context.Context, p *languageunderstandingMinersProxy, name string) (string, *platformclientv2.APIResponse, bool, error) {
	miners, resp, err := p.languageUnderstandingApi.GetLanguageunderstandingMiners()
	if err != nil {
		return "", resp, false, err
	}

	if miners.Entities == nil || len(*miners.Entities) == 0 {
		return "", resp, true, err
	}

	for _, miner := range *miners.Entities {
		if *miner.Name == name {
			log.Printf("Retrieved the languageunderstanding miners id %s by name %s", *miner.Id, name)
			return *miner.Id, resp, false, nil
		}
	}

	return "", resp, true, fmt.Errorf("Unable to find languageunderstanding miners with name %s", name)
}

// getLanguageunderstandingMinersByIdFn is an implementation of the function to get a Genesys Cloud languageunderstanding miners by Id
func getLanguageunderstandingMinersByIdFn(ctx context.Context, p *languageunderstandingMinersProxy, id string) (*platformclientv2.Miner, *platformclientv2.APIResponse, error) {
	return p.languageUnderstandingApi.GetLanguageunderstandingMiner(id)
}

// deleteLanguageunderstandingMinersFn is an implementation function for deleting a Genesys Cloud languageunderstanding miners
func deleteLanguageunderstandingMinersFn(ctx context.Context, p *languageunderstandingMinersProxy, id string) (*platformclientv2.APIResponse, error) {
	return p.languageUnderstandingApi.DeleteLanguageunderstandingMiner(id)
}
