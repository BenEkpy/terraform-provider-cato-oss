package provider

// import (
// 	"context"
// 	"fmt"

// 	"github.com/hashicorp/terraform-plugin-framework/datasource"
// 	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
// 	"github.com/hashicorp/terraform-plugin-framework/types"
// 	cato_models "github.com/routebyintuition/cato-go-sdk/models"
// )

// // Ensure the implementation satisfies the expected interfaces.
// var (
// 	_ datasource.DataSource              = &InternetFwPolicyDataSource{}
// 	_ datasource.DataSourceWithConfigure = &InternetFwPolicyDataSource{}
// )

// func NewInternetFwPolicyDataSource() datasource.DataSource {
// 	return &InternetFwPolicyDataSource{}
// }

// type InternetFwPolicyDataSource struct {
// 	client *catoClientData
// }

// func (d *InternetFwPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
// 	if req.ProviderData == nil {
// 		return
// 	}

// 	//dataConfig := req.ProviderData.(*catoClientData)
// 	client, ok := req.ProviderData.(*catoClientData)
// 	if !ok {
// 		resp.Diagnostics.AddError(
// 			"Unexpected Data Source Configure Type",
// 			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
// 		)

// 		return
// 	}

// 	d.client = client
// }

// func (d *InternetFwPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
// 	resp.TypeName = req.ProviderTypeName + "_internet_fw_policy"
// }

// func (d *InternetFwPolicyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
// 	resp.Schema = schema.Schema{
// 		Attributes: map[string]schema.Attribute{
// 			"enabled": schema.BoolAttribute{
// 				Computed: true,
// 			},
// 			"sections": schema.ListNestedAttribute{
// 				Computed: true,
// 				NestedObject: schema.NestedAttributeObject{
// 					Attributes: map[string]schema.Attribute{
// 						"audit": schema.SingleNestedAttribute{
// 							Computed: true,
// 							Attributes: map[string]schema.Attribute{
// 								"updatedtime": schema.StringAttribute{
// 									Computed: true,
// 								},
// 								"updatedby": schema.StringAttribute{
// 									Computed: true,
// 								},
// 							},
// 						},
// 						"section": schema.SingleNestedAttribute{
// 							Computed: true,
// 							Attributes: map[string]schema.Attribute{
// 								"id": schema.StringAttribute{
// 									Computed: true,
// 								},
// 								"name": schema.StringAttribute{
// 									Computed: true,
// 								},
// 							},
// 						},
// 						"properties": schema.ListAttribute{
// 							Computed:    true,
// 							ElementType: types.StringType,
// 						},
// 					},
// 				},
// 			},
// 			"audit": schema.SingleNestedAttribute{
// 				Computed: true,
// 				Attributes: map[string]schema.Attribute{
// 					"publishedtime": schema.StringAttribute{
// 						Computed: true,
// 					},
// 					"publishedby": schema.StringAttribute{
// 						Computed: true,
// 					},
// 				},
// 			},
// 			"revision": schema.SingleNestedAttribute{
// 				Computed: true,
// 				Attributes: map[string]schema.Attribute{
// 					"id": schema.StringAttribute{
// 						Computed: true,
// 					},
// 					"name": schema.StringAttribute{
// 						Computed: true,
// 					},
// 					"description": schema.StringAttribute{
// 						Computed: true,
// 					},
// 					"changes": schema.Int64Attribute{
// 						Computed: true,
// 					},
// 					"createdtime": schema.StringAttribute{
// 						Computed: true,
// 					},
// 					"updatedtime": schema.StringAttribute{
// 						Computed: true,
// 					},
// 				},
// 			},
// 			"rules": schema.ListNestedAttribute{
// 				Computed: true,
// 				NestedObject: schema.NestedAttributeObject{
// 					Attributes: map[string]schema.Attribute{
// 						"audit": schema.SingleNestedAttribute{
// 							Computed: true,
// 							Attributes: map[string]schema.Attribute{
// 								"updatedtime": schema.StringAttribute{
// 									Computed: true,
// 								},
// 								"updatedby": schema.StringAttribute{
// 									Computed: true,
// 								},
// 							},
// 						},
// 						"rule": schema.SingleNestedAttribute{
// 							Computed: true,
// 							Attributes: map[string]schema.Attribute{
// 								"id": schema.StringAttribute{
// 									Computed: true,
// 								},
// 								"name": schema.StringAttribute{
// 									Computed: true,
// 								},
// 								"description": schema.StringAttribute{
// 									Computed: true,
// 								},
// 								"index": schema.Int64Attribute{
// 									Computed: true,
// 								},
// 								"section": schema.SingleNestedAttribute{
// 									Computed: true,
// 									Attributes: map[string]schema.Attribute{
// 										"id": schema.StringAttribute{
// 											Computed: true,
// 										},
// 										"name": schema.StringAttribute{
// 											Computed: true,
// 										},
// 									},
// 								},
// 								"enabled": schema.BoolAttribute{
// 									Computed: true,
// 								},
// 								"source": schema.SingleNestedAttribute{
// 									Optional: true,
// 									Attributes: map[string]schema.Attribute{
// 										"ip": schema.ListAttribute{
// 											Computed:    true,
// 											ElementType: types.StringType,
// 										},
// 										"host": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"site": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"subnet": schema.ListAttribute{
// 											Computed:    true,
// 											ElementType: types.StringType,
// 										},
// 										"iprange": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"from": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"to": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"globaliprange": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"networkinterface": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"sitenetworksubnet": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"floatingsubnet": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"user": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"usersgroup": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"group": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"systemgroup": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 								"connectionorigin": schema.StringAttribute{
// 									Computed: true,
// 									Optional: true,
// 								},
// 								"country": schema.ListNestedAttribute{
// 									Computed: true,
// 									NestedObject: schema.NestedAttributeObject{
// 										Attributes: map[string]schema.Attribute{
// 											"id": schema.StringAttribute{
// 												Computed: true,
// 											},
// 											"name": schema.StringAttribute{
// 												Computed: true,
// 											},
// 										},
// 									},
// 								},
// 								"device": schema.ListNestedAttribute{
// 									Computed: true,
// 									NestedObject: schema.NestedAttributeObject{
// 										Attributes: map[string]schema.Attribute{
// 											"id": schema.StringAttribute{
// 												Computed: true,
// 											},
// 											"name": schema.StringAttribute{
// 												Computed: true,
// 											},
// 										},
// 									},
// 								},
// 								"deviceos": schema.ListAttribute{
// 									Computed:    true,
// 									ElementType: types.StringType,
// 									Optional:    true,
// 									Required:    false,
// 								},
// 								"destination": schema.SingleNestedAttribute{
// 									Computed: true,
// 									Optional: true,
// 									Attributes: map[string]schema.Attribute{
// 										"application": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"customapp": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"appcategory": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"customcategory": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"sanctionedappscategory": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"country": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"domain": schema.ListAttribute{
// 											Computed:    true,
// 											ElementType: types.StringType,
// 											Required:    false,
// 										},
// 										"fqdn": schema.ListAttribute{
// 											Computed:    true,
// 											ElementType: types.StringType,
// 											Required:    false,
// 										},
// 										"ip": schema.ListAttribute{
// 											Computed:    true,
// 											ElementType: types.StringType,
// 											Required:    false,
// 										},
// 										"subnet": schema.ListAttribute{
// 											Computed:    true,
// 											ElementType: types.StringType,
// 											Required:    false,
// 										},
// 										"iprange": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"from": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"to": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"globaliprange": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"remoteasn": schema.ListAttribute{
// 											Computed:    true,
// 											ElementType: types.StringType,
// 											Required:    false,
// 										},
// 									},
// 								},
// 								"service": schema.SingleNestedAttribute{
// 									Optional: true,
// 									Attributes: map[string]schema.Attribute{
// 										"standard": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 										"custom": schema.ListNestedAttribute{
// 											Computed: true,
// 											NestedObject: schema.NestedAttributeObject{
// 												Attributes: map[string]schema.Attribute{
// 													"id": schema.StringAttribute{
// 														Computed: true,
// 													},
// 													"name": schema.StringAttribute{
// 														Computed: true,
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 								"action": schema.StringAttribute{
// 									Computed: true,
// 									Optional: true,
// 								},
// 								"tracking": schema.SingleNestedAttribute{
// 									Computed: true,
// 									Optional: true,
// 									Attributes: map[string]schema.Attribute{
// 										"event": schema.SingleNestedAttribute{
// 											Computed: true,

// 											Attributes: map[string]schema.Attribute{
// 												"enabled": schema.BoolAttribute{
// 													Computed: true,
// 												},
// 											},
// 										},
// 										"alert": schema.SingleNestedAttribute{
// 											Computed: true,
// 											Attributes: map[string]schema.Attribute{
// 												"enabled": schema.BoolAttribute{
// 													Computed: true,
// 												},
// 												"frequency": schema.StringAttribute{
// 													Computed: true,
// 												},
// 												"subscriptiongroup": schema.ListNestedAttribute{
// 													Computed: true,
// 													NestedObject: schema.NestedAttributeObject{
// 														Attributes: map[string]schema.Attribute{
// 															"id": schema.StringAttribute{
// 																Computed: true,
// 															},
// 															"name": schema.StringAttribute{
// 																Computed: true,
// 															},
// 														},
// 													},
// 												},
// 												"webhook": schema.ListNestedAttribute{
// 													Computed: true,
// 													NestedObject: schema.NestedAttributeObject{
// 														Attributes: map[string]schema.Attribute{
// 															"id": schema.StringAttribute{
// 																Computed: true,
// 															},
// 															"name": schema.StringAttribute{
// 																Computed: true,
// 															},
// 														},
// 													},
// 												},
// 												"mailinglist": schema.ListNestedAttribute{
// 													Computed: true,
// 													NestedObject: schema.NestedAttributeObject{
// 														Attributes: map[string]schema.Attribute{
// 															"id": schema.StringAttribute{
// 																Computed: true,
// 															},
// 															"name": schema.StringAttribute{
// 																Computed: true,
// 															},
// 														},
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 								"schedule": schema.SingleNestedAttribute{
// 									Computed: true,
// 									Optional: true,
// 									Attributes: map[string]schema.Attribute{
// 										"activeon": schema.StringAttribute{
// 											Computed: true,
// 										},
// 										"customtimeframe": schema.StringAttribute{
// 											Computed: true,
// 										},
// 										"customrecurring": schema.StringAttribute{
// 											Computed: true,
// 										},
// 									},
// 								},
// 								"exceptions": schema.ListNestedAttribute{
// 									Computed: true,
// 									NestedObject: schema.NestedAttributeObject{
// 										Attributes: map[string]schema.Attribute{
// 											"name": schema.StringAttribute{
// 												Computed: true,
// 											},
// 											"deviceos": schema.ListAttribute{
// 												Computed:    true,
// 												Optional:    true,
// 												ElementType: types.StringType,
// 											},
// 											"source": schema.SingleNestedAttribute{
// 												Computed: true,
// 												Optional: true,
// 												Attributes: map[string]schema.Attribute{
// 													"ip": schema.ListAttribute{
// 														Computed:    true,
// 														ElementType: types.StringType,
// 													},
// 													"subnet": schema.ListAttribute{
// 														Computed:    true,
// 														ElementType: types.StringType,
// 													},
// 												},
// 											},
// 											"country": schema.ListNestedAttribute{
// 												Computed: true,
// 												NestedObject: schema.NestedAttributeObject{
// 													Attributes: map[string]schema.Attribute{
// 														"id": schema.StringAttribute{
// 															Computed: true,
// 														},
// 														"name": schema.StringAttribute{
// 															Computed: true,
// 														},
// 													},
// 												},
// 											},
// 											"device": schema.ListNestedAttribute{
// 												Computed: true,
// 												NestedObject: schema.NestedAttributeObject{
// 													Attributes: map[string]schema.Attribute{
// 														"id": schema.StringAttribute{
// 															Computed: true,
// 														},
// 														"name": schema.StringAttribute{
// 															Computed: true,
// 														},
// 													},
// 												},
// 											},
// 											"destination": schema.SingleNestedAttribute{
// 												Computed: true,
// 												Attributes: map[string]schema.Attribute{
// 													"domain": schema.ListAttribute{
// 														Computed:    true,
// 														ElementType: types.StringType,
// 													},
// 													"fqdn": schema.ListAttribute{
// 														Computed:    true,
// 														ElementType: types.StringType,
// 													},
// 													"ip": schema.ListAttribute{
// 														Computed:    true,
// 														ElementType: types.StringType,
// 													},
// 													"subnet": schema.ListAttribute{
// 														Computed:    true,
// 														ElementType: types.StringType,
// 													},
// 													"remoteasn": schema.ListAttribute{
// 														Computed:    true,
// 														ElementType: types.StringType,
// 													},
// 												},
// 											},
// 											"connectionorigin": schema.StringAttribute{
// 												Computed: true,
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 						"properties": schema.ListAttribute{
// 							Computed:    true,
// 							Optional:    true,
// 							ElementType: types.StringType,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// type internetFwPolicyList struct {
// 	ID      types.String                    "tfsdk:\"id\" graphql:\"id\""
// 	Enabled types.Bool                      `tfsdk:"enabled"`
// 	Rules   []InternetFirewall_Policy_Rules `tfsdk:"rules"`
// 	// change rules type
// 	Sections []InternetFirewall_Policy_Sections `tfsdk:"sections"`
// 	// change rules type
// 	Audit InternetFirewall_Policy_Audit `tfsdk:"audit"`
// 	// change rules type
// 	Revision InternetFirewall_Policy_Revision `tfsdk:"revision"`
// }

// func (d *InternetFwPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
// 	var state internetFwPolicyList

// 	queryPolicy := &cato_models.InternetFirewallPolicyInput{}
// 	body, err := d.client.catov2.Policy(ctx, queryPolicy, d.client.AccountId)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Catov2 API error in (d *InternetFwPolicyDataSource) Read",
// 			err.Error(),
// 		)
// 		return
// 	}

// 	state.Enabled = types.BoolValue(body.Policy.InternetFirewall.Policy.Enabled)
// 	state.Rules = make([]InternetFirewall_Policy_Rules, 0)
// 	state.Sections = make([]InternetFirewall_Policy_Sections, 0)

// 	auditData := InternetFirewall_Policy_Audit{
// 		PublishedTime: types.StringValue(body.GetPolicy().InternetFirewall.Policy.Audit.GetPublishedTime()),
// 		PublishedBy:   types.StringValue(body.GetPolicy().InternetFirewall.Policy.Audit.GetPublishedBy()),
// 	}
// 	state.Audit = auditData

// 	revisionData := InternetFirewall_Policy_Revision{
// 		ID:          types.StringValue(body.GetPolicy().InternetFirewall.Policy.Revision.GetID()),
// 		Name:        types.StringValue(body.GetPolicy().InternetFirewall.Policy.Revision.GetName()),
// 		Description: types.StringValue(body.GetPolicy().InternetFirewall.Policy.Revision.GetDescription()),
// 		Changes:     types.Int64Value(body.GetPolicy().InternetFirewall.Policy.Revision.GetChanges()),
// 		CreatedTime: types.StringValue(body.GetPolicy().InternetFirewall.Policy.Revision.GetCreatedTime()),
// 		UpdatedTime: types.StringValue(body.GetPolicy().InternetFirewall.Policy.Revision.GetUpdatedTime()),
// 	}

// 	state.Revision = revisionData

// 	for _, sectionEntry := range body.GetPolicy().InternetFirewall.Policy.GetSections() {
// 		sectionAuditItem := Policy_Policy_InternetFirewall_Policy_Sections_Audit{}
// 		sectionSectionItem := Policy_Policy_InternetFirewall_Policy_Sections_Section{}

// 		sectionItem := InternetFirewall_Policy_Sections{
// 			Properties: makeStringListFromStringSlice(sectionEntry.GetProperties()),
// 		}

// 		sectionItem.Audit = sectionAuditItem
// 		sectionItem.Section = sectionSectionItem
// 		state.Sections = append(state.Sections, sectionItem)
// 	}

// 	for _, ifRule := range body.GetPolicy().InternetFirewall.Policy.GetRules() {
// 		ipRuleItemEntry := InternetFirewall_Policy_Rules{}
// 		ipRuleItemEntryAudit := Policy_Policy_InternetFirewall_Policy_Rules_Audit{}
// 		ipRuleItemEntryRule := Policy_Policy_InternetFirewall_Policy_Rules_Rule{
// 			Source:      Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source{},
// 			Destination: Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination{},
// 		}

// 		ipRuleItemEntryRule.ID = types.StringValue(ifRule.Rule.GetID())
// 		ipRuleItemEntryRule.Name = types.StringValue(ifRule.Rule.GetName())
// 		ipRuleItemEntryRule.Description = types.StringValue(ifRule.Rule.GetDescription())
// 		ipRuleItemEntryRule.Index = types.Int64Value(ifRule.Rule.GetIndex())
// 		ipRuleItemEntryRule.Enabled = types.BoolValue(ifRule.Rule.GetEnabled())

// 		// DESTINATION_STARTS_HERE
// 		ipRuleItemEntryRule.Destination.Domain = makeStringListFromStringSlice(ifRule.Rule.Destination.GetDomain())
// 		ipRuleItemEntryRule.Destination.Fqdn = makeStringListFromStringSlice(ifRule.Rule.Destination.GetFqdn())
// 		ipRuleItemEntryRule.Destination.IP = makeStringListFromStringSlice(ifRule.Rule.Destination.GetIP())
// 		ipRuleItemEntryRule.Destination.Subnet = makeStringListFromStringSlice(ifRule.Rule.Destination.GetSubnet())
// 		ipRuleItemEntryRule.Destination.RemoteAsn = makeStringListFromStringSlice(ifRule.Rule.Destination.GetRemoteAsn())

// 		getRuleDestinationApplication := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Application{}
// 		for _, val := range ifRule.Rule.Destination.GetApplication() {
// 			getRuleDestinationApplication = append(getRuleDestinationApplication, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Application{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Destination.Application = getRuleDestinationApplication

// 		getRuleDestinationCustomApp := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomApp{}
// 		for _, val := range ifRule.Rule.Destination.GetCustomApp() {
// 			getRuleDestinationCustomApp = append(getRuleDestinationCustomApp, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomApp{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Destination.CustomApp = getRuleDestinationCustomApp

// 		getRuleDestinationAppCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_AppCategory{}
// 		for _, val := range ifRule.Rule.Destination.GetAppCategory() {
// 			getRuleDestinationAppCategory = append(getRuleDestinationAppCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_AppCategory{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Destination.AppCategory = getRuleDestinationAppCategory

// 		getRuleDestinationCustomCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomCategory{}
// 		for _, val := range ifRule.Rule.Destination.GetCustomCategory() {
// 			getRuleDestinationCustomCategory = append(getRuleDestinationCustomCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomCategory{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Destination.CustomCategory = getRuleDestinationCustomCategory

// 		getRuleDestinationSanctionedAppsCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_SanctionedAppsCategory{}
// 		for _, val := range ifRule.Rule.Destination.GetSanctionedAppsCategory() {
// 			getRuleDestinationSanctionedAppsCategory = append(getRuleDestinationSanctionedAppsCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_SanctionedAppsCategory{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Destination.SanctionedAppsCategory = getRuleDestinationSanctionedAppsCategory

// 		getRuleDestinationCountry := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Country{}
// 		for _, val := range ifRule.Rule.Destination.GetCountry() {
// 			getRuleDestinationCountry = append(getRuleDestinationCountry, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Country{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Destination.Country = getRuleDestinationCountry

// 		getRuleDestinationIPRange := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_IPRange{}
// 		for _, val := range ifRule.Rule.Destination.GetIPRange() {
// 			getRuleDestinationIPRange = append(getRuleDestinationIPRange, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_IPRange{
// 				From: types.StringValue(val.GetFrom()),
// 				To:   types.StringValue(val.GetTo()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Destination.IPRange = getRuleDestinationIPRange

// 		getRuleDestinationGlobalIPRange := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_GlobalIPRange{}
// 		for _, val := range ifRule.Rule.Destination.GetGlobalIPRange() {
// 			getRuleDestinationGlobalIPRange = append(getRuleDestinationGlobalIPRange, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_GlobalIPRange{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Destination.GlobalIPRange = getRuleDestinationGlobalIPRange

// 		// START_OF_SOURCE_DEFINITIONS:

// 		ipRuleItemEntryRule.Source.IP = makeStringListFromStringSlice(ifRule.Rule.Source.GetIP())
// 		ipRuleItemEntryRule.Source.Subnet = makeStringListFromStringSlice(ifRule.Rule.Source.GetSubnet())

// 		getRuleSourceHost := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Host{}
// 		for _, val := range ifRule.Rule.Source.GetHost() {
// 			getRuleSourceHost = append(getRuleSourceHost, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Host{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.Host = getRuleSourceHost

// 		getRuleSourceSite := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Site{}
// 		for _, val := range ifRule.Rule.Source.GetSite() {
// 			getRuleSourceSite = append(getRuleSourceSite, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Site{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.Site = getRuleSourceSite

// 		getRuleSourceIPRange := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_IPRange{}
// 		for _, val := range ifRule.Rule.Source.GetIPRange() {
// 			getRuleSourceIPRange = append(getRuleSourceIPRange, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_IPRange{
// 				From: types.StringValue(val.GetFrom()),
// 				To:   types.StringValue(val.GetTo()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.IPRange = getRuleSourceIPRange

// 		getRuleSourceGlobalIpRange := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_GlobalIPRange{}
// 		for _, val := range ifRule.Rule.Source.GetGlobalIPRange() {
// 			getRuleSourceGlobalIpRange = append(getRuleSourceGlobalIpRange, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_GlobalIPRange{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.GlobalIPRange = getRuleSourceGlobalIpRange

// 		getRuleSourceNetworkInterface := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_NetworkInterface{}
// 		for _, val := range ifRule.Rule.Source.GetNetworkInterface() {
// 			getRuleSourceNetworkInterface = append(getRuleSourceNetworkInterface, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_NetworkInterface{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.NetworkInterface = getRuleSourceNetworkInterface

// 		getSourceSiteNetworkSubnet := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_SiteNetworkSubnet{}
// 		for _, val := range ifRule.Rule.Source.GetSiteNetworkSubnet() {
// 			getSourceSiteNetworkSubnet = append(getSourceSiteNetworkSubnet, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_SiteNetworkSubnet{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.SiteNetworkSubnet = getSourceSiteNetworkSubnet

// 		getSourceFloatingSubnet := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_FloatingSubnet{}
// 		for _, val := range ifRule.Rule.Source.GetFloatingSubnet() {
// 			getSourceFloatingSubnet = append(getSourceFloatingSubnet, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_FloatingSubnet{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.FloatingSubnet = getSourceFloatingSubnet

// 		getSourceUser := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_User{}
// 		for _, val := range ifRule.Rule.Source.GetUser() {
// 			getSourceUser = append(getSourceUser, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_User{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.User = getSourceUser

// 		getSourceUsersGroup := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_UsersGroup{}
// 		for _, val := range ifRule.Rule.Source.GetUsersGroup() {
// 			getSourceUsersGroup = append(getSourceUsersGroup, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_UsersGroup{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.UsersGroup = getSourceUsersGroup

// 		getSourceGroup := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Group{}
// 		for _, val := range ifRule.Rule.Source.GetGroup() {
// 			getSourceGroup = append(getSourceGroup, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Group{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.Group = getSourceGroup

// 		getSourceSystemGroup := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_SystemGroup{}
// 		for _, val := range ifRule.Rule.Source.GetSystemGroup() {
// 			getSourceSystemGroup = append(getSourceSystemGroup, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_SystemGroup{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Source.SystemGroup = getSourceSystemGroup

// 		ipRuleItemEntry.Properties = makeStringListFromStringSlice(ifRule.GetProperties())

// 		ipRuleItemEntryAudit.UpdatedBy = types.StringValue(ifRule.Audit.GetUpdatedBy())
// 		ipRuleItemEntryAudit.UpdatedTime = types.StringValue(ifRule.Audit.GetUpdatedTime())

// 		ipRuleItemEntryRule.Action = types.StringValue(ifRule.Rule.Action.String())
// 		ipRuleItemEntryRule.ConnectionOrigin = types.StringValue(ifRule.Rule.ConnectionOrigin.String())

// 		getRuleCountryList := ifRule.Rule.GetCountry()
// 		getRuleCountryType := []Policy_Policy_InternetFirewall_Policy_Rules_Rule_Country{}

// 		for _, val := range getRuleCountryList {
// 			getRuleCountryType = append(getRuleCountryType, Policy_Policy_InternetFirewall_Policy_Rules_Rule_Country{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}

// 		ipRuleItemEntryRule.Country = getRuleCountryType

// 		// ... destination needed

// 		deviceListEnum := []Policy_Policy_InternetFirewall_Policy_Rules_Rule_Device{}
// 		for _, val := range ifRule.Rule.GetDevice() {
// 			deviceListEnum = append(deviceListEnum, Policy_Policy_InternetFirewall_Policy_Rules_Rule_Device{
// 				ID:   types.StringValue(val.GetID()),
// 				Name: types.StringValue(val.GetName()),
// 			})
// 		}
// 		ipRuleItemEntryRule.Device = deviceListEnum

// 		ipRuleItemEntryRule.DeviceOs = makeStringListFromStringSlice(ifRule.Rule.GetDeviceOs())

// 		ipRuleItemEntry.Audit = ipRuleItemEntryAudit
// 		ipRuleItemEntry.Rule = ipRuleItemEntryRule
// 		//ipRuleItemEntry.Rule.Source = sourceSourceEntry
// 		state.Rules = append(state.Rules, ipRuleItemEntry)
// 	}

// 	// Set state
// 	diags := resp.State.Set(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}
// }
