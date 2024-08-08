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
	_ resource.Resource              = &adminResource{}
	_ resource.ResourceWithConfigure = &adminResource{}
)

func NewAdminResource() resource.Resource {
	return &adminResource{}
}

type adminResource struct {
	client *catogo.Client
}

type Admin struct {
	Id                   types.String           `tfsdk:"id"`
	FirstName            types.String           `tfsdk:"firstname"`
	LastName             types.String           `tfsdk:"lastname"`
	Email                types.String           `tfsdk:"email"`
	PasswordNeverExpires types.Bool             `tfsdk:"password_never_expires"`
	MfaEnabled           types.Bool             `tfsdk:"mfa_enabled"`
	ManagedRoles         []UpdateAdminRoleInput `tfsdk:"managed_roles"`
	// ResellerRoles        []UpdateAdminRoleInput `tfsdk:"reseller_roles"`
}

type UpdateAdminRoleInput struct {
	Role            UpdateAccountRoleInput `tfsdk:"role"`
	AllowedEntities []EntityInput          `tfsdk:"allowed_entities"`
	AllowedAccounts []types.String         `tfsdk:"allowed_accounts"`
}

type UpdateAccountRoleInput struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type EntityInput struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func (r *adminResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_admin"
}

func (r *adminResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for the site",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"firstname": schema.StringAttribute{
				Description: "Admin firstname",
				Required:    true,
			},
			"lastname": schema.StringAttribute{
				Description: "Admin lastname",
				Required:    true,
			},
			"email": schema.StringAttribute{
				Description: "Admin email",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"password_never_expires": schema.BoolAttribute{
				Description: "Admin password never expires",
				Required:    true,
			},
			"mfa_enabled": schema.BoolAttribute{
				Description: "Enables MFA for the admin auth",
				Required:    true,
			},
			"managed_roles": schema.ListNestedAttribute{
				Description: "List of managed roles",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"role": schema.SingleNestedAttribute{
							Description: "Admin Role",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Admin Role id",
									Required:    true,
								},
								"name": schema.StringAttribute{
									Description: "Admin Role name",
									Optional:    true,
								},
							},
						},
						"allowed_entities": schema.ListNestedAttribute{
							Description: "List of entities allowed",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Entity ID",
										Required:    true,
									},
									"name": schema.StringAttribute{
										Description: "Entity Name",
										Optional:    true,
									},
									"type": schema.StringAttribute{
										Description: "Entity Type",
										Required:    true,
									},
								},
							},
						},
						"allowed_accounts": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (d *adminResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*catogo.Client)
}

func (r *adminResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan Admin
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.AddAdminInput{
		FirstName:            plan.FirstName.ValueString(),
		LastName:             plan.LastName.ValueString(),
		Email:                plan.Email.ValueString(),
		PasswordNeverExpires: plan.PasswordNeverExpires.ValueBoolPointer(),
		MfaEnabled:           plan.MfaEnabled.ValueBoolPointer(),
	}

	// append managed_roles
	for index_managed_role, managed_role := range plan.ManagedRoles {
		input.ManagedRoles = append(input.ManagedRoles, catogo.UpdateAdminRoleInput{
			Role: catogo.UpdateAccountRoleInput{
				Id:   managed_role.Role.Id.ValueString(),
				Name: managed_role.Role.Name.ValueStringPointer(),
			},
		})

		// append allowed_entites
		for _, allowed_entitie := range plan.ManagedRoles[index_managed_role].AllowedEntities {
			input.ManagedRoles[index_managed_role].AllowedEntities = append(input.ManagedRoles[index_managed_role].AllowedEntities, catogo.EntityInput{
				Id:   allowed_entitie.Id.ValueString(),
				Name: allowed_entitie.Name.ValueStringPointer(),
				Type: allowed_entitie.Type.String(),
			},
			)
		}

		// append allowed_account
		for _, allowed_account := range plan.ManagedRoles[index_managed_role].AllowedAccounts {
			input.ManagedRoles[index_managed_role].AllowedAccounts = append(input.ManagedRoles[index_managed_role].AllowedAccounts, allowed_account.ValueString())
		}
	}

	body, err := r.client.AddAdmin(input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	plan.Id = types.StringValue(body.AdminID)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *adminResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state Admin

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	body, err := r.client.GetAdmin(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Cato API error",
			err.Error(),
		)
		return
	}

	state = Admin{
		Id:                   types.StringValue(body.Id),
		FirstName:            types.StringValue(body.FirstName),
		LastName:             types.StringValue(body.LastName),
		Email:                types.StringValue(body.Email),
		PasswordNeverExpires: types.BoolValue(body.PasswordNeverExpires),
		MfaEnabled:           types.BoolValue(body.MfaEnabled),
	}

	// append managed_roles
	for index_managed_role, managed_role := range body.ManagedRoles {
		state.ManagedRoles = append(state.ManagedRoles, UpdateAdminRoleInput{
			Role: UpdateAccountRoleInput{
				Id:   types.StringValue(managed_role.Role.Id),
				Name: types.StringValue(managed_role.Role.Name),
			},
		})

		// append allowed_entites
		for _, allowed_entitie := range body.ManagedRoles[index_managed_role].AllowedEntities {
			state.ManagedRoles[index_managed_role].AllowedEntities = append(state.ManagedRoles[index_managed_role].AllowedEntities, EntityInput{
				Id:   types.StringValue(allowed_entitie.Id),
				Name: types.StringPointerValue(allowed_entitie.Name),
				Type: types.StringValue(allowed_entitie.Type),
			},
			)
		}

		// append allowed_account
		for _, allowed_account := range body.ManagedRoles[index_managed_role].AllowedAccounts {
			state.ManagedRoles[index_managed_role].AllowedAccounts = append(state.ManagedRoles[index_managed_role].AllowedAccounts, types.StringValue(allowed_account))
		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *adminResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan Admin
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := catogo.UpdateAdminInput{
		FirstName:            plan.FirstName.ValueStringPointer(),
		LastName:             plan.LastName.ValueStringPointer(),
		PasswordNeverExpires: plan.PasswordNeverExpires.ValueBoolPointer(),
		MfaEnabled:           plan.MfaEnabled.ValueBoolPointer(),
	}

	// append managed_roles
	for index_managed_role, managed_role := range plan.ManagedRoles {
		input.ManagedRoles = append(input.ManagedRoles, catogo.UpdateAdminRoleInput{
			Role: catogo.UpdateAccountRoleInput{
				Id:   managed_role.Role.Id.ValueString(),
				Name: managed_role.Role.Name.ValueStringPointer(),
			},
		})

		// append allowed_entites
		for _, allowed_entitie := range plan.ManagedRoles[index_managed_role].AllowedEntities {
			input.ManagedRoles[index_managed_role].AllowedEntities = append(input.ManagedRoles[index_managed_role].AllowedEntities, catogo.EntityInput{
				Id:   allowed_entitie.Id.ValueString(),
				Name: allowed_entitie.Name.ValueStringPointer(),
				Type: allowed_entitie.Type.String(),
			},
			)
		}

		// append allowed_account
		for _, allowed_account := range plan.ManagedRoles[index_managed_role].AllowedAccounts {
			input.ManagedRoles[index_managed_role].AllowedAccounts = append(input.ManagedRoles[index_managed_role].AllowedAccounts, allowed_account.ValueString())
		}
	}
	// }

	_, err := r.client.UpdateAdmin(plan.Id.ValueString(), input)
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

func (r *adminResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state Admin
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.RemoveAdmin(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to connect or request the Cato API",
			err.Error(),
		)
		return
	}
}
