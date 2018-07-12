package hierarchy

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SKF/go-eventsource/eventsource"

	"github.com/SKF/go-enlight-sdk/services/hierarchy/grpcapi"
)

func (c *client) SaveNode(request grpcapi.SaveNodeInput) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.SaveNode(ctx, &request)
	return resp.GetValue(), err
}

func (c *client) GetNode(uuid string) (node grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.GetNode(ctx, &grpcapi.PrimitiveString{Value: uuid})
	if resp != nil {
		node = *resp
	}
	return
}

func (c *client) GetNodes(parentID string) (nodes []grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.GetNodes(ctx, &grpcapi.PrimitiveString{Value: parentID})
	if resp != nil {
		for _, node := range resp.Nodes {
			if node != nil {
				nodes = append(nodes, *node)
			}
		}
	}
	return
}

func (c *client) GetChildNodes(parentID string) (nodes []grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.GetChildNodes(ctx, &grpcapi.PrimitiveString{Value: parentID})
	if resp != nil {
		for _, node := range resp.Nodes {
			if node != nil {
				nodes = append(nodes, *node)
			}
		}
	}
	return
}

func (c *client) DeleteNode(request grpcapi.DeleteNodeInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := c.api.DeleteNode(ctx, &request)
	return err
}

func (c *client) GetParentNode(nodeID string) (node grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.GetParentNode(ctx, &grpcapi.PrimitiveString{Value: nodeID})
	if resp != nil {
		node = *resp
	}
	return
}

func (c *client) GetAncestors(nodeID string) (nodes []grpcapi.AncestorNode, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.GetAncestors(ctx, &grpcapi.GetAncestorsInput{NodeId: nodeID})
	if resp != nil {
		for _, node := range resp.Nodes {
			if node != nil {
				nodes = append(nodes, *node)
			}
		}
	}
	return
}

func (c *client) GetEvents(since int, limit *int32) (events []eventsource.Record, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	input := grpcapi.GetEventsInput{Since: int64(since)}
	if limit != nil {
		input.Limit = &grpcapi.PrimitiveInt32{Value: *limit}
	}

	output, err := c.api.GetEvents(ctx, &input)
	if err != nil {
		return
	}

	err = json.Unmarshal(output.Events, &events)
	return
}
