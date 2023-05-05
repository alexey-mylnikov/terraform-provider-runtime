package provider

import (
	"context"
	"runtime"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = (*runtimeDataSource)(nil)
)

func NewRuntimeDataSource() datasource.DataSource {
	return &runtimeDataSource{}
}

type runtimeDataSource struct{}

type runtimeDataSourceModel struct {
	OS   types.String `tfsdk:"os"`
	ARCH types.String `tfsdk:"arch"`
	ID   types.String `tfsdk:"id"`
}

func (n *runtimeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName
}

func (n *runtimeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The `runtime` data source allows determine some information about host operating system.",
		Attributes: map[string]schema.Attribute{
			"os": schema.StringAttribute{
				Description: "The running program's architecture target: one of 386, amd64, arm, s390x, and so on.",
				Computed:    true,
			},
			"arch": schema.StringAttribute{
				Description: "The running program's operating system target: one of darwin, freebsd, linux, and so on.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "The id of the data source. This will always be set to `-`",
				Computed:    true,
			},
		},
	}
}

func (n *runtimeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config runtimeDataSourceModel

	config.OS = types.StringValue(runtime.GOOS)
	config.ARCH = types.StringValue(runtime.GOARCH)
	config.ID = types.StringValue("-")

	diags := resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
}
