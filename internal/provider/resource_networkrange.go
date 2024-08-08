package provider

import (
	"context"
	"encoding/json"

	"github.com/BenEkpy/terraform-provider-cato-oss/internal/catogo"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &networkRangeResource{}
	_ resource.ResourceWithConfigure = &networkRangeResource{}
)

func NewNetworkRangeResource() resource.Resource {
	return &networkRangeResource{}
}

type networkRangeResource struct {
	client *catogo.Client
}

type NetworkRange struct {
	Id               types.String `tfsdk:"id"`
	InterfaceId      types.String `tfsdk:"interface_id"`
	SiteId           types.String `tfsdk:"site_id"`
	Name             types.String `tfsdk:"name"`
	RangeType        types.String `tfsdk:"range_type"`
	Subnet           types.String `tfsdk:"subnet"`
	TranslatedSubnet types.String `tfsdk:"translated_subnet"`
	LocalIp          types.String `tfsdk:"local_ip"`
	Gateway          types.String `tfsdk:"gateway"`
	Vlan             types.Int64  `tfsdk:"vlan"`
	DhcpSettings     types.Object `tfsdk:"dhcp_settings"`
}

func (r *networkRangeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_range"
}

func (r *networkRangeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Network Range id",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_id": schema.StringAttribute{
				Description: "Network Interface id",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				Description: "Host Site ID",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Network range name",
				Required:    true,
			},
			"range_type": schema.StringAttribute{
				Description: "Network range type",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"subnet": schema.StringAttribute{
				Description: "Network range (CIDR)",
				Required:    true,
			},
			"local_ip": schema.StringAttribute{
				Description: "Network range local ip",
				Optional:    true,
			},
			"translated_subnet": schema.StringAttribute{
				Description: "Network range translated native IP range (CIDR)",
				Optional:    true,
			},
			"gateway": schema.StringAttribute{
				Description: "Network range gateway",
				Optional:    true,
			},
			"vlan": schema.Int64Attribute{
				Description: "Network range VLAN ID",
				Optional:    true,
			},
			"dhcp_settings": schema.SingleNestedAttribute{
				Description: "Site native range DHCP settings",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"dhcp_type": schema.StringAttribute{
						Description: "Network range dhcp type",
						Required:    true,
					},
					"ip_range": schema.StringAttribute{
						Description: "Network range dhcp range",
						Optional:    true,
					},
					"relay_group_id": schema.StringAttribute{
						Description: "Network range dhcp relay group id",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *networkRangeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*catogo.Client)
}

func (r *networkRangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan NetworkRange
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.AddNetworkRangeInput{
		Name:             plan.Name.ValueString(),
		RangeType:        plan.RangeType.ValueString(),
		Subnet:           plan.Subnet.ValueString(),
		LocalIp:          plan.LocalIp.ValueStringPointer(),
		TranslatedSubnet: plan.TranslatedSubnet.ValueStringPointer(),
		Gateway:          plan.Gateway.ValueStringPointer(),
		Vlan:             plan.Vlan.ValueInt64Pointer(),
	}

	// get planned DHCP settings Object value, or set default value if null (for VLAN Type)
	var DhcpSettings DhcpSettings
	if plan.RangeType == types.StringValue("VLAN") {
		if plan.DhcpSettings.IsNull() {
			DhcpSettings.DhcpType = types.StringValue("DHCP_DISABLED")
		} else {
			diags = plan.DhcpSettings.As(ctx, &DhcpSettings, basetypes.ObjectAsOptions{})
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
		input.DhcpSettings = &catogo.NetworkDhcpSettingsInput{
			DhcpType:     *DhcpSettings.DhcpType.ValueStringPointer(),
			IpRange:      DhcpSettings.IpRange.ValueStringPointer(),
			RelayGroupId: DhcpSettings.RelayGroupId.ValueStringPointer(),
		}
	}

	// retrieving native-network range ID to update native range
	lan_interface, err := r.client.GetLanSocketInterfaceId(plan.SiteId.ValueString(), "LAN 01")
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	// DEBUGGGG
	jsonlog, _ := json.Marshal(input)
	tflog.Info(ctx, string(jsonlog))

	body, err := r.client.AddNetworkRange(lan_interface.Id, input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	plan.InterfaceId = types.StringValue(lan_interface.Id)
	plan.Id = types.StringValue(body.NetworkRangeId)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *networkRangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *networkRangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan NetworkRange
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.UpdateNetworkRangeInput{
		Name:             plan.Name.ValueStringPointer(),
		RangeType:        plan.RangeType.ValueStringPointer(),
		Subnet:           plan.Subnet.ValueStringPointer(),
		LocalIp:          plan.LocalIp.ValueStringPointer(),
		TranslatedSubnet: plan.TranslatedSubnet.ValueStringPointer(),
		Gateway:          plan.Gateway.ValueStringPointer(),
		Vlan:             plan.Vlan.ValueInt64Pointer(),
	}

	// get planned DHCP settings Object value, or set default value if null (for VLAN Type)
	var DhcpSettings DhcpSettings
	if plan.RangeType == types.StringValue("VLAN") {
		if plan.DhcpSettings.IsNull() {
			DhcpSettings.DhcpType = types.StringValue("DHCP_DISABLED")
		} else {
			diags = plan.DhcpSettings.As(ctx, &DhcpSettings, basetypes.ObjectAsOptions{})
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
		input.DhcpSettings = &catogo.NetworkDhcpSettingsInput{
			DhcpType:     *DhcpSettings.DhcpType.ValueStringPointer(),
			IpRange:      DhcpSettings.IpRange.ValueStringPointer(),
			RelayGroupId: DhcpSettings.RelayGroupId.ValueStringPointer(),
		}
	}

	// DEBUGGGG
	jsonlog, _ := json.Marshal(input)
	tflog.Info(ctx, string(jsonlog))

	_, err := r.client.UpdateNetworkRange(plan.Id.ValueString(), input)
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

func (r *networkRangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state NetworkRange
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// check if site exist before removing
	siteExists, err := r.client.SiteExists(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to connect or request the Cato API",
			err.Error(),
		)
		return
	}

	if siteExists {
		_, err := r.client.RemoveNetworkRange(state.Id.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to connect or request the Cato API",
				err.Error(),
			)
			return
		}
	}

}
