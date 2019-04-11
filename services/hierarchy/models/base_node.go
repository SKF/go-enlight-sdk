package models

import (
	"github.com/SKF/go-utility/uuid"
)

// Node represent a hierarchy node
type BaseNode struct {
	ID          uuid.UUID        `json:"id"`
	ParentID    uuid.UUID        `json:"parentId"`
	Type        NodeType         `json:"nodeType"`
	SubType     NodeSubType      `json:"nodeSubType"`
	Industry    *IndustrySegment `json:"industrySegment,omitempty"`
	Origin      *Origin          `json:"origin,omitempty"`
	Label       string           `json:"label"`
	Description string           `json:"description"`
	Position    *int64           `json:"position"`
	Tags        *string          `json:"tags"`
}
