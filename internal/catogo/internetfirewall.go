package catogo

import (
	"encoding/json"
	"fmt"
)

type PolicyResult struct {
	InternetFirewall struct {
		Policy                InternetFirewallPolicy                `json:"policy,omitempty"`
		AddRule               InternetFirewallRuleMutationPayload   `json:"addRule,omitempty"`
		UpdateRule            InternetFirewallRuleMutationPayload   `json:"updateRule,omitempty"`
		RemoveRule            InternetFirewallRuleMutationPayload   `json:"removeRule,omitempty"`
		PublishPolicyRevision InternetFirewallPolicyMutationPayload `json:"publishPolicyRevision,omitempty"`
		DiscardPolicyRevision InternetFirewallPolicyMutationPayload `json:"discardPolicyRevision,omitempty"`
	} `json:"internetFirewall,omitempty"`
}

type InternetFirewall struct {
	InternetFirewall InternetFirewallPolicy `json:"internetFirewall,omitempty"`
}

type InternetFirewallPolicy struct {
	Enabled  bool                          `json:"enabled,omitempty"`
	Rules    []InternetFirewallRulePayload `json:"rules,omitempty"`
	Sections []PolicySectionPayload        `json:"sections,omitempty"`
	Audit    *PolicyAudit                  `json:"audit,omitempty"`
	Revision *PolicyRevision               `json:"revision,omitempty"`
}

type InternetFirewallRule struct {
	Id               string                          `json:"id,omitempty"`
	Name             string                          `json:"name,omitempty"`
	Description      string                          `json:"description,omitempty"`
	Index            int64                           `json:"index,omitempty"`
	Section          PolicySectionInfo               `json:"section,omitempty"`
	Enabled          bool                            `json:"enabled,omitempty"`
	Source           InternetFirewallSource          `json:"source,omitempty"`
	ConnectionOrigin string                          `json:"connectionOrigin,omitempty"`
	Country          []ObjectRef                     `json:"country,omitempty"`
	Device           []ObjectRef                     `json:"device,omitempty"`
	DeviceOS         []string                        `json:"deviceOS,omitempty"`
	Destination      InternetFirewallDestination     `json:"destination,omitempty"`
	Service          InternetFirewallServiceType     `json:"service,omitempty"`
	Action           string                          `json:"action,omitempty"`
	Tracking         PolicyTracking                  `json:"tracking,omitempty"`
	Schedule         PolicySchedule                  `json:"schedule,omitempty"`
	Exceptions       []InternetFirewallRuleException `json:"exceptions,omitempty"`
}

type InternetFirewallRuleMutationPayload struct {
	Rule   *InternetFirewallRulePayload `json:"rule,omitempty"`
	Status string                       `json:"status,omitempty"`
	Errors []PolicyMutationError        `json:"errors,omitempty"`
}

type InternetFirewallRulePayload struct {
	Audit      PolicyElementAudit   `json:"audit,omitempty"`
	Rule       InternetFirewallRule `json:"rule,omitempty"`
	Properties []string             `json:"properties,omitempty"`
}

type InternetFirewallPolicyMutationPayload struct {
	// Policy *InternetFirewallPolicy `json:"policy"`
	Status string                `json:"status"`
	Errors []PolicyMutationError `json:"errors"`
}

type PolicyMutationError struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty"`
}

type ObjectRef struct {
	By    string `json:"by,omitempty"`
	Input string `json:"input,omitempty"`
}

type IpAddressRange struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type PortRange struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type CustomService struct {
	Port      []string   `json:"port,omitempty"`
	PortRange *PortRange `json:"portRange,omitempty"`
	Protocol  string     `json:"protocol,omitempty"`
}

type InternetFirewallDestination struct {
	Application            []ObjectRef      `json:"application,omitempty"`
	CustomApp              []ObjectRef      `json:"customApp,omitempty"`
	AppCategory            []ObjectRef      `json:"appCategory,omitempty"`
	CustomCategory         []ObjectRef      `json:"customCategory,omitempty"`
	SanctionedAppsCategory []ObjectRef      `json:"sanctionedAppsCategory,omitempty"`
	Country                []ObjectRef      `json:"country,omitempty"`
	Domain                 []string         `json:"domain,omitempty"`
	Fqdn                   []string         `json:"fqdn,omitempty"`
	Ip                     []string         `json:"ip,omitempty"`
	Subnet                 []string         `json:"subnet,omitempty"`
	IpRange                []IpAddressRange `json:"ipRange,omitempty"`
	GlobalIpRange          []ObjectRef      `json:"globalIpRange,omitempty"`
	RemoteAsn              []string         `json:"remoteAsn,omitempty"`
}

type InternetFirewallRuleException struct {
	Name             string                      `json:"name,omitempty"`
	Source           InternetFirewallSource      `json:"source,omitempty"`
	DeviceOS         []string                    `json:"deviceOS,omitempty"`
	Country          []ObjectRef                 `json:"country,omitempty"`
	Device           []ObjectRef                 `json:"device,omitempty"`
	Destination      InternetFirewallDestination `json:"destination,omitempty"`
	Service          InternetFirewallServiceType `json:"service,omitempty"`
	ConnectionOrigin string                      `json:"connectionOrigin,omitempty"`
}

type InternetFirewallServiceType struct {
	Standard []ObjectRef     `json:"standard,omitempty"`
	Custom   []CustomService `json:"custom,omitempty"`
}

type InternetFirewallSource struct {
	Ip                []string         `json:"ip,omitempty"`
	Host              []ObjectRef      `json:"host,omitempty"`
	Site              []ObjectRef      `json:"site,omitempty"`
	Subnet            []string         `json:"subnet,omitempty"`
	IpRange           []IpAddressRange `json:"ipRange,omitempty"`
	GlobalIpRange     []ObjectRef      `json:"globalIpRange,omitempty"`
	NetworkInterface  []ObjectRef      `json:"networkInterface,omitempty"`
	SiteNetworkSubnet []ObjectRef      `json:"siteNetworkSubnet,omitempty"`
	FloatingSubnet    []ObjectRef      `json:"floatingSubnet,omitempty"`
	User              []ObjectRef      `json:"user,omitempty"`
	UsersGroup        []ObjectRef      `json:"usersGroup,omitempty"`
	Group             []ObjectRef      `json:"group,omitempty"`
	SystemGroup       []ObjectRef      `json:"systemGroup,omitempty"`
}

type PolicyAudit struct {
	PublishedTime string `json:"publishedTime,omitempty"`
	PublishedBy   string `json:"publishedBy,omitempty"`
}

type PolicyCustomRecurring struct {
	From string   `json:"from,omitempty"`
	To   string   `json:"to,omitempty"`
	Days []string `json:"days,omitempty"`
}

type PolicyCustomTimeframe struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type PolicyElementAudit struct {
	UpdatedTime string `json:"updatedTime,omitempty"`
	UpdatedBy   string `json:"updatedBy,omitempty"`
}

type PolicyRevision struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Changes     int64  `json:"changes,omitempty"`
	CreatedTime string `json:"createdTime,omitempty"`
	UpdatedTime string `json:"updatedTime,omitempty"`
}

type PolicyRuleTrackingAlert struct {
	Enabled           bool        `json:"enabled,omitempty"`
	Frequency         string      `json:"frequency,omitempty"`
	SubscriptionGroup []ObjectRef `json:"subscriptionGroup,omitempty"`
	Webhook           []ObjectRef `json:"webhook,omitempty"`
	MailingList       []ObjectRef `json:"mailingList,omitempty"`
}

type PolicyRuleTrackingEvent struct {
	Enabled bool `json:"enabled,omitempty"`
}

type PolicySchedule struct {
	ActiveOn        string                 `json:"activeOn,omitempty"`
	CustomTimeframe *PolicyCustomTimeframe `json:"customTimeframe,omitempty"`
	CustomRecurring *PolicyCustomRecurring `json:"customRecurring,omitempty"`
}

type PolicySectionInfo struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PolicySectionPayload struct {
	Audit      PolicyElementAudit `json:"audit,omitempty"`
	Section    PolicySectionInfo  `json:"section,omitempty"`
	Properties []string           `json:"properties,omitempty"`
}

type PolicyTracking struct {
	Event PolicyRuleTrackingEvent `json:"event,omitempty"`
	Alert PolicyRuleTrackingAlert `json:"alert,omitempty"`
}

type InternetFirewallAddRuleDataInput struct {
	Enabled          bool                                 `json:"enabled,omitempty"`
	Name             string                               `json:"name,omitempty"`
	Description      string                               `json:"description,omitempty"`
	Source           *InternetFirewallSourceInput         `json:"source,omitempty"`
	ConnectionOrigin string                               `json:"connectionOrigin,omitempty"`
	Country          []ObjectRefInput                     `json:"country,omitempty"`
	Device           []ObjectRefInput                     `json:"device,omitempty"`
	DeviceOS         []string                             `json:"deviceOS,omitempty"`
	Destination      *InternetFirewallDestinationInput    `json:"destination,omitempty"`
	Service          *InternetFirewallServiceTypeInput    `json:"service,omitempty"`
	Action           string                               `json:"action,omitempty"`
	Tracking         *PolicyTrackingInput                 `json:"tracking,omitempty"`
	Schedule         *PolicyScheduleInput                 `json:"schedule,omitempty"`
	Exceptions       []InternetFirewallRuleExceptionInput `json:"exceptions,omitempty"`
}

type InternetFirewallSourceInput struct {
	Ip                []string              `json:"ip,omitempty"`
	Host              []ObjectRefInput      `json:"host,omitempty"`
	Site              []ObjectRefInput      `json:"site,omitempty"`
	Subnet            []string              `json:"subnet,omitempty"`
	IpRange           []IpAddressRangeInput `json:"ipRange,omitempty"`
	GlobalIpRange     []ObjectRefInput      `json:"globalIpRange,omitempty"`
	NetworkInterface  []ObjectRefInput      `json:"networkInterface,omitempty"`
	SiteNetworkSubnet []ObjectRefInput      `json:"siteNetworkSubnet,omitempty"`
	FloatingSubnet    []ObjectRefInput      `json:"floatingSubnet,omitempty"`
	User              []ObjectRefInput      `json:"user,omitempty"`
	UsersGroup        []ObjectRefInput      `json:"usersGroup,omitempty"`
	Group             []ObjectRefInput      `json:"group,omitempty"`
	SystemGroup       []ObjectRefInput      `json:"systemGroup,omitempty"`
}

type InternetFirewallDestinationInput struct {
	Application            []ObjectRefInput      `json:"application,omitempty"`
	CustomApp              []ObjectRefInput      `json:"customApp,omitempty"`
	AppCategory            []ObjectRefInput      `json:"appCategory,omitempty"`
	CustomCategory         []ObjectRefInput      `json:"customCategory,omitempty"`
	SanctionedAppsCategory []ObjectRefInput      `json:"sanctionedAppsCategory,omitempty"`
	Country                []ObjectRefInput      `json:"country,omitempty"`
	Domain                 []string              `json:"domain,omitempty"`
	Fqdn                   []string              `json:"fqdn,omitempty"`
	Ip                     []string              `json:"ip,omitempty"`
	Subnet                 []string              `json:"subnet,omitempty"`
	IpRange                []IpAddressRangeInput `json:"ipRange,omitempty"`
	GlobalIpRange          []ObjectRefInput      `json:"globalIpRange,omitempty"`
	RemoteAsn              []string              `json:"remoteAsn,omitempty"`
}

type InternetFirewallRuleExceptionInput struct {
	Name             string                           `json:"name,omitempty"`
	Source           InternetFirewallSourceInput      `json:"source,omitempty"`
	DeviceOS         []string                         `json:"deviceOS,omitempty"`
	Country          []ObjectRefInput                 `json:"country,omitempty"`
	Device           []ObjectRefInput                 `json:"device,omitempty"`
	Destination      InternetFirewallDestinationInput `json:"destination,omitempty"`
	Service          InternetFirewallServiceTypeInput `json:"service,omitempty"`
	ConnectionOrigin string                           `json:"connectionOrigin,omitempty"`
}

type InternetFirewallServiceTypeInput struct {
	Standard []ObjectRefInput     `json:"standard,omitempty"`
	Custom   []CustomServiceInput `json:"custom,omitempty"`
}

type PolicyCustomRecurringInput struct {
	From string   `json:"from,omitempty"`
	To   string   `json:"to,omitempty"`
	Days []string `json:"days,omitempty"`
}

type PolicyCustomTimeframeInput struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type PolicyRulePositionInput struct {
	Position *string `json:"position,omitempty"`
	Ref      *string `json:"ref,omitempty"`
}

type InternetFirewallAddRuleInput struct {
	Rule InternetFirewallAddRuleDataInput `json:"rule"`
	At   *PolicyRulePositionInput         `json:"at"`
}

type PolicyRuleTrackingAlertInput struct {
	Enabled           bool             `json:"enabled,omitempty"`
	Frequency         string           `json:"frequency,omitempty"`
	SubscriptionGroup []ObjectRefInput `json:"subscriptionGroup,omitempty"`
	Webhook           []ObjectRefInput `json:"webhook,omitempty"`
	MailingList       []ObjectRefInput `json:"mailingList,omitempty"`
}

type PolicyRuleTrackingEventInput struct {
	Enabled bool `json:"enabled,omitempty"`
}

type PolicyScheduleInput struct {
	ActiveOn        string                      `json:"activeOn,omitempty"`
	CustomTimeframe *PolicyCustomTimeframeInput `json:"customTimeframe,omitempty"`
	CustomRecurring *PolicyCustomRecurringInput `json:"customRecurring,omitempty"`
}

type PolicyTrackingInput struct {
	Event PolicyRuleTrackingEventInput `json:"event,omitempty"`
	Alert PolicyRuleTrackingAlertInput `json:"alert,omitempty"`
}

type IpAddressRangeInput struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type PortRangeInput struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type CustomServiceInput struct {
	Port      []string        `json:"port,omitempty"`
	PortRange *PortRangeInput `json:"portRange,omitempty"`
	Protocol  string          `json:"protocol,omitempty"`
}

type ObjectRefInput struct {
	By    string `json:"by,omitempty"`
	Input string `json:"input,omitempty"`
}

type InternetFirewallUpdateRuleInput struct {
	Id   string                              `json:"id"`
	Rule InternetFirewallUpdateRuleDataInput `json:"rule"`
}

type InternetFirewallUpdateRuleDataInput struct {
	Enabled          bool                                       `json:"enabled"`
	Name             string                                     `json:"name"`
	Description      string                                     `json:"description"`
	Source           *InternetFirewallUpdateSourceInput         `json:"source"`
	ConnectionOrigin string                                     `json:"connectionOrigin,omitempty"`
	Country          []ObjectRefUpdateInput                     `json:"country"`
	Device           []ObjectRefUpdateInput                     `json:"device"`
	DeviceOS         []string                                   `json:"deviceOS"`
	Destination      *InternetFirewallUpdateDestinationInput    `json:"destination"`
	Service          *InternetFirewallUpdateServiceTypeInput    `json:"service"`
	Action           string                                     `json:"action"`
	Tracking         *PolicyTrackingUpdateInput                 `json:"tracking"`
	Schedule         *PolicyScheduleUpdateInput                 `json:"schedule"`
	Exceptions       []InternetFirewallUpdateRuleExceptionInput `json:"exceptions"`
}

type InternetFirewallUpdateSourceInput struct {
	Ip                []string                    `json:"ip"`
	Host              []ObjectRefUpdateInput      `json:"host"`
	Site              []ObjectRefUpdateInput      `json:"site"`
	Subnet            []string                    `json:"subnet"`
	IpRange           []IpAddressRangeUpdateInput `json:"ipRange"`
	GlobalIpRange     []ObjectRefUpdateInput      `json:"globalIpRange"`
	NetworkInterface  []ObjectRefUpdateInput      `json:"networkInterface"`
	SiteNetworkSubnet []ObjectRefUpdateInput      `json:"siteNetworkSubnet"`
	FloatingSubnet    []ObjectRefUpdateInput      `json:"floatingSubnet"`
	User              []ObjectRefUpdateInput      `json:"user"`
	UsersGroup        []ObjectRefUpdateInput      `json:"usersGroup"`
	Group             []ObjectRefUpdateInput      `json:"group"`
	SystemGroup       []ObjectRefUpdateInput      `json:"systemGroup"`
}

type InternetFirewallUpdateDestinationInput struct {
	Application            []ObjectRefUpdateInput      `json:"application"`
	CustomApp              []ObjectRefUpdateInput      `json:"customApp"`
	AppCategory            []ObjectRefUpdateInput      `json:"appCategory"`
	CustomCategory         []ObjectRefUpdateInput      `json:"customCategory"`
	SanctionedAppsCategory []ObjectRefUpdateInput      `json:"sanctionedAppsCategory"`
	Country                []ObjectRefUpdateInput      `json:"country"`
	Domain                 []string                    `json:"domain"`
	Fqdn                   []string                    `json:"fqdn"`
	Ip                     []string                    `json:"ip"`
	Subnet                 []string                    `json:"subnet"`
	IpRange                []IpAddressRangeUpdateInput `json:"ipRange"`
	GlobalIpRange          []ObjectRefUpdateInput      `json:"globalIpRange"`
	RemoteAsn              []string                    `json:"remoteAsn"`
}

type InternetFirewallUpdateRuleExceptionInput struct {
	Name             string                                 `json:"name"`
	Source           InternetFirewallUpdateSourceInput      `json:"source"`
	DeviceOS         []string                               `json:"deviceOS"`
	Country          []ObjectRefUpdateInput                 `json:"country"`
	Device           []ObjectRefUpdateInput                 `json:"device"`
	Destination      InternetFirewallUpdateDestinationInput `json:"destination"`
	Service          InternetFirewallUpdateServiceTypeInput `json:"service"`
	ConnectionOrigin string                                 `json:"connectionOrigin"`
}

type InternetFirewallUpdateServiceTypeInput struct {
	Standard []ObjectRefInput           `json:"standard"`
	Custom   []CustomServiceUpdateInput `json:"custom"`
}

type PolicyCustomRecurringUpdateInput struct {
	From string   `json:"from"`
	To   string   `json:"to"`
	Days []string `json:"days"`
}

type PolicyCustomTimeframeUpdateInput struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type PolicyRuleTrackingAlertUpdateInput struct {
	Enabled           bool                   `json:"enabled"`
	Frequency         string                 `json:"frequency"`
	SubscriptionGroup []ObjectRefUpdateInput `json:"subscriptionGroup"`
	Webhook           []ObjectRefUpdateInput `json:"webhook"`
	MailingList       []ObjectRefUpdateInput `json:"mailingList"`
}

type PolicyRuleTrackingEventUpdateInput struct {
	Enabled bool `json:"enabled"`
}

type PolicyScheduleUpdateInput struct {
	ActiveOn        string                            `json:"activeOn"`
	CustomTimeframe *PolicyCustomTimeframeUpdateInput `json:"customTimeframe"`
	CustomRecurring *PolicyCustomRecurringUpdateInput `json:"customRecurring"`
}

type PolicyTrackingUpdateInput struct {
	Event PolicyRuleTrackingEventUpdateInput `json:"event"`
	Alert PolicyRuleTrackingAlertUpdateInput `json:"alert"`
}

type IpAddressRangeUpdateInput struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type PortRangeUpdateInput struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type CustomServiceUpdateInput struct {
	Port      []string              `json:"port"`
	PortRange *PortRangeUpdateInput `json:"portRange"`
	Protocol  string                `json:"protocol"`
}

type ObjectRefUpdateInput struct {
	By    string `json:"by"`
	Input string `json:"input"`
}

type InternetFirewallRemoveRuleInput struct {
	Id string `json:"id"`
}

func (c *Client) GetInternetFirewallPolicy() (*InternetFirewallPolicy, error) {

	query := graphQLRequest{
		Query: `query InternetFirewall($accountId: ID!) {
					policy(accountId: $accountId) {
						internetFirewall {
						policy {
							audit {
							publishedBy
							publishedTime
							}
							enabled
							revision {
							changes
							createdTime
							description
							id
							name
							updatedTime
							}
							rules {
							properties
							audit {
								updatedBy
								updatedTime
							}
							rule {
								action
								connectionOrigin
								country {
								id
								name
								}
								source {
								ip
								subnet
								ipRange {
									from
									to
								}
								floatingSubnet {
									id
									name
								}
								group {
									id
									name
								}
								site {
									id
									name
								}
								host {
									id
									name
								}
								usersGroup {
									id
									name
								}
								user {
									id
									name
								}
								systemGroup {
									id
									name
								}
								}
								section {
								id
								name
								}
								schedule {
								activeOn
								}
								name
								index
								id
								description
								destination {
								appCategory {
									id
									name
								}
								application {
									id
									name
								}
								domain
								fqdn
								ip
								subnet
								remoteAsn
								}
								enabled
								deviceOS
								device {
								id
								name
								}
								tracking {
								alert {
									enabled
								}
								event {
									enabled
								}
								}
							}
							}
							sections {
							properties
							audit {
								updatedBy
								updatedTime
							}
							section {
								id
								name
							}
							}
						}
						}
					}
        }`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct{ Policy PolicyResult }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Policy.InternetFirewall.Policy, nil
}

func (c *Client) GetInternetFirewallRuleByName(name string) (*InternetFirewallRule, error) {

	policy, _ := c.GetInternetFirewallPolicy()

	rule := InternetFirewallRule{}

	for _, item := range policy.Rules {

		if item.Rule.Name == name {
			rule = item.Rule
		}

	}

	return &rule, nil
}

func (c *Client) GetInternetFirewallRuleById(Id string) (*InternetFirewallRule, error) {

	policy, _ := c.GetInternetFirewallPolicy()

	rule := InternetFirewallRule{}

	for _, item := range policy.Rules {

		if item.Rule.Id == Id {
			rule = item.Rule
		}

	}

	return &rule, nil
}

func (c *Client) CreateInternetFirewallRule(rule InternetFirewallAddRuleInput) (*InternetFirewallRuleMutationPayload, error) {

	query := graphQLRequest{
		Query: `mutation AddInternetFirewallRule($accountId: ID!, $input: InternetFirewallAddRuleInput!) {
					policy(accountId: $accountId) {
					internetFirewall {
						addRule( input: $input ) {
							status
							errors {
								errorCode
								errorMessage
							}
							rule {
								rule {

								action
								connectionOrigin
								country {
								id
								name
								}
								source {
								ip
								subnet
								ipRange {
									from
									to
								}
								floatingSubnet {
									id
									name
								}
								group {
									id
									name
								}
								site {
									id
									name
								}
								host {
									id
									name
								}
								usersGroup {
									id
									name
								}
								user {
									id
									name
								}
								systemGroup {
									id
									name
								}
								}
								section {
								id
								name
								}
								schedule {
								activeOn
								}
								name
								index
								id
								description
								destination {
								appCategory {
									id
									name
								}
								application {
									id
									name
								}
								domain
								fqdn
								ip
								subnet
								remoteAsn
								}
								enabled
								deviceOS
								device {
								id
								name
								}
								tracking {
								alert {
									enabled
								}
								event {
									enabled
								}
								}
							}
							}
						}
					}
					}
				}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"input":     rule,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct{ Policy PolicyResult }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Policy.InternetFirewall.AddRule, nil
}

func (c *Client) UpdateInternetFirewallRule(input InternetFirewallUpdateRuleInput) (*InternetFirewallRuleMutationPayload, error) {

	query := graphQLRequest{
		Query: `mutation UpdateInternetFirewallRule($accountId: ID!, $input: InternetFirewallUpdateRuleInput!) {
					policy(accountId: $accountId) {
					internetFirewall {
						updateRule( input: $input) {
							status
							errors {
								errorCode
								errorMessage
							}
						}
					}
					}
				}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"input":     input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct{ Policy PolicyResult }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Policy.InternetFirewall.UpdateRule, nil
}

func (c *Client) RemoveInternetFirewallRule(input InternetFirewallRemoveRuleInput) (*InternetFirewallRuleMutationPayload, error) {

	query := graphQLRequest{
		Query: `mutation RemoveInternetFirewallRule($accountId: ID!, $input: InternetFirewallRemoveRuleInput!) {
					policy(accountId: $accountId) {
						internetFirewall {
							removeRule( input: $input ) {
							status
							errors {
								errorCode
								errorMessage
							}
						}
					}
					}
				}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"input":     input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var response struct{ Policy PolicyResult }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Policy.InternetFirewall.RemoveRule, nil
}

func (c *Client) PublishInternetFirewallDefaultPolicyRevision() (*InternetFirewallPolicyMutationPayload, error) {

	query := graphQLRequest{
		Query: `mutation PublishFirewalRevision($accountId: ID!) {
			policy(accountId: $accountId) {
				internetFirewall {
				publishPolicyRevision {
					status
					errors {
						errorCode
						errorMessage
					}
				}
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct{ Policy PolicyResult }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Policy.InternetFirewall.PublishPolicyRevision, nil
}

func (c *Client) DiscardInternetFirewallDefaultPolicyRevision() (*InternetFirewallPolicyMutationPayload, error) {

	query := graphQLRequest{
		Query: `
		mutation DiscardPolicyRevision($accountId: ID!) {
			policy(accountId: $accountId) {
				internetFirewall {
					discardPolicyRevision {
						status
					}
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var response struct{ Policy PolicyResult }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Policy.InternetFirewall.DiscardPolicyRevision, nil
}
