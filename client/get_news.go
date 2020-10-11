package client

import (
	"encoding/json"
	"net/http"

	"github.com/anVlad11/mozambiquehere-go/v4/domain/models"
	"github.com/anVlad11/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetNews(lang models.NewsLanguage) (response.GetNewsResponse, error) {
	var resp response.GetNewsResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathNews, map[string]string{
		"lang": string(lang),
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
