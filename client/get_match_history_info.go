package client

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/anVlad11/mozambiquehere-go/v4/domain/models"
	"github.com/anVlad11/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetMatchHistoryInfo(platform models.Platform, usernames []string) (response.GetMatchHistoryInfoResponse, error) {
	var resp response.GetMatchHistoryInfoResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathBridge, map[string]string{
		"platform": string(platform),
		"player":   strings.Join(usernames, ","),
		"action":   "info",
		"history":  "1",
	})

	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}
