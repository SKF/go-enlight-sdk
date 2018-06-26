package models

import "github.com/SKF/go-utility/uuid"

// Relation represents a relation between two nodes
type Relation struct {
	ID     uuid.UUID `json:"id"`
	FromID uuid.UUID `json:"fromId"`
	ToID   uuid.UUID `json:"toId"`
}
