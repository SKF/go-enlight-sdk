package models

import (
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
