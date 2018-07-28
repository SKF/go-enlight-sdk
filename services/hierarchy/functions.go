package hierarchy

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SKF/go-eventsource/eventsource"

	"github.com/SKF/go-enlight-sdk/services/hierarchy/grpcapi"
)

// SaveNode will add the node if it this not exist and otherwise
// create it.
func (c *Client) SaveNode(request grpcapi.SaveNodeInput) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.SaveNode(ctx, &request)
	return resp.GetValue(), err
}

// GetNode takes an id of a node and returns the node.
func (c *Client) GetNode(uuid string) (node grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.GetNode(ctx, &grpcapi.PrimitiveString{Value: uuid})
	if resp != nil {
		node = *resp
	}
	return
}

// GetNodes will get all child nodes for the node id it takes as
// an argument.
func (c *Client) GetNodes(parentID string) (nodes []grpcapi.Node, err error) {
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

// GetChildNodes will get all child nodes for the node id it takes as
// an argument.
func (c *Client) GetChildNodes(parentID string) (nodes []grpcapi.Node, err error) {
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

// DeleteNode will remove the node of the node id it takes as
// an argument.
func (c *Client) DeleteNode(request grpcapi.DeleteNodeInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := c.api.DeleteNode(ctx, &request)
	return err
}

// GetParentNode will return the parent of the node id it takes as
// an argument.
func (c *Client) GetParentNode(nodeID string) (node grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.api.GetParentNode(ctx, &grpcapi.PrimitiveString{Value: nodeID})
	if resp != nil {
		node = *resp
	}
	return
}

// GetAncestors will return all ancestors to the top of the node id
// it takes as an argument.
func (c *Client) GetAncestors(nodeID string) (nodes []grpcapi.AncestorNode, err error) {
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

// GetEvents will return all events that has occured in the Hierarchy
// Management Service.
func (c *Client) GetEvents(since int, limit *int32) (events []eventsource.Record, err error) {
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
