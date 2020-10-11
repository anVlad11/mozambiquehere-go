package client

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/anVlad11/mozambiquehere-go/v4/domain/models"
	"github.com/anVlad11/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetMatchHistoryGet(platform models.Platform, usernames []string) (response.GetMatchHistoryGetResponse, error) {
	var resp response.GetMatchHistoryGetResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathBridge, map[string]string{
		"platform": string(platform),
		"player":   strings.Join(usernames, ","),
		"action":   "get",
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
