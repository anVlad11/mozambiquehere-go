package client

import (
	"encoding/json"
	"net/http"

	"github.com/anVlad11/mozambiquehere-go/v4/domain/response"
)

const serverStatusEndpoint = "https://apexlegendsstatus.com/servers.json"

func (c *clientImplementation) GetServerStatus() (response.GetServerStatusResponse, error) {
	var resp response.GetServerStatusResponse
	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodGet, serverStatusEndpoint, nil)
	if err != nil {
		return resp, err
	}

	var body []byte
	body, err = c.doRequest(req)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}
