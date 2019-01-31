package hierarchy

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SKF/go-eventsource/eventsource"
	"github.com/SKF/proto/common"
	hierarchy_grpcapi "github.com/SKF/proto/hierarchy"
)

// SaveNode will add the node if it this not exist and otherwise
// create it.
func (c *Client) SaveNode(request hierarchy_grpcapi.SaveNodeInput) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.SaveNodeWithContext(ctx, request)
}

// SaveNode will add the node if it this not exist and otherwise
// create it.
func (c *Client) SaveNodeWithContext(ctx context.Context, request hierarchy_grpcapi.SaveNodeInput) (string, error) {
	resp, err := c.api.SaveNode(ctx, &request)
	return resp.GetValue(), err
}

// GetNode takes an id of a node and returns the node.
func (c *Client) GetNode(uuid string) (node hierarchy_grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetNodeWithContext(ctx, uuid)
}

// GetNodeWithContext takes an id of a node and returns the node.
func (c *Client) GetNodeWithContext(ctx context.Context, uuid string) (node hierarchy_grpcapi.Node, err error) {
	resp, err := c.api.GetNode(ctx, &common.PrimitiveString{Value: uuid})
	if resp != nil {
		node = *resp
	}
	return
}

// GetNodeIDByOrigin takes an origin and returns the Enlight ID.
func (c *Client) GetNodeIDByOrigin(origin common.Origin) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetNodeIDByOriginWithContext(ctx, origin)
}

// GetNodeIDByOriginWithContext takes an origin and returns the Enlight ID.
func (c *Client) GetNodeIDByOriginWithContext(ctx context.Context, origin common.Origin) (string, error) {
	resp, err := c.api.GetNodeIdByOrigin(ctx, &origin)
	return resp.GetValue(), err
}

// GetNodes will get all child nodes for the node id it takes as
// an argument.
func (c *Client) GetNodes(parentID string) (nodes []hierarchy_grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetNodesWithContext(ctx, parentID)
}

// GetNodesWithContext will get all child nodes for the node id it takes as
// an argument.
func (c *Client) GetNodesWithContext(ctx context.Context, parentID string) (nodes []hierarchy_grpcapi.Node, err error) {
	resp, err := c.api.GetNodes(ctx, &common.PrimitiveString{Value: parentID})
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
func (c *Client) GetChildNodes(parentID string) (nodes []hierarchy_grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetChildNodesWithContext(ctx, parentID)
}

// GetChildNodesWithContext will get all child nodes for the node id it takes as
// an argument.
func (c *Client) GetChildNodesWithContext(ctx context.Context, parentID string) (nodes []hierarchy_grpcapi.Node, err error) {
	resp, err := c.api.GetChildNodes(ctx, &common.PrimitiveString{Value: parentID})
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
func (c *Client) DeleteNode(request hierarchy_grpcapi.DeleteNodeInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.DeleteNodeWithContext(ctx, request)
}

// DeleteNodeWithContext will remove the node of the node id it takes as
// an argument.
func (c *Client) DeleteNodeWithContext(ctx context.Context, request hierarchy_grpcapi.DeleteNodeInput) error {
	_, err := c.api.DeleteNode(ctx, &request)
	return err
}

// GetParentNode will return the parent of the node id it takes as
// an argument.
func (c *Client) GetParentNode(nodeID string) (node hierarchy_grpcapi.Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetParentNodeWithContext(ctx, nodeID)
}

// GetParentNodeWithContext will return the parent of the node id it takes as
// an argument.
func (c *Client) GetParentNodeWithContext(ctx context.Context, nodeID string) (node hierarchy_grpcapi.Node, err error) {
	resp, err := c.api.GetParentNode(ctx, &common.PrimitiveString{Value: nodeID})
	if resp != nil {
		node = *resp
	}
	return
}

// GetAncestors will return all ancestors to the top of the node id
// it takes as an argument.
func (c *Client) GetAncestors(nodeID string) (nodes []hierarchy_grpcapi.AncestorNode, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetAncestorsWithContext(ctx, nodeID)
}

// GetAncestorsWithContext will return all ancestors to the top of the node id
// it takes as an argument.
func (c *Client) GetAncestorsWithContext(ctx context.Context, nodeID string) (nodes []hierarchy_grpcapi.AncestorNode, err error) {
	resp, err := c.api.GetAncestors(ctx, &hierarchy_grpcapi.GetAncestorsInput{NodeId: nodeID})
	if resp != nil {
		for _, node := range resp.Nodes {
			if node != nil {
				nodes = append(nodes, *node)
			}
		}
	}
	return
}

// GetEvents will return all events that has occurred in the Hierarchy
// Management Service.
func (c *Client) GetEvents(since int, limit *int32) (events []eventsource.Record, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetEventsWithContext(ctx, since, limit)
}

// GetEventsWithContext will return all events that has occurred in the Hierarchy
// Management Service.
func (c *Client) GetEventsWithContext(ctx context.Context, since int, limit *int32) (events []eventsource.Record, err error) {
	input := hierarchy_grpcapi.GetEventsInput{Since: int64(since)}
	if limit != nil {
		input.Limit = &common.PrimitiveInt32{Value: *limit}
	}

	output, err := c.api.GetEvents(ctx, &input)
	if err != nil {
		return
	}

	err = json.Unmarshal(output.Events, &events)
	return
}
