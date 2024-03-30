package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Temporary Cato Client, to be externalised

type Client struct {
	httpclient *http.Client
	token      string
	baseurl    string
}

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// to be improve
type Response struct {
	Data   interface{}   `json:"data,omitempty"`
	Errors []interface{} `json:"errors,omitempty"`
}

func CatoClient(baseurl string, token string) *Client {

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	return &Client{
		httpclient: client,
		baseurl:    baseurl,
		token:      token,
	}
}

func (c *Client) do(reqBody graphQLRequest) ([]byte, error) {

	var respBody Response

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseurl, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", c.token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

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

	return body, nil

}
