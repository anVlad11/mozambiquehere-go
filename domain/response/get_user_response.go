package response

import "github.com/anVlad11/mozambiquehere-go/v4/domain/models"

type GetUsersResponse []GetUserResponse

type GetUserResponse struct {
	Global struct {
		Name                string          `json:"name"`
		UID                 int64           `json:"uid"`
		Platform            models.Platform `json:"platform"`
		Level               int             `json:"level"`
		ToNextLevelPercent  int             `json:"toNextLevelPercent"`
		InternalUpdateCount int             `json:"internalUpdateCount"`
		Bans                struct {
			IsActive         bool   `json:"isActive"`
			RemainingSeconds int    `json:"remainingSeconds"`
			LastBanReason    string `json:"last_banReason"`
		} `json:"bans"`
		Rank struct {
			RankScore    int    `json:"rankScore"`
			RankName     string `json:"rankName"`
			RankDiv      int    `json:"rankDiv"`
			LadderPos    int    `json:"ladderPos"`
			RankImg      string `json:"rankImg"`
			RankedSeason string `json:"rankedSeason"`
		} `json:"rank"`
		Battlepass struct {
			Level   string         `json:"level"`
			History map[string]int `json:"history"`
		} `json:"battlepass"`
	} `json:"global"`
	Realtime struct {
		LobbyState     string        `json:"lobbyState"`
		IsOnline       int           `json:"isOnline"`
		IsInGame       int           `json:"isInGame"`
		CanJoin        int           `json:"canJoin"`
		PartyFull      int           `json:"partyFull"`
		SelectedLegend models.Legend `json:"selectedLegend"`
	} `json:"realtime"`
	Legends struct {
		Selected struct {
			LegendName models.Legend `json:"LegendName"`
			Data       []struct {
				Name  string `json:"name"`
				Value int    `json:"value"`
				Key   string `json:"key"`
			} `json:"data"`
			ImgAssets struct {
				Icon   string `json:"icon"`
				Banner string `json:"banner"`
			} `json:"ImgAssets"`
		} `json:"selected"`
		All map[models.Legend]struct {
			ImgAssets struct {
				Icon   string `json:"icon"`
				Banner string `json:"banner"`
			} `json:"ImgAssets"`
			Data []struct {
				Name  string `json:"name"`
				Value int    `json:"value"`
				Key   string `json:"key"`
			} `json:"data,omitempty"`
		} `json:"all"`
	} `json:"legends"`
	MozambiquehereInternal struct {
		IsNewToDB     bool   `json:"isNewToDB"`
		ClaimedBy     string `json:"claimedBy"`
		APIAccessType string `json:"APIAccessType"`
		ClusterID     string `json:"ClusterID"`
		RateLimit     struct {
			MaxPerSecond int    `json:"max_per_second"`
			CurrentReq   string `json:"current_req"`
		} `json:"rate_limit"`
	} `json:"mozambiquehere_internal"`
	Total struct { //map[string]struct{}?
		Kills struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"kills"`
		Headshots struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"headshots"`
		Damage struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"damage"`
		DetectedBreaches struct {
			Name  interface{} `json:"name"`
			Value int         `json:"value"`
		} `json:"detected_breaches"`
		ShieldsCharged struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"shields_charged"`
		OrdnanceDestroyed struct {
			Name  interface{} `json:"name"`
			Value int         `json:"value"`
		} `json:"ordnance_destroyed"`
		Kd struct {
			Value int    `json:"value"`
			Name  string `json:"name"`
		} `json:"kd"`
	} `json:"total"`
}
