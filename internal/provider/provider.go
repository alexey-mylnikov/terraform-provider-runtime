package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*runtimeProvider)(nil)

func New() provider.Provider {
	return &runtimeProvider{}
}

type runtimeProvider struct{}

func (p *runtimeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "runtime"
}

func (p *runtimeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *runtimeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewRuntimeDataSource,
	}
}

func (p *runtimeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *runtimeProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {
}
