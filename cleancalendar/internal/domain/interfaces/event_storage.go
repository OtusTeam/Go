package interfaces

import (
	"context"
	"github.com/otusteam/go/cleancalendar/internal/domain/models"
	"time"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event *models.Event) error
	GetEventById(ctx context.Context, id string) (*models.Event, error)
	GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) []*models.Event
}
