package repositories

import (
	"fmt"
	"time"

	"github.com/vasialek/testserver/models"
)

// TableEventsRepository provides access to table events
type TableEventsRepository struct {
	tableEvents []models.TableEvent
}

// NewTableEventsRepository returns instance of TableEventsRepository
func NewTableEventsRepository() *TableEventsRepository {
	return &TableEventsRepository{}
}

func (ter *TableEventsRepository) AddWaiterEvent(table *models.Table, eventID string) error {
	index := ter.getIndex(table.Events, eventID)
	if index == -1 {
		fmt.Println("adding new waiter event...")
		table.Events = append(table.Events, models.TableEvent{
			EventID:   eventID,
			CreatedAt: time.Now(),
		})
		return nil
	}
	fmt.Println("such waiter event already exists...")
	if table.Events[index].IsHandled {
		table.Events[index].IsHandled = false
		return nil
	}
	return nil
}

func (ter *TableEventsRepository) getIndex(events []models.TableEvent, eventID string) int {
	for index, event := range events {
		if event.EventID == eventID {
			return index
		}
	}

	return -1
}
