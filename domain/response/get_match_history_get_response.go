package response

import (
	"github.com/anVlad11/mozambiquehere-go/v4/domain/errors"
	"github.com/anVlad11/mozambiquehere-go/v4/domain/models"
)

type GetMatchHistoryGetResponseItem struct {
	UID          string           `json:"uid"`
	Player       string           `json:"player"`
	Timestamp    string           `json:"timestamp"`
	EventType    models.EventType `json:"eventType"`
	XpProgress   string           `json:"xpProgress,omitempty"`
	GameLength   string           `json:"gameLength,omitempty"`
	LegendPlayed models.Legend    `json:"legendPlayed,omitempty"`
	Event        interface{}      `json:"event"`
}

type GetMatchHistoryGetResponse []GetMatchHistoryGetResponseItem

//EventType == Game
type EventTypeGameData []EventTypeGameDataItem

type EventTypeGameDataItem struct {
	Value int     `json:"value"`
	Key   *string `json:"key"` //sql.NullString maybe? Too complicated?
	Name  *string `json:"name"`
}

func (r *GetMatchHistoryGetResponseItem) GetEventTypeGameData() (*EventTypeGameData, error) {
	if r.EventType != models.EventTypeGame {
		return nil, errors.GetUnexpectedEventTypeError(models.EventTypeLevel, r.EventType)
	}

	eventDataRaw := r.Event.([]map[string]interface{})
	eventData := make(EventTypeGameData, len(eventDataRaw))
	for _, v := range eventDataRaw {
		eventData = append(eventData, EventTypeGameDataItem{
			Value: v["value"].(int),
			Key:   v["key"].(*string),
			Name:  v["name"].(*string),
		})
	}

	return &eventData, nil
}

//EventType == Level
type EventTypeLevelData struct {
	NewLevel string `json:"newLevel"`
}

func (r *GetMatchHistoryGetResponseItem) GetEventTypeLevelData() (*EventTypeLevelData, error) {
	if r.EventType != models.EventTypeLevel {
		return nil, errors.GetUnexpectedEventTypeError(models.EventTypeLevel, r.EventType)
	}

	eventDataRaw := r.Event.(map[string]string)
	eventData := EventTypeLevelData{
		NewLevel: eventDataRaw["newLevel"],
	}

	return &eventData, nil
}

//EventType == Session
type EventTypeActionData struct {
	Action          models.EventSessionActionType `json:"action"`
	SessionDuration *string                       `json:"sessionDuration,omitempty"`
}

func (r *GetMatchHistoryGetResponseItem) GetEventTypeActionData() (*EventTypeActionData, error) {
	if r.EventType != models.EventTypeSession {
		return nil, errors.GetUnexpectedEventTypeError(models.EventTypeSession, r.EventType)
	}

	eventDataRaw := r.Event.(map[string]string)
	eventData := EventTypeActionData{
		Action: models.EventSessionActionType(eventDataRaw["action"]),
	}

	if _, exists := models.ValidSessionTypes[eventData.Action]; !exists {
		return nil, errors.GetInvalidSessionActionTypeError(eventData.Action)

	}

	if duration, exists := eventDataRaw["session_duration"]; exists {
		eventData.SessionDuration = &duration
	}

	return &eventData, nil
}
