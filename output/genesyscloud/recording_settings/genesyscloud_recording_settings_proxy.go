package recording_settings

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v165/platformclientv2"
)

/*
The genesyscloud_recording_settings_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *recordingSettingsProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type getRecordingSettingsByIdFunc func(ctx context.Context, p *recordingSettingsProxy, id string) (*platformclientv2.Recordingsettings, *platformclientv2.APIResponse, error)
type updateRecordingSettingsFunc func(ctx context.Context, p *recordingSettingsProxy, id string, recordingSettings *platformclientv2.Recordingsettings) (*platformclientv2.Recordingsettings, *platformclientv2.APIResponse, error)

// recordingSettingsProxy contains all of the methods that call genesys cloud APIs.
type recordingSettingsProxy struct {
	clientConfig                 *platformclientv2.Configuration
	recordingApi                 *platformclientv2.RecordingApi
	getRecordingSettingsByIdAttr getRecordingSettingsByIdFunc
	updateRecordingSettingsAttr  updateRecordingSettingsFunc
}

// newRecordingSettingsProxy initializes the recording settings proxy with all of the data needed to communicate with Genesys Cloud
func newRecordingSettingsProxy(clientConfig *platformclientv2.Configuration) *recordingSettingsProxy {
	api := platformclientv2.NewRecordingApiWithConfig(clientConfig)
	return &recordingSettingsProxy{
		clientConfig:                 clientConfig,
		recordingApi:                 api,
		getRecordingSettingsByIdAttr: getRecordingSettingsByIdFn,
		updateRecordingSettingsAttr:  updateRecordingSettingsFn,
	}
}

// getRecordingSettingsProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getRecordingSettingsProxy(clientConfig *platformclientv2.Configuration) *recordingSettingsProxy {
	if internalProxy == nil {
		internalProxy = newRecordingSettingsProxy(clientConfig)
	}

	return internalProxy
}

// getRecordingSettingsById returns a single Genesys Cloud recording settings by Id
func (p *recordingSettingsProxy) getRecordingSettingsById(ctx context.Context, id string) (*platformclientv2.Recordingsettings, *platformclientv2.APIResponse, error) {
	return p.getRecordingSettingsByIdAttr(ctx, p, id)
}

// updateRecordingSettings updates a Genesys Cloud recording settings
func (p *recordingSettingsProxy) updateRecordingSettings(ctx context.Context, id string, recordingSettings *platformclientv2.Recordingsettings) (*platformclientv2.Recordingsettings, *platformclientv2.APIResponse, error) {
	return p.updateRecordingSettingsAttr(ctx, p, id, recordingSettings)
}

// getRecordingSettingsByIdFn is an implementation of the function to get a Genesys Cloud recording settings by Id
func getRecordingSettingsByIdFn(ctx context.Context, p *recordingSettingsProxy, id string) (*platformclientv2.Recordingsettings, *platformclientv2.APIResponse, error) {
	return p.recordingApi.GetRecordingSettings(id)
}

// updateRecordingSettingsFn is an implementation of the function to update a Genesys Cloud recording settings
func updateRecordingSettingsFn(ctx context.Context, p *recordingSettingsProxy, id string, recordingSettings *platformclientv2.Recordingsettings) (*platformclientv2.Recordingsettings, *platformclientv2.APIResponse, error) {
	return p.recordingApi.PutRecordingSettings(id, *recordingSettings)
}
