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

func (c *client) AvailableDSKFStream(dc chan<- mhubapi.AvailableDSKFStreamOutput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := c.api.AvailableDSKFStream(ctx, &mhubapi.AvailableDSKFStreamInput{})
	if err != nil {
		return
	}

	for {
		var output *mhubapi.AvailableDSKFStreamOutput
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
