package iot

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
)

func (c *client) CreateTask(task iotgrpcapi.InitialTaskDescription) (taskID string, err error) {
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

	input := iotgrpcapi.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.DeleteTask(ctx, &input)
	return
}

func (c *client) SetTaskCompleted(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	input := iotgrpcapi.TaskUser{UserId: userID, TaskId: taskID}
	_, err = c.api.SetTaskCompleted(ctx, &input)
	return
}

func (c *client) GetAllTasks(userID string) (out []iotgrpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	input := iotgrpcapi.PrimitiveString{Value: userID}
	output, err := c.api.GetAllTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *client) GetUncompletedTasks(userID string) (out []iotgrpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	input := iotgrpcapi.PrimitiveString{Value: userID}
	output, err := c.api.GetUncompletedTasks(ctx, &input)

	if output != nil {
		for _, desc := range output.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *client) GetUncompletedTasksByHierarchy(nodeID string) (out []iotgrpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	input := iotgrpcapi.PrimitiveString{Value: nodeID}
	tasks, err := c.api.GetUncompletedTasksByHierarchy(ctx, &input)
	if tasks != nil {
		for _, desc := range tasks.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}
