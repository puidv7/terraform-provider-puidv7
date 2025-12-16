// Copyright 2025 puidv7
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	puidv7 "github.com/puidv7/puidv7-go"
)

var _ datasource.DataSource = &EncodeDataSource{}

type EncodeDataSource struct{}

type EncodeDataSourceModel struct {
	Uuid   types.String `tfsdk:"uuid"`
	Prefix types.String `tfsdk:"prefix"`
	Id     types.String `tfsdk:"id"`
}

func NewEncodeDataSource() datasource.DataSource {
	return &EncodeDataSource{}
}

func (d *EncodeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_encode"
}

func (d *EncodeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Encodes an existing UUID into puidv7 format (prefixed Crockford Base32).",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "The UUID to encode (with or without hyphens).",
				Required:    true,
			},
			"prefix": schema.StringAttribute{
				Description: "A 3-character lowercase alphabetic prefix (a-z) for the identifier.",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "The encoded puidv7 identifier.",
				Computed:    true,
			},
		},
	}
}

func (d *EncodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data EncodeDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	prefix := data.Prefix.ValueString()
	uuid := data.Uuid.ValueString()
	id, err := puidv7.Encode(uuid, prefix)
	if err != nil {
		resp.Diagnostics.AddError("Failed to encode UUID", err.Error())
		return
	}

	data.Id = types.StringValue(id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
