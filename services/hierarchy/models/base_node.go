package models

import (
	"fmt"

	"github.com/SKF/go-utility/uuid"
)

// Node represent a hierarchy node
type BaseNode struct {
	ID          uuid.UUID   `json:"id"`
	Type        NodeType    `json:"nodeType"`
	SubType     NodeSubType `json:"nodeSubType"`
	Origin      *Origin     `json:"origin,omitempty"`
	Label       string      `json:"label"`
	Description string      `json:"description"`
}

// EnlightRootNodeUUID is the base/root node that all other nodes attach to, must not be deleted or updated
const EnlightRootNodeUUID uuid.UUID = uuid.UUID("df3214a6-2db7-11e8-b467-0ed5f89f718b")

func (node BaseNode) Validate() error {
	if node.Label == "" {
		return fmt.Errorf("Required field 'label' cannot be empty")
	}

	if err := node.Type.Validate(); err != nil {
		return fmt.Errorf("Required field 'type' is invalid: %+v", err)
	}

	if err := node.SubType.Validate(); err != nil {
		return fmt.Errorf("Required field 'subtype' is invalid: %+v", err)
	}

	if !node.SubType.IsTypeOf(node.Type) {
		return fmt.Errorf("'%s' doesnt not have a subtype '%s'", node.Type, node.SubType)
	}

	if node.Origin != nil {
		if err := node.Origin.Validate(); err != nil {
			return fmt.Errorf("Required field 'origin' is invalid: %+v", err)
		}
	}

	return nil
}
