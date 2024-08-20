package catogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpclient *http.Client
	token      string
	baseurl    string
	accountId  string
	tfversion  string
}

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type Response struct {
	Data struct {
		Admin           interface{} `json:"admin,omitempty"`
		Admins          interface{} `json:"admins,omitempty"`
		AccountRoles    interface{} `json:"accountRoles,omitempty"`
		AccountSnapshot interface{} `json:"accountSnapshot,omitempty"`
		EntityLookup    interface{} `json:"entityLookup,omitempty"`
		Site            interface{} `json:"site,omitempty"`
		Policy          interface{} `json:"policy,omitempty"`
	} `json:"data,omitempty"`
	Errors []interface{} `json:"errors,omitempty"`
}

func CatoClient(baseurl string, token string, accountId string, tfversion string) *Client {

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	return &Client{
		httpclient: client,
		baseurl:    baseurl,
		token:      token,
		accountId:  accountId,
		tfversion:  tfversion,
	}
}

func (c *Client) do(reqBody graphQLRequest) ([]byte, error) {
	var respBody Response

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] API Request: %s\n", string(jsonReqBody))

	req, err := http.NewRequest("POST", c.baseurl, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "cato-terraform-"+c.tfversion)

	res, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] API Response: %s\n", string(body))
	if res.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &respBody)
		if err != nil {
			return nil, err
		}
		if respBody.Errors != nil {
			json_error, _ := json.Marshal(respBody.Errors)
			return nil, fmt.Errorf(string(json_error))
		}
	} else {
		return nil, fmt.Errorf("unknown error - " + http.StatusText(res.StatusCode))
	}

	return getByteData(body)
}

// function that extract the value of "data" key in JSON response from API
func getByteData(body []byte) ([]byte, error) {

	response := Response{}
	byteData := []byte{}

	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	byteData, err = json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	return byteData, nil
}
