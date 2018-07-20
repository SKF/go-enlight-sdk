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

	output, err := c.api.CreateTask(ctx, &task)
	if output != nil {
		taskID = output.Value
	}
	return
}

func (c *client) DeleteTask(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	input := api.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.DeleteTask(ctx, &input)
	return
}

func (c *client) SetTaskCompleted(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	input := api.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.SetTaskCompleted(ctx, &input)
	return
}

func (c *client) GetAllTasks(userID string) (out []api.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

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
	input := api.IngestNodeDataInput{
		NodeId:   nodeID,
		NodeData: &nodeData,
	}
	_, err = c.api.IngestNodeData(ctx, &input)
	return
}

func (c *client) IngestNodeDataStream(inputChannel <-chan api.IngestNodeDataStreamInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

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

func (c *client) GetNodeData(input api.GetNodeDataInput) (out []api.NodeData, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	nodeDataList, err := c.api.GetNodeData(ctx, &input)
	if nodeDataList != nil {
		for _, elem := range nodeDataList.NodeDataList {
			out = append(out, *elem)
		}
	}
	return
}

func (c *client) GetNodeDataStream(input api.GetNodeDataStreamInput, dc chan<- api.GetNodeDataStreamOutput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

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
