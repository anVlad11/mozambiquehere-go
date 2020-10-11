package client

import (
	"encoding/json"
	"net/http"

	"github.com/anVlad11/mozambiquehere-go/v4/domain/models"
	"github.com/anVlad11/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetUserByName(platform models.Platform, username string) (response.GetUserResponse, error) {
	var resp response.GetUserResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathBridge, map[string]string{
		"platform": string(platform),
		"player":   username,
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

func (c *clientImplementation) GetUserByUID(platform models.Platform, uid string) (response.GetUserResponse, error) {
	var resp response.GetUserResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathBridge, map[string]string{
		"platform": string(platform),
		"uid":      uid,
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
