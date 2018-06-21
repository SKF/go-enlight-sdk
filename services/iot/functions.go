package iot

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
)

func (c *client) CreateTask(task iotgrpcapi.InitialTaskDescription) (taskID string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	id, err := c.api.CreateTask(ctx, &task)
	if id != nil {
		taskID = id.Value
	}
	return
}

func (c *client) DeleteTask(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	input := iotgrpcapi.TaskUser{
		UserId: userID,
		TaskId: taskID,
	}
	_, err = c.api.DeleteTask(ctx, &input)
	return
}

func (c *client) SetTaskCompleted(userID, taskID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	input := iotgrpcapi.TaskUser{
		UserId: userID,
		TaskId: taskID,
	}
	_, err = c.api.SetTaskCompleted(ctx, &input)
	return
}

func (c *client) GetAllTasks(userID string) (out []iotgrpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	input := iotgrpcapi.PrimitiveString{Value: userID}
	tasks, err := c.api.GetAllTasks(ctx, &input)
	if tasks != nil {
		for _, desc := range tasks.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}

func (c *client) GetUncompletedTasks(userID string) (out []iotgrpcapi.TaskDescription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	input := iotgrpcapi.PrimitiveString{Value: userID}
	tasks, err := c.api.GetUncompletedTasks(ctx, &input)
	if tasks != nil {
		for _, desc := range tasks.TaskDescriptionArr {
			out = append(out, *desc)
		}
	}
	return
}
