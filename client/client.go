package client

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/anVlad11/mozambiquehere-go/v4/domain/models"
	"github.com/anVlad11/mozambiquehere-go/v4/domain/response"
)

type Client interface {
	//Player stats API
	GetUserByName(platform models.Platform, username string) (response.GetUserResponse, error)
	GetUserByUID(platform models.Platform, uid string) (response.GetUserResponse, error)

	GetUsersByNames(platform models.Platform, usernames []string) (response.GetUserResponse, error)
	GetUsersByUIDs(platform models.Platform, uids []string) (response.GetUserResponse, error)

	//Match history API
	AddMatchHistoryUsers(platform models.Platform, usernames []string) (response.GetMatchHistoryInfoResponse, error)
	DeleteMatchHistoryUsers(platform models.Platform, usernames []string) (response.GetMatchHistoryInfoResponse, error)
	GetMatchHistoryGet(platform models.Platform, usernames []string) (response.GetMatchHistoryGetResponse, error)
	GetMatchHistoryInfo(platform models.Platform, usernames []string) (response.GetMatchHistoryInfoResponse, error)

	//News API
	GetNews(lang models.NewsLanguage) (response.GetNewsResponse, error)

	//Server status API
	GetServerStatus() (response.GetServerStatusResponse, error)

	/**
		TODO: Game data API

		GetGameData() (response.GetGameDataResponse, error)

		or

		GetGameDataAssaultRifles() (response.GetGameDataAssaultRiflesResponse, error)
		GetGameDataAttachments() (response.GetGameDataAttachmentsResponse, error)
		GetGameDataConsumables() (response.GetGameDataConsumablesResponse, error)
		GetGameDataEquipment() (response.GetGameDataEquipment, error)
		GetGameDataGrenades() (response.GetGameDataGrenades, error)
		GetGameDataLegends() (response.GetGameDataLegends, error)
		GetGameDataLightMachineGuns() (response.GetGameDataLightMachineGuns, error)
		GetGameDataPistols() (response.GetGameDataPistols, error)
		GetGameDataShotguns() (response.GetGameDataShotguns, error)
		GetGameDataSniperRifles() (response.GetGameDataSniperRifles, error)
		GetGameDataSubMachineGuns() (response.GetGameDataSubMachineGuns, error)

		hmm :thonk:
	**/

	//TODO: Map rotation API
	//GetMapRotation() (response.GetMapRotationResponse, error)

	//TODO: Shop API
	//GetShop() (response.GetShop, error)
}

type clientImplementation struct {
	token      string
	endpoint   string
	httpClient *http.Client
}

const apiEndpoint = "api.mozambiquehe.re"
const apiVersion = "4"

type path string

//nolint:varcheck,deadcode
const (
	pathBridge      = path("bridge")
	pathNews        = path("news")
	pathGamedata    = path("gamedata")
	pathMapRotation = path("maprotation")
	pathShop        = path("shop")
)

func NewClient(token string, endpoint string, client *http.Client) Client {
	return &clientImplementation{
		token:      token,
		endpoint:   endpoint,
		httpClient: client,
	}
}

func DefaultClient(token string) Client {
	return &clientImplementation{
		token:      token,
		endpoint:   apiEndpoint,
		httpClient: http.DefaultClient,
	}
}

func (c *clientImplementation) doEndpointRequest(method string, path path, params map[string]string) ([]byte, error) {
	var err error

	u := url.URL{}
	u.Scheme = "https"
	u.Host = c.endpoint
	u.Path = string(path)

	if params == nil {
		params = map[string]string{}
	}

	if _, exists := params["auth"]; !exists {
		params["auth"] = c.token
	}

	if _, exists := params["version"]; !exists {
		params["version"] = apiVersion
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}

	u.RawQuery = q.Encode()
	var req *http.Request
	req, err = http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	var body []byte
	body, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func (c *clientImplementation) doRequest(req *http.Request) ([]byte, error) {
	var err error

	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "mozambiquehere-go")
	}

	var res *http.Response
	res, err = c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
