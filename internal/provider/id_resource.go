// Copyright 2025 puidv7
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	puidv7 "github.com/puidv7/puidv7-go"
)

var _ resource.Resource = &IdResource{}

type IdResource struct{}

type IdResourceModel struct {
	Prefix types.String `tfsdk:"prefix"`
	Id     types.String `tfsdk:"id"`
	Uuid   types.String `tfsdk:"uuid"`
}

func NewIdResource() resource.Resource {
	return &IdResource{}
}

func (r *IdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_id"
}

func (r *IdResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Generates a new puidv7 identifier (UUIDv7 with a 3-character prefix, encoded in Crockford Base32).",
		Attributes: map[string]schema.Attribute{
			"prefix": schema.StringAttribute{
				Description: "A 3-character lowercase alphabetic prefix (a-z) for the identifier.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Description: "The generated puidv7 identifier.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"uuid": schema.StringAttribute{
				Description: "The underlying UUIDv7 in standard format.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *IdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IdResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	prefix := data.Prefix.ValueString()

	id, err := puidv7.New(prefix)
	if err != nil {
		resp.Diagnostics.AddError("Failed to generate puidv7", err.Error())
		return
	}

	uuid, err := puidv7.Decode(id, prefix)
	if err != nil {
		resp.Diagnostics.AddError("Failed to decode puidv7", err.Error())
		return
	}

	data.Id = types.StringValue(id)
	data.Uuid = types.StringValue(uuid)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IdResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data IdResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IdResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}
