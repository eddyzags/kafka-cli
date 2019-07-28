package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/eddyzags/kafkactl/models"
)

type BrokerStartResponse struct {
	Brokers []*models.Broker `json:"brokers"`
	Status  string           `json:"status"`
}

func (c *Client) BrokerStart(expr, timeout string) ([]*models.Broker, error) {
	values := url.Values{}
	if timeout != "" {
		values.Set("timeout", timeout)
	}
	values.Set("broker", expr)

	url := c.URL + path.Join("/api", "broker", "start") + "?" + values.Encode()

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return nil, errors.New("Failed to retreived response: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &BrokerStartResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.Status != "started" {
		return nil, errors.New("Internal error occured: starting failed")
	}

	return response.Brokers, nil
}
