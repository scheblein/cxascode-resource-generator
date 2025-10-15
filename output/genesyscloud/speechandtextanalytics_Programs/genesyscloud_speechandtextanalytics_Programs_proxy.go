package speechandtextanalytics_programs

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
	"log"
)

/*
The genesyscloud_speechandtextanalytics_programs_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *speechandtextanalyticsProgramsProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createSpeechandtextanalyticsProgramsFunc func(ctx context.Context, p *speechandtextanalyticsProgramsProxy, program *platformclientv2.Program) (*platformclientv2.Program, *platformclientv2.APIResponse, error)
type getAllSpeechandtextanalyticsProgramsFunc func(ctx context.Context, p *speechandtextanalyticsProgramsProxy) (*[]platformclientv2.Program, *platformclientv2.APIResponse, error)
type getSpeechandtextanalyticsProgramsIdByNameFunc func(ctx context.Context, p *speechandtextanalyticsProgramsProxy, name string) (string, *platformclientv2.APIResponse, bool, error)
type getSpeechandtextanalyticsProgramsByIdFunc func(ctx context.Context, p *speechandtextanalyticsProgramsProxy, id string) (*platformclientv2.Program, *platformclientv2.APIResponse, error)
type updateSpeechandtextanalyticsProgramsFunc func(ctx context.Context, p *speechandtextanalyticsProgramsProxy, id string, program *platformclientv2.Program) (*platformclientv2.Program, *platformclientv2.APIResponse, error)
type deleteSpeechandtextanalyticsProgramsFunc func(ctx context.Context, p *speechandtextanalyticsProgramsProxy, id string) (*platformclientv2.APIResponse, error)

// speechandtextanalyticsProgramsProxy contains all of the methods that call genesys cloud APIs.
type speechandtextanalyticsProgramsProxy struct {
	clientConfig                                  *platformclientv2.Configuration
	speechTextAnalyticsApi                        *platformclientv2.SpeechTextAnalyticsApi
	createSpeechandtextanalyticsProgramsAttr      createSpeechandtextanalyticsProgramsFunc
	getAllSpeechandtextanalyticsProgramsAttr      getAllSpeechandtextanalyticsProgramsFunc
	getSpeechandtextanalyticsProgramsIdByNameAttr getSpeechandtextanalyticsProgramsIdByNameFunc
	getSpeechandtextanalyticsProgramsByIdAttr     getSpeechandtextanalyticsProgramsByIdFunc
	updateSpeechandtextanalyticsProgramsAttr      updateSpeechandtextanalyticsProgramsFunc
	deleteSpeechandtextanalyticsProgramsAttr      deleteSpeechandtextanalyticsProgramsFunc
}

// newSpeechandtextanalyticsProgramsProxy initializes the speechandtextanalytics programs proxy with all of the data needed to communicate with Genesys Cloud
func newSpeechandtextanalyticsProgramsProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsProgramsProxy {
	api := platformclientv2.NewSpeechTextAnalyticsApiWithConfig(clientConfig)
	return &speechandtextanalyticsProgramsProxy{
		clientConfig:                                  clientConfig,
		speechTextAnalyticsApi:                        api,
		createSpeechandtextanalyticsProgramsAttr:      createSpeechandtextanalyticsProgramsFn,
		getAllSpeechandtextanalyticsProgramsAttr:      getAllSpeechandtextanalyticsProgramsFn,
		getSpeechandtextanalyticsProgramsIdByNameAttr: getSpeechandtextanalyticsProgramsIdByNameFn,
		getSpeechandtextanalyticsProgramsByIdAttr:     getSpeechandtextanalyticsProgramsByIdFn,
		updateSpeechandtextanalyticsProgramsAttr:      updateSpeechandtextanalyticsProgramsFn,
		deleteSpeechandtextanalyticsProgramsAttr:      deleteSpeechandtextanalyticsProgramsFn,
	}
}

// getSpeechandtextanalyticsProgramsProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getSpeechandtextanalyticsProgramsProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsProgramsProxy {
	if internalProxy == nil {
		internalProxy = newSpeechandtextanalyticsProgramsProxy(clientConfig)
	}

	return internalProxy
}

// createSpeechandtextanalyticsPrograms creates a Genesys Cloud speechandtextanalytics programs
func (p *speechandtextanalyticsProgramsProxy) createSpeechandtextanalyticsPrograms(ctx context.Context, speechandtextanalyticsPrograms *platformclientv2.Program) (*platformclientv2.Program, *platformclientv2.APIResponse, error) {
	return p.createSpeechandtextanalyticsProgramsAttr(ctx, p, speechandtextanalyticsPrograms)
}

// getSpeechandtextanalyticsPrograms retrieves all Genesys Cloud speechandtextanalytics programs
func (p *speechandtextanalyticsProgramsProxy) getAllSpeechandtextanalyticsPrograms(ctx context.Context) (*[]platformclientv2.Program, *platformclientv2.APIResponse, error) {
	return p.getAllSpeechandtextanalyticsProgramsAttr(ctx, p)
}

// getSpeechandtextanalyticsProgramsIdByName returns a single Genesys Cloud speechandtextanalytics programs by a name
func (p *speechandtextanalyticsProgramsProxy) getSpeechandtextanalyticsProgramsIdByName(ctx context.Context, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return p.getSpeechandtextanalyticsProgramsIdByNameAttr(ctx, p, name)
}

// getSpeechandtextanalyticsProgramsById returns a single Genesys Cloud speechandtextanalytics programs by Id
func (p *speechandtextanalyticsProgramsProxy) getSpeechandtextanalyticsProgramsById(ctx context.Context, id string) (*platformclientv2.Program, *platformclientv2.APIResponse, error) {
	return p.getSpeechandtextanalyticsProgramsByIdAttr(ctx, p, id)
}

// updateSpeechandtextanalyticsPrograms updates a Genesys Cloud speechandtextanalytics programs
func (p *speechandtextanalyticsProgramsProxy) updateSpeechandtextanalyticsPrograms(ctx context.Context, id string, speechandtextanalyticsPrograms *platformclientv2.Program) (*platformclientv2.Program, *platformclientv2.APIResponse, error) {
	return p.updateSpeechandtextanalyticsProgramsAttr(ctx, p, id, speechandtextanalyticsPrograms)
}

// deleteSpeechandtextanalyticsPrograms deletes a Genesys Cloud speechandtextanalytics programs by Id
func (p *speechandtextanalyticsProgramsProxy) deleteSpeechandtextanalyticsPrograms(ctx context.Context, id string) (*platformclientv2.APIResponse, error) {
	return p.deleteSpeechandtextanalyticsProgramsAttr(ctx, p, id)
}

// createSpeechandtextanalyticsProgramsFn is an implementation function for creating a Genesys Cloud speechandtextanalytics programs
func createSpeechandtextanalyticsProgramsFn(ctx context.Context, p *speechandtextanalyticsProgramsProxy, speechandtextanalyticsPrograms *platformclientv2.Program) (*platformclientv2.Program, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PostSpeechandtextanalyticsPrograms(*speechandtextanalyticsPrograms)
}

// getAllSpeechandtextanalyticsProgramsFn is the implementation for retrieving all speechandtextanalytics programs in Genesys Cloud
func getAllSpeechandtextanalyticsProgramsFn(ctx context.Context, p *speechandtextanalyticsProgramsProxy) (*[]platformclientv2.Program, *platformclientv2.APIResponse, error) {
	var allPrograms []platformclientv2.Program
	const pageSize = 100

	programs, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsPrograms()
	if err != nil {
		return nil, resp, err
	}
	if programs.Entities == nil || len(*programs.Entities) == 0 {
		return &allPrograms, resp, nil
	}
	for _, program := range *programs.Entities {
		allPrograms = append(allPrograms, program)
	}

	for pageNum := 2; pageNum <= *programs.PageCount; pageNum++ {
		programs, _, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsPrograms()
		if err != nil {
			return nil, resp, err
		}

		if programs.Entities == nil || len(*programs.Entities) == 0 {
			break
		}

		for _, program := range *programs.Entities {
			allPrograms = append(allPrograms, program)
		}
	}

	return &allPrograms, resp, nil
}

// getSpeechandtextanalyticsProgramsIdByNameFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics programs by name
func getSpeechandtextanalyticsProgramsIdByNameFn(ctx context.Context, p *speechandtextanalyticsProgramsProxy, name string) (string, *platformclientv2.APIResponse, bool, error) {
	programs, resp, err := p.speechTextAnalyticsApi.GetSpeechandtextanalyticsPrograms()
	if err != nil {
		return "", resp, false, err
	}

	if programs.Entities == nil || len(*programs.Entities) == 0 {
		return "", resp, true, err
	}

	for _, program := range *programs.Entities {
		if *program.Name == name {
			log.Printf("Retrieved the speechandtextanalytics programs id %s by name %s", *program.Id, name)
			return *program.Id, resp, false, nil
		}
	}

	return "", resp, true, fmt.Errorf("Unable to find speechandtextanalytics programs with name %s", name)
}

// getSpeechandtextanalyticsProgramsByIdFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics programs by Id
func getSpeechandtextanalyticsProgramsByIdFn(ctx context.Context, p *speechandtextanalyticsProgramsProxy, id string) (*platformclientv2.Program, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.GetSpeechandtextanalyticsProgram(id)
}

// updateSpeechandtextanalyticsProgramsFn is an implementation of the function to update a Genesys Cloud speechandtextanalytics programs
func updateSpeechandtextanalyticsProgramsFn(ctx context.Context, p *speechandtextanalyticsProgramsProxy, id string, speechandtextanalyticsPrograms *platformclientv2.Program) (*platformclientv2.Program, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PutSpeechandtextanalyticsProgram(id, *speechandtextanalyticsPrograms)
}

// deleteSpeechandtextanalyticsProgramsFn is an implementation function for deleting a Genesys Cloud speechandtextanalytics programs
func deleteSpeechandtextanalyticsProgramsFn(ctx context.Context, p *speechandtextanalyticsProgramsProxy, id string) (*platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.DeleteSpeechandtextanalyticsProgram(id)
}
