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

type BrokerStopResponse struct {
	Brokers []*models.Broker `json:"brokers"`
	Status  string           `json:"status"`
}

func (c *Client) BrokerStop(expr string) ([]*models.Broker, error) {
	values := url.Values{}

	values.Set("broker", expr)

	url := c.URL + path.Join("/api", "broker", "stop") + "?" + values.Encode()

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
		return nil, errors.New("Failed to retreive response: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &BrokerStopResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.Status != "stopped" {
		return nil, errors.New("An internal error occured: status=" + response.Status)
	}

	return response.Brokers, nil
}
