package micrologproxyhub

import (
	"context"
	"io"
	"time"

	api "github.com/SKF/go-enlight-sdk/services/micrologproxyhub/grpcapi"
	"github.com/SKF/go-utility/log"
	"github.com/SKF/go-utility/uuid"
)

func (c *client) SetTaskStatus(taskID, userID uuid.UUID, status api.TaskStatus) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	_, err = c.api.SetTaskStatus(ctx, &api.SetTaskStatusInput{
		TaskId: taskID.String(),
		Status: status,
	})
	return
}

func (c *client) GetTasksStream(dc chan<- api.GetTasksStreamOutput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := c.api.GetTasksStream(ctx, &api.GetTasksStreamInput{})
	if err != nil {
		return
	}

	for {
		var output *api.GetTasksStreamOutput
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
