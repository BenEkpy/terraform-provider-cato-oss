package provider

import (
	"context"

	"github.com/BenEkpy/terraform-provider-cato-oss/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cato_models "github.com/routebyintuition/cato-go-sdk/models"
)

var (
	_ resource.Resource              = &staticHostResource{}
	_ resource.ResourceWithConfigure = &staticHostResource{}
)

func NewStaticHostResource() resource.Resource {
	return &staticHostResource{}
}

type staticHostResource struct {
	client *catoClientData
}

type StaticHost struct {
	Id         types.String `tfsdk:"id"`
	SiteId     types.String `tfsdk:"site_id"`
	Name       types.String `tfsdk:"name"`
	Ip         types.String `tfsdk:"ip"`
	MacAddress types.String `tfsdk:"mac_address"`
}

func (r *staticHostResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_static_host"
}

func (r *staticHostResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				Description: "",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Description: "",
				Required:    true,
			},
			"ip": schema.StringAttribute{
				Description: "",
				Required:    true,
			},
			"mac_address": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (r *staticHostResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*catoClientData)
}

func (r *staticHostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan StaticHost
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// setting input
	input := cato_models.AddStaticHostInput{
		Name:       plan.Name.ValueString(),
		IP:         plan.Ip.ValueString(),
		MacAddress: plan.MacAddress.ValueStringPointer(),
	}

	tflog.Debug(ctx, "static_host create", map[string]interface{}{
		"input": utils.InterfaceToJSONString(input),
	})

	body, err := r.client.catosdk.SiteAddStaticHost(ctx, plan.SiteId.ValueString(), input, r.client.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API error",
			err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// overiding state with static host id
	resp.State.SetAttribute(
		ctx,
		path.Empty().AtName("id"),
		body.Site.GetAddStaticHost().HostID,
	)
}

func (r *staticHostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state StaticHost
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// check if site exist, else remove resource
	querySiteResult, err := r.client.catosdk.EntityLookup(ctx, r.client.AccountId, cato_models.EntityType("site"), nil, nil, nil, nil, []string{state.SiteId.ValueString()}, nil, nil, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API error",
			err.Error(),
		)
		return
	}

	if len(querySiteResult.EntityLookup.GetItems()) == 1 {
		tflog.Warn(ctx, "site not found, static host resource removed")
		resp.State.RemoveResource(ctx)
		return
	}

	// check if host exist before removing
	queryHostResult, err := r.client.catosdk.EntityLookup(ctx, r.client.AccountId, cato_models.EntityType("host"), nil, nil, nil, nil, []string{state.Id.ValueString()}, nil, nil, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API error",
			err.Error(),
		)
		return
	}

	if len(queryHostResult.EntityLookup.GetItems()) == 1 {
		tflog.Warn(ctx, "static host found, resource removed")
		resp.State.RemoveResource(ctx)
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *staticHostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan StaticHost
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// setting input
	input := cato_models.UpdateStaticHostInput{
		Name:       plan.Name.ValueStringPointer(),
		IP:         plan.Ip.ValueStringPointer(),
		MacAddress: plan.MacAddress.ValueStringPointer(),
	}

	tflog.Debug(ctx, "static_host update", map[string]interface{}{
		"input": utils.InterfaceToJSONString(input),
	})

	_, err := r.client.catosdk.SiteUpdateStaticHost(ctx, plan.Id.ValueString(), input, r.client.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API error",
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

	querySiteResult, err := r.client.catosdk.EntityLookup(ctx, r.client.AccountId, cato_models.EntityType("site"), nil, nil, nil, nil, []string{state.SiteId.ValueString()}, nil, nil, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API error",
			err.Error(),
		)
		return
	}

	// check if site exist before removing
	if len(querySiteResult.EntityLookup.GetItems()) == 1 {

		queryHostResult, err := r.client.catosdk.EntityLookup(ctx, r.client.AccountId, cato_models.EntityType("host"), nil, nil, nil, nil, []string{state.Id.ValueString()}, nil, nil, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Catov2 API error",
				err.Error(),
			)
			return
		}

		// check if host exist before removing
		if len(queryHostResult.EntityLookup.GetItems()) == 1 {

			_, err := r.client.catosdk.SiteRemoveStaticHost(ctx, state.Id.ValueString(), r.client.AccountId)
			if err != nil {
				resp.Diagnostics.AddError(
					"Catov2 API error",
					err.Error(),
				)
				return
			}
		}
	}

}
