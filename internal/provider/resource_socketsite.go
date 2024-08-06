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
	_ resource.Resource              = &socketSiteResource{}
	_ resource.ResourceWithConfigure = &socketSiteResource{}
)

func NewSocketSiteResource() resource.Resource {
	return &socketSiteResource{}
}

type socketSiteResource struct {
	client *catogo.Client
}

type SocketSite struct {
	Id                   types.String         `tfsdk:"id"`
	Name                 types.String         `tfsdk:"name"`
	ConnectionType       types.String         `tfsdk:"connection_type"`
	SiteType             types.String         `tfsdk:"site_type"`
	Description          types.String         `tfsdk:"description"`
	NativeNetworkRange   types.String         `tfsdk:"native_network_range"`
	NativeNetworkRangeId types.String         `tfsdk:"native_network_range_id"`
	LocalIp              types.String         `tfsdk:"local_ip"`
	TranslatedSubnet     types.String         `tfsdk:"translated_subnet"`
	SiteLocation         AddSiteLocationInput `tfsdk:"site_location"`
}

type AddSiteLocationInput struct {
	CountryCode types.String `tfsdk:"country_code"`
	StateCode   types.String `tfsdk:"state_code"`
	Timezone    types.String `tfsdk:"timezone"`
	// Address     types.String `tfsdk:"address"`
	// City        types.String `tfsdk:"city"`
}

func (r *socketSiteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_socketsite"
}

func (r *socketSiteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for the site",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Site name",
				Required:    true,
			},
			"connection_type": schema.StringAttribute{
				Description: "Connection type for the site (SOCKET_X1500, SOCKET_AWS1500, SOCKET_AZ1500, ...)",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_type": schema.StringAttribute{
				Description: "Site type (BRANCH, DATACENTER, ...)",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Site description",
				Optional:    true,
			},
			"native_network_range": schema.StringAttribute{
				Description: "Site native IP range (CIDR)",
				Required:    true,
			},
			"native_network_range_id": schema.StringAttribute{
				Description: "Site native IP range ID (for update purpose)",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"local_ip": schema.StringAttribute{
				Description: "Site native range local ip",
				Required:    true,
			},
			"translated_subnet": schema.StringAttribute{
				Description: "Site translated native IP range (CIDR)",
				Optional:    true,
			},
			"site_location": schema.SingleNestedAttribute{
				Description: "Site location",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"country_code": schema.StringAttribute{
						Description: "Site country code (can be retrieve from entityLookup)",
						Required:    true,
					},
					"state_code": schema.StringAttribute{
						Description: "Optionnal site state code(can be retrieve from entityLookup)",
						Optional:    true,
					},
					"timezone": schema.StringAttribute{
						Description: "Site timezone (can be retrieve from entityLookup)",
						Required:    true,
					},
					// "city": schema.StringAttribute{
					// 	Description: "Optionnal city",
					// 	Optional:    true,
					// },
					// "address": schema.StringAttribute{
					// 	Description: "Optionnal address",
					// 	Optional:    true,
					// },
				},
			},
		},
	}
}

func (d *socketSiteResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*catogo.Client)
}

func (r *socketSiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan SocketSite
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.AddSocketSiteInput{
		Name:               plan.Name.ValueString(),
		ConnectionType:     plan.ConnectionType.ValueString(),
		SiteType:           plan.SiteType.ValueString(),
		Description:        plan.Description.ValueStringPointer(),
		NativeNetworkRange: plan.NativeNetworkRange.ValueString(),
		TranslatedSubnet:   plan.TranslatedSubnet.ValueStringPointer(),
		SiteLocation: catogo.AddSiteLocationInput{
			CountryCode: plan.SiteLocation.CountryCode.ValueString(),
			StateCode:   plan.SiteLocation.StateCode.ValueStringPointer(),
			Timezone:    plan.SiteLocation.Timezone.ValueString(),
		},
	}

	body, err := r.client.AddSocketSite(input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	// retrieving native-network range id to update local IP
	network_range_id, err := r.client.GetSocketSiteNativeRangeId(body.SiteId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	// update local IP on Native range
	input_network_range := catogo.UpdateNetworkRangeInput{
		LocalIp: plan.LocalIp.ValueStringPointer(),
	}
	_, err = r.client.UpdateNetworkRange(network_range_id.Id, input_network_range)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	plan.Id = types.StringValue(body.SiteId)
	plan.NativeNetworkRangeId = types.StringValue(network_range_id.Id)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *socketSiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *socketSiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan SocketSite
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update general details
	input_update_general := catogo.UpdateSiteGeneralDetailsInput{
		Name:        plan.Name.ValueStringPointer(),
		SiteType:    plan.SiteType.ValueStringPointer(),
		Description: plan.Description.ValueStringPointer(),
		SiteLocation: &catogo.UpdateSiteLocationInput{
			CountryCode: plan.SiteLocation.CountryCode.ValueStringPointer(),
			StateCode:   plan.SiteLocation.StateCode.ValueStringPointer(),
			Timezone:    plan.SiteLocation.Timezone.ValueStringPointer(),
		},
	}

	_, err := r.client.UpdateSiteGeneralDetails(plan.Id.ValueString(), input_update_general)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	// Update network settings
	input_update_network := catogo.UpdateNetworkRangeInput{
		Subnet:           plan.NativeNetworkRange.ValueStringPointer(),
		TranslatedSubnet: plan.TranslatedSubnet.ValueStringPointer(),
		LocalIp:          plan.LocalIp.ValueStringPointer(),
	}

	_, err = r.client.UpdateNetworkRange(plan.NativeNetworkRangeId.ValueString(), input_update_network)
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

func (r *socketSiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state SocketSite
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.RemoveSite(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to connect or request the Cato API",
			err.Error(),
		)
		return
	}
}
