package catogo

import "encoding/json"

type AddNetworkRangeInput struct {
	Name             string                    `json:"name,omitempty"`
	RangeType        string                    `json:"rangeType,omitempty"`
	Subnet           string                    `json:"subnet,omitempty"`
	TranslatedSubnet *string                   `json:"translatedSubnet,omitempty"`
	LocalIp          *string                   `json:"localIp,omitempty"`
	Gateway          *string                   `json:"gateway,omitempty"`
	Vlan             *int64                    `json:"vlan,omitempty"`
	AzureFloatingIp  *string                   `json:"azureFloatingIp,omitempty"`
	DhcpSettings     *NetworkDhcpSettingsInput `json:"dhcpSettings,omitempty"`
}

type NetworkDhcpSettingsInput struct {
	DhcpType     string  `json:"dhcpType,omitempty"`
	IpRange      *string `json:"ipRange,omitempty"`
	RelayGroupId *string `json:"relayGroupId,omitempty"`
}

type AddNetworkRangePayload struct {
	NetworkRangeId string `json:"networkRangeId,omitempty"`
}

type UpdateNetworkRangeInput struct {
	Name             *string                   `json:"name,omitempty"`
	RangeType        *string                   `json:"rangeType,omitempty"`
	Subnet           *string                   `json:"subnet,omitempty"`
	TranslatedSubnet *string                   `json:"translatedSubnet,omitempty"`
	LocalIp          *string                   `json:"localIp,omitempty"`
	Gateway          *string                   `json:"gateway,omitempty"`
	Vlan             *int64                    `json:"vlan,omitempty"`
	AzureFloatingIp  *string                   `json:"azureFloatingIp,omitempty"`
	DhcpSettings     *NetworkDhcpSettingsInput `json:"dhcpSettings,omitempty"`
}

type UpdateNetworkRangePayload struct {
	NetworkRangeId string `json:"networkRangeId,omitempty"`
}

type RemoveNetworkRangePayload struct {
	NetworkRangeId string `json:"networkRangeId,omitempty"`
}

func (c *Client) AddNetworkRange(lanSocketInterfaceId string, input AddNetworkRangeInput) (*AddNetworkRangePayload, error) {
	query := graphQLRequest{
		Query: `
		mutation addNetworkRange($accountId: ID!, $lanSocketInterfaceId: ID!, $input: AddNetworkRangeInput!) {
		site(accountId: $accountId) {
			addNetworkRange(lanSocketInterfaceId: $lanSocketInterfaceId, input: $input) {
			networkRangeId
			}
		}
		}`,
		Variables: map[string]interface{}{
			"accountId":            c.accountId,
			"lanSocketInterfaceId": lanSocketInterfaceId,
			"input":                input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Site struct {
			AddNetworkRange AddNetworkRangePayload
		}
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.AddNetworkRange, nil
}

func (c *Client) UpdateNetworkRange(networkRangeId string, input UpdateNetworkRangeInput) (*UpdateNetworkRangePayload, error) {

	query := graphQLRequest{
		Query: `
		mutation updateNetworkRange($accountId: ID!, $networkRangeId: ID!, $input: UpdateNetworkRangeInput!) {
		site(accountId: $accountId) {
			updateNetworkRange(networkRangeId: $networkRangeId, input: $input) {
			networkRangeId
			}
		}
		}`,
		Variables: map[string]interface{}{
			"accountId":      c.accountId,
			"networkRangeId": networkRangeId,
			"input":          input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Site struct {
			UpdateNetworkRange UpdateNetworkRangePayload
		}
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.UpdateNetworkRange, nil
}

func (c *Client) RemoveNetworkRange(networkRangeId string) (*RemoveNetworkRangePayload, error) {
	query := graphQLRequest{
		Query: `
		mutation removeNetworkRange($accountId: ID!, $networkRangeId: ID!) {
		site(accountId: $accountId) {
			removeNetworkRange(networkRangeId: $networkRangeId) {
			networkRangeId
			}
		}
		}`,
		Variables: map[string]interface{}{
			"accountId":      c.accountId,
			"networkRangeId": networkRangeId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Site struct{ RemoveNetworkRange RemoveNetworkRangePayload }
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.RemoveNetworkRange, nil
}
