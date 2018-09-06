package iot

import (
	"context"
	"io"
	"time"

	proto_common "github.com/SKF/proto/common"
	proto_iot "github.com/SKF/proto/iot"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) CreateTask(task proto_iot.InitialTaskDescription) (taskID string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.CreateTaskWithContext(ctx, task)
}
func (c *Client) CreateTaskWithContext(ctx context.Context, task proto_iot.InitialTaskDescription) (taskID string, err error) {
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
	input := proto_iot.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.DeleteTask(ctx, &input)
	return
}

func (c *Client) SetTaskCompleted(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetTaskCompletedWithContext(ctx, userID, taskID)
}
func (c *Client) SetTaskCompletedWithContext(ctx context.Context, userID, taskID string) (err error) {
	input := proto_iot.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.SetTaskCompleted(ctx, &input)
	return
}

func (c *Client) GetAllTasks(userID string) (out []proto_iot.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetAllTasksWithContext(ctx, userID)
}
func (c *Client) GetAllTasksWithContext(ctx context.Context, userID string) (out []proto_iot.TaskDescription, err error) {
	input := proto_common.PrimitiveString{Value: userID}
	output, err := c.api.GetAllTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *Client) GetUncompletedTasks(userID string) (out []proto_iot.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetUncompletedTasksWithContext(ctx, userID)
}
func (c *Client) GetUncompletedTasksWithContext(ctx context.Context, userID string) (out []proto_iot.TaskDescription, err error) {
	input := proto_common.PrimitiveString{Value: userID}
	output, err := c.api.GetUncompletedTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *Client) GetUncompletedTasksByHierarchy(nodeID string) (out []proto_iot.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetUncompletedTasksByHierarchyWithContext(ctx, nodeID)
}
func (c *Client) GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) (out []proto_iot.TaskDescription, err error) {
	input := proto_common.PrimitiveString{Value: nodeID}
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
func (c *Client) SetTaskStatus(input proto_iot.SetTaskStatusInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetTaskStatusWithContext(ctx, input)
}
func (c *Client) SetTaskStatusWithContext(ctx context.Context, input proto_iot.SetTaskStatusInput) (err error) {
	_, err = c.api.SetTaskStatus(ctx, &input)
	return
}

func (c *Client) IngestNodeData(input proto_iot.IngestNodeDataInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.IngestNodeDataWithContext(ctx, input)
}
func (c *Client) IngestNodeDataWithContext(ctx context.Context, input proto_iot.IngestNodeDataInput) (err error) {
	_, err = c.api.IngestNodeData(ctx, &input)
	return
}

func (c *Client) IngestNodeDataStream(inputChannel <-chan proto_iot.IngestNodeDataStreamInput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.IngestNodeDataStreamWithContext(ctx, inputChannel)
}
func (c *Client) IngestNodeDataStreamWithContext(ctx context.Context, inputChannel <-chan proto_iot.IngestNodeDataStreamInput) (err error) {
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

func (c *Client) GetTaskByUUID(input string) (output *proto_iot.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetTaskByUUIDWithContext(ctx, input)
}
func (c *Client) GetTaskByUUIDWithContext(ctx context.Context, input string) (output *proto_iot.TaskDescription, err error) {
	grpcInput := proto_iot.GetTaskByUUIDInput{
		TaskId: input,
	}
	response, err := c.api.GetTaskByUUID(ctx, &grpcInput)
	if response != nil {
		output = response.Task
	}
	return
}

func (c *Client) GetTaskByLongId(input int64) (output *proto_iot.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetTaskByLongIdWithContext(ctx, input)
}
func (c *Client) GetTaskByLongIdWithContext(ctx context.Context, input int64) (output *proto_iot.TaskDescription, err error) {
	grpcInput := proto_iot.GetTaskByLongIdInput{
		TaskId: input,
	}
	response, err := c.api.GetTaskByLongId(ctx, &grpcInput)
	if response != nil {
		output = response.Task
	}
	return
}

func (c *Client) GetLatestNodeData(input proto_iot.GetLatestNodeDataInput) (*proto_iot.NodeData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetLatestNodeDataWithContext(ctx, input)
}

func (c *Client) GetLatestNodeDataWithContext(ctx context.Context, input proto_iot.GetLatestNodeDataInput) (nodeData *proto_iot.NodeData, err error) {
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

func (c *Client) GetNodeData(input proto_iot.GetNodeDataInput) (out []proto_iot.NodeData, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetNodeDataWithContext(ctx, input)
}
func (c *Client) GetNodeDataWithContext(ctx context.Context, input proto_iot.GetNodeDataInput) (out []proto_iot.NodeData, err error) {
	nodeDataList, err := c.api.GetNodeData(ctx, &input)
	if nodeDataList != nil {
		for _, elem := range nodeDataList.NodeDataList {
			out = append(out, *elem)
		}
	}
	return
}

func (c *Client) GetNodeDataStream(input proto_iot.GetNodeDataStreamInput, dc chan<- proto_iot.GetNodeDataStreamOutput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetNodeDataStreamWithContext(ctx, input, dc)
}
func (c *Client) GetNodeDataStreamWithContext(ctx context.Context, input proto_iot.GetNodeDataStreamInput, dc chan<- proto_iot.GetNodeDataStreamOutput) (err error) {
	stream, err := c.api.GetNodeDataStream(ctx, &input)
	if err != nil {
		return
	}

	for {
		var nodeData *proto_iot.GetNodeDataStreamOutput
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

func (c *Client) GetMedia(input proto_iot.GetMediaInput) (proto_iot.Media, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetMediaWithContext(ctx, input)
}
func (c *Client) GetMediaWithContext(ctx context.Context, input proto_iot.GetMediaInput) (media proto_iot.Media, err error) {
	output, err := c.api.GetMedia(ctx, &input)
	if output != nil && output.Media != nil {
		media = *output.Media
	}
	return
}

func (c *Client) GetTaskStream(input proto_iot.GetTaskStreamInput, dc chan<- proto_iot.GetTaskStreamOutput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetTaskStreamWithContext(ctx, input, dc)
}
func (c *Client) GetTaskStreamWithContext(ctx context.Context, input proto_iot.GetTaskStreamInput, dc chan<- proto_iot.GetTaskStreamOutput) (err error) {
	stream, err := c.api.GetTaskStream(ctx, &input)
	if err != nil {
		return
	}

	for {
		var nodeData *proto_iot.GetTaskStreamOutput
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

func (c *Client) GetTasksByStatus(input proto_iot.GetTasksByStatusInput) ([]*proto_iot.TaskDescription, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetTasksByStatusWithContext(ctx, input)
}
func (c *Client) GetTasksByStatusWithContext(ctx context.Context, input proto_iot.GetTasksByStatusInput) (tasks []*proto_iot.TaskDescription, err error) {
	result, err := c.api.GetTasksByStatus(ctx, &input)
	if result != nil {
		tasks = result.TaskList
	}
	return
}
