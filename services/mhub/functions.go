package mhub

import (
	"context"
	"io"
	"time"

	"github.com/SKF/go-enlight-sdk/services/mhub/mhubapi"
	"github.com/SKF/go-utility/log"
	"github.com/SKF/go-utility/uuid"
)

func (c *client) SetTaskStatus(taskID, userID uuid.UUID, status mhubapi.TaskStatus) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	_, err = c.api.SetTaskStatus(ctx, &mhubapi.SetTaskStatusInput{
		TaskId: taskID.String(),
		Status: status,
	})
	return
}

func (c *client) GetTasksStream(dc chan<- mhubapi.GetTasksStreamOutput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := c.api.GetTasksStream(ctx, &mhubapi.GetTasksStreamInput{})
	if err != nil {
		return
	}

	for {
		var output *mhubapi.GetTasksStreamOutput
		output, err = stream.Recv()
		if err == io.EOF {
			err = nil
			return
		}
		if err != nil {
			log.WithError(err).Errorf("stream.Recv")
			return
		}
		dc <- *output
	}
}
