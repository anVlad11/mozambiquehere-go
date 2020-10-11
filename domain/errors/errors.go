package errors

import (
	"fmt"

	"github.com/anVlad11/mozambiquehere-go/v4/domain/models"
)

func GetUnexpectedEventTypeError(expected models.EventType, got models.EventType) error {
	return fmt.Errorf("unexpected event type: expected '%s', got '%s'", expected, got)
}

func GetInvalidSessionActionTypeError(got models.EventSessionActionType) error {
	return fmt.Errorf("invalid session event action type: '%s'", got)
}
