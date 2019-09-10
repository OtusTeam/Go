package services

import (
	"context"
	"github.com/otusteam/go/cleancalendar/internal/domain/interfaces"
	"github.com/otusteam/go/cleancalendar/internal/domain/models"
	"github.com/satori/go.uuid"
	"time"
)

type EventService struct {
	EventStorage interfaces.EventStorage
}

func (es *EventService) CreateEvent(ctx context.Context, owner, title, text string, startTime *time.Time, endTime *time.Time) (*models.Event, error) {
	// TODO: persistence, validation
	event := &models.Event{
		Id:        uuid.NewV4(),
		Owner:     owner,
		Title:     title,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err := es.EventStorage.SaveEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}
