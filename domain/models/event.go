package models

type EventType string

const (
	EventTypeGame    = EventType("Game")
	EventTypeLevel   = EventType("Level")
	EventTypeSession = EventType("Session")
)

type EventSessionActionType string

const (
	EventSessionActionTypeJoin  = EventSessionActionType("join")
	EventSessionActionTypeLeave = EventSessionActionType("leave")
)

var ValidSessionTypes = map[EventSessionActionType]interface{}{
	EventSessionActionTypeJoin:  nil,
	EventSessionActionTypeLeave: nil,
}
