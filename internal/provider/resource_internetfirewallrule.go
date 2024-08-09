package provider

import (
	"context"
	"fmt"

	"github.com/BenEkpy/terraform-provider-cato-oss/internal/catogo"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &internetFirewallRuleResource{}
	_ resource.ResourceWithConfigure = &internetFirewallRuleResource{}
)

func NewInternetFirewallRuleResource() resource.Resource {
	return &internetFirewallRuleResource{}
}

type internetFirewallRuleResource struct {
	client *catogo.Client
}

type InternetFirewallRule struct {
	Id          types.String `tfsdk:"id"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Action      types.String `tfsdk:"action"`
	Position    types.String `tfsdk:"position"`
	Source      types.Object `tfsdk:"source"`
}

type InternetFirewallRuleSourceInput struct {
	Ip         types.List                       `tfsdk:"ip"`
	Subnet     types.List                       `tfsdk:"subnet"`
	UsersGroup []InternetFirewallObjectRefInput `tfsdk:"users_group"`
	// UserGroups types.List `tfsdk:"users_group"`
	// UserGroups []types.Object `tfsdk:"users_group"`
}

type InternetFirewallObjectRefInput struct {
	By    types.String `tfsdk:"by"`
	Input types.String `tfsdk:"input"`
}

func (r *internetFirewallRuleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_internet_firewall_rule"
}

func (r *internetFirewallRuleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier Internet Firewall rule",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Internet Firewall state",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Internet Firewall rule name",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Internet Firewall description",
				Optional:    true,
			},
			"source": schema.SingleNestedAttribute{
				Description: "Internet Firewall source",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"ip": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "Internet Firewall source IPs",
						Optional:    true,
					},
					"subnet": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "Internet Firewall source subnets",
						Optional:    true,
					},
					"users_group": schema.ListNestedAttribute{
						Description: "Internet Firewall source user groups",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"by": schema.StringAttribute{
									Description: "Internet Firewall source user groups type",
									Required:    true,
								},
								"input": schema.StringAttribute{
									Description: "Internet Firewall source user groups type entity name",
									Required:    true,
								},
							},
						},
					},
				},
			},
			"action": schema.StringAttribute{
				Description: "Internet Firewall rule action",
				Required:    true,
			},
			"position": schema.StringAttribute{
				Description: "Internet Firewall rule position",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (d *internetFirewallRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*catogo.Client)
}

func (r *internetFirewallRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan InternetFirewallRule
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var SourceInput InternetFirewallRuleSourceInput
	var SourceInputIp []string
	var SourceInputSubnet []string

	diags = plan.Source.As(ctx, &SourceInput, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = SourceInput.Ip.ElementsAs(ctx, &SourceInputIp, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = SourceInput.Subnet.ElementsAs(ctx, &SourceInputSubnet, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var items []catogo.ObjectRefInput
	for _, item := range SourceInput.UsersGroup {
		items = append(items, catogo.ObjectRefInput{
			By:    item.By.ValueString(),
			Input: item.Input.ValueString(),
		})
	}

	input := catogo.InternetFirewallAddRuleInput{
		Rule: catogo.InternetFirewallAddRuleDataInput{
			Enabled:     plan.Enabled.ValueBool(),
			Name:        plan.Name.ValueString(),
			Description: plan.Description.ValueString(),
			Action:      plan.Action.ValueString(),
			Source: &catogo.InternetFirewallSourceInput{
				Ip:         SourceInputIp,
				Subnet:     SourceInputSubnet,
				UsersGroup: items,
			},
		},
		At: &catogo.PolicyRulePositionInput{
			Position: plan.Position.ValueStringPointer(),
		},
	}

	body, err := r.client.CreateInternetFirewallRule(input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	_, err = r.client.PublishInternetFirewallDefaultPolicyRevision()
	if err != nil {
		fmt.Println("error:", err)
	}

	plan.Id = types.StringValue(body.Rule.Rule.Id)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *internetFirewallRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *internetFirewallRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan InternetFirewallRule
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var SourceInput InternetFirewallRuleSourceInput
	var SourceInputIp []string
	var SourceInputSubnet []string

	diags = SourceInput.Ip.ElementsAs(ctx, &SourceInputIp, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = SourceInput.Subnet.ElementsAs(ctx, &SourceInputSubnet, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var items []catogo.ObjectRefInput
	for _, item := range SourceInput.UsersGroup {
		items = append(items, catogo.ObjectRefInput{
			By:    item.By.ValueString(),
			Input: item.Input.ValueString(),
		})
	}

	input := catogo.InternetFirewallUpdateRuleInput{
		Id: plan.Id.ValueString(),
		Rule: catogo.InternetFirewallAddRuleDataInput{
			Enabled:     plan.Enabled.ValueBool(),
			Name:        plan.Name.ValueString(),
			Description: plan.Description.ValueString(),
			Action:      plan.Action.ValueString(),
			Source: &catogo.InternetFirewallSourceInput{
				Ip:         SourceInputIp,
				Subnet:     SourceInputSubnet,
				UsersGroup: items,
			},
		},
	}

	_, err := r.client.UpdateInternetFirewallRule(input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	_, err = r.client.PublishInternetFirewallDefaultPolicyRevision()
	if err != nil {
		fmt.Println("error:", err)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *internetFirewallRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state InternetFirewallRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.InternetFirewallRemoveRuleInput{
		Id: state.Id.ValueString(),
	}

	_, err := r.client.RemoveInternetFirewallRule(input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to connect or request the Cato API",
			err.Error(),
		)
		return
	}

	_, err = r.client.PublishInternetFirewallDefaultPolicyRevision()
	if err != nil {
		fmt.Println("error:", err)
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
