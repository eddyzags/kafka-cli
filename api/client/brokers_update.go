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

type BrokerUpdateResponse struct {
	Brokers []*models.Broker `json:"brokers"`
}

func (c *Client) BrokerUpdate(in map[string]string) ([]*models.Broker, error) {
	values := url.Values{}

	for k, v := range in {
		if v == "" {
			continue
		}

		values.Set(k, v)
	}

	url := c.URL + path.Join("/api", "broker", "add") + "?" + values.Encode()

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
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

	response := &BrokerUpdateResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Brokers, nil
}
