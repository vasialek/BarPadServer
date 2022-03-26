package helpers

import (
	"fmt"

	"github.com/vasialek/testserver/models"
)

// DumpTables outputs list of tables
func DumpTables(tables []models.Table) {
	for _, table := range tables {
		DumpTable(&table)
	}
}

// DumpTable outputs information about table
func DumpTable(table *models.Table) {
	fmt.Printf("Table ID:           %s\n", table.TableID)
	fmt.Printf("Last heartbeat at:  %s\n", table.LastHeartbeatAt)
	for _, event := range table.Events {
		DumpTableEvent(&event)
	}
}

// DumpTableEvent outputs information about table event
func DumpTableEvent(event *models.TableEvent) {
	fmt.Printf("  Event ID:         %s\n", event.EventID)
	fmt.Printf("  Created at:       %s\n", event.CreatedAt)
	fmt.Printf("  Is handled:       %t\n", event.IsHandled)
}
