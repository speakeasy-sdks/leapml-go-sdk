package leapml

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/utils"
	"net/http"
	"time"
)

var ServerList = []string{
	"https://api.leapml.dev",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Leapml struct {
	FineTuning       *fineTuning
	GeneratingImages *generatingImages
	ImageEditing     *imageEditing

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient

	_serverURL  string
	_language   string
	_sdkVersion string
	_genVersion string
}

type SDKOption func(*Leapml)

func WithServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Leapml) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Leapml) {
		sdk._defaultClient = client
	}
}

func New(opts ...SDKOption) *Leapml {
	sdk := &Leapml{
		_language:   "go",
		_sdkVersion: "1.4.0",
		_genVersion: "1.5.0",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		sdk._securityClient = sdk._defaultClient

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.FineTuning = newFineTuning(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.GeneratingImages = newGeneratingImages(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.ImageEditing = newImageEditing(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
