// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/gerbil/terraform-provider-Metadefender/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &MetadefenderProvider{}

type MetadefenderProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// MetadefenderProviderModel describes the provider data model.
type MetadefenderProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
}

func (p *MetadefenderProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "Metadefender"
	resp.Version = p.version
}

func (p *MetadefenderProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `MetaDefender Core: ## Developer Guide` + "\n" +
			`` + "\n" +
			`----` + "\n" +
			`` + "\n" +
			`This is the API documentation for *MetaDefender Core Public API*.  If you` + "\n" +
			`would like to evaluate or have any questions about this documentation,` + "\n" +
			`please contact us via our [Contact Us](https://www.opswat.com/contact) form.` + "\n" +
			`` + "\n" +
			`----` + "\n" +
			`` + "\n" +
			`## How to Interact with MetaDefender Core using REST` + "\n" +
			`` + "\n" +
			`Beginning with MetaDefender Core 4.x, OPSWAT recommends using the JSON-based` + "\n" +
			`REST API. The available methods are documented below.` + "\n" +
			`` + "\n" +
			`> _**Note**:_ MetaDefender API doesn't support chunk upload, however is` + "\n" +
			`recommended to stream the files to MetaDefender Core as part of the upload` + "\n" +
			`process. ` + "\n" +
			`` + "\n" +
			`---` + "\n" +
			`` + "\n" +
			`## File Analysis Process` + "\n" +
			`` + "\n" +
			`  MetaDefender's main functionality is to analyze large volumes with a high throughput. Depending on the configuration and licensed technologies, the analysis times can vary. ` + "\n" +
			`  Below is a brief description of the API integration flow:` + "\n" +
			`` + "\n" +
			`  1. Upload a file for analysis (` + "`" + `POST /file` + "`" + `), which returns the ` + "`" + `data_id` + "`" + `: [File Analysis](/mdcore/metadefender-core/ref#fileanalysispost)).` + "\n" +
			`    ` + "\n" +
			`    > _**Note**:_ The performance depends on:` + "\n" +
			`    > - number of nodes (scaling)` + "\n" +
			`    > - number of engines per node` + "\n" +
			`    > - type of file to be scanned` + "\n" +
			`    > - Metadefender Core and nodes' hardware` + "\n" +
			`  ` + "\n" +
			`` + "\n" +
			`  2. You have 2 ways to retrieve the analysis report: ` + "\n" +
			`    - **Polling**: Fetch the result with previously received data_id (` + "`" + `GET /file/{data_id}` + "`" + ` resource) until scan result belonging to data_id doesn't reach the 100 percent progress_percentage: ([Fetch analysis result](/mdcore/metadefender-core/ref#fileanalysisget))` + "\n" +
			`  ` + "\n" +
			`    > _**Note**:_ Too many data_id requests can reduce performance. It is enough to just check every few hundred milliseconds.` + "\n" +
			`    ` + "\n" +
			`    - **Callbackurl**: Specify a callbackurl that will be called once the analysis is complete. ` + "\n" +
			`` + "\n" +
			`  3. Retrieve the analysis results anytime after the analysis is completed with hash for files (md5, sha1, sha256) by calling [Fetch analysis result by hash](/mdcore/metadefender-core/ref#hashget).` + "\n" +
			`    - The hash can be found in the scan results` + "\n" +
			`` + "\n" +
			`  4. Retrieve processed file (sanitized, redacted, watermarked, etc.) after the analysis is complete. ` + "\n" +
			`    > _**Note**:_ Based on the configured retention policy, the files might be available for retrieval at a later time. ` + "\n" +
			`` + "\n" +
			`---` + "\n" +
			`` + "\n" +
			`OPSWAT provides some sample codes on [GitHub](https://github.com/OPSWAT) to` + "\n" +
			`make it easier to understand how the MetaDefender REST API works.` + "\n" +
			``,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to http://localhost:8008)",
				Optional:            true,
				Required:            false,
			},
		},
	}
}

func (p *MetadefenderProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data MetadefenderProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "http://localhost:8008"
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *MetadefenderProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *MetadefenderProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MetadefenderProvider{
			version: version,
		}
	}
}
