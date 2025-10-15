package speechandtextanalytics_settings

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
)

/*
The genesyscloud_speechandtextanalytics_settings_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *speechandtextanalyticsSettingsProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type getSpeechandtextanalyticsSettingsByIdFunc func(ctx context.Context, p *speechandtextanalyticsSettingsProxy, id string) (*platformclientv2.Speechtextanalyticssettingsresponse, *platformclientv2.APIResponse, error)
type updateSpeechandtextanalyticsSettingsFunc func(ctx context.Context, p *speechandtextanalyticsSettingsProxy, id string, speechTextAnalyticsSettingsResponse *platformclientv2.Speechtextanalyticssettingsresponse) (*platformclientv2.Speechtextanalyticssettingsresponse, *platformclientv2.APIResponse, error)

// speechandtextanalyticsSettingsProxy contains all of the methods that call genesys cloud APIs.
type speechandtextanalyticsSettingsProxy struct {
	clientConfig                              *platformclientv2.Configuration
	speechTextAnalyticsApi                    *platformclientv2.SpeechTextAnalyticsApi
	getSpeechandtextanalyticsSettingsByIdAttr getSpeechandtextanalyticsSettingsByIdFunc
	updateSpeechandtextanalyticsSettingsAttr  updateSpeechandtextanalyticsSettingsFunc
}

// newSpeechandtextanalyticsSettingsProxy initializes the speechandtextanalytics settings proxy with all of the data needed to communicate with Genesys Cloud
func newSpeechandtextanalyticsSettingsProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsSettingsProxy {
	api := platformclientv2.NewSpeechTextAnalyticsApiWithConfig(clientConfig)
	return &speechandtextanalyticsSettingsProxy{
		clientConfig:           clientConfig,
		speechTextAnalyticsApi: api,
		getSpeechandtextanalyticsSettingsByIdAttr: getSpeechandtextanalyticsSettingsByIdFn,
		updateSpeechandtextanalyticsSettingsAttr:  updateSpeechandtextanalyticsSettingsFn,
	}
}

// getSpeechandtextanalyticsSettingsProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getSpeechandtextanalyticsSettingsProxy(clientConfig *platformclientv2.Configuration) *speechandtextanalyticsSettingsProxy {
	if internalProxy == nil {
		internalProxy = newSpeechandtextanalyticsSettingsProxy(clientConfig)
	}

	return internalProxy
}

// getSpeechandtextanalyticsSettingsById returns a single Genesys Cloud speechandtextanalytics settings by Id
func (p *speechandtextanalyticsSettingsProxy) getSpeechandtextanalyticsSettingsById(ctx context.Context, id string) (*platformclientv2.Speechtextanalyticssettingsresponse, *platformclientv2.APIResponse, error) {
	return p.getSpeechandtextanalyticsSettingsByIdAttr(ctx, p, id)
}

// updateSpeechandtextanalyticsSettings updates a Genesys Cloud speechandtextanalytics settings
func (p *speechandtextanalyticsSettingsProxy) updateSpeechandtextanalyticsSettings(ctx context.Context, id string, speechandtextanalyticsSettings *platformclientv2.Speechtextanalyticssettingsresponse) (*platformclientv2.Speechtextanalyticssettingsresponse, *platformclientv2.APIResponse, error) {
	return p.updateSpeechandtextanalyticsSettingsAttr(ctx, p, id, speechandtextanalyticsSettings)
}

// getSpeechandtextanalyticsSettingsByIdFn is an implementation of the function to get a Genesys Cloud speechandtextanalytics settings by Id
func getSpeechandtextanalyticsSettingsByIdFn(ctx context.Context, p *speechandtextanalyticsSettingsProxy, id string) (*platformclientv2.Speechtextanalyticssettingsresponse, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.GetSpeechandtextanalyticsSettings(id)
}

// updateSpeechandtextanalyticsSettingsFn is an implementation of the function to update a Genesys Cloud speechandtextanalytics settings
func updateSpeechandtextanalyticsSettingsFn(ctx context.Context, p *speechandtextanalyticsSettingsProxy, id string, speechandtextanalyticsSettings *platformclientv2.Speechtextanalyticssettingsresponse) (*platformclientv2.Speechtextanalyticssettingsresponse, *platformclientv2.APIResponse, error) {
	return p.speechTextAnalyticsApi.PutSpeechandtextanalyticsSettings(id, *speechandtextanalyticsSettings)
}
