package mhub

import (
	"context"
	"io"
	"time"

	"github.com/SKF/go-enlight-sdk/services/mhub/mhubapi"
)

func (c *client) SetTaskStatus(id int64, status mhubapi.TaskStatus) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	return c.SetTaskStatusWithContext(ctx, id, status)
}

func (c *client) SetTaskStatusWithContext(ctx context.Context, id int64, status mhubapi.TaskStatus) (err error) {
	_, err = c.api.SetTaskStatus(ctx, &mhubapi.SetTaskStatusInput{
		Id:     id,
		Status: status,
	})
	return
}

func (c *client) AvailableDSKFStream(dc chan<- mhubapi.AvailableDSKFStreamOutput) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.AvailableDSKFStreamWithContext(ctx, dc)
}

func (c *client) AvailableDSKFStreamWithContext(ctx context.Context, dc chan<- mhubapi.AvailableDSKFStreamOutput) (err error) {
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
			return
		}
		dc <- *output
	}
}
