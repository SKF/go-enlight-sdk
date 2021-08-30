package pas

import (
	"context"
	"fmt"
	"time"

	pas_api "github.com/SKF/proto/v2/pas"
)

// SetPointAlarmThreshold sets the alarm threshold for a specific point
func (c *Client) SetPointAlarmThreshold(input pas_api.SetPointAlarmThresholdInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetPointAlarmThresholdWithContext(ctx, input)
}

// SetPointAlarmThresholdWithContext sets the alarm threshold for a specific point
func (c *Client) SetPointAlarmThresholdWithContext(ctx context.Context, input pas_api.SetPointAlarmThresholdInput) error {
	_, err := c.api.SetPointAlarmThreshold(ctx, &input)
	return err
}

// GetPointAlarmThreshold gets the alarm threshold for a specific point
func (c *Client) GetPointAlarmThreshold(nodeID string) (pas_api.GetPointAlarmThresholdOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointAlarmThresholdWithContext(ctx, nodeID)
}

// GetPointAlarmThresholdWithContext gets the alarm threshold for a specific point
func (c *Client) GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) (output pas_api.GetPointAlarmThresholdOutput, err error) {
	input := pas_api.GetPointAlarmThresholdInput{NodeId: nodeID}
	resp, err := c.api.GetPointAlarmThreshold(ctx, &input)
	if resp != nil {
		output = *resp
	}
	return
}

// SetPointAlarmStatus sets the alarm status for a specific point
func (c *Client) SetPointAlarmStatus(input pas_api.SetPointAlarmStatusInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SetPointAlarmStatusWithContext(ctx, input)
}

// SetPointAlarmStatusWithContext sets the alarm status for a specific point
func (c *Client) SetPointAlarmStatusWithContext(ctx context.Context, input pas_api.SetPointAlarmStatusInput) error {
	_, err := c.api.SetPointAlarmStatus(ctx, &input)
	return err
}

// GetPointAlarmStatus gets the alarm status for a specific point
func (c *Client) GetPointAlarmStatus(input pas_api.GetPointAlarmStatusInput) (pas_api.AlarmStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointAlarmStatusWithContext(ctx, input)
}

// GetPointAlarmStatusWithContext gets the alarm status for a specific point
func (c *Client) GetPointAlarmStatusWithContext(ctx context.Context, input pas_api.GetPointAlarmStatusInput) (status pas_api.AlarmStatus, err error) {
	output, err := c.api.GetPointAlarmStatus(ctx, &input)
	if output != nil {
		status = output.AlarmStatus
	}
	return
}

func (c *Client) GetPointAlarmStatusV2(input pas_api.GetPointAlarmStatusInput) (*pas_api.GetPointAlarmStatusOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetPointAlarmStatusV2WithContext(ctx, input)
}

func (c *Client) GetPointAlarmStatusV2WithContext(ctx context.Context, input pas_api.GetPointAlarmStatusInput) (*pas_api.GetPointAlarmStatusOutput, error) {
	output, err := c.api.GetPointAlarmStatus(ctx, &input)
	if err != nil {
		return nil, err
	}

	return output, err
}

// GetPointAlarmStatusEventLog get all alarm events after a given sequence ID
func (c *Client) GetPointAlarmStatusEventLog(seqID string) (events pas_api.GetPointAlarmStatusEventLogOutput, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.GetPointAlarmStatusEventLogWithContext(ctx, seqID)
}

// GetPointAlarmStatusEventLogWithContext get all alarm events after a given sequence ID
func (c *Client) GetPointAlarmStatusEventLogWithContext(ctx context.Context, seqID string) (events pas_api.GetPointAlarmStatusEventLogOutput, err error) {
	return pas_api.GetPointAlarmStatusEventLogOutput{}, fmt.Errorf("this api is deprecated")
}

// CalculateAndSetPointAlarmStatus calculates and sets new PAS based on input data
func (c *Client) CalculateAndSetPointAlarmStatus(input pas_api.CalculateAndSetPointAlarmStatusInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.CalculateAndSetPointAlarmStatusWithContext(ctx, input)
}

// CalculateAndSetPointAlarmStatusWithContext calculates and sets new PAS based on input data
func (c *Client) CalculateAndSetPointAlarmStatusWithContext(ctx context.Context, input pas_api.CalculateAndSetPointAlarmStatusInput) error {
	_, err := c.api.CalculateAndSetPointAlarmStatus(ctx, &input)
	return err
}
