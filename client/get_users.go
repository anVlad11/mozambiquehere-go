package client

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/anVlad11/mozambiquehere-go/v4/domain/models"
	"github.com/anVlad11/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetUsersByNames(platform models.Platform, usernames []string) (response.GetUserResponse, error) {
	var resp response.GetUserResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathBridge, map[string]string{
		"platform": string(platform),
		"player":   strings.Join(usernames, ","),
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

func (c *clientImplementation) GetUsersByUIDs(platform models.Platform, uids []string) (response.GetUserResponse, error) {
	var resp response.GetUserResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathBridge, map[string]string{
		"platform": string(platform),
		"uid":      strings.Join(uids, ","),
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
