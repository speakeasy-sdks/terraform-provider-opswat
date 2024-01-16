// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"fmt"
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"http://localhost:8008",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient  HTTPClient
	SecurityClient HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Metadefender - MetaDefender Core: ## Developer Guide
//
// ----
//
// This is the API documentation for *MetaDefender Core Public API*.  If you
// would like to evaluate or have any questions about this documentation,
// please contact us via our [Contact Us](https://www.opswat.com/contact) form.
//
// ----
//
// ## How to Interact with MetaDefender Core using REST
//
// Beginning with MetaDefender Core 4.x, OPSWAT recommends using the JSON-based
// REST API. The available methods are documented below.
//
// > _**Note**:_ MetaDefender API doesn't support chunk upload, however is
// recommended to stream the files to MetaDefender Core as part of the upload
// process.
//
// ---
//
// ## File Analysis Process
//
//	MetaDefender's main functionality is to analyze large volumes with a high throughput. Depending on the configuration and licensed technologies, the analysis times can vary.
//	Below is a brief description of the API integration flow:
//
//	1. Upload a file for analysis (`POST /file`), which returns the `data_id`: [File Analysis](/mdcore/metadefender-core/ref#fileanalysispost)).
//
//	  > _**Note**:_ The performance depends on:
//	  > - number of nodes (scaling)
//	  > - number of engines per node
//	  > - type of file to be scanned
//	  > - Metadefender Core and nodes' hardware
//
//
//	2. You have 2 ways to retrieve the analysis report:
//	  - **Polling**: Fetch the result with previously received data_id (`GET /file/{data_id}` resource) until scan result belonging to data_id doesn't reach the 100 percent progress_percentage: ([Fetch analysis result](/mdcore/metadefender-core/ref#fileanalysisget))
//
//	  > _**Note**:_ Too many data_id requests can reduce performance. It is enough to just check every few hundred milliseconds.
//
//	  - **Callbackurl**: Specify a callbackurl that will be called once the analysis is complete.
//
//	3. Retrieve the analysis results anytime after the analysis is completed with hash for files (md5, sha1, sha256) by calling [Fetch analysis result by hash](/mdcore/metadefender-core/ref#hashget).
//	  - The hash can be found in the scan results
//
//	4. Retrieve processed file (sanitized, redacted, watermarked, etc.) after the analysis is complete.
//	  > _**Note**:_ Based on the configured retention policy, the files might be available for retrieval at a later time.
//
// ---
//
// OPSWAT provides some sample codes on [GitHub](https://github.com/OPSWAT) to
// make it easier to understand how the MetaDefender REST API works.
//
// https://onlinehelp.opswat.com/corev4 - Official product documentation
type Metadefender struct {
	// Configure the product through APIs (especially the Settings). Will require admin apikey.
	//
	Config *Config
	// Yara engine configuration and source management APIs.
	Yara *Yara
	// Enable/disable or pin/unpin the engines through API.
	Engines *Engines
	// Admin specific API requests.
	Admin *Admin
	// Activate the product or get licensing information. Will require admin apikey.
	//
	License *License
	// ### File analysis APIs
	// Submit each file to MetaDefender Core individually or group them in batches. Each file submission will return a `data_id` which will be the unique identifier used to retrieve the analysis results.
	// **Important**: Even though one file is being submitted, if MetaDefender Core is configured to extract the files, all compound file types (archives, Office documents, etc.) will be extracted and each file within will be analyzed as a separate entry.
	//   - This means that if you submit an archive with 100 files in it, MetaDefender Core will process 101 files: original file as it is and each of the 100 child files
	//     - Note that by opening the files, detection ratio can increase even by 30%.
	//
	// > _**Note**:_ MetaDefender API doesn't support chunk upload. You shouldn't load the file in memory, is recommended to stream the files to MetaDefender Core as part of the upload process.
	//
	Analysis *Analysis
	// Group the analysis requests in batches.
	Batch *Batch
	// ### Authentication APIs
	// User authentication is done via username & password.
	// Additional integrations are available within the product:
	//   - **LDAP** integration
	//   - **Active Directory** integration
	//   - **SAML** integration (starting with v4.18.0)
	//
	Auth *Auth
	// Health check and statistics about MetaDefender Core usage.
	Stats *Stats

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Metadefender)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Metadefender) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Metadefender) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Metadefender) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Metadefender) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Metadefender) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Metadefender {
	sdk := &Metadefender{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "v5.6.1",
			SDKVersion:        "0.11.2",
			GenVersion:        "2.233.2",
			UserAgent:         "speakeasy-sdk/go 0.11.2 2.233.2 v5.6.1 Metadefender",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
	}

	sdk.Config = newConfig(sdk.sdkConfiguration)

	sdk.Yara = newYara(sdk.sdkConfiguration)

	sdk.Engines = newEngines(sdk.sdkConfiguration)

	sdk.Admin = newAdmin(sdk.sdkConfiguration)

	sdk.License = newLicense(sdk.sdkConfiguration)

	sdk.Analysis = newAnalysis(sdk.sdkConfiguration)

	sdk.Batch = newBatch(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.Stats = newStats(sdk.sdkConfiguration)

	return sdk
}
