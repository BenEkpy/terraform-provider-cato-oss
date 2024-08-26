package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cato_models "github.com/routebyintuition/cato-go-sdk/models"
)

var (
	_ resource.Resource              = &internetFwPolicyResource{}
	_ resource.ResourceWithConfigure = &internetFwPolicyResource{}
)

func NewInternetFwPolicyResource() resource.Resource {
	return &internetFwPolicyResource{}
}

type internetFwPolicyResource struct {
	info *catoClientData
}

type InternetFirewallRuleModel struct {
	Name  types.String `tfsdk:"name"`
	Id    types.String `tfsdk:"id"`
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type InternetFirewall_Policy_Sections struct {
	Audit   Policy_Policy_InternetFirewall_Policy_Sections_Audit   "tfsdk:\"audit\" graphql:\"audit\""
	Section Policy_Policy_InternetFirewall_Policy_Sections_Section "tfsdk:\"section\" graphql:\"section\""
	// need to update properties for enum PolicyElementPropertiesEnum
	Properties []types.String "tfsdk:\"properties\" graphql:\"properties\""
}

type Policy_Policy_InternetFirewall_Policy_Sections_Audit struct {
	UpdatedTime types.String "tfsdk:\"updatedtime\" graphql:\"updatedTime\""
	UpdatedBy   types.String "tfsdk:\"updatedby\" graphql:\"updatedBy\""
}

type Policy_Policy_InternetFirewall_Policy_Sections_Section struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type InternetFirewall_Policy_Audit struct {
	PublishedTime types.String "tfsdk:\"publishedtime\" graphql:\"publishedTime\""
	PublishedBy   types.String "tfsdk:\"publishedby\" graphql:\"publishedBy\""
}

type InternetFirewall_Policy_Revision struct {
	ID          types.String "tfsdk:\"id\" graphql:\"id\""
	Name        types.String "tfsdk:\"name\" graphql:\"name\""
	Description types.String "tfsdk:\"description\" graphql:\"description\""
	Changes     types.Int64  "tfsdk:\"changes\" graphql:\"changes\""
	CreatedTime types.String "tfsdk:\"createdtime\" graphql:\"createdTime\""
	UpdatedTime types.String "tfsdk:\"updatedtime\" graphql:\"updatedTime\""
}

type InternetFirewallAddRuleInput struct {
	ID   types.String                  "tfsdk:\"id\" graphql:\"id\""
	Rule InternetFirewall_Policy_Rules `tfsdk:"rule"`
	At   *PolicyRulePositionInput      `tfsdk:"at"`
}

type InternetFirewallCreateRuleInput struct {
	Rule       Policy_Policy_InternetFirewall_Policy_Rules_Rule `tfsdk:"rule"`
	At         *PolicyRulePositionInput                         `tfsdk:"at"`
	Publish    types.Bool                                       `tfsdk:"publish"`
	SdkKeyName types.String                                     `tfsdk:"sdk_key_name"`
	// ID         types.String                                     `tfsdk:"id"`
}

type PolicyRulePositionInput struct {
	// this needs to be an emum PolicyRulePositionEnum
	Position types.String `tfsdk:"position"`
	Ref      types.String `tfsdk:"ref"`
}

type InternetFirewall_Policy_Rules struct {
	Audit Policy_Policy_InternetFirewall_Policy_Rules_Audit "tfsdk:\"audit\" graphql:\"audit\""
	Rule  Policy_Policy_InternetFirewall_Policy_Rules_Rule  "tfsdk:\"rule\" graphql:\"rule\""
	// need to switch properties to enum slice
	Properties []types.String "tfsdk:\"properties\" graphql:\"properties\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Audit struct {
	UpdatedTime types.String "tfsdk:\"updatedtime\" graphql:\"updatedTime\""
	UpdatedBy   types.String "tfsdk:\"updatedby\" graphql:\"updatedBy\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule struct {
	ID          types.String                                             "tfsdk:\"id\" graphql:\"id\""
	Name        types.String                                             "tfsdk:\"name\" graphql:\"name\""
	Description types.String                                             "tfsdk:\"description\" graphql:\"description\""
	Index       types.Int64                                              "tfsdk:\"index\" graphql:\"index\""
	Section     Policy_Policy_InternetFirewall_Policy_Rules_Rule_Section "tfsdk:\"section\" graphql:\"section\""
	Enabled     types.Bool                                               "tfsdk:\"enabled\" graphql:\"enabled\""
	Source      Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source  "tfsdk:\"source\" graphql:\"source\""
	// need to switch to enums
	ConnectionOrigin types.String                                               "tfsdk:\"connectionorigin\" graphql:\"connectionOrigin\""
	Country          []Policy_Policy_InternetFirewall_Policy_Rules_Rule_Country "tfsdk:\"country\" graphql:\"country\""
	Device           []Policy_Policy_InternetFirewall_Policy_Rules_Rule_Device  "tfsdk:\"device\" graphql:\"device\""
	// needs to be enum OperatingSystem
	DeviceOs    []types.String                                               "tfsdk:\"deviceos\" graphql:\"deviceOS\""
	Destination Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination "tfsdk:\"destination\" graphql:\"destination\""
	Service     Policy_Policy_InternetFirewall_Policy_Rules_Rule_Service     "tfsdk:\"service\" graphql:\"service\""
	// needs to be enum InternetFirewallActionEnum
	Action     types.String                                                   "tfsdk:\"action\" graphql:\"action\""
	Tracking   Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking      "tfsdk:\"tracking\" graphql:\"tracking\""
	Schedule   Policy_Policy_InternetFirewall_Policy_Rules_Rule_Schedule      "tfsdk:\"schedule\" graphql:\"schedule\""
	Exceptions []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions "tfsdk:\"exceptions\" graphql:\"exceptions\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Country struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Device struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination struct {
	Application            []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Application            "tfsdk:\"application\" graphql:\"application\""
	CustomApp              []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomApp              "tfsdk:\"customapp\" graphql:\"customApp\""
	AppCategory            []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_AppCategory            "tfsdk:\"appcategory\" graphql:\"appCategory\""
	CustomCategory         []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomCategory         "tfsdk:\"customcategory\" graphql:\"customCategory\""
	SanctionedAppsCategory []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_SanctionedAppsCategory "tfsdk:\"sanctionedappscategory\" graphql:\"sanctionedAppsCategory\""
	Country                []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Country                "tfsdk:\"country\" graphql:\"country\""
	Domain                 []types.String                                                                         "tfsdk:\"domain\" graphql:\"domain\""
	Fqdn                   []types.String                                                                         "tfsdk:\"fqdn\" graphql:\"fqdn\""
	IP                     []types.String                                                                         "tfsdk:\"ip\" graphql:\"ip\""
	Subnet                 []types.String                                                                         "tfsdk:\"subnet\" graphql:\"subnet\""
	IPRange                []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_IPRange                "tfsdk:\"iprange\" graphql:\"ipRange\""
	GlobalIPRange          []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_GlobalIPRange          "tfsdk:\"globaliprange\" graphql:\"globalIpRange\""
	RemoteAsn              []types.String                                                                         "tfsdk:\"remoteasn\" graphql:\"remoteAsn\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Application struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomApp struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_AppCategory struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomCategory struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_SanctionedAppsCategory struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Country struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_IPRange struct {
	From  types.String "tfsdk:\"from\" graphql:\"from\""
	To    types.String "tfsdk:\"to\" graphql:\"to\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_GlobalIPRange struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source struct {
	IP                []types.String                                                               "tfsdk:\"ip\" graphql:\"ip\""
	Host              []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Host              "tfsdk:\"host\" graphql:\"host\""
	Site              []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Site              "tfsdk:\"site\" graphql:\"site\""
	Subnet            []types.String                                                               "tfsdk:\"subnet\" graphql:\"subnet\""
	IPRange           []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_IPRange           "tfsdk:\"iprange\" graphql:\"ipRange\""
	GlobalIPRange     []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_GlobalIPRange     "tfsdk:\"globaliprange\" graphql:\"globalIpRange\""
	NetworkInterface  []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_NetworkInterface  "tfsdk:\"networkinterface\" graphql:\"networkInterface\""
	SiteNetworkSubnet []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_SiteNetworkSubnet "tfsdk:\"sitenetworksubnet\" graphql:\"siteNetworkSubnet\""
	FloatingSubnet    []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_FloatingSubnet    "tfsdk:\"floatingsubnet\" graphql:\"floatingSubnet\""
	User              []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_User              "tfsdk:\"user\" graphql:\"user\""
	UsersGroup        []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_UsersGroup        "tfsdk:\"usersgroup\" graphql:\"usersGroup\""
	Group             []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Group             "tfsdk:\"group\" graphql:\"group\""
	SystemGroup       []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_SystemGroup       "tfsdk:\"systemgroup\" graphql:\"systemGroup\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Host struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Site struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_IPRange struct {
	From  types.String "tfsdk:\"from\" graphql:\"from\""
	To    types.String "tfsdk:\"to\" graphql:\"to\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_GlobalIPRange struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_NetworkInterface struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_SiteNetworkSubnet struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_User struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_UsersGroup struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_Group struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_SystemGroup struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Source_FloatingSubnet struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Service struct {
	Standard []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Service_Standard "tfsdk:\"standard\" graphql:\"standard\""
	Custom   []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Service_Custom   "tfsdk:\"custom\" graphql:\"custom\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Service_Standard struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Service_Custom struct {
	Port      []types.String                                                             "tfsdk:\"port\" graphql:\"port\""
	PortRange *Policy_Policy_InternetFirewall_Policy_Rules_Rule_Service_Custom_PortRange "tfsdk:\"portrange\" graphql:\"portRange\""
	Protocol  types.String                                                               "tfsdk:\"protocol\" graphql:\"protocol\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Service_Custom_PortRange struct {
	From types.String "tfsdk:\"from\" graphql:\"from\""
	To   types.String "tfsdk:\"to\" graphql:\"to\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking struct {
	Event Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Event "tfsdk:\"event\" graphql:\"event\""
	Alert Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Alert "tfsdk:\"alert\" graphql:\"alert\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Event struct {
	Enabled types.Bool "tfsdk:\"enabled\" graphql:\"enabled\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Alert struct {
	Enabled types.Bool "tfsdk:\"enabled\" graphql:\"enabled\""
	// needs to be an enum PolicyRuleTrackingFrequencyEnum
	Frequency         types.String                                                                         "tfsdk:\"frequency\" graphql:\"frequency\""
	SubscriptionGroup []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Alert_SubscriptionGroup "tfsdk:\"subscriptiongroup\" graphql:\"subscriptionGroup\""
	Webhook           []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Alert_Webhook           "tfsdk:\"webhook\" graphql:\"webhook\""
	MailingList       []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Alert_MailingList       "tfsdk:\"mailinglist\" graphql:\"mailingList\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Alert_SubscriptionGroup struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Alert_Webhook struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Tracking_Alert_MailingList struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Schedule struct {
	// needs to be enum PolicyActiveOnEnum
	ActiveOn        types.String                                                               "tfsdk:\"activeon\" graphql:\"activeOn\""
	CustomTimeframe *Policy_Policy_InternetFirewall_Policy_Rules_Rule_Schedule_CustomTimeframe "tfsdk:\"customtimeframe\" graphql:\"customTimeframe\""
	CustomRecurring *Policy_Policy_InternetFirewall_Policy_Rules_Rule_Schedule_CustomRecurring "tfsdk:\"customrecurring\" graphql:\"customRecurring\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Schedule_CustomTimeframe struct {
	From types.String "tfsdk:\"from\" graphql:\"from\""
	To   types.String "tfsdk:\"to\" graphql:\"to\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Schedule_CustomRecurring struct {
	From types.String "tfsdk:\"from\" graphql:\"from\""
	To   types.String "tfsdk:\"to\" graphql:\"to\""
	// needs to be an enum cato_query_models.DayOfWeek
	Days []DayOfWeek "tfsdk:\"days\" graphql:\"days\""
}

type DayOfWeek types.String

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions struct {
	Name   types.String                                                       "tfsdk:\"name\" graphql:\"name\""
	Source Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Source "tfsdk:\"source\" graphql:\"source\""
	// neeeds to be enum
	DeviceOs    []OperatingSystem                                                       "tfsdk:\"deviceos\" graphql:\"deviceOS\""
	Country     []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Country  "tfsdk:\"country\" graphql:\"country\""
	Device      []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Device   "tfsdk:\"device\" graphql:\"device\""
	Destination Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Destination "tfsdk:\"destination\" graphql:\"destination\""
	// needs to be enum ConnectionOriginEnum
	ConnectionOrigin types.String "tfsdk:\"connectionorigin\" graphql:\"connectionOrigin\""
}

type OperatingSystem types.String

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Country struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Source struct {
	IP     []types.String "tfsdk:\"ip\" graphql:\"ip\""
	Subnet []types.String "tfsdk:\"subnet\" graphql:\"subnet\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Device struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Destination struct {
	Domain    []types.String "tfsdk:\"domain\" graphql:\"domain\""
	Fqdn      []types.String "tfsdk:\"fqdn\" graphql:\"fqdn\""
	IP        []types.String "tfsdk:\"ip\" graphql:\"ip\""
	Subnet    []types.String "tfsdk:\"subnet\" graphql:\"subnet\""
	RemoteAsn []types.String "tfsdk:\"remoteasn\" graphql:\"remoteAsn\""
}

type Policy_Policy_InternetFirewall_Policy_Rules_Rule_Section struct {
	ID    types.String "tfsdk:\"id\" graphql:\"id\""
	Name  types.String "tfsdk:\"name\" graphql:\"name\""
	By    types.String "tfsdk:\"by\" graphql:\"by\""
	Input types.String "tfsdk:\"input\" graphql:\"input\""
}

func (r *internetFwPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_if_policy"
}

func (r *internetFwPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"at": schema.SingleNestedAttribute{
				Description: "at",
				Required:    false,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"position": schema.StringAttribute{
						Description: "position",
						Required:    false,
						Optional:    true,
					},
					"ref": schema.StringAttribute{
						Description: "ref",
						Required:    false,
						Optional:    true,
					},
				},
			},
			"publish": schema.BoolAttribute{
				Description: "publish",
				Required:    true,
				Optional:    false,
			},
			"sdk_key_name": schema.StringAttribute{
				Description: "sdk key name",
				Required:    false,
				Optional:    true,
			},
			"rule": schema.SingleNestedAttribute{
				Description: "rule item",
				Required:    true,

				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: false,
					},
					"name": schema.StringAttribute{
						Description: "Rule name",
						Required:    true,
					},
					"description": schema.StringAttribute{
						Description: "Rule description",
						Required:    true,
					},
					"index": schema.Int64Attribute{
						Description: "Rule index",
						Required:    false,
						Optional:    true,
					},
					"section": schema.SingleNestedAttribute{
						Required: false,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "id",
								Required:    false,
								Optional:    true,
							},
							"name": schema.StringAttribute{
								Description: "name",
								Required:    false,
								Optional:    true,
							},
							"by": schema.StringAttribute{
								Description: "by",
								Required:    false,
								Optional:    true,
							},
							"input": schema.StringAttribute{
								Description: "input",
								Required:    false,
								Optional:    true,
							},
						},
					},
					"enabled": schema.BoolAttribute{
						Description: "enabled",
						Required:    true,
					},
					"source": schema.SingleNestedAttribute{
						Required: true,
						Optional: false,
						Attributes: map[string]schema.Attribute{
							"ip": schema.ListAttribute{
								Description: "ip",
								ElementType: types.StringType,
								Required:    true,
							},
							"host": schema.ListNestedAttribute{
								Required: false,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "id",
											Required:    false,
											Optional:    true,
										},
										"name": schema.StringAttribute{
											Description: "name",
											Required:    false,
											Optional:    true,
										},
									},
								},
							},
							"site": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "site",
								Required:    false,
								Optional:    true,
							},
							"subnet": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "subnet",
								Required:    false,
								Optional:    true,
							},
							"iprange": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "ipRange",
								Required:    false,
								Optional:    true,
							},
							"globaliprange": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "globalIpRange",
								Required:    false,
								Optional:    true,
							},
							"networkinterface": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "networkInterface",
								Required:    false,
								Optional:    true,
							},
							"sitenetworksubnet": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "siteNetworkSubnet",
								Required:    false,
								Optional:    true,
							},
							"floatingsubnet": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "floatingSubnet",
								Required:    false,
								Optional:    true,
							},
							"user": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "user",
								Required:    false,
								Optional:    true,
							},
							"usersgroup": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "usersGroup",
								Required:    false,
								Optional:    true,
							},
							"group": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "group",
								Required:    false,
								Optional:    true,
							},
							"systemgroup": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "systemGroup",
								Required:    false,
								Optional:    true,
							},
						},
					},
					"connectionorigin": schema.StringAttribute{
						Description: "connectionOrigin",
						Required:    true,
					},
					"country": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "country",
						Required:    true,
					},
					"device": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "device",
						Required:    true,
					},
					"deviceos": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "deviceOS",
						Required:    true,
					},
					"destination": schema.SingleNestedAttribute{
						Optional: false,
						Required: true,

						Attributes: map[string]schema.Attribute{
							"application": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "application",
								Required:    true,
							},
							"customapp": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "customApp",
								Required:    true,
							},
							"appcategory": schema.ListNestedAttribute{
								Required: false,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"by": schema.StringAttribute{
											Description: "by",
											Required:    false,
											Optional:    true,
										},
										"input": schema.StringAttribute{
											Description: "input",
											Required:    false,
											Optional:    true,
										},
										"id": schema.StringAttribute{
											Description: "by",
											Required:    false,
											Optional:    true,
										},
										"name": schema.StringAttribute{
											Description: "input",
											Required:    false,
											Optional:    true,
										},
									},
								},
							},
							"customcategory": schema.ListNestedAttribute{
								Description: "customCategory",
								Required:    false,
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"by": schema.StringAttribute{
											Description: "by",
											Required:    false,
											Optional:    true,
										},
										"input": schema.StringAttribute{
											Description: "input",
											Required:    false,
											Optional:    true,
										},
										"id": schema.StringAttribute{
											Description: "by",
											Required:    false,
											Optional:    true,
										},
										"name": schema.StringAttribute{
											Description: "input",
											Required:    false,
											Optional:    true,
										},
									},
								},
							},
							"sanctionedappscategory": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "sanctionedAppsCategory",
								Required:    true,
							},
							"country": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "country",
								Required:    true,
							},
							"domain": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "domain",
								Required:    true,
							},
							"fqdn": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "fqdn",
								Required:    true,
							},
							"ip": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "ip",
								Required:    true,
							},
							"subnet": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "subnet",
								Required:    true,
							},
							"iprange": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "ipRange",
								Required:    true,
							},
							"globaliprange": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "globalIpRange",
								Required:    true,
							},
							"remoteasn": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "remoteAsn",
								Required:    true,
							},
						},
					},
					"service": schema.SingleNestedAttribute{
						Optional: true,
						Required: false,

						Attributes: map[string]schema.Attribute{
							"standard": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "standard",
								Required:    false,
								Optional:    true,
							},
							"custom": schema.ListAttribute{
								ElementType: types.StringType,
								Description: "custom",
								Required:    false,
								Optional:    true,
							},
						},
					},
					"action": schema.StringAttribute{
						Description: "action",
						Required:    true,
					},
					"tracking": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"event": schema.SingleNestedAttribute{
								Description: "event",
								Required:    true,

								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "enabled",
										Required:    true,
									},
								},
							},
							"alert": schema.SingleNestedAttribute{
								Description: "alert",
								Required:    true,

								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "enabled",
										Required:    true,
									},
									"frequency": schema.StringAttribute{
										Description: "frequency",
										Required:    true,
									},
									"subscriptiongroup": schema.ListAttribute{
										ElementType: types.StringType,
										Description: "subscriptionGroup",
										Required:    false,
										Optional:    true,
									},
									"webhook": schema.ListAttribute{
										ElementType: types.StringType,
										Description: "webhook",
										Required:    false,
										Optional:    true,
									},
									"mailinglist": schema.ListAttribute{
										ElementType: types.StringType,
										Description: "mailingList",
										Required:    false,
										Optional:    true,
									},
								},
							},
						},
					},
					"schedule": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"activeon": schema.StringAttribute{
								Description: "activeOn",
								Required:    false,
								Optional:    true,
							},
							"customtimeframe": schema.StringAttribute{
								Description: "customtimeframe",
								Required:    false,
								Optional:    true,
							},
							"customrecurring": schema.StringAttribute{
								Description: "customrecurring",
								Required:    false,
								Optional:    true,
							},
						},
					},
					"exceptions": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "exceptions",
						Required:    false,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *internetFwPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.info = req.ProviderData.(*catoClientData)

}

func makeStringSliceFromStringList(src []types.String) []string {
	res := []string{}

	for _, element := range src {
		res = append(res, element.ValueString())
	}

	return res
}

func makeStringListFromStringSlice(src any) []types.String {
	res := []types.String{}

	srcToStringSlice, ok := src.([]string)
	if !ok {
		return res
	}

	for _, element := range srcToStringSlice {
		res = append(res, types.StringValue(element))
	}

	return res
}

func (r *internetFwPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan InternetFirewallCreateRuleInput

	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	position := cato_models.PolicyRulePositionEnum(plan.At.Position.ValueString())

	hostSourceRefInput := []*cato_models.HostRefInput{}
	for _, val := range plan.Rule.Source.Host {
		hostSourceRefInput = append(hostSourceRefInput, &cato_models.HostRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	siteSourceRefInput := []*cato_models.SiteRefInput{}
	for _, val := range plan.Rule.Source.Site {
		siteSourceRefInput = append(siteSourceRefInput, &cato_models.SiteRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	ipSourcerange := []*cato_models.IPAddressRangeInput{}
	for _, val := range plan.Rule.Source.IPRange {
		ipSourcerange = append(ipSourcerange, &cato_models.IPAddressRangeInput{
			From: val.From.ValueString(),
			To:   val.To.ValueString(),
		})
	}

	globalSourceIpRange := []*cato_models.GlobalIPRangeRefInput{}
	for _, val := range plan.Rule.Source.GlobalIPRange {
		globalSourceIpRange = append(globalSourceIpRange, &cato_models.GlobalIPRangeRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	networkSourceInterfaceRefInput := []*cato_models.NetworkInterfaceRefInput{}
	for _, val := range plan.Rule.Source.NetworkInterface {
		networkSourceInterfaceRefInput = append(networkSourceInterfaceRefInput, &cato_models.NetworkInterfaceRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	siteSourceNetworkSubnetRefInput := []*cato_models.SiteNetworkSubnetRefInput{}
	for _, val := range plan.Rule.Source.SiteNetworkSubnet {
		siteSourceNetworkSubnetRefInput = append(siteSourceNetworkSubnetRefInput, &cato_models.SiteNetworkSubnetRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	floatingSourceSubnetRefInput := []*cato_models.FloatingSubnetRefInput{}
	for _, val := range plan.Rule.Source.FloatingSubnet {
		floatingSourceSubnetRefInput = append(floatingSourceSubnetRefInput, &cato_models.FloatingSubnetRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	userSourceRefInput := []*cato_models.UserRefInput{}
	for _, val := range plan.Rule.Source.User {
		userSourceRefInput = append(userSourceRefInput, &cato_models.UserRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	usersSourceGroupRefInput := []*cato_models.UsersGroupRefInput{}
	for _, val := range plan.Rule.Source.UsersGroup {
		usersSourceGroupRefInput = append(usersSourceGroupRefInput, &cato_models.UsersGroupRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	groupSourceRefInput := []*cato_models.GroupRefInput{}
	for _, val := range plan.Rule.Source.Group {
		groupSourceRefInput = append(groupSourceRefInput, &cato_models.GroupRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	systemSouceGroupRefInput := []*cato_models.SystemGroupRefInput{}
	for _, val := range plan.Rule.Source.SystemGroup {
		systemSouceGroupRefInput = append(systemSouceGroupRefInput, &cato_models.SystemGroupRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	countryInput := []*cato_models.CountryRefInput{}
	for _, val := range plan.Rule.Country {
		countryInput = append(countryInput, &cato_models.CountryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	countryDestInput := []*cato_models.CountryRefInput{}
	for _, val := range plan.Rule.Destination.Country {
		countryInput = append(countryInput, &cato_models.CountryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	connectionOrigin := cato_models.ConnectionOriginEnum(plan.Rule.ConnectionOrigin.ValueString())
	actionEnum := cato_models.InternetFirewallActionEnum(plan.Rule.Action.ValueString())

	domainDestList := makeStringSliceFromStringList(plan.Rule.Destination.Domain)
	fqdnDestList := makeStringSliceFromStringList(plan.Rule.Destination.Fqdn)
	subnetSourceRefInput := makeStringSliceFromStringList(plan.Rule.Source.Subnet)

	deviceInput := []*cato_models.DeviceProfileRefInput{}
	for _, val := range plan.Rule.Device {
		deviceInput = append(deviceInput, &cato_models.DeviceProfileRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	deviceOsInput := []cato_models.OperatingSystem{}
	for _, val := range plan.Rule.DeviceOs {
		deviceOsInput = append(deviceOsInput, cato_models.OperatingSystem(val.ValueString()))
	}

	ipDestRange := []*cato_models.IPAddressRangeInput{}
	for _, val := range plan.Rule.Destination.IPRange {
		ipDestRange = append(ipDestRange, &cato_models.IPAddressRangeInput{
			From: val.From.ValueString(),
			To:   val.To.ValueString(),
		})
	}

	globalDestIpRange := []*cato_models.GlobalIPRangeRefInput{}
	for _, val := range plan.Rule.Destination.GlobalIPRange {
		globalDestIpRange = append(globalDestIpRange, &cato_models.GlobalIPRangeRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	applicationDestInput := []*cato_models.ApplicationRefInput{}
	for _, val := range plan.Rule.Destination.Application {
		applicationDestInput = append(applicationDestInput, &cato_models.ApplicationRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	customAppDestInput := []*cato_models.CustomApplicationRefInput{}
	for _, val := range plan.Rule.Destination.CustomApp {
		customAppDestInput = append(customAppDestInput, &cato_models.CustomApplicationRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	appDestCategoryInput := []*cato_models.ApplicationCategoryRefInput{}
	for _, val := range plan.Rule.Destination.AppCategory {
		appDestCategoryInput = append(appDestCategoryInput, &cato_models.ApplicationCategoryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	customDestCategory := []*cato_models.CustomCategoryRefInput{}
	for _, val := range plan.Rule.Destination.CustomCategory {
		customDestCategory = append(customDestCategory, &cato_models.CustomCategoryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	sanctionedDestAppsCategory := []*cato_models.SanctionedAppsCategoryRefInput{}
	for _, val := range plan.Rule.Destination.SanctionedAppsCategory {
		sanctionedDestAppsCategory = append(sanctionedDestAppsCategory, &cato_models.SanctionedAppsCategoryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	serviceInput := &cato_models.InternetFirewallServiceTypeInput{}

	ruleTrackingAlertSubscriptionGroup := []*cato_models.SubscriptionGroupRefInput{}
	for _, val := range plan.Rule.Tracking.Alert.SubscriptionGroup {
		ruleTrackingAlertSubscriptionGroup = append(ruleTrackingAlertSubscriptionGroup, &cato_models.SubscriptionGroupRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	ruleTrackingAlertMailingList := []*cato_models.SubscriptionMailingListRefInput{}
	for _, val := range plan.Rule.Tracking.Alert.MailingList {
		ruleTrackingAlertMailingList = append(ruleTrackingAlertMailingList, &cato_models.SubscriptionMailingListRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	ruleTrackingAlertSubscriptionWebhook := []*cato_models.SubscriptionWebhookRefInput{}
	for _, val := range plan.Rule.Tracking.Alert.Webhook {
		ruleTrackingAlertSubscriptionWebhook = append(ruleTrackingAlertSubscriptionWebhook, &cato_models.SubscriptionWebhookRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	input := cato_models.InternetFirewallAddRuleInput{
		At: &cato_models.PolicyRulePositionInput{
			Position: &position,
		},
		Rule: &cato_models.InternetFirewallAddRuleDataInput{
			Enabled:     plan.Rule.Enabled.ValueBool(),
			Name:        plan.Rule.Name.ValueString(),
			Description: plan.Rule.Description.ValueString(),
			Source: &cato_models.InternetFirewallSourceInput{
				IP:                makeStringSliceFromStringList(plan.Rule.Source.IP),
				Host:              hostSourceRefInput,
				Site:              siteSourceRefInput,
				Subnet:            subnetSourceRefInput,
				IPRange:           ipSourcerange,
				GlobalIPRange:     globalSourceIpRange,
				NetworkInterface:  networkSourceInterfaceRefInput,
				SiteNetworkSubnet: siteSourceNetworkSubnetRefInput,
				FloatingSubnet:    floatingSourceSubnetRefInput,
				User:              userSourceRefInput,
				UsersGroup:        usersSourceGroupRefInput,
				Group:             groupSourceRefInput,
				SystemGroup:       systemSouceGroupRefInput,
			},
			ConnectionOrigin: connectionOrigin,
			Country:          countryInput,
			Device:           deviceInput,
			DeviceOs:         deviceOsInput,
			Destination: &cato_models.InternetFirewallDestinationInput{
				Application:            applicationDestInput,
				CustomApp:              customAppDestInput,
				AppCategory:            appDestCategoryInput,
				CustomCategory:         customDestCategory,
				SanctionedAppsCategory: sanctionedDestAppsCategory,
				Country:                countryDestInput,
				Domain:                 domainDestList,
				Fqdn:                   fqdnDestList,
				IP:                     makeStringSliceFromStringList(plan.Rule.Destination.IP),
				Subnet:                 makeStringSliceFromStringList(plan.Rule.Destination.Subnet),
				IPRange:                ipDestRange,
				GlobalIPRange:          globalDestIpRange,
				RemoteAsn:              makeStringSliceFromStringList(plan.Rule.Destination.RemoteAsn),
			},
			Service: serviceInput,
			Action:  actionEnum,
			Schedule: &cato_models.PolicyScheduleInput{
				ActiveOn: cato_models.PolicyActiveOnEnum(plan.Rule.Schedule.ActiveOn.ValueString()),
			},
			Tracking: &cato_models.PolicyTrackingInput{
				Event: &cato_models.PolicyRuleTrackingEventInput{
					Enabled: plan.Rule.Tracking.Event.Enabled.ValueBool(),
				},
				Alert: &cato_models.PolicyRuleTrackingAlertInput{
					Enabled:           plan.Rule.Tracking.Alert.Enabled.ValueBool(),
					Frequency:         cato_models.PolicyRuleTrackingFrequencyEnum(plan.Rule.Tracking.Alert.Frequency.ValueString()),
					SubscriptionGroup: ruleTrackingAlertSubscriptionGroup,
					MailingList:       ruleTrackingAlertMailingList,
					Webhook:           ruleTrackingAlertSubscriptionWebhook,
				},
			},
		},
	}

	b, _ := json.Marshal(input)

	tflog.Info(ctx, "initial create input")
	tflog.Info(ctx, string(b))

	policyChange, err := r.info.catov2.PolicyInternetFirewallAddRule(ctx, input, r.info.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API PolicyInternetFirewallAddRule error",
			err.Error(),
		)
		return
	}

	tflog.Warn(ctx, "after PolicyInternetFirewallAddRule")

	plan.Rule.ID = types.StringValue(policyChange.GetPolicy().GetInternetFirewall().GetAddRule().Rule.GetRule().ID)

	tflog.Warn(ctx, fmt.Sprintf("Set Value plan.ID: %s", plan.Rule.ID))

	if plan.Publish.ValueBool() {
		tflog.Info(ctx, "publishing new rule")
		publishDataIfEnabled := &cato_models.PolicyPublishRevisionInput{
			Name: plan.SdkKeyName.ValueStringPointer(),
		}
		_, err := r.info.catov2.PolicyInternetFirewallPublishPolicyRevision(ctx, &cato_models.InternetFirewallPolicyMutationInput{}, publishDataIfEnabled, r.info.AccountId)
		if err != nil {
			resp.Diagnostics.AddError(
				"Catov2 API PolicyInternetFirewallPublishPolicyRevision error",
				err.Error(),
			)
			return
		}
	} else {
		tflog.Info(ctx, "NOT publishing new rule")
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *internetFwPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state InternetFirewallCreateRuleInput

	tflog.Info(ctx, "initial Read_CALL")

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	queryPolicy := &cato_models.InternetFirewallPolicyInput{}
	body, err := r.info.catov2.Policy(ctx, queryPolicy, r.info.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API error",
			err.Error(),
		)
		return
	}

	ruleList := body.GetPolicy().InternetFirewall.Policy.GetRules()
	for _, ruleListItem := range ruleList {
		if ruleListItem.GetRule().ID == state.Rule.ID.ValueString() {

			state.Rule.ID = types.StringValue(ruleListItem.GetRule().ID)
			state.Rule.Action = types.StringValue(ruleListItem.Rule.Action.String())
			state.Rule.ConnectionOrigin = types.StringValue(ruleListItem.Rule.ConnectionOrigin.String())
			state.Rule.Description = types.StringValue(ruleListItem.Rule.GetDescription())
			state.Rule.Name = types.StringValue(ruleListItem.GetRule().Name)
			countryRuleItem := []Policy_Policy_InternetFirewall_Policy_Rules_Rule_Country{}
			for _, val := range ruleListItem.Rule.GetCountry() {
				countryRuleItem = append(countryRuleItem, Policy_Policy_InternetFirewall_Policy_Rules_Rule_Country{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			state.Rule.Country = countryRuleItem

			state.Rule.Destination = Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination{}

			dstAppCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_AppCategory{}
			for _, val := range ruleListItem.Rule.Destination.AppCategory {
				dstAppCategory = append(dstAppCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_AppCategory{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.GetName()),
				})
			}
			state.Rule.Destination.AppCategory = dstAppCategory

			dstApplication := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Application{}
			for _, val := range ruleListItem.Rule.Destination.Application {
				dstApplication = append(dstApplication, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Application{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			state.Rule.Destination.Application = dstApplication

			dstCountry := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Country{}
			for _, val := range ruleListItem.Rule.Destination.Country {
				dstCountry = append(dstCountry, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Country{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			state.Rule.Destination.Country = dstCountry

			dstCustomApp := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomApp{}
			for _, val := range ruleListItem.Rule.Destination.CustomApp {
				dstCustomApp = append(dstCustomApp, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomApp{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			state.Rule.Destination.CustomApp = dstCustomApp

			dstCustomCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomCategory{}
			for _, val := range ruleListItem.Rule.Destination.CustomCategory {
				dstCustomCategory = append(dstCustomCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomCategory{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			state.Rule.Destination.CustomCategory = dstCustomCategory

			dstDomain := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.Domain {
				dstDomain = append(dstDomain, types.StringValue(val))
			}
			state.Rule.Destination.Domain = dstDomain

			dstFqdn := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.Fqdn {
				dstDomain = append(dstDomain, types.StringValue(val))
			}
			state.Rule.Destination.Fqdn = dstFqdn

			dstGlobalRange := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_GlobalIPRange{}
			for _, val := range ruleListItem.Rule.Destination.GlobalIPRange {
				dstGlobalRange = append(dstGlobalRange, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_GlobalIPRange{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			state.Rule.Destination.GlobalIPRange = dstGlobalRange

			dstIp := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.IP {
				dstIp = append(dstIp, types.StringValue(val))
			}
			state.Rule.Destination.IP = dstIp

			dstIPRange := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_IPRange{}
			for _, val := range ruleListItem.Rule.Destination.IPRange {
				dstIPRange = append(dstIPRange, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_IPRange{
					From: types.StringValue(val.From),
					To:   types.StringValue(val.To),
				})
			}
			state.Rule.Destination.IPRange = dstIPRange

			dstRemoteAsn := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.RemoteAsn {
				dstRemoteAsn = append(dstRemoteAsn, types.StringValue(val))
			}
			state.Rule.Destination.RemoteAsn = dstRemoteAsn

			dstSanctionedAppsCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_SanctionedAppsCategory{}
			for _, val := range ruleListItem.Rule.Destination.SanctionedAppsCategory {
				dstSanctionedAppsCategory = append(dstSanctionedAppsCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_SanctionedAppsCategory{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			state.Rule.Destination.SanctionedAppsCategory = dstSanctionedAppsCategory

			dstSSubnet := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.Subnet {
				dstSSubnet = append(dstSSubnet, types.StringValue(val))
			}
			state.Rule.Destination.Subnet = dstSSubnet

			deviceList := []Policy_Policy_InternetFirewall_Policy_Rules_Rule_Device{}
			for _, val := range ruleListItem.Rule.Device {
				deviceList = append(deviceList, Policy_Policy_InternetFirewall_Policy_Rules_Rule_Device{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			state.Rule.Device = deviceList

			deviceOSList := []types.String{}
			for _, val := range ruleListItem.Rule.DeviceOs {
				deviceOSList = append(deviceOSList, types.StringValue(string(val)))
			}
			state.Rule.DeviceOs = deviceOSList

			exceptionsList := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions{}
			for _, val := range ruleListItem.Rule.Exceptions {
				exceptionsList = append(exceptionsList, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions{
					Name:             types.StringValue(val.Name),
					Source:           Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Source{},
					DeviceOs:         []OperatingSystem{},
					Country:          []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Country{},
					Device:           []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Device{},
					Destination:      Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Destination{},
					ConnectionOrigin: types.StringValue(""),
				})
			}
			state.Rule.Exceptions = exceptionsList

			// state.Rule.Schedule.ActiveOn = types.StringValue(ruleListItem.Schedule.ActiveOn)

		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *internetFwPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan InternetFirewallCreateRuleInput
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state InternetFirewallCreateRuleInput
	diagState := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diagState...)
	if resp.Diagnostics.HasError() {
		return
	}

	mutationInput := &cato_models.InternetFirewallPolicyMutationInput{}

	hostSourceRefInput := []*cato_models.HostRefInput{}
	for _, val := range plan.Rule.Source.Host {
		hostSourceRefInput = append(hostSourceRefInput, &cato_models.HostRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	siteSourceRefInput := []*cato_models.SiteRefInput{}
	for _, val := range plan.Rule.Source.Site {
		siteSourceRefInput = append(siteSourceRefInput, &cato_models.SiteRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Name.ValueString(),
		})
	}

	ipSourcerange := []*cato_models.IPAddressRangeInput{}
	for _, val := range plan.Rule.Source.IPRange {
		ipSourcerange = append(ipSourcerange, &cato_models.IPAddressRangeInput{
			From: val.From.ValueString(),
			To:   val.To.ValueString(),
		})
	}

	globalSourceIpRange := []*cato_models.GlobalIPRangeRefInput{}
	for _, val := range plan.Rule.Source.GlobalIPRange {
		globalSourceIpRange = append(globalSourceIpRange, &cato_models.GlobalIPRangeRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	networkSourceInterfaceRefInput := []*cato_models.NetworkInterfaceRefInput{}
	for _, val := range plan.Rule.Source.NetworkInterface {
		networkSourceInterfaceRefInput = append(networkSourceInterfaceRefInput, &cato_models.NetworkInterfaceRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	siteSourceNetworkSubnetRefInput := []*cato_models.SiteNetworkSubnetRefInput{}
	for _, val := range plan.Rule.Source.SiteNetworkSubnet {
		siteSourceNetworkSubnetRefInput = append(siteSourceNetworkSubnetRefInput, &cato_models.SiteNetworkSubnetRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	floatingSourceSubnetRefInput := []*cato_models.FloatingSubnetRefInput{}
	for _, val := range plan.Rule.Source.FloatingSubnet {
		floatingSourceSubnetRefInput = append(floatingSourceSubnetRefInput, &cato_models.FloatingSubnetRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	userSourceRefInput := []*cato_models.UserRefInput{}
	for _, val := range plan.Rule.Source.User {
		userSourceRefInput = append(userSourceRefInput, &cato_models.UserRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	usersSourceGroupRefInput := []*cato_models.UsersGroupRefInput{}
	for _, val := range plan.Rule.Source.UsersGroup {
		usersSourceGroupRefInput = append(usersSourceGroupRefInput, &cato_models.UsersGroupRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	groupSourceRefInput := []*cato_models.GroupRefInput{}
	for _, val := range plan.Rule.Source.Group {
		groupSourceRefInput = append(groupSourceRefInput, &cato_models.GroupRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	systemSouceGroupRefInput := []*cato_models.SystemGroupRefInput{}
	for _, val := range plan.Rule.Source.Group {
		systemSouceGroupRefInput = append(systemSouceGroupRefInput, &cato_models.SystemGroupRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	countryInput := []*cato_models.CountryRefInput{}
	for _, val := range plan.Rule.Country {
		countryInput = append(countryInput, &cato_models.CountryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	countryDestInput := []*cato_models.CountryRefInput{}
	for _, val := range plan.Rule.Destination.Country {
		countryInput = append(countryInput, &cato_models.CountryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	connectionOrigin := cato_models.ConnectionOriginEnum(plan.Rule.ConnectionOrigin.ValueString())
	actionEnum := cato_models.InternetFirewallActionEnum(plan.Rule.Action.ValueString())
	domainDestList := makeStringSliceFromStringList(plan.Rule.Destination.Domain)
	fqdnDestList := makeStringSliceFromStringList(plan.Rule.Destination.Fqdn)
	subnetSourceRefInput := makeStringSliceFromStringList(plan.Rule.Source.Subnet)

	deviceInput := []*cato_models.DeviceProfileRefInput{}
	for _, val := range plan.Rule.Device {
		deviceInput = append(deviceInput, &cato_models.DeviceProfileRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	deviceOsInput := []cato_models.OperatingSystem{}
	for _, val := range plan.Rule.Device {
		deviceOsInput = append(deviceOsInput, cato_models.OperatingSystem(val.ID.ValueString()))
	}

	ipDestRange := []*cato_models.IPAddressRangeInput{}
	for _, val := range plan.Rule.Destination.IPRange {
		ipDestRange = append(ipDestRange, &cato_models.IPAddressRangeInput{
			From: val.From.ValueString(),
			To:   val.To.ValueString(),
		})
	}

	globalDestIpRange := []*cato_models.GlobalIPRangeRefInput{}
	for _, val := range plan.Rule.Destination.GlobalIPRange {
		globalDestIpRange = append(globalDestIpRange, &cato_models.GlobalIPRangeRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	applicationDestInput := []*cato_models.ApplicationRefInput{}
	for _, val := range plan.Rule.Destination.Application {
		applicationDestInput = append(applicationDestInput, &cato_models.ApplicationRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	customAppDestInput := []*cato_models.CustomApplicationRefInput{}
	for _, val := range plan.Rule.Destination.Application {
		customAppDestInput = append(customAppDestInput, &cato_models.CustomApplicationRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	appDestCategoryInput := []*cato_models.ApplicationCategoryRefInput{}
	for _, val := range plan.Rule.Destination.AppCategory {
		appDestCategoryInput = append(appDestCategoryInput, &cato_models.ApplicationCategoryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	customDestCategory := []*cato_models.CustomCategoryRefInput{}
	for _, val := range plan.Rule.Destination.CustomCategory {
		customDestCategory = append(customDestCategory, &cato_models.CustomCategoryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	sanctionedDestAppsCategory := []*cato_models.SanctionedAppsCategoryRefInput{}
	for _, val := range plan.Rule.Destination.SanctionedAppsCategory {
		sanctionedDestAppsCategory = append(sanctionedDestAppsCategory, &cato_models.SanctionedAppsCategoryRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	serviceInput := &cato_models.InternetFirewallServiceTypeUpdateInput{}

	ruleTrackingAlertSubscriptionGroup := []*cato_models.SubscriptionGroupRefInput{}
	for _, val := range plan.Rule.Destination.Application {
		ruleTrackingAlertSubscriptionGroup = append(ruleTrackingAlertSubscriptionGroup, &cato_models.SubscriptionGroupRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	ruleTrackingAlertMailingList := []*cato_models.SubscriptionMailingListRefInput{}
	for _, val := range plan.Rule.Tracking.Alert.MailingList {
		ruleTrackingAlertMailingList = append(ruleTrackingAlertMailingList, &cato_models.SubscriptionMailingListRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	ruleTrackingAlertSubscriptionWebhook := []*cato_models.SubscriptionWebhookRefInput{}
	for _, val := range plan.Rule.Destination.Application {
		ruleTrackingAlertSubscriptionWebhook = append(ruleTrackingAlertSubscriptionWebhook, &cato_models.SubscriptionWebhookRefInput{
			By:    cato_models.ObjectRefBy(val.By.ValueString()),
			Input: val.Input.ValueString(),
		})
	}

	activeOnUpdate := cato_models.PolicyActiveOnEnum(plan.Rule.Schedule.ActiveOn.ValueString())
	frequencyUpdate := cato_models.PolicyRuleTrackingFrequencyEnum(plan.Rule.Tracking.Alert.Frequency.ValueString())
	sourceIpSlice := makeStringSliceFromStringList(plan.Rule.Source.IP)
	destIpSlice := makeStringSliceFromStringList(plan.Rule.Destination.IP)
	destSubnetSlice := makeStringSliceFromStringList(plan.Rule.Destination.Subnet)
	destRemoteAsn := makeStringSliceFromStringList(plan.Rule.Destination.RemoteAsn)

	updateInput := cato_models.InternetFirewallUpdateRuleInput{
		ID: state.Rule.ID.ValueString(),
		Rule: &cato_models.InternetFirewallUpdateRuleDataInput{
			Enabled:     state.Rule.Enabled.ValueBoolPointer(),
			Name:        state.Rule.Name.ValueStringPointer(),
			Description: state.Rule.Description.ValueStringPointer(),
			Source: &cato_models.InternetFirewallSourceUpdateInput{
				IP:                sourceIpSlice,
				Host:              hostSourceRefInput,
				Site:              siteSourceRefInput,
				Subnet:            subnetSourceRefInput,
				IPRange:           ipSourcerange,
				GlobalIPRange:     globalSourceIpRange,
				NetworkInterface:  networkSourceInterfaceRefInput,
				SiteNetworkSubnet: siteSourceNetworkSubnetRefInput,
				FloatingSubnet:    floatingSourceSubnetRefInput,
				User:              userSourceRefInput,
				UsersGroup:        usersSourceGroupRefInput,
				Group:             groupSourceRefInput,
				SystemGroup:       systemSouceGroupRefInput,
			},
			ConnectionOrigin: &connectionOrigin,
			Country:          countryInput,
			Device:           deviceInput,
			DeviceOs:         deviceOsInput,
			Destination: &cato_models.InternetFirewallDestinationUpdateInput{
				Application:            applicationDestInput,
				CustomApp:              customAppDestInput,
				AppCategory:            appDestCategoryInput,
				CustomCategory:         customDestCategory,
				SanctionedAppsCategory: sanctionedDestAppsCategory,
				Country:                countryDestInput,
				Domain:                 domainDestList,
				Fqdn:                   fqdnDestList,
				IP:                     destIpSlice,
				Subnet:                 destSubnetSlice,
				IPRange:                ipDestRange,
				GlobalIPRange:          globalDestIpRange,
				RemoteAsn:              destRemoteAsn,
			},
			Service: serviceInput,
			Action:  &actionEnum,
			Schedule: &cato_models.PolicyScheduleUpdateInput{
				ActiveOn: &activeOnUpdate,
			},
			Tracking: &cato_models.PolicyTrackingUpdateInput{
				Event: &cato_models.PolicyRuleTrackingEventUpdateInput{
					Enabled: plan.Rule.Tracking.Event.Enabled.ValueBoolPointer(),
				},
				Alert: &cato_models.PolicyRuleTrackingAlertUpdateInput{
					Enabled:           plan.Rule.Tracking.Alert.Enabled.ValueBoolPointer(),
					Frequency:         &frequencyUpdate,
					SubscriptionGroup: ruleTrackingAlertSubscriptionGroup,
					MailingList:       ruleTrackingAlertMailingList,
					Webhook:           ruleTrackingAlertSubscriptionWebhook,
				},
			},
		},
	}

	_, err := r.info.catov2.PolicyInternetFirewallUpdateRule(ctx, mutationInput, updateInput, r.info.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API PolicyInternetFirewallUpdateRule error",
			err.Error(),
		)
		return
	}

	if plan.Publish.ValueBool() {
		publishDataIfEnabled := &cato_models.PolicyPublishRevisionInput{
			Name: plan.SdkKeyName.ValueStringPointer(),
		}
		_, err := r.info.catov2.PolicyInternetFirewallPublishPolicyRevision(ctx, &cato_models.InternetFirewallPolicyMutationInput{}, publishDataIfEnabled, r.info.AccountId)
		if err != nil {
			resp.Diagnostics.AddError(
				"Catov2 API PolicyInternetFirewallPublishPolicyRevision error",
				err.Error(),
			)
			return
		}
	}

	// Submission of plan results to state

	queryPolicy := &cato_models.InternetFirewallPolicyInput{}
	queryResult, err := r.info.catov2.Policy(ctx, queryPolicy, r.info.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Catov2 API Policy(ctx, queryPolicy, r.info.AccountId) error",
			err.Error(),
		)
		return
	}
	ruleList := queryResult.GetPolicy().InternetFirewall.Policy.GetRules()
	for _, ruleListItem := range ruleList {
		if ruleListItem.GetRule().ID == plan.Rule.ID.ValueString() {
			plan.Rule.ID = types.StringValue(ruleListItem.GetRule().ID)
			plan.Rule.Action = types.StringValue(ruleListItem.Rule.Action.String())
			plan.Rule.ConnectionOrigin = types.StringValue(ruleListItem.Rule.ConnectionOrigin.String())
			plan.Rule.Description = types.StringValue(ruleListItem.Rule.GetDescription())
			plan.Rule.Name = types.StringValue(ruleListItem.GetRule().Name)
			countryRuleItem := []Policy_Policy_InternetFirewall_Policy_Rules_Rule_Country{}
			for _, val := range ruleListItem.Rule.GetCountry() {
				countryRuleItem = append(countryRuleItem, Policy_Policy_InternetFirewall_Policy_Rules_Rule_Country{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Country = countryRuleItem

			plan.Rule.Destination = Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination{}

			dstAppCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_AppCategory{}
			for _, val := range ruleListItem.Rule.Destination.AppCategory {
				dstAppCategory = append(dstAppCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_AppCategory{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Destination.AppCategory = dstAppCategory

			dstApplication := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Application{}
			for _, val := range ruleListItem.Rule.Destination.Application {
				dstApplication = append(dstApplication, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Application{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Destination.Application = dstApplication

			dstCountry := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Country{}
			for _, val := range ruleListItem.Rule.Destination.Country {
				dstCountry = append(dstCountry, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_Country{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Destination.Country = dstCountry

			dstCustomApp := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomApp{}
			for _, val := range ruleListItem.Rule.Destination.CustomApp {
				dstCustomApp = append(dstCustomApp, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomApp{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Destination.CustomApp = dstCustomApp

			dstCustomCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomCategory{}
			for _, val := range ruleListItem.Rule.Destination.CustomCategory {
				dstCustomCategory = append(dstCustomCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_CustomCategory{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Destination.CustomCategory = dstCustomCategory

			dstDomain := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.Domain {
				dstDomain = append(dstDomain, types.StringValue(val))
			}
			plan.Rule.Destination.Domain = dstDomain

			dstFqdn := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.Fqdn {
				dstDomain = append(dstDomain, types.StringValue(val))
			}
			plan.Rule.Destination.Fqdn = dstFqdn

			dstGlobalRange := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_GlobalIPRange{}
			for _, val := range ruleListItem.Rule.Destination.GlobalIPRange {
				dstGlobalRange = append(dstGlobalRange, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_GlobalIPRange{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Destination.GlobalIPRange = dstGlobalRange

			dstIp := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.IP {
				dstIp = append(dstIp, types.StringValue(val))
			}
			plan.Rule.Destination.IP = dstIp

			dstIPRange := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_IPRange{}
			for _, val := range ruleListItem.Rule.Destination.IPRange {
				dstIPRange = append(dstIPRange, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_IPRange{
					From: types.StringValue(val.From),
					To:   types.StringValue(val.To),
				})
			}
			plan.Rule.Destination.IPRange = dstIPRange

			dstRemoteAsn := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.RemoteAsn {
				dstRemoteAsn = append(dstRemoteAsn, types.StringValue(val))
			}
			plan.Rule.Destination.RemoteAsn = dstRemoteAsn

			dstSanctionedAppsCategory := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_SanctionedAppsCategory{}
			for _, val := range ruleListItem.Rule.Destination.SanctionedAppsCategory {
				dstSanctionedAppsCategory = append(dstSanctionedAppsCategory, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Destination_SanctionedAppsCategory{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Destination.SanctionedAppsCategory = dstSanctionedAppsCategory

			dstSSubnet := []types.String{}
			for _, val := range ruleListItem.Rule.Destination.Subnet {
				dstSSubnet = append(dstSSubnet, types.StringValue(val))
			}
			plan.Rule.Destination.Subnet = dstSSubnet

			deviceList := []Policy_Policy_InternetFirewall_Policy_Rules_Rule_Device{}
			for _, val := range ruleListItem.Rule.Device {
				deviceList = append(deviceList, Policy_Policy_InternetFirewall_Policy_Rules_Rule_Device{
					By:    types.StringValue("NAME"),
					Input: types.StringValue(val.Name),
				})
			}
			plan.Rule.Device = deviceList

			deviceOSList := []types.String{}
			for _, val := range ruleListItem.Rule.DeviceOs {
				deviceOSList = append(deviceOSList, types.StringValue(string(val)))
			}
			plan.Rule.DeviceOs = deviceOSList

			exceptionsList := []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions{}
			for _, val := range ruleListItem.Rule.Exceptions {
				exceptionsList = append(exceptionsList, &Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions{
					Name:             types.StringValue(val.Name),
					Source:           Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Source{},
					DeviceOs:         []OperatingSystem{},
					Country:          []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Country{},
					Device:           []*Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Device{},
					Destination:      Policy_Policy_InternetFirewall_Policy_Rules_Rule_Exceptions_Destination{},
					ConnectionOrigin: types.StringValue(""),
				})
			}
			plan.Rule.Exceptions = exceptionsList
		}
	}

	plan.Rule.ID = state.Rule.ID

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *internetFwPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state InternetFirewallCreateRuleInput
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	removeMutations := &cato_models.InternetFirewallPolicyMutationInput{}
	removeRule := cato_models.InternetFirewallRemoveRuleInput{
		ID: state.Rule.ID.ValueString(),
	}

	_, err := r.info.catov2.PolicyInternetFirewallRemoveRule(ctx, removeMutations, removeRule, r.info.AccountId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to connect or request the Catov2 API",
			err.Error(),
		)
		return
	}

	publishDataIfEnabled := &cato_models.PolicyPublishRevisionInput{
		Name: state.SdkKeyName.ValueStringPointer(),
	}
	_, errSdk := r.info.catov2.PolicyInternetFirewallPublishPolicyRevision(ctx, &cato_models.InternetFirewallPolicyMutationInput{}, publishDataIfEnabled, r.info.AccountId)
	if errSdk != nil {
		resp.Diagnostics.AddError(
			"Catov2 API Delete/PolicyInternetFirewallPublishPolicyRevision error",
			errSdk.Error(),
		)
		return
	}

}
