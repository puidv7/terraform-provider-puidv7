// Copyright 2025 puidv7
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &Puidv7Provider{}

type Puidv7Provider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Puidv7Provider{
			version: version,
		}
	}
}

func (p *Puidv7Provider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "puidv7"
	resp.Version = p.version
}

func (p *Puidv7Provider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provider for generating and encoding puidv7 identifiers (prefixed UUIDv7 in Crockford Base32).",
	}
}

func (p *Puidv7Provider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p *Puidv7Provider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewIdResource,
	}
}

func (p *Puidv7Provider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewEncodeDataSource,
		NewDecodeDataSource,
	}
}
