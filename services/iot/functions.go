package iot

import (
	"context"
	"io"
	"time"

	api "github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) CreateTask(task api.InitialTaskDescription) (taskID string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.CreateTaskWithContext(ctx, task)
}
func (c *Client) CreateTaskWithContext(ctx context.Context, task api.InitialTaskDescription) (taskID string, err error) {
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
	input := api.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.DeleteTask(ctx, &input)
	return
}

func (c *Client) SetTaskCompleted(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetTaskCompletedWithContext(ctx, userID, taskID)
}
func (c *Client) SetTaskCompletedWithContext(ctx context.Context, userID, taskID string) (err error) {
	input := api.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.SetTaskCompleted(ctx, &input)
	return
}

func (c *Client) GetAllTasks(userID string) (out []api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetAllTasksWithContext(ctx, userID)
}
func (c *Client) GetAllTasksWithContext(ctx context.Context, userID string) (out []api.TaskDescription, err error) {
	input := api.PrimitiveString{Value: userID}
	output, err := c.api.GetAllTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *Client) GetUncompletedTasks(userID string) (out []api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetUncompletedTasksWithContext(ctx, userID)
}
func (c *Client) GetUncompletedTasksWithContext(ctx context.Context, userID string) (out []api.TaskDescription, err error) {
	input := api.PrimitiveString{Value: userID}
	output, err := c.api.GetUncompletedTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *Client) GetUncompletedTasksByHierarchy(nodeID string) (out []api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetUncompletedTasksByHierarchyWithContext(ctx, nodeID)
}
func (c *Client) GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) (out []api.TaskDescription, err error) {
	input := api.PrimitiveString{Value: nodeID}
	tasks, err := c.api.GetUncompletedTasksByHierarchy(ctx, &input)
	if tasks != nil {
		for _, desc := range tasks.TaskDescriptionArr {
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
func (c *Client) SetTaskStatus(input api.SetTaskStatusInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetTaskStatusWithContext(ctx, input)
}
func (c *Client) SetTaskStatusWithContext(ctx context.Context, input api.SetTaskStatusInput) (err error) {
	_, err = c.api.SetTaskStatus(ctx, &input)
	return
}

func (c *Client) IngestNodeData(input api.IngestNodeDataInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.IngestNodeDataWithContext(ctx, input)
}
func (c *Client) IngestNodeDataWithContext(ctx context.Context, input api.IngestNodeDataInput) (err error) {
	_, err = c.api.IngestNodeData(ctx, &input)
	return
}

func (c *Client) IngestNodeDataStream(inputChannel <-chan api.IngestNodeDataStreamInput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.IngestNodeDataStreamWithContext(ctx, inputChannel)
}
func (c *Client) IngestNodeDataStreamWithContext(ctx context.Context, inputChannel <-chan api.IngestNodeDataStreamInput) (err error) {
	stream, err := c.api.IngestNodeDataStream(ctx)
	if err != nil {
		return
	}
	for nodeData := range inputChannel {
		nd := nodeData
		if err = stream.Send(&nd); err != nil {
			return
		}
	}

	_, err = stream.CloseAndRecv()
	return
}

func (c *Client) GetTaskByUUID(input api.PrimitiveString) (output *api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetTaskByUUIDWithContext(ctx, input)
}
func (c *Client) GetTaskByUUIDWithContext(ctx context.Context, input api.PrimitiveString) (output *api.TaskDescription, err error) {
	response, err := c.api.GetTaskByUUID(ctx, &input)
	if response != nil {
		output = response
	}
	return
}

func (c *Client) GetTaskByLongId(input api.PrimitiveLong) (output *api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetTaskByLongIdWithContext(ctx, input)
}
func (c *Client) GetTaskByLongIdWithContext(ctx context.Context, input api.PrimitiveLong) (output *api.TaskDescription, err error) {
	response, err := c.api.GetTaskByLongId(ctx, &input)
	if response != nil {
		output = response
	}
	return
}

func (c *Client) GetLatestNodeData(input api.GetLatestNodeDataInput) (*api.NodeData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetLatestNodeDataWithContext(ctx, input)
}

func (c *Client) GetLatestNodeDataWithContext(ctx context.Context, input api.GetLatestNodeDataInput) (nodeData *api.NodeData, err error) {
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

func (c *Client) GetNodeData(input api.GetNodeDataInput) (out []api.NodeData, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetNodeDataWithContext(ctx, input)
}
func (c *Client) GetNodeDataWithContext(ctx context.Context, input api.GetNodeDataInput) (out []api.NodeData, err error) {
	nodeDataList, err := c.api.GetNodeData(ctx, &input)
	if nodeDataList != nil {
		for _, elem := range nodeDataList.NodeDataList {
			out = append(out, *elem)
		}
	}
	return
}

func (c *Client) GetNodeDataStream(input api.GetNodeDataStreamInput, dc chan<- api.GetNodeDataStreamOutput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetNodeDataStreamWithContext(ctx, input, dc)
}
func (c *Client) GetNodeDataStreamWithContext(ctx context.Context, input api.GetNodeDataStreamInput, dc chan<- api.GetNodeDataStreamOutput) (err error) {
	stream, err := c.api.GetNodeDataStream(ctx, &input)
	if err != nil {
		return
	}

	for {
		var nodeData *api.GetNodeDataStreamOutput
		nodeData, err = stream.Recv()
		if err == io.EOF {
			err = nil
			return
		}
		if err != nil {
			return
		}
		dc <- *nodeData
	}
}

func (c *Client) GetMedia(input api.GetMediaInput) (api.Media, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetMediaWithContext(ctx, input)
}
func (c *Client) GetMediaWithContext(ctx context.Context, input api.GetMediaInput) (media api.Media, err error) {
	output, err := c.api.GetMedia(ctx, &input)
	if output != nil && output.Media != nil {
		media = *output.Media
	}
	return
}

func (c *Client) GetTaskStream(input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetTaskStreamWithContext(ctx, input, dc)
}
func (c *Client) GetTaskStreamWithContext(ctx context.Context, input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error) {
	stream, err := c.api.GetTaskStream(ctx, &input)
	if err != nil {
		return
	}

	for {
		var nodeData *api.GetTaskStreamOutput
		nodeData, err = stream.Recv()
		if err == io.EOF {
			err = nil
			return
		}
		if err != nil {
			return
		}
		dc <- *nodeData
	}
}

func (c *Client) GetTasksByStatus(input api.GetTasksByStatusInput) ([]*api.TaskDescription, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetTasksByStatusWithContext(ctx, input)
}
func (c *Client) GetTasksByStatusWithContext(ctx context.Context, input api.GetTasksByStatusInput) (tasks []*api.TaskDescription, err error) {
	result, err := c.api.GetTasksByStatus(ctx, &input)
	if result != nil {
		tasks = result.TaskList
	}
	return
}
