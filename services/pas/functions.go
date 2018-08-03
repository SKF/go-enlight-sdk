package pas

import (
	"context"
	"io"
	"time"

	"github.com/SKF/go-enlight-sdk/services/pas/pasapi"
	"github.com/SKF/go-utility/log"
)

// SetPointAlarmThreshold sets the alarm threshold for a specific point
func (c *Client) SetPointAlarmThreshold(input pasapi.SetPointAlarmThresholdInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetPointAlarmThresholdWithContext(ctx, input)
}

// SetPointAlarmThresholdWithContext sets the alarm threshold for a specific point
func (c *Client) SetPointAlarmThresholdWithContext(ctx context.Context, input pasapi.SetPointAlarmThresholdInput) error {
	_, err := c.api.SetPointAlarmThreshold(ctx, &input)
	return err
}

// GetPointAlarmThreshold gets the alarm threshold for a specific point
func (c *Client) GetPointAlarmThreshold(nodeID string) (pasapi.GetPointAlarmThresholdOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointAlarmThresholdWithContext(ctx, nodeID)
}

// GetPointAlarmThresholdWithContext gets the alarm threshold for a specific point
func (c *Client) GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) (output pasapi.GetPointAlarmThresholdOutput, err error) {
	input := pasapi.GetPointAlarmThresholdInput{NodeId: nodeID}
	resp, err := c.api.GetPointAlarmThreshold(ctx, &input)
	if resp != nil {
		output = *resp
	}
	return
}

// SetPointAlarmStatus sets the alarm status for a specific point
func (c *Client) SetPointAlarmStatus(input pasapi.SetPointAlarmStatusInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetPointAlarmStatusWithContext(ctx, input)
}

// SetPointAlarmStatusWithContext sets the alarm status for a specific point
func (c *Client) SetPointAlarmStatusWithContext(ctx context.Context, input pasapi.SetPointAlarmStatusInput) error {
	_, err := c.api.SetPointAlarmStatus(ctx, &input)
	return err
}

// GetPointAlarmStatus gets the alarm status for a specific point
func (c *Client) GetPointAlarmStatus(input pasapi.GetPointAlarmStatusInput) (pasapi.AlarmStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointAlarmStatusWithContext(ctx, input)
}

// GetPointAlarmStatusWithContext gets the alarm status for a specific point
func (c *Client) GetPointAlarmStatusWithContext(ctx context.Context, input pasapi.GetPointAlarmStatusInput) (status pasapi.AlarmStatus, err error) {
	output, err := c.api.GetPointAlarmStatus(ctx, &input)
	if output != nil {
		status = output.AlarmStatus
	}
	return
}

// GetPointAlarmStatusStream gets a stream of alarm status updates for all points
func (c *Client) GetPointAlarmStatusStream(dc chan<- pasapi.GetPointAlarmStatusStreamOutput) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetPointAlarmStatusStreamWithContext(ctx, dc)

}

// GetPointAlarmStatusStreamWithContext gets a stream of alarm status updates for all points
func (c *Client) GetPointAlarmStatusStreamWithContext(ctx context.Context, dc chan<- pasapi.GetPointAlarmStatusStreamOutput) (err error) {
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
