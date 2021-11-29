package iot

import (
	"context"
	"fmt"
	"time"

	"github.com/SKF/proto/v2/common"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	iot_grpcapi "github.com/SKF/proto/v2/iot"
)

func (c *Client) CreateTask(task iot_grpcapi.InitialTaskDescription) (taskID string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.CreateTaskWithContext(ctx, task)
}

func (c *Client) CreateTaskWithContext(ctx context.Context, task iot_grpcapi.InitialTaskDescription) (taskID string, err error) {
	output, err := c.api.CreateTask(ctx, &task)
	if output != nil {
		taskID = output.Value
	}
	return
}

func (c *Client) DeleteTask(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.DeleteTaskWithContext(ctx, userID, taskID)
}
func (c *Client) DeleteTaskWithContext(ctx context.Context, userID, taskID string) (err error) {
	input := iot_grpcapi.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.DeleteTask(ctx, &input)
	return
}

func (c *Client) GetAllTasks(userID string) (out []iot_grpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetAllTasksWithContext(ctx, userID)
}
func (c *Client) GetAllTasksWithContext(ctx context.Context, userID string) (out []iot_grpcapi.TaskDescription, err error) {
	input := common.PrimitiveString{Value: userID}
	output, err := c.api.GetAllTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *Client) GetUncompletedTasks(userID string) (out []iot_grpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetUncompletedTasksWithContext(ctx, userID)
}
func (c *Client) GetUncompletedTasksWithContext(ctx context.Context, userID string) (out []iot_grpcapi.TaskDescription, err error) {
	input := common.PrimitiveString{Value: userID}
	output, err := c.api.GetUncompletedTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *Client) GetUncompletedTasksByHierarchy(nodeID string) (out []iot_grpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetUncompletedTasksByHierarchyWithContext(ctx, nodeID)
}
func (c *Client) GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) (out []iot_grpcapi.TaskDescription, err error) {
	input := common.PrimitiveString{Value: nodeID}
	tasks, err := c.api.GetUncompletedTasksByHierarchy(ctx, &input)
	if tasks != nil {
		for _, desc := range tasks.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *Client) GetActiveTasks(userID string) (out []iot_grpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetActiveTasksWithContext(ctx, userID)
}
func (c *Client) GetActiveTasksWithContext(ctx context.Context, userID string) (out []iot_grpcapi.TaskDescription, err error) {
	input := common.PrimitiveString{Value: userID}
	output, err := c.api.GetActiveTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

// SetTaskStatus will set the status of the task.
//   SetTaskStatusInput:
//     task_id uuid (required)
//     user_id uuid  (required)
//     status TaskStatus (required)
//     - allowed values: NOT_SENT, SENT, RECEIVED, IN_PROGRESS, COMPLETED
//     updated_at int (optional)
//     - UNIX timestamp in ms
//
func (c *Client) SetTaskStatus(input iot_grpcapi.SetTaskStatusInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetTaskStatusWithContext(ctx, input)
}
func (c *Client) SetTaskStatusWithContext(ctx context.Context, input iot_grpcapi.SetTaskStatusInput) (err error) {
	_, err = c.api.SetTaskStatus(ctx, &input)
	return
}

//IngestNodesData ingest nodes, this operation can success or fail. Fail means that no data is stored.
func (c *Client) IngestNodesData(input iot_grpcapi.IngestNodesDataInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.IngestNodesDataWithContext(ctx, input)
}

//IngestNodesDataWithContext ingest nodes with context, this operation can success or fail. Fail means that no data is stored.
func (c *Client) IngestNodesDataWithContext(ctx context.Context, input iot_grpcapi.IngestNodesDataInput) error {
	if len(input.Nodes) == 0 {
		return fmt.Errorf("IngestNodesData missing nodes to insert")
	}

	res, err := c.api.IngestNodesData(ctx, &input)
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("IngestNodesData failed")
	}
	return nil
}

func (c *Client) IngestNodeData(input iot_grpcapi.IngestNodeDataInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.IngestNodeDataWithContext(ctx, input)
}
func (c *Client) IngestNodeDataWithContext(ctx context.Context, input iot_grpcapi.IngestNodeDataInput) (err error) {
	_, err = c.api.IngestNodeData(ctx, &input)
	return
}

func (c *Client) GetTaskByUUID(input string) (output *iot_grpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetTaskByUUIDWithContext(ctx, input)
}
func (c *Client) GetTaskByUUIDWithContext(ctx context.Context, input string) (output *iot_grpcapi.TaskDescription, err error) {
	grpcInput := iot_grpcapi.GetTaskByUUIDInput{
		TaskId: input,
	}
	response, err := c.api.GetTaskByUUID(ctx, &grpcInput)
	if response != nil {
		output = response.Task
	}
	return
}

func (c *Client) GetTaskByLongId(input int64) (output *iot_grpcapi.TaskDescription, err error) { // nolint: golint
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetTaskByLongIdWithContext(ctx, input)
}
func (c *Client) GetTaskByLongIdWithContext(ctx context.Context, input int64) (output *iot_grpcapi.TaskDescription, err error) { // nolint: golint
	grpcInput := iot_grpcapi.GetTaskByLongIdInput{
		TaskId: input,
	}
	response, err := c.api.GetTaskByLongId(ctx, &grpcInput)
	if response != nil {
		output = response.Task
	}
	return
}

func (c *Client) GetLatestNodeData(input iot_grpcapi.GetLatestNodeDataInput) (*iot_grpcapi.NodeData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetLatestNodeDataWithContext(ctx, input)
}

func (c *Client) GetLatestNodeDataWithContext(ctx context.Context, input iot_grpcapi.GetLatestNodeDataInput) (nodeData *iot_grpcapi.NodeData, err error) {
	resp, err := c.api.GetLatestNodeData(ctx, &input)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			err = nil
		}

		return
	}

	nodeData = resp.NodeData
	return
}

func (c *Client) GetNodeData(input iot_grpcapi.GetNodeDataInput) (out []iot_grpcapi.NodeData, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetNodeDataWithContext(ctx, input)
}
func (c *Client) GetNodeDataWithContext(ctx context.Context, input iot_grpcapi.GetNodeDataInput) (out []iot_grpcapi.NodeData, err error) {
	nodeDataList, err := c.api.GetNodeData(ctx, &input)
	if nodeDataList != nil {
		for _, elem := range nodeDataList.NodeDataList {
			out = append(out, *elem)
		}
	}
	return
}

// Delete Node data functions
func (c *Client) DeleteNodeData(input iot_grpcapi.DeleteNodeDataInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.DeleteNodeDataWithContext(ctx, input)
}

func (c *Client) DeleteNodeDataWithContext(ctx context.Context, input iot_grpcapi.DeleteNodeDataInput) (err error) {
	_, err = c.api.DeleteNodeData(ctx, &input)
	return
}

func (c *Client) GetMedia(input iot_grpcapi.GetMediaInput) (iot_grpcapi.Media, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetMediaWithContext(ctx, input)
}
func (c *Client) GetMediaWithContext(ctx context.Context, input iot_grpcapi.GetMediaInput) (media iot_grpcapi.Media, err error) {
	output, err := c.api.GetMedia(ctx, &input)
	if output != nil && output.Media != nil {
		media = *output.Media
	}
	return
}

//RequestGetMediaSignedURLWithContext if successful returns a url where media data can be accessed.
func (c *Client) RequestGetMediaSignedURLWithContext(ctx context.Context, in *iot_grpcapi.GetMediaSignedUrlInput) (*iot_grpcapi.GetMediaSignedUrlOutput, error) {
	return c.api.RequestGetMediaSignedUrl(ctx, in)
}

//RequestGetMediaSignedURL if successful returns a url where media data can be accessed.
func (c *Client) RequestGetMediaSignedURL(in *iot_grpcapi.GetMediaSignedUrlInput) (*iot_grpcapi.GetMediaSignedUrlOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.RequestGetMediaSignedURLWithContext(ctx, in)
}

//RequestPutMediaSignedURLWithContext if successful returns a url where media data can be uploaded.
func (c *Client) RequestPutMediaSignedURLWithContext(ctx context.Context, in *iot_grpcapi.PutMediaSignedUrlInput) (*iot_grpcapi.PutMediaSignedUrlOutput, error) {
	return c.api.RequestPutMediaSignedUrl(ctx, in)
}

//RequestPutMediaSignedURL if successful returns a url where media data can be uploaded.
func (c *Client) RequestPutMediaSignedURL(in *iot_grpcapi.PutMediaSignedUrlInput) (*iot_grpcapi.PutMediaSignedUrlOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.RequestPutMediaSignedURLWithContext(ctx, in)
}

func (c *Client) GetTasksByStatus(input iot_grpcapi.GetTasksByStatusInput) ([]*iot_grpcapi.TaskDescription, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetTasksByStatusWithContext(ctx, input)
}
func (c *Client) GetTasksByStatusWithContext(ctx context.Context, input iot_grpcapi.GetTasksByStatusInput) (tasks []*iot_grpcapi.TaskDescription, err error) {
	result, err := c.api.GetTasksByStatus(ctx, &input)
	if result != nil {
		tasks = result.TaskList
	}
	return
}

func (c *Client) GetTasksModifiedSinceTimestamp(input iot_grpcapi.GetTasksModifiedSinceTimestampInput) (*iot_grpcapi.GetTasksModifiedSinceTimestampOutput, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetTasksModifiedSinceTimestampWithContext(ctx, input)

}
func (c *Client) GetTasksModifiedSinceTimestampWithContext(ctx context.Context, input iot_grpcapi.GetTasksModifiedSinceTimestampInput) (output *iot_grpcapi.GetTasksModifiedSinceTimestampOutput, err error) {
	output, err = c.api.GetTasksModifiedSinceTimestamp(ctx, &input)
	return
}

func (c *Client) GetNodeEventLog(input iot_grpcapi.GetNodeEventLogInput) (output *iot_grpcapi.GetNodeEventLogOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetNodeEventLogWithContext(ctx, input)
}
func (c *Client) GetNodeEventLogWithContext(ctx context.Context, input iot_grpcapi.GetNodeEventLogInput) (output *iot_grpcapi.GetNodeEventLogOutput, err error) {
	output, err = c.api.GetNodeEventLog(ctx, &input)
	return
}
