package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cato_models "github.com/routebyintuition/cato-go-sdk/models"
)

var (
	_ resource.Resource              = &siteIpsecResource{}
	_ resource.ResourceWithConfigure = &siteIpsecResource{}
)

func NewSiteIpsecResource() resource.Resource {
	return &siteIpsecResource{}
}

type siteIpsecResource struct {
	info *catoClientData
}

type SiteIpsecIkeV2 struct {
	ID                 types.String               `tfsdk:"id"`
	Name               types.String               `tfsdk:"name"`
	SiteType           types.String               `tfsdk:"sitetype"`
	Description        types.String               `tfsdk:"description"`
	NativeNetworkRange types.String               `tfsdk:"nativenetworkrange"`
	SiteLocation       *AddIpsecSiteLocationInput `tfsdk:"sitelocation"`
}

type AddIpsecSiteLocationInput struct {
	CountryCode types.String `tfsdk:"countrycode"`
	StateCode   types.String `tfsdk:"statecode"`
	Timezone    types.String `tfsdk:"timezone"`
	Address     types.String `tfsdk:"address"`
	City        types.String `tfsdk:"city"`
}

func (r *siteIpsecResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsec_site"
}

func (r *siteIpsecResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for Ipsec Site",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Ipsec Site Name",
				Required:    true,
			},
			"sitetype": schema.StringAttribute{
				Description: "Valid values are: BRANCH, HEADQUARTERS, CLOUD_DC, and DATACENTER.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "description",
				Required:    true,
			},
			"nativenetworkrange": schema.StringAttribute{
				Description: "NativeNetworkRange",
				Required:    true,
			},
			"sitelocation": schema.SingleNestedAttribute{
				Description: "SiteLocation",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"countrycode": schema.StringAttribute{
						Description: "Country Code",
						Required:    true,
					},
					"statecode": schema.StringAttribute{
						Description: "State Code",
						Required:    true,
					},
					"timezone": schema.StringAttribute{
						Description: "Timezone",
						Required:    true,
					},
					"address": schema.StringAttribute{
						Description: "Address",
						Required:    false,
						Optional:    true,
					},
					"city": schema.StringAttribute{
						Description: "City",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *siteIpsecResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.info = req.ProviderData.(*catoClientData)
}

func (r *siteIpsecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan SiteIpsecIkeV2
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := cato_models.AddIpsecIkeV2SiteInput{
		Name:               plan.Name.ValueString(),
		SiteType:           cato_models.SiteType(plan.SiteType.ValueString()),
		Description:        plan.Description.ValueStringPointer(),
		NativeNetworkRange: plan.NativeNetworkRange.ValueString(),
		SiteLocation: &cato_models.AddSiteLocationInput{
			CountryCode: plan.SiteLocation.CountryCode.ValueString(),
			StateCode:   plan.SiteLocation.StateCode.ValueStringPointer(),
			Timezone:    plan.SiteLocation.Timezone.ValueString(),
			Address:     plan.SiteLocation.Address.ValueStringPointer(),
			City:        plan.SiteLocation.City.ValueStringPointer(),
		},
	}

	body, err := r.info.catov2.SiteAddIpsecIkeV2Site(ctx, input, r.info.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(body.GetSite().GetAddIpsecIkeV2Site().SiteID)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteIpsecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SiteIpsecIkeV2
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	entityIds := []string{state.ID.ValueString()}

	queryResult, err := r.info.catov2.EntityLookup(ctx, r.info.AccountId, cato_models.EntityType("site"), nil, nil, nil, nil, entityIds, nil, nil, nil)
	if err != nil {
		tflog.Error(ctx, "Read: did not receive a response in read for EntityLookup")
		resp.Diagnostics.AddError(
			"Catov2 API EntityLookup error",
			err.Error(),
		)
		return
	}

	if len(queryResult.EntityLookup.GetItems()) == 1 {
		entryList := queryResult.EntityLookup.GetItems()
		entry := entryList[0]
		helperFields := entry.GetHelperFields()
		state.ID = types.StringValue(entry.Entity.ID)
		state.Name = types.StringValue(*entry.Entity.Name)
		state.SiteType = types.StringValue(helperFields["type"].(string))
		state.Description = types.StringValue(helperFields["description"].(string))

	} else {
		tflog.Error(ctx, "Read: did not receive a response in read for EntityLookup")
		resp.Diagnostics.AddError(
			"Catov2 API EntityLookup error",
			"more than one value returned in EntityLookup for SiteIpsecIkeV2",
		)
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteIpsecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan SiteIpsecIkeV2
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	currentId := plan.ID.ValueString()
	siteType := cato_models.SiteType(plan.SiteType.ValueString())

	input := cato_models.UpdateSiteGeneralDetailsInput{
		Name:        plan.Name.ValueStringPointer(),
		SiteType:    &siteType,
		Description: plan.Description.ValueStringPointer(),
		SiteLocation: &cato_models.UpdateSiteLocationInput{
			CountryCode: plan.SiteLocation.CountryCode.ValueStringPointer(),
			StateCode:   plan.SiteLocation.StateCode.ValueStringPointer(),
			Timezone:    plan.SiteLocation.Timezone.ValueStringPointer(),
			Address:     plan.SiteLocation.Address.ValueStringPointer(),
		},
	}

	_, err := r.info.catov2.SiteUpdateSiteGeneralDetails(ctx, currentId, input, r.info.AccountId)
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

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteIpsecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state SiteIpsecIkeV2
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	logIdEntry := "DELETING SITEIPSEC ID: " + state.ID.ValueString()
	tflog.Warn(ctx, logIdEntry)
	_, err := r.info.catov2.SiteRemoveSite(ctx, state.ID.ValueString(), r.info.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API v2 error in SiteRemoveSite",
			err.Error(),
		)
		return
	}

}
