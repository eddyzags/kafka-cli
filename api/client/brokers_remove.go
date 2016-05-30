package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type BrokerRemoveResponse struct {
	Ids string `json:"ids"`
}

func (c *Client) BrokerRemove(expr string) (string, error) {
	values := url.Values{}

	values.Set("broker", expr)

	url := c.URL + path.Join("/api", "broker", "remove") + "?" + values.Encode()

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return "", errors.New("Failed to retreived response: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	response := &BrokerRemoveResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.Ids, nil
}
