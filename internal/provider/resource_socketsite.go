package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	client *Client
}

type socketSiteResourceModel struct {
	AccountID          types.String              `tfsdk:"account_id"`
	ID                 types.String              `tfsdk:"id"`
	Name               types.String              `tfsdk:"name"`
	ConnectionType     types.String              `tfsdk:"connection_type"`
	SiteType           types.String              `tfsdk:"site_type"`
	Description        types.String              `tfsdk:"description"`
	NativeNetworkRange types.String              `tfsdk:"native_network_range"`
	SiteLocation       siteLocationResourceModel `tfsdk:"site_location"`
}

type siteLocationResourceModel struct {
	CountryCode types.String `tfsdk:"country_code"`
	StateCode   types.String `tfsdk:"state_code"`
	Timezone    types.String `tfsdk:"timezone"`
}

func (r *socketSiteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_socketsite"
}

func (r *socketSiteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cato Account ID (can be found into the URL on the CMA)",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier for the site",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Site name",
				Required:    true,
			},
			"connection_type": schema.StringAttribute{
				Description: "Connection type for the site (SOCKET_X1500, SOCKET_AWS1500, SOCKET_AZ1500, ...)",
				Required:    true,
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
				},
			},
		},
	}
}

func (d *socketSiteResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*Client)
}

func (r *socketSiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan socketSiteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	query := graphQLRequest{
		Query: `mutation addSocketSite($accountId:ID!, $input:AddSocketSiteInput!){
			site(accountId:$accountId) {
				addSocketSite(input:$input) {
					siteId
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": string(plan.AccountID.ValueString()),
			"input": &socketSite{
				Name:               string(plan.Name.ValueString()),
				Description:        string(plan.Description.ValueString()),
				SiteType:           string(plan.SiteType.ValueString()),
				NativeNetworkRange: string(plan.NativeNetworkRange.ValueString()),
				ConnectionType:     string(plan.ConnectionType.ValueString()),
				SiteLocation: siteLocation{
					CountryCode: string(plan.SiteLocation.CountryCode.ValueString()),
					StateCode:   string(plan.SiteLocation.StateCode.ValueString()),
					Timezone:    string(plan.SiteLocation.Timezone.ValueString()),
				},
			},
		},
	}

	// Cato Client logic, to be externalized
	body, err := r.client.do(query)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	// Cato Client logic, to be externalized
	var response mutationSSResponseModel
	err = json.Unmarshal(body, &response)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to process Cato API Response",
			err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(response.Data.Site.AddSocketSite.SiteID)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *socketSiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *socketSiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *socketSiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state socketSiteResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	query := graphQLRequest{
		Query: `mutation removeSite ($accountId: ID!, $siteId: ID!) { site(accountId: $accountId) { removeSite (siteId: $siteId) { siteId } } }`,
		Variables: map[string]interface{}{
			"accountId": state.AccountID.ValueString(),
			"siteId":    state.ID.ValueString(),
		},
	}
	_, err := r.client.do(query)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to connect or request the Cato API",
			err.Error(),
		)
		return
	}
}
