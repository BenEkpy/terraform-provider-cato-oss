package provider

import (
	"context"

	"github.com/BenEkpy/cato-go-client-oss/catogo"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &staticHostResource{}
	_ resource.ResourceWithConfigure = &staticHostResource{}
)

func NewStaticHostResource() resource.Resource {
	return &staticHostResource{}
}

type staticHostResource struct {
	client *catogo.Client
}

type StaticHost struct {
	Id         types.String `tfsdk:"id"`
	SiteId     types.String `tfsdk:"site_id"`
	Name       types.String `tfsdk:"name"`
	Ip         types.String `tfsdk:"ip"`
	MacAddress types.String `tfsdk:"mac_address"`
}

func (r *staticHostResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_statichost"
}

func (r *staticHostResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for the host",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				Description: "Host Site ID",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Host name",
				Required:    true,
			},
			"ip": schema.StringAttribute{
				Description: "Host IP Address",
				Required:    true,
			},
			"mac_address": schema.StringAttribute{
				Description: "Host Mac Address",
				Optional:    true,
			},
		},
	}
}

func (d *staticHostResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*catogo.Client)
}

func (r *staticHostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan StaticHost
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.AddStaticHostInput{
		Name:       plan.Name.ValueString(),
		Ip:         plan.Ip.ValueString(),
		MacAddress: plan.MacAddress.ValueStringPointer(),
	}

	body, err := r.client.AddStaticHost(plan.SiteId.ValueString(), input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	plan.Id = types.StringValue(body.HostId)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *staticHostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *staticHostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan StaticHost
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.UpdateStaticHostInput{
		Name:       plan.Name.ValueString(),
		Ip:         plan.Ip.ValueString(),
		MacAddress: plan.MacAddress.ValueStringPointer(),
	}

	_, err := r.client.UpdateStaticHost(plan.SiteId.ValueString(), plan.Id.ValueString(), input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *staticHostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state StaticHost
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// check if site exist before removing static host
	siteExists, err := r.client.SiteExists(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to connect or request the Cato API",
			err.Error(),
		)
		return
	}

	if siteExists {
		_, err := r.client.RemoveStaticHost(state.SiteId.ValueString(), state.Id.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to connect or request the Cato API",
				err.Error(),
			)
			return
		}
	}

}
