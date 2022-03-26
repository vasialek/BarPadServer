package services

import (
	"fmt"
	"time"

	"github.com/vasialek/testserver/helpers"
	"github.com/vasialek/testserver/models"
	"github.com/vasialek/testserver/repositories"
)

// TableService provides managemenet of physical tables
type TableService struct {
	tableRepository  *repositories.TableRepository
	eventsRepository *repositories.TableEventsRepository
}

// NewTableService returns instance of TableService
func NewTableService() *TableService {
	return &TableService{
		tableRepository:  repositories.NewTableRepository(),
		eventsRepository: repositories.NewTableEventsRepository(),
	}
}

// HandleStatusRequest returns response with information about tables
func (ts *TableService) HandleStatusRequest(rq *models.StatusRequest) *models.StatusResponse {
	resp := models.StatusResponse{
		Status: false,
	}
	if helpers.IsValidUniqueID(rq.RequestID) == false {
		resp.Message = "Incorrect request ID."
		return &resp
	}

	tables, err := ts.tableRepository.GetListOfTables()
	if err != nil {
		resp.Message = "Error getting list of Tables for status request."
		return &resp
	}
	resp.Status = true
	resp.Tables = tables
	return &resp
}

// HandleHeartbeatRequest creates/updates information about table
func (ts *TableService) HandleHeartbeatRequest(rq *models.HeartbeatRequest) *models.HeartbeatResponse {
	err := ts.validateHeartbeatRequest(rq)
	resp := models.HeartbeatResponse{
		Status:  false,
		TableID: rq.TableID,
	}
	if err != nil {
		fmt.Printf("Error validating Heartbeat request: %v\n", err)
		resp.Message = "Incorrect Heartbeat request."
		return &resp
	}

	table, err := ts.tableRepository.GetByTableID(rq.TableID)
	if err != nil {
		resp.Message = "Incorrect Table ID."
		return &resp
	}

	if table == nil {
		fmt.Println("Got Heartbeat from new Table, let's create it.")
		table, err = ts.tableRepository.CreateFromTableID(rq.TableID)
		if err != nil {
			resp.Message = "Can't create Table for this Heartbeat."
			return &resp
		}
		resp.Message = "Table is created for this Heartbeat."
	}

	table.LastHeartbeatAt = time.Now()
	fmt.Printf("Set last heartbeat to %s\n", table.LastHeartbeatAt)
	if err = ts.tableRepository.UpdateTable(table.TableID, table); err != nil {
		resp.Message = "Error updating heartbeat information."
		return &resp
	}

	resp.Status = true
	return &resp
}

// HandleWaiterRequest handles request for waiter
func (ts *TableService) HandleWaiterRequest(rq *models.WaiterRequest) *models.WaiterResponse {
	resp := models.WaiterResponse{
		TableID:   rq.TableID,
		RequestID: rq.RequestID,
		Status:    false,
	}

	table, err := ts.tableRepository.GetByTableID(rq.TableID)
	if err != nil {
		resp.Message = "Incorrect Table ID."
		return &resp
	}
	if table == nil {
		resp.Message = "Non-existing Table ID."
		return &resp
	}

	ts.addWaiterEvent(table, rq.RequestID)
	helpers.DumpTable(table)
	if err = ts.tableRepository.UpdateTable(table.TableID, table); err != nil {
		resp.Message = "Error updating events for table."
		return &resp
	}

	// resp.Message = fmt.Sprintf("Waiter request event (ID: %s) is added for Table.", waiterRequestEvent.EventID)
	resp.Message = fmt.Sprintf("Waiter request event is added for Table.")
	resp.Status = true

	return &resp
}

func (ts *TableService) addWaiterEvent(table *models.Table, waiterRequestID string) error {
	return ts.eventsRepository.AddWaiterEvent(table, waiterRequestID)
	// waiterRequestEvent := models.TableEvent{
	// 	EventID:   waiterRequestID,
	// 	CreatedAt: time.Now(),
	// 	IsHandled: false,
	// }

	// table.Events = append(table.Events, waiterRequestEvent)
	// ts.tableRepository.UpdateTable(table.TableID, table)
	// return nil
}

func (ts *TableService) validateHeartbeatRequest(rq *models.HeartbeatRequest) error {
	if helpers.IsValidUniqueID(rq.TableID) == false {
		return fmt.Errorf("bad length of Table ID. Must be 32 received %d", len(rq.TableID))
	}

	return nil
}
