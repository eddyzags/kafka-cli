package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type BrokerStartResponse struct {
	Success bool   `json:"success"`
	Ids     string `json:"ids"`
}

func (c *Client) BrokerStart(expr, timeout string) (string, error) {
	values := url.Values{}
	if timeout != "" {
		values.Set("timeout", timeout)
	}
	values.Set("broker", expr)

	url := c.URL + path.Join("/api", "broker", "start") + "?" + values.Encode()

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return "", errors.New("Failed to retreived response: ", fmt.Sprintf("%d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	response := &BrokerStartResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if response.Success != true {
		return "", errors.New("Internal error occured: starting failed")
	}

	return response.Ids, nil
}
