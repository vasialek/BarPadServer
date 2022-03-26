package repositories

import (
	"fmt"

	"github.com/vasialek/testserver/models"
)

// NewTableRepository returns instance of TableRepository
func NewTableRepository() *TableRepository {
	tr := TableRepository{}

	// tr.tables = make(0, models.Table)
	tr.tables = append(tr.tables, models.Table{TableID: "FakeTableId"})

	return &tr
}

// TableRepository provides access to tables
type TableRepository struct {
	tables []models.Table
}

// GetListOfTables returns list of all tables
func (tr *TableRepository) GetListOfTables() ([]models.Table, error) {
	return tr.tables, nil
}

// GetByTableID returns Table or error if not found
func (tr *TableRepository) GetByTableID(tableID string) (*models.Table, error) {
	fmt.Printf("Table ID to search: `%s`\n", tableID)
	for _, table := range tr.tables {
		fmt.Printf("  table ID: `%s`\n", table.TableID)
		if table.TableID == tableID {
			return &table, nil
		}
	}
	return nil, nil
}

// CreateFromTableID creates Table from ID
func (tr *TableRepository) CreateFromTableID(tableID string) (*models.Table, error) {
	table := models.Table{
		TableID: tableID,
	}
	tr.tables = append(tr.tables, table)

	return &table, nil
}

// UpdateTable updates existing table
func (tr *TableRepository) UpdateTable(tableID string, table *models.Table) error {
	for index, t := range tr.tables {
		if t.TableID == tableID {
			tr.tables[index].LastHeartbeatAt = table.LastHeartbeatAt
			tr.tables[index].Events = table.Events
			return nil
		}
	}
	return nil
}
