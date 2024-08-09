package provider

import (
	"context"

	"github.com/BenEkpy/terraform-provider-cato-oss/internal/catogo"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &wanInterfaceResource{}
	_ resource.ResourceWithConfigure = &wanInterfaceResource{}
)

func NewWanInterfaceResource() resource.Resource {
	return &wanInterfaceResource{}
}

type wanInterfaceResource struct {
	client *catogo.Client
}

type WanInterface struct {
	SiteId              types.String `tfsdk:"site_id"`
	InterfaceID         types.String `tfsdk:"interface_id"`
	Name                types.String `tfsdk:"name"`
	UpstreamBandwidth   types.Int64  `tfsdk:"upstream_bandwidth"`
	DownstreamBandwidth types.Int64  `tfsdk:"downstream_bandwidth"`
	Role                types.String `tfsdk:"role"`
	Precedence          types.String `tfsdk:"precedence"`
}

func (r *wanInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wan_interface"
}

func (r *wanInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"interface_id": schema.StringAttribute{
				Description: "WAN Interface id",
				Required:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Host Site ID",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Description: "WAN interface name",
				Required:    true,
			},
			"upstream_bandwidth": schema.Int64Attribute{
				Description: "WAN interface upstream bandwitdh",
				Required:    true,
			},
			"downstream_bandwidth": schema.Int64Attribute{
				Description: "WAN interface downstream bandwitdh",
				Required:    true,
			},
			"role": schema.StringAttribute{
				Description: "WAN interface role",
				Required:    true,
			},
			"precedence": schema.StringAttribute{
				Description: "WAN interface precedence",
				Required:    true,
			},
		},
	}
}

func (d *wanInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*catogo.Client)
}

func (r *wanInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan WanInterface
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.UpdateSocketInterfaceInput{
		DestType: "CATO",
		Name:     plan.Name.ValueStringPointer(),
		Bandwidth: &catogo.SocketInterfaceBandwidthInput{
			UpstreamBandwidth:   *plan.UpstreamBandwidth.ValueInt64Pointer(),
			DownstreamBandwidth: *plan.DownstreamBandwidth.ValueInt64Pointer(),
		},
		Wan: &catogo.SocketInterfaceWanInput{
			Role:       *plan.Role.ValueStringPointer(),
			Precedence: *plan.Precedence.ValueStringPointer(),
		},
	}
	_, err := r.client.UpdateSocketInterface(plan.SiteId.ValueString(), plan.InterfaceID.ValueString(), input)
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

func (r *wanInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *wanInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan WanInterface
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.UpdateSocketInterfaceInput{
		DestType: "CATO",
		Name:     plan.Name.ValueStringPointer(),
		Bandwidth: &catogo.SocketInterfaceBandwidthInput{
			UpstreamBandwidth:   *plan.UpstreamBandwidth.ValueInt64Pointer(),
			DownstreamBandwidth: *plan.DownstreamBandwidth.ValueInt64Pointer(),
		},
		Wan: &catogo.SocketInterfaceWanInput{
			Role:       *plan.Role.ValueStringPointer(),
			Precedence: *plan.Precedence.ValueStringPointer(),
		},
	}

	_, err := r.client.UpdateSocketInterface(plan.SiteId.ValueString(), plan.InterfaceID.ValueString(), input)
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

func (r *wanInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state WanInterface
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

	var input = catogo.UpdateSocketInterfaceInput{}
	if siteExists {

		// Check if there is only one WAN interface & rewrite the input with default one
		wanInterfaceList, err := r.client.GetSocketWanInterfacelist(state.SiteId.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to connect or request the Cato API",
				err.Error(),
			)
			return
		}

		if len(wanInterfaceList) == 1 {
			defaultName := "Default WAN Interface"
			input = catogo.UpdateSocketInterfaceInput{
				DestType: "CATO",
				Name:     &defaultName,
				Bandwidth: &catogo.SocketInterfaceBandwidthInput{
					UpstreamBandwidth:   10,
					DownstreamBandwidth: 10,
				},
				Wan: &catogo.SocketInterfaceWanInput{
					Role:       "wan_1",
					Precedence: "ACTIVE",
				},
			}
		} else {
			// Disabled interface to "remove" an interface
			input = catogo.UpdateSocketInterfaceInput{
				DestType: "INTERFACE_DISABLED",
			}
		}

		_, err = r.client.UpdateSocketInterface(state.SiteId.ValueString(), state.InterfaceID.ValueString(), input)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to connect or request the Cato API",
				err.Error(),
			)
			return
		}
	}
}
