package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/SKF/go-utility/log"
	"github.com/SKF/go-utility/uuid"
	iotgrpcapi "github.com/SKF/proto/v2/iot"
)

var (
	ErrTaskInvalidDueDate              = errors.New("Invalid due date, not unix timestamp in ms")
	ErrTaskInvalidUserID               = errors.New("Required field 'userId' is invalid")
	ErrTaskInvalidExternalID           = errors.New("Required field 'externalId' is invalid")
	ErrTaskFLIDsAndNodeIdsEmpty        = errors.New("'FunctionalLocationIds' or 'NodeIds' must contain at least one value")
	ErrTaskFLIDsAndNodeIdsFilledIn     = errors.New("'FunctionalLocationIds' and 'NodeIds' cannot contain values")
	ErrTaskFLAndNodeIdsNilEmpty        = errors.New("Both 'functionalLocationIds' and 'NodeIDs' cannot be nil/empty")
	ErrTaskInvalidHierarchy            = errors.New("Required field 'hierarchyId' is invalid")
	ErrTaskInvalidID                   = errors.New("Required field 'id' is invalid")
	ErrTaskInvalidkTaskStatus          = fmt.Errorf("Field 'taskstatus' is invalid, valid statuses: %v", validTaskStatuses)
	ErrTaskInvalidkTaskStatusUpdatedAt = errors.New("Invalid taskStatusUpdateAt, not unix timestamp in ms")
	ErrTaskNodeTypeEmpty               = errors.New("Nodetype cannot be empty when 'NodeIds' is used")
	year                               = time.Hour * 24 * 365
)

type Task struct {
	ID                    uuid.UUID   `json:"id"`
	Name                  string      `json:"name"`
	HierarchyID           uuid.UUID   `json:"hierarchyId"`
	DueDate               int64       `json:"dueDate"`
	IsCompleted           bool        `json:"isCompleted"`
	UserID                uuid.UUID   `json:"userId"`
	FunctionalLocationIDs []uuid.UUID `json:"functionalLocationIds"`
	ExternalID            uuid.UUID   `json:"externalId"`
	TaskStatus            string      `json:"taskStatus"`
	TaskStatusUpdatedAt   int64       `json:"taskStatusUpdatedAt"`
	LongID                int64       `json:"longId"`
	Nodes                 []Node      `json:"nodeIds"`
}

type Node struct {
	NodeID   uuid.UUID `json:"nodeId"`
	NodeType string    `json:"nodeType"`
}

var validTaskStatuses = map[string]bool{
	"NOT_SENT":            true,
	"SENT":                true,
	"COMPLETED":           true,
	"RECEIVED":            true,
	"IN_PROGRESS":         true,
	"MISSED":              true,
	"PARTIALLY_COLLECTED": true,
}

func (task Task) Validate() (err error) {
	if !task.ID.IsValid() {
		return ErrTaskInvalidID
	}
	if !task.HierarchyID.IsValid() {
		return ErrTaskInvalidHierarchy
	}
	if (task.FunctionalLocationIDs == nil && task.Nodes == nil) ||
		(len(task.FunctionalLocationIDs) == 0 && len(task.Nodes) == 0) ||
		(task.FunctionalLocationIDs == nil && len(task.Nodes) == 0) ||
		(len(task.FunctionalLocationIDs) == 0 && task.Nodes == nil) {
		return ErrTaskFLAndNodeIdsNilEmpty
	}
	if len(task.FunctionalLocationIDs) == 0 && len(task.Nodes) == 0 {
		return ErrTaskFLIDsAndNodeIdsEmpty
	}
	if task.Nodes != nil && len(task.Nodes) > 0 {
		for _, node := range task.Nodes {
			if len(node.NodeType) == 0 {
				return ErrTaskNodeTypeEmpty
			}
		}

	}
	if len(task.FunctionalLocationIDs) > 0 && len(task.Nodes) > 0 {
		return ErrTaskFLIDsAndNodeIdsFilledIn
	}
	if !task.ExternalID.IsValid() {
		return ErrTaskInvalidExternalID
	}
	if !task.UserID.IsValid() {
		return ErrTaskInvalidUserID
	}
	if !ValidTaskStatus(task.TaskStatus) {
		return ErrTaskInvalidkTaskStatus
	}

	now := time.Now()
	if task.DueDate < unixTimestampInMs(now.Add(year*-10)) || task.DueDate > unixTimestampInMs(now.Add(year*10)) {
		return ErrTaskInvalidDueDate
	}
	if !validTaskStatusUpdatedAt(task.TaskStatusUpdatedAt) {
		return ErrTaskInvalidkTaskStatusUpdatedAt
	}

	for index, id := range task.FunctionalLocationIDs {
		if !id.IsValid() {
			return fmt.Errorf("Invalid 'functionalLocationIds' at index %+v for '%+v'", index, id)
		}
	}

	return
}

func unixTimestampInMs(t time.Time) int64 {
	return t.UnixNano() / 1e+6
}

func ValidTaskStatus(status string) bool {
	return validTaskStatuses[status]
}

func validTaskStatusUpdatedAt(updatedAt int64) bool {
	return updatedAt > unixTimestampInMs(time.Now().Add(year*-10))
}

func ValidateSetTaskStatusInput(input *iotgrpcapi.SetTaskStatusInput) error {
	if !uuid.UUID(input.TaskId).IsValid() {
		return ErrTaskInvalidID
	}
	if !uuid.UUID(input.UserId).IsValid() {
		return ErrTaskInvalidUserID
	}
	if !ValidTaskStatus(input.Status.String()) {
		return ErrTaskInvalidkTaskStatus
	}
	if !validTaskStatusUpdatedAt(input.UpdatedAt) {
		return fmt.Errorf("%+v [%d]", ErrTaskInvalidkTaskStatusUpdatedAt, input.UpdatedAt)
	}
	return nil
}

func (task *Task) FromGRPC(td *iotgrpcapi.TaskDescription) {
	task.ID = uuid.UUID(td.TaskId)
	task.Name = td.TaskName
	task.HierarchyID = uuid.UUID(td.HierarchyId)
	task.DueDate = td.DueDateTimestamp
	task.IsCompleted = td.IsCompleted
	task.UserID = uuid.UUID(td.UserId)
	if td.FunctionalLocationIds != nil {
		task.FunctionalLocationIDs = stringToUUIDArray(td.FunctionalLocationIds.IdArr)
	}
	task.LongID = td.LongId
}

func (task Task) ToGRPC() iotgrpcapi.TaskDescription {
	var nodes []*iotgrpcapi.Node
	for _, n := range task.Nodes {
		node := iotgrpcapi.Node{
			NodeId:   n.NodeID.String(),
			NodeType: n.NodeType,
		}
		nodes = append(nodes, &node)
	}
	return iotgrpcapi.TaskDescription{
		TaskId:           task.ID.String(),
		TaskName:         task.Name,
		HierarchyId:      task.HierarchyID.String(),
		DueDateTimestamp: task.DueDate,
		IsCompleted:      task.IsCompleted,
		UserId:           task.UserID.String(),
		FunctionalLocationIds: &iotgrpcapi.FunctionalLocationIds{
			IdArr: uuidToStringArray(task.FunctionalLocationIDs),
		},
		Status:          parseGRPCTaskStatus(task.TaskStatus, task.ID),
		StatusUpdatedAt: task.TaskStatusUpdatedAt,
		ExternalTaskId:  task.ExternalID.String(),
		LongId:          task.LongID,
		Nodes:           nodes,
	}
}

func parseGRPCTaskStatus(enumString string, taskID uuid.UUID) iotgrpcapi.TaskStatus {
	value, ok := iotgrpcapi.TaskStatus_value[enumString]
	if !ok {
		log.WithField("enumString", enumString).
			WithField("taskId", taskID).
			Infof("Invalid enumString from datalayer")
		return iotgrpcapi.TaskStatus_NOT_SENT
	}

	return iotgrpcapi.TaskStatus(value)
}

func (task *Task) FromGRPCInitial(request *iotgrpcapi.InitialTaskDescription) (err error) {
	status := request.Status
	if request.Status == iotgrpcapi.TaskStatus_NOT_SENT {
		status = iotgrpcapi.TaskStatus_SENT
	}

	task.Name = request.TaskName
	task.HierarchyID = uuid.UUID(request.HierarchyId)
	task.DueDate = request.DueDateTimestamp
	task.IsCompleted = false
	task.UserID = uuid.UUID(request.UserId)
	if (request.Nodes != nil) && (len(request.Nodes) > 0) {
		//Check that all nodes have a nodetype
		for _, node := range request.Nodes {
			if len(node.NodeType) == 0 {
				err = fmt.Errorf("Nodetype is missing")
			}
		}

		//Add nodes
		for _, node := range request.Nodes {
			var tempNode Node
			tempNode.NodeID = uuid.UUID(node.NodeId)
			tempNode.NodeType = node.NodeType
			task.Nodes = append(task.Nodes, tempNode)
		}
	}

	if request.FunctionalLocationIds != nil {
		task.FunctionalLocationIDs = stringToUUIDArray(request.FunctionalLocationIds.IdArr)
	}
	task.ExternalID = uuid.UUID(request.ExternalTaskId)
	task.TaskStatus = status.String()
	task.TaskStatusUpdatedAt = unixTimestampInMs(time.Now())
	return
}

func (task Task) ToGRPCInitial() iotgrpcapi.InitialTaskDescription {
	return iotgrpcapi.InitialTaskDescription{
		TaskName:         task.Name,
		HierarchyId:      task.HierarchyID.String(),
		DueDateTimestamp: task.DueDate,
		UserId:           task.UserID.String(),
		ExternalTaskId:   task.ExternalID.String(),
		FunctionalLocationIds: &iotgrpcapi.FunctionalLocationIds{
			IdArr: uuidToStringArray(task.FunctionalLocationIDs),
		},
	}
}
