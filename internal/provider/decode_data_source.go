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

var _ datasource.DataSource = &DecodeDataSource{}

type DecodeDataSource struct{}

type DecodeDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Prefix types.String `tfsdk:"prefix"`
	Uuid   types.String `tfsdk:"uuid"`
}

func NewDecodeDataSource() datasource.DataSource {
	return &DecodeDataSource{}
}

func (d *DecodeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_decode"
}

func (d *DecodeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Decodes a puidv7 identifier back to its underlying UUID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The puidv7 identifier to decode.",
				Required:    true,
			},
			"prefix": schema.StringAttribute{
				Description: "Optional: verify the prefix matches this value. If empty, any prefix is accepted.",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "The decoded UUID in standard format (with hyphens).",
				Computed:    true,
			},
		},
	}
}

func (d *DecodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DecodeDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	prefix := data.Prefix.ValueString()

	uuid, err := puidv7.Decode(id, prefix)
	if err != nil {
		resp.Diagnostics.AddError("Failed to decode puidv7", err.Error())
		return
	}

	data.Uuid = types.StringValue(uuid)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
