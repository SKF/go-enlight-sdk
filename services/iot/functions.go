package iot

import (
	"context"
	"io"
	"time"

	api "github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
	"github.com/SKF/go-utility/log"
)

func (c *client) CreateTask(task api.InitialTaskDescription) (taskID string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.CreateTaskWithContext(ctx, task)
}
func (c *client) CreateTaskWithContext(ctx context.Context, task api.InitialTaskDescription) (taskID string, err error) {
	output, err := c.api.CreateTask(ctx, &task)
	if output != nil {
		taskID = output.Value
	}
	return
}

func (c *client) DeleteTask(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.DeleteTaskWithContext(ctx, userID, taskID)
}
func (c *client) DeleteTaskWithContext(ctx context.Context, userID, taskID string) (err error) {
	input := api.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.DeleteTask(ctx, &input)
	return
}

func (c *client) SetTaskCompleted(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetTaskCompletedWithContext(ctx, userID, taskID)
}
func (c *client) SetTaskCompletedWithContext(ctx context.Context, userID, taskID string) (err error) {
	input := api.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.SetTaskCompleted(ctx, &input)
	return
}

func (c *client) GetAllTasks(userID string) (out []api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetAllTasksWithContext(ctx, userID)
}
func (c *client) GetAllTasksWithContext(ctx context.Context, userID string) (out []api.TaskDescription, err error) {
	input := api.PrimitiveString{Value: userID}
	output, err := c.api.GetAllTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *client) GetUncompletedTasks(userID string) (out []api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetUncompletedTasksWithContext(ctx, userID)
}
func (c *client) GetUncompletedTasksWithContext(ctx context.Context, userID string) (out []api.TaskDescription, err error) {
	input := api.PrimitiveString{Value: userID}
	output, err := c.api.GetUncompletedTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *client) GetUncompletedTasksByHierarchy(nodeID string) (out []api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetUncompletedTasksByHierarchyWithContext(ctx, nodeID)
}
func (c *client) GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) (out []api.TaskDescription, err error) {
	input := api.PrimitiveString{Value: nodeID}
	tasks, err := c.api.GetUncompletedTasksByHierarchy(ctx, &input)
	if tasks != nil {
		for _, desc := range tasks.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *client) SetTaskStatus(taskID, userID string, status api.TaskStatus) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetTaskStatusWithContext(ctx, taskID, userID, status)
}
func (c *client) SetTaskStatusWithContext(ctx context.Context, taskID, userID string, status api.TaskStatus) (err error) {
	input := api.SetTaskStatusInput{
		TaskId: taskID,
		UserId: userID,
		Status: status,
	}
	_, err = c.api.SetTaskStatus(ctx, &input)
	return
}

func (c *client) IngestNodeData(nodeID string, nodeData api.NodeData) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.IngestNodeDataWithContext(ctx, nodeID, nodeData)
}
func (c *client) IngestNodeDataWithContext(ctx context.Context, nodeID string, nodeData api.NodeData) (err error) {
	input := api.IngestNodeDataInput{
		NodeId:   nodeID,
		NodeData: &nodeData,
	}
	_, err = c.api.IngestNodeData(ctx, &input)
	return
}

func (c *client) IngestNodeDataStream(inputChannel <-chan api.IngestNodeDataStreamInput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.IngestNodeDataStreamWithContext(ctx, inputChannel)
}
func (c *client) IngestNodeDataStreamWithContext(ctx context.Context, inputChannel <-chan api.IngestNodeDataStreamInput) (err error) {
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

func (c *client) GetLatestNodeData(input api.GetLatestNodeDataInput) (api.NodeData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetLatestNodeDataWithContext(ctx, input)
}

func (c *client) GetLatestNodeDataWithContext(ctx context.Context, input api.GetLatestNodeDataInput) (nodeData api.NodeData, err error) {
	resp, err := c.api.GetLatestNodeData(ctx, &input)
	if err != nil {
		return
	}

	nodeData = *resp.NodeData
	return
}

func (c *client) GetNodeData(input api.GetNodeDataInput) (out []api.NodeData, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetNodeDataWithContext(ctx, input)
}
func (c *client) GetNodeDataWithContext(ctx context.Context, input api.GetNodeDataInput) (out []api.NodeData, err error) {
	nodeDataList, err := c.api.GetNodeData(ctx, &input)
	if nodeDataList != nil {
		for _, elem := range nodeDataList.NodeDataList {
			out = append(out, *elem)
		}
	}
	return
}

func (c *client) GetNodeDataStream(input api.GetNodeDataStreamInput, dc chan<- api.GetNodeDataStreamOutput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetNodeDataStreamWithContext(ctx, input, dc)
}
func (c *client) GetNodeDataStreamWithContext(ctx context.Context, input api.GetNodeDataStreamInput, dc chan<- api.GetNodeDataStreamOutput) (err error) {
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
			log.WithError(err).Errorf("stream.Recv")
			return
		}
		dc <- *nodeData
	}
}

func (c *client) GetTaskStream(input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetTaskStreamWithContext(ctx, input, dc)
}
func (c *client) GetTaskStreamWithContext(ctx context.Context, input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error) {
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
			log.WithError(err).Errorf("stream.Recv")
			return
		}
		dc <- *nodeData
	}
}
