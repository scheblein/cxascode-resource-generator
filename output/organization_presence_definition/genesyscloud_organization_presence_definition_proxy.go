package organization_presence_definition

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
	"log"
)

/*
The genesyscloud_organization_presence_definition_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *organizationPresenceDefinitionProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createOrganizationPresenceDefinitionFunc func(ctx context.Context, p *organizationPresenceDefinitionProxy, organizationPresenceDefinition *platformclientv2.Organizationpresencedefinition) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error)
type getAllOrganizationPresenceDefinitionFunc func(ctx context.Context, p *organizationPresenceDefinitionProxy) (*[]platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error)
type getOrganizationPresenceDefinitionIdByNameFunc func(ctx context.Context, p *organizationPresenceDefinitionProxy, name string) (string, *platformclientv2.APIResponse, bool, error)
type getOrganizationPresenceDefinitionByIdFunc func(ctx context.Context, p *organizationPresenceDefinitionProxy, id string) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error)
type updateOrganizationPresenceDefinitionFunc func(ctx context.Context, p *organizationPresenceDefinitionProxy, id string, organizationPresenceDefinition *platformclientv2.Organizationpresencedefinition) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error)
type deleteOrganizationPresenceDefinitionFunc func(ctx context.Context, p *organizationPresenceDefinitionProxy, id string) (*platformclientv2.APIResponse, error)

// organizationPresenceDefinitionProxy contains all of the methods that call genesys cloud APIs.
type organizationPresenceDefinitionProxy struct {
	clientConfig                                  *platformclientv2.Configuration
	presenceApi                                   *platformclientv2.PresenceApi
	createOrganizationPresenceDefinitionAttr      createOrganizationPresenceDefinitionFunc
	getAllOrganizationPresenceDefinitionAttr      getAllOrganizationPresenceDefinitionFunc
	getOrganizationPresenceDefinitionIdByNameAttr getOrganizationPresenceDefinitionIdByNameFunc
	getOrganizationPresenceDefinitionByIdAttr     getOrganizationPresenceDefinitionByIdFunc
	updateOrganizationPresenceDefinitionAttr      updateOrganizationPresenceDefinitionFunc
	deleteOrganizationPresenceDefinitionAttr      deleteOrganizationPresenceDefinitionFunc
}

// newOrganizationPresenceDefinitionProxy initializes the organization presence definition proxy with all of the data needed to communicate with Genesys Cloud
func newOrganizationPresenceDefinitionProxy(clientConfig *platformclientv2.Configuration) *organizationPresenceDefinitionProxy {
	api := platformclientv2.NewPresenceApiWithConfig(clientConfig)
	return &organizationPresenceDefinitionProxy{
		clientConfig:                                  clientConfig,
		presenceApi:                                   api,
		createOrganizationPresenceDefinitionAttr:      createOrganizationPresenceDefinitionFn,
		getAllOrganizationPresenceDefinitionAttr:      getAllOrganizationPresenceDefinitionFn,
		getOrganizationPresenceDefinitionIdByNameAttr: getOrganizationPresenceDefinitionIdByNameFn,
		getOrganizationPresenceDefinitionByIdAttr:     getOrganizationPresenceDefinitionByIdFn,
		updateOrganizationPresenceDefinitionAttr:      updateOrganizationPresenceDefinitionFn,
		deleteOrganizationPresenceDefinitionAttr:      deleteOrganizationPresenceDefinitionFn,
	}
}

// getOrganizationPresenceDefinitionProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getOrganizationPresenceDefinitionProxy(clientConfig *platformclientv2.Configuration) *organizationPresenceDefinitionProxy {
	if internalProxy == nil {
		internalProxy = newOrganizationPresenceDefinitionProxy(clientConfig)
	}

	return internalProxy
}

// createOrganizationPresenceDefinition creates a Genesys Cloud organization presence definition
func (p *organizationPresenceDefinitionProxy) createOrganizationPresenceDefinition(ctx context.Context, organizationPresenceDefinition *platformclientv2.Organizationpresencedefinition) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error) {
	return p.createOrganizationPresenceDefinitionAttr(ctx, p, organizationPresenceDefinition)
}

// getOrganizationPresenceDefinition retrieves all Genesys Cloud organization presence definition
func (p *organizationPresenceDefinitionProxy) getAllOrganizationPresenceDefinition(ctx context.Context) (*[]platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error) {
	return p.getAllOrganizationPresenceDefinitionAttr(ctx, p)
}

// getOrganizationPresenceDefinitionIdByName returns a single Genesys Cloud organization presence definition by a name
func (p *organizationPresenceDefinitionProxy) getOrganizationPresenceDefinitionIdByName(ctx context.Context, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return p.getOrganizationPresenceDefinitionIdByNameAttr(ctx, p, name)
}

// getOrganizationPresenceDefinitionById returns a single Genesys Cloud organization presence definition by Id
func (p *organizationPresenceDefinitionProxy) getOrganizationPresenceDefinitionById(ctx context.Context, id string) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error) {
	return p.getOrganizationPresenceDefinitionByIdAttr(ctx, p, id)
}

// updateOrganizationPresenceDefinition updates a Genesys Cloud organization presence definition
func (p *organizationPresenceDefinitionProxy) updateOrganizationPresenceDefinition(ctx context.Context, id string, organizationPresenceDefinition *platformclientv2.Organizationpresencedefinition) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error) {
	return p.updateOrganizationPresenceDefinitionAttr(ctx, p, id, organizationPresenceDefinition)
}

// deleteOrganizationPresenceDefinition deletes a Genesys Cloud organization presence definition by Id
func (p *organizationPresenceDefinitionProxy) deleteOrganizationPresenceDefinition(ctx context.Context, id string) (*platformclientv2.APIResponse, error) {
	return p.deleteOrganizationPresenceDefinitionAttr(ctx, p, id)
}

// createOrganizationPresenceDefinitionFn is an implementation function for creating a Genesys Cloud organization presence definition
func createOrganizationPresenceDefinitionFn(ctx context.Context, p *organizationPresenceDefinitionProxy, organizationPresenceDefinition *platformclientv2.Organizationpresencedefinition) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error) {
	return p.presenceApi.PostPresenceDefinitions(*organizationPresenceDefinition)
}

// getAllOrganizationPresenceDefinitionFn is the implementation for retrieving all organization presence definition in Genesys Cloud
func getAllOrganizationPresenceDefinitionFn(ctx context.Context, p *organizationPresenceDefinitionProxy) (*[]platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error) {
	var allOrganizationPresenceDefinitions []platformclientv2.Organizationpresencedefinition
	const pageSize = 100

	organizationPresenceDefinitions, resp, err := p.presenceApi.GetPresenceDefinitions()
	if err != nil {
		return nil, resp, err
	}
	if organizationPresenceDefinitions.Entities == nil || len(*organizationPresenceDefinitions.Entities) == 0 {
		return &allOrganizationPresenceDefinitions, resp, nil
	}
	for _, organizationPresenceDefinition := range *organizationPresenceDefinitions.Entities {
		allOrganizationPresenceDefinitions = append(allOrganizationPresenceDefinitions, organizationPresenceDefinition)
	}

	for pageNum := 2; pageNum <= *organizationPresenceDefinitions.PageCount; pageNum++ {
		organizationPresenceDefinitions, _, err := p.presenceApi.GetPresenceDefinitions()
		if err != nil {
			return nil, resp, err
		}

		if organizationPresenceDefinitions.Entities == nil || len(*organizationPresenceDefinitions.Entities) == 0 {
			break
		}

		for _, organizationPresenceDefinition := range *organizationPresenceDefinitions.Entities {
			allOrganizationPresenceDefinitions = append(allOrganizationPresenceDefinitions, organizationPresenceDefinition)
		}
	}

	return &allOrganizationPresenceDefinitions, resp, nil
}

// getOrganizationPresenceDefinitionIdByNameFn is an implementation of the function to get a Genesys Cloud organization presence definition by name
func getOrganizationPresenceDefinitionIdByNameFn(ctx context.Context, p *organizationPresenceDefinitionProxy, name string) (string, *platformclientv2.APIResponse, bool, error) {
	organizationPresenceDefinitions, resp, err := p.presenceApi.GetPresenceDefinitions()
	if err != nil {
		return "", resp, false, err
	}

	if organizationPresenceDefinitions.Entities == nil || len(*organizationPresenceDefinitions.Entities) == 0 {
		return "", resp, true, err
	}

	for _, organizationPresenceDefinition := range *organizationPresenceDefinitions.Entities {
		if *organizationPresenceDefinition.Name == name {
			log.Printf("Retrieved the organization presence definition id %s by name %s", *organizationPresenceDefinition.Id, name)
			return *organizationPresenceDefinition.Id, resp, false, nil
		}
	}

	return "", resp, true, fmt.Errorf("Unable to find organization presence definition with name %s", name)
}

// getOrganizationPresenceDefinitionByIdFn is an implementation of the function to get a Genesys Cloud organization presence definition by Id
func getOrganizationPresenceDefinitionByIdFn(ctx context.Context, p *organizationPresenceDefinitionProxy, id string) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error) {
	return p.presenceApi.GetPresenceDefinition(id)
}

// updateOrganizationPresenceDefinitionFn is an implementation of the function to update a Genesys Cloud organization presence definition
func updateOrganizationPresenceDefinitionFn(ctx context.Context, p *organizationPresenceDefinitionProxy, id string, organizationPresenceDefinition *platformclientv2.Organizationpresencedefinition) (*platformclientv2.Organizationpresencedefinition, *platformclientv2.APIResponse, error) {
	return p.presenceApi.PutPresenceDefinition(id, *organizationPresenceDefinition)
}

// deleteOrganizationPresenceDefinitionFn is an implementation function for deleting a Genesys Cloud organization presence definition
func deleteOrganizationPresenceDefinitionFn(ctx context.Context, p *organizationPresenceDefinitionProxy, id string) (*platformclientv2.APIResponse, error) {
	return p.presenceApi.DeletePresenceDefinition(id)
}
