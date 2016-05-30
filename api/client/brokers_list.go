package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/eddyzags/kafkactl/models"
)

type BrokersListResponse struct {
	Brokers []*models.Broker `json:"brokers"`
}

func (c *Client) BrokerList() ([]*models.Broker, error) {
	url := c.URL + path.Join("/api", "broker", "list")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return nil, errors.New("Failed to retreive response: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &BrokersListResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Brokers, nil
}
