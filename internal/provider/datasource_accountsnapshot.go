package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &accountSnapshotDataSource{}
	_ datasource.DataSourceWithConfigure = &accountSnapshotDataSource{}
)

func NewAccountSnapshotDataSource() datasource.DataSource {
	return &accountSnapshotDataSource{}
}

type accountSnapshotDataSource struct {
	client *Client
}

type accountSnapshotDataSourceModel struct {
	AccountID types.String            `tfsdk:"account_id"`
	SiteID    types.String            `tfsdk:"site_id"`
	Sites     []itemASDataSourceModel `tfsdk:"sites"`
}

type itemASDataSourceModel struct {
	Id   types.String          `tfsdk:"id"`
	Info infoASDataSourceModel `tfsdk:"info"`
}

type infoASDataSourceModel struct {
	Name    types.String              `tfsdk:"name"`
	Type    types.String              `tfsdk:"type"`
	Sockets []socketASDataSourceModel `tfsdk:"sockets"`
}

type socketASDataSourceModel struct {
	Serial types.String `tfsdk:"serial"`
}

func (d *accountSnapshotDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_accountSnapshot"
}

func (d *accountSnapshotDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cato Account ID (can be found into the URL on the CMA)",
				Required:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Identifier for the site",
				Required:    true,
			},
			"sites": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier for the site",
							Computed:    true,
						},
						"info": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "Site Name",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "Site Type",
									Computed:    true,
								},
								"sockets": schema.ListNestedAttribute{
									Description: "List of sockets attached to the site",
									Computed:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"serial": schema.StringAttribute{
												Description: "Socket serial number",
												Computed:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *accountSnapshotDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*Client)
}

func (d *accountSnapshotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state accountSnapshotDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	query := graphQLRequest{
		Query: `query accountSnapshot ($accountID: ID!, $siteIDs: [ID!] ) {
			accountSnapshot (accountID: $accountID) {
				sites (siteIDs: $siteIDs) {
					id
					info {
						name
						type
						sockets {
							id
							serial
						}
					}
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountID": string(state.AccountID.ValueString()),
			"siteIDs":   []string{state.SiteID.ValueString()},
		},
	}

	// Cato Client logic, to be externalized
	body, err := d.client.do(query)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	// Cato Client logic, to be externalized
	var response queryASResponseModel
	// var response queryResponseModel
	err = json.Unmarshal(body, &response)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to process Cato API Response",
			err.Error(),
		)
		return
	}

	// Cato Client logic, to be externalized
	for _, item := range response.Data.AccountSnapshot.Sites {

		itemState := itemASDataSourceModel{
			Id: types.StringValue(item.Id),
			Info: infoASDataSourceModel{
				Name: types.StringValue(item.Info.Name),
				Type: types.StringValue(item.Info.Type),
			},
		}

		for _, socket := range item.Info.Sockets {
			itemState.Info.Sockets = append(itemState.Info.Sockets, socketASDataSourceModel{
				Serial: types.StringValue(socket.Serial),
			})
		}

		state.Sites = append(state.Sites, itemState)

	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
