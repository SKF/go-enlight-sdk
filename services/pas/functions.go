package pas

import (
	"context"
	"io"
	"time"

	"github.com/SKF/go-enlight-sdk/services/pas/pasapi"
	"github.com/SKF/go-utility/log"
)

func (c *client) SetPointAlarmThreshold(input pasapi.SetPointAlarmThresholdInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetPointAlarmThresholdWithContext(ctx, input)
}
func (c *client) SetPointAlarmThresholdWithContext(ctx context.Context, input pasapi.SetPointAlarmThresholdInput) error {
	_, err := c.api.SetPointAlarmThreshold(ctx, &input)
	return err
}

func (c *client) GetPointAlarmThreshold(nodeID string) ([]pasapi.AlarmStatusInterval, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointAlarmThresholdWithContext(ctx, nodeID)
}
func (c *client) GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) (intervals []pasapi.AlarmStatusInterval, err error) {
	input := pasapi.GetPointAlarmThresholdInput{NodeId: nodeID}
	output, err := c.api.GetPointAlarmThreshold(ctx, &input)
	if output != nil {
		for _, interval := range output.Intervals.List {
			intervals = append(intervals, *interval)
		}
	}
	return
}

func (c *client) SetPointAlarmStatus(input pasapi.SetPointAlarmStatusInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetPointAlarmStatusWithContext(ctx, input)
}
func (c *client) SetPointAlarmStatusWithContext(ctx context.Context, input pasapi.SetPointAlarmStatusInput) error {
	_, err := c.api.SetPointAlarmStatus(ctx, &input)
	return err
}

func (c *client) GetPointAlarmStatus(input pasapi.GetPointAlarmStatusInput) (pasapi.AlarmStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointAlarmStatusWithContext(ctx, input)
}
func (c *client) GetPointAlarmStatusWithContext(ctx context.Context, input pasapi.GetPointAlarmStatusInput) (status pasapi.AlarmStatus, err error) {
	output, err := c.api.GetPointAlarmStatus(ctx, &input)
	if output != nil {
		status = output.AlarmStatus
	}
	return
}

func (c *client) GetPointAlarmStatusStream(dc chan<- pasapi.GetPointAlarmStatusStreamOutput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointAlarmStatusStreamWithContext(ctx, dc)

}
func (c *client) GetPointAlarmStatusStreamWithContext(ctx context.Context, dc chan<- pasapi.GetPointAlarmStatusStreamOutput) (err error) {
	stream, err := c.api.GetPointAlarmStatusStream(ctx, &pasapi.GetPointAlarmStatusStreamInput{})
	if err != nil {
		return
	}

	for {
		var output *pasapi.GetPointAlarmStatusStreamOutput
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
