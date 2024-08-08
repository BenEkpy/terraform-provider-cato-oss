package catogo

import "encoding/json"

type AddSocketSiteInput struct {
	Name               string               `json:"name,omitempty"`
	ConnectionType     string               `json:"connectionType,omitempty"`
	SiteType           string               `json:"siteType,omitempty"`
	Description        *string              `json:"description,omitempty"`
	NativeNetworkRange string               `json:"nativeNetworkRange,omitempty"`
	TranslatedSubnet   *string              `json:"translatedSubnet,omitempty"`
	SiteLocation       AddSiteLocationInput `json:"siteLocation,omitempty"`
}

type AddSiteLocationInput struct {
	CountryCode string  `json:"countryCode,omitempty"`
	StateCode   *string `json:"stateCode,omitempty"`
	Timezone    string  `json:"timezone,omitempty"`
	Address     *string `json:"address,omitempty"`
	City        *string `json:"city,omitempty"`
}

type AddSocketSitePayload struct {
	SiteId string `json:"siteId,omitempty"`
}

type RemoveSitePayload struct {
	SiteId string `json:"siteId,omitempty"`
}

type UpdateSiteGeneralDetailsPayload struct {
	SiteId string `json:"siteId,omitempty"`
}

type UpdateSiteGeneralDetailsInput struct {
	Name         *string                  `json:"name,omitempty"`
	SiteType     *string                  `json:"siteType,omitempty"`
	Description  *string                  `json:"description,omitempty"`
	SiteLocation *UpdateSiteLocationInput `json:"siteLocation,omitempty"`
}

type UpdateSiteLocationInput struct {
	CountryCode *string `json:"countryCode,omitempty"`
	StateCode   *string `json:"stateCode,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
	Address     *string `json:"address,omitempty"`
}

type UpdateSocketInterfaceInput struct {
	DestType  string                         `json:"destType,omitempty"`
	Name      *string                        `json:"name,omitempty"`
	Lan       *SocketInterfaceLanInput       `json:"lan,omitempty"`
	Bandwidth *SocketInterfaceBandwidthInput `json:"bandwidth,omitempty"`
	Wan       *SocketInterfaceWanInput       `json:"wan,omitempty"`
	OffCloud  *SocketInterfaceOffCloudInput  `json:"offCloud,omitempty"`
	AltWan    *SocketInterfaceAltWanInput    `json:"altWan,omitempty"`
	Lag       *SocketInterfaceLagInput       `json:"lag,omitempty"`
	Vrrp      *SocketInterfaceVrrpInput      `json:"vrrp,omitempty"`
}

type SocketInterfaceBandwidthInput struct {
	UpstreamBandwidth   int64 `json:"upstreamBandwidth,omitempty"`
	DownstreamBandwidth int64 `json:"downstreamBandwidth,omitempty"`
}

type SocketInterfaceWanInput struct {
	Role       string `json:"role,omitempty"`
	Precedence string `json:"precedence,omitempty"`
}

type SocketInterfaceLanInput struct {
	Subnet           string  `json:"subnet,omitempty"`
	TranslatedSubnet *string `json:"translatedSubnet,omitempty"`
	LocalIp          string  `json:"localIp,omitempty"`
}

type SocketInterfaceOffCloudInput struct {
	Enabled          bool    `json:"enabled,omitempty"`
	PublicIp         *string `json:"publicIp,omitempty"`
	PublicStaticPort *int64  `json:"publicStaticPort,omitempty"`
}

type SocketInterfaceAltWanInput struct {
	PrivateInterfaceIp string  `json:"privateInterfaceIp,omitempty"`
	PrivateNetwork     string  `json:"privateNetwork,omitempty"`
	PrivateGatewayIp   string  `json:"privateGatewayIp,omitempty"`
	PrivateVlanTag     *int64  `json:"privateVlanTag,omitempty"`
	PublicInterfaceIp  *string `json:"publicInterfaceIp,omitempty"`
	PublicNetwork      *string `json:"publicNetwork,omitempty"`
	PublicGatewayIp    *string `json:"publicGatewayIp,omitempty"`
	PublicVlanTag      *int64  `json:"publicVlanTag,omitempty"`
}

type SocketInterfaceLagInput struct {
	MinLinks int64 `json:"minLinks,omitempty"`
}

type SocketInterfaceVrrpInput struct {
	VrrpType *string `json:"vrrpType,omitempty"`
}

type UpdateSocketInterfacePayload struct {
	SiteId            string `json:"siteId,omitempty"`
	SocketInterfaceId string `json:"socketInterfaceId,omitempty"`
}

func (c *Client) AddSocketSite(input AddSocketSiteInput) (*AddSocketSitePayload, error) {

	query := graphQLRequest{
		Query: `
		mutation addSocketSite($accountId:ID!, $input:AddSocketSiteInput!){
		site(accountId:$accountId) {
			addSocketSite(input:$input) {
			siteId
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

	var response struct {
		Site struct{ AddSocketSite AddSocketSitePayload }
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.AddSocketSite, nil
}

func (c *Client) RemoveSite(siteId string) (*RemoveSitePayload, error) {

	query := graphQLRequest{
		Query: `
		mutation removeSite($accountId:ID!, $siteId:ID!){
			site(accountId:$accountId) {
				removeSite(siteId:$siteId) {
					siteId
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"siteId":    siteId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Site struct{ RemoveSite RemoveSitePayload }
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.RemoveSite, nil
}

func (c *Client) UpdateSiteGeneralDetails(siteId string, input UpdateSiteGeneralDetailsInput) (*UpdateSiteGeneralDetailsPayload, error) {

	query := graphQLRequest{
		Query: `
		mutation updateSiteGeneralDetails($accountId:ID!, $siteId:ID!, $input:UpdateSiteGeneralDetailsInput!) {
		site(accountId: $accountId){
			updateSiteGeneralDetails(siteId:$siteId, input:$input) {
			siteId
			}
		}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"siteId":    siteId,
			"input":     input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Site struct {
			UpdateSiteGeneralDetails UpdateSiteGeneralDetailsPayload
		}
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.UpdateSiteGeneralDetails, nil
}

func (c *Client) UpdateSocketInterface(siteId string, socketInterfaceId string, input UpdateSocketInterfaceInput) (*UpdateSocketInterfacePayload, error) {

	query := graphQLRequest{
		Query: `
		mutation updateSocketInterface($accountId: ID!, $siteId: ID!, $socketInterfaceId: SocketInterfaceIDEnum!, $input: UpdateSocketInterfaceInput!) {
		site(accountId: $accountId) {
			updateSocketInterface(
			siteId: $siteId
			socketInterfaceId: $socketInterfaceId
			input: $input
			) {
			siteId
			socketInterfaceId
			}
		}
		}`,
		Variables: map[string]interface{}{
			"accountId":         c.accountId,
			"siteId":            siteId,
			"socketInterfaceId": socketInterfaceId,
			"input":             input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Site struct {
			UpdateSocketInterface UpdateSocketInterfacePayload
		}
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.UpdateSocketInterface, nil
}
