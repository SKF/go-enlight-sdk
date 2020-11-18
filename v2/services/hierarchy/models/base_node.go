package models

import (
	"fmt"

	"github.com/SKF/go-utility/v2/uuid"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

// Node represent a hierarchy node
type BaseNode struct {
	// ID of node, as a UUID
	ID uuid.UUID `json:"id" swaggertype:"string" format:"uuid" example:"7bcd1711-21bd-4eb7-8349-b053d6d5226f"`
	// ID of parent node, as a UUID
	ParentID uuid.UUID `json:"parentId" swaggertype:"string" format:"uuid" example:"8f7551c5-3357-406d-ab82-bcb138d0b13f"`
	// Type of node
	Type NodeType `json:"nodeType" swaggertype:"string" example:"asset" enums:"root,company,site,plant,system,functional_location,asset,measurement_point,inspection_point,lubrication_point,unknown"`
	// Type of node
	SubType NodeSubType `json:"nodeSubType" swaggertype:"string" example:"asset" enums:"root,company,site,plant,ship,system,functional_location,asset,measurement_point,inspection_point,lubrication_point"`
	// Industry segment of this node
	Industry *IndustrySegment `json:"industrySegment,omitempty" swaggertype:"string" example:"metal" enums:"agriculture,construction,food_and_beverage,hydrocarbon_processing,machine_tool,marine,metal,mining,power_generation,pulp_and_paper,renewable,undefined"`
	// Origin of node, if imported from another system
	Origin *Origin `json:"origin,omitempty"`
	// Descriptive name of the node
	Label string `json:"label" example:"01AA DE"`
	// Description of the node
	Description string `json:"description" example:"First bearing, driven end"`
	// Relative position of node in the Enlight Centre UI
	Position *int64 `json:"position"`
	// Comma separated list of free form tags on this node
	Tags *string `json:"tags" example:"tag1,tag2=value2"`
	// Which country the node is in
	Country *string `json:"country,omitempty" example:"SWE"`
	// Metadata with keys and optional values
	MetaData *NodeMetaData `json:"metaData,omitempty"`
}

// EnlightRootNodeUUID is the base/root node that all other nodes attach to, must not be deleted or updated
const EnlightRootNodeUUID uuid.UUID = uuid.UUID("df3214a6-2db7-11e8-b467-0ed5f89f718b")

func (node BaseNode) Validate() error {
	if !node.ID.IsValid() {
		return fmt.Errorf("Required field 'id' contains an invalid UUID")
	}
	if node.ParentID != "" && !node.ParentID.IsValid() {
		return fmt.Errorf("Required field 'parentId' contains an invalid UUID")
	}
	if node.Label == "" {
		return fmt.Errorf("Required field 'label' cannot be empty")
	}
	if node.Position != nil && *node.Position < 0 {
		return fmt.Errorf("Field 'position' cannot be negative")
	}
	if node.Country != nil {
		_, err := language.ParseRegion(*node.Country)
		return errors.Wrapf(err, "Field 'country' contains illegal value %s", *node.Country)
	}

	if node.Industry != nil {
		if err := node.Industry.Validate(); err != nil {
			return errors.Wrap(err, "Optional field 'industrySegment' is invalid")
		}
	}

	if err := node.Type.Validate(); err != nil {
		return errors.Wrap(err, "Required field 'nodeType' is invalid")
	}

	if err := node.SubType.Validate(); err != nil {
		return errors.Wrap(err, "Required field 'nodeSubType' is invalid")
	}

	if !node.SubType.IsTypeOf(node.Type) {
		return fmt.Errorf("Node type '%s' doesn't have a subtype '%s'", node.Type, node.SubType)
	}

	if node.Origin != nil {
		if err := node.Origin.Validate(); err != nil {
			return errors.Wrap(err, "Optional field 'origin' is invalid")
		}
	}

	if node.MetaData != nil {
		if err := node.MetaData.Validate(); err != nil {
			return errors.Wrap(err, "Optional field 'metadata' is invalid")
		}
	}

	return nil
}
