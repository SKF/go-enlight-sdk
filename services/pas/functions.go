package pas

import (
	"context"
	"io"
	"time"

	"github.com/SKF/go-enlight-sdk/services/pas/pasapi"
	"github.com/SKF/go-utility/log"
)

func (c *client) SetPointThreshold(input pasapi.SetPointThresholdInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetPointThresholdWithContext(ctx, input)
}
func (c *client) SetPointThresholdWithContext(ctx context.Context, input pasapi.SetPointThresholdInput) error {
	_, err := c.api.SetPointThreshold(ctx, &input)
	return err
}

func (c *client) GetPointThreshold(nodeID string) ([]pasapi.AlarmStatusInterval, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointThresholdWithContext(ctx, nodeID)
}
func (c *client) GetPointThresholdWithContext(ctx context.Context, nodeID string) (intervals []pasapi.AlarmStatusInterval, err error) {
	input := pasapi.GetPointThresholdInput{NodeId: nodeID}
	output, err := c.api.GetPointThreshold(ctx, &input)
	if output != nil {
		for _, interval := range output.Intervals {
			intervals = append(intervals, *interval)
		}
	}
	return
}

func (c *client) SetPointStatus(input pasapi.SetPointStatusInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetPointStatusWithContext(ctx, input)
}
func (c *client) SetPointStatusWithContext(ctx context.Context, input pasapi.SetPointStatusInput) error {
	_, err := c.api.SetPointStatus(ctx, &input)
	return err
}

func (c *client) GetPointStatus(input pasapi.GetPointStatusInput) (pasapi.AlarmStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointStatusWithContext(ctx, input)
}
func (c *client) GetPointStatusWithContext(ctx context.Context, input pasapi.GetPointStatusInput) (status pasapi.AlarmStatus, err error) {
	output, err := c.api.GetPointStatus(ctx, &input)
	if output != nil {
		status = output.AlarmStatus
	}
	return
}

func (c *client) GetPointStatusStream(dc chan<- pasapi.GetPointStatusStreamOutput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointStatusStreamWithContext(ctx, dc)

}
func (c *client) GetPointStatusStreamWithContext(ctx context.Context, dc chan<- pasapi.GetPointStatusStreamOutput) (err error) {
	stream, err := c.api.GetPointStatusStream(ctx, &pasapi.GetPointStatusStreamInput{})
	if err != nil {
		return
	}

	for {
		var output *pasapi.GetPointStatusStreamOutput
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
