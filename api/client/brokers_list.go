package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/eddyzags/kafkactl/models"
)

type BrokersListResponse struct {
	Brokers []*models.Broker `json:"brokers"`
}

func (c *Client) BrokerList() ([]*models.Broker, error) {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != 201 && resp.Response.StatusCode != 200 {
		return nil, errors.New("Failed to retreive response")
	}

	body, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	response := &BrokersListResponse{}
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.Brokers, nil
}
