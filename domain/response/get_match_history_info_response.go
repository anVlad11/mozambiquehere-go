package response

type GetMatchHistoryInfoResponse struct {
	Registered int `json:"registered"`
	Data       []struct {
		Name     string `json:"name"`
		Platform string `json:"platform"`
	} `json:"data"`
}
