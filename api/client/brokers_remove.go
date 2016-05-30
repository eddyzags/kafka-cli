package client

import (
	"errors"
	"net/http"
	"net/url"
	"path"

	"github.com/eddyzags/kafkactl/models"
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

	response := &BrokerRemoveResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Ids, nil
}
