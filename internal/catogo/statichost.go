package catogo

import "encoding/json"

type AddStaticHostPayload struct {
	HostId string `json:"hostId,omitempty"`
}

type UpdateStaticHostPayload struct {
	HostId string `json:"hostId"`
}

type RemoveStaticHostPayload struct {
	HostId string `json:"hostId,omitempty"`
}

type AddStaticHostInput struct {
	Name       string  `json:"name,omitempty"`
	Ip         string  `json:"ip,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
}

type UpdateStaticHostInput struct {
	Name       string  `json:"name,omitempty"`
	Ip         string  `json:"ip,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
}

func (c *Client) AddStaticHost(siteId string, input AddStaticHostInput) (*AddStaticHostPayload, error) {
	query := graphQLRequest{
		Query: `
		mutation addStaticHost($accountId:ID!,$siteId: ID!, $input: AddStaticHostInput!) {
		site(accountId:$accountId){
			addStaticHost(siteId:$siteId, input:$input){
			hostId
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
			AddStaticHost AddStaticHostPayload
		}
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.AddStaticHost, nil
}

func (c *Client) UpdateStaticHost(siteId string, hostId string, input UpdateStaticHostInput) (*UpdateStaticHostPayload, error) {
	query := graphQLRequest{
		Query: `
		mutation updateStaticHost($accountId:ID!,$hostId: ID!, $input: UpdateStaticHostInput!) {
		site(accountId:$accountId){
			updateStaticHost(hostId:$hostId, input:$input){
			hostId
			}
		}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"siteId":    siteId,
			"hostId":    hostId,
			"input":     input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Site struct {
			UpdateStaticHost UpdateStaticHostPayload
		}
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.UpdateStaticHost, nil
}

func (c *Client) RemoveStaticHost(siteId string, hostId string) (*RemoveStaticHostPayload, error) {
	query := graphQLRequest{
		Query: `
		mutation removeStaticHost($accountId:ID!,$hostId: ID!) {
		site(accountId:$accountId){
			removeStaticHost(hostId:$hostId){
			hostId
			}
		}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"siteId":    siteId,
			"hostId":    hostId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Site struct{ RemoveStaticHost RemoveStaticHostPayload }
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Site.RemoveStaticHost, nil
}
