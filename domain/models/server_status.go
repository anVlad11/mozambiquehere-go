package models

import "time"

type ServiceName string

const (
	ServiceNameOriginLogin            = ServiceName("Origin_login")
	ServiceNameEANovaFusion           = ServiceName("EA_novafusion")
	ServiceNameEAAccounts             = ServiceName("EA_accounts")
	ServiceNameApexOauthPC            = ServiceName("ApexOauth_PC")
	ServiceNameApexOauthPS4           = ServiceName("ApexOauth_PS4")
	ServiceNameApexOauthX1            = ServiceName("ApexOauth_X1")
	ServiceNameMozambiqueHereStatsAPI = ServiceName("Mozambiquehere_StatsAPI")
)

type Service map[ServerName]ServerStatus

type ServerName string

const (
	ServerNameAsia         = ServerName("Asia")
	ServerNameEUEast       = ServerName("EU-East")
	ServerNameEUWest       = ServerName("EU-West")
	ServerNameSouthAmerica = ServerName("SouthAmerica")
	ServerNameUSCentral    = ServerName("US-Central")
	ServerNameUSEast       = ServerName("US-East")
	ServerNameUSWest       = ServerName("US-West")
)

type ServerStatus struct {
	Status         ServerStatusType `json:"Status"`
	HTTPCode       int              `json:"HTTPCode"`
	ResponseTime   int              `json:"ResponseTime"`
	QueryTimestamp int64            `json:"QueryTimestamp"`
}

type ServerStatusType string

const (
	ServerStatusUp     = ServerStatusType("UP")
	ServerStatusDown   = ServerStatusType("DOWN")
	ServerStatusNoData = ServerStatusType("NO DATA")
)

func (s *ServerStatus) GetQueryTime() time.Time {
	return time.Unix(s.QueryTimestamp, 0)
}
