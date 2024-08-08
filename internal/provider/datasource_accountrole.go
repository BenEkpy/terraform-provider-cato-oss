package provider

import (
	"context"

	"github.com/BenEkpy/terraform-provider-cato-oss/internal/catogo"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource              = &AccountRoleDataSource{}
	_ datasource.DataSourceWithConfigure = &AccountRoleDataSource{}
)

func NewAccountRoleDataSource() datasource.DataSource {
	return &AccountRoleDataSource{}
}

type AccountRoleDataSource struct {
	client *catogo.Client
}

type RBACRole struct {
	Id          string `tfsdk:"id"`
	Name        string `tfsdk:"name"`
	Description string `tfsdk:"description"`
}

func (d *AccountRoleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_accountRole"
}

func (d *AccountRoleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Account Role ID",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Account Role name",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Account Role description",
				Optional:    true,
			},
		},
	}
}

func (d *AccountRoleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*catogo.Client)
}

func (d *AccountRoleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state RBACRole
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	body, err := d.client.GetAccountRoleByName(string(state.Name))
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	state = RBACRole{
		Id:          body.Id,
		Name:        body.Name,
		Description: *body.Description,
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
