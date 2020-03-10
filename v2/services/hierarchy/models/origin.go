package models

import (
	"fmt"

	"github.com/SKF/proto/common"
)

type Origin struct {
	// Origin identity
	ID string `json:"id" example:"d932a2f2-bd5e-4803-831a-0f32d50c5b8e"`
	// Origin type
	Type string `json:"type" example:"TREEELEM"`
	// Origin provider
	Provider string `json:"provider" example:"8dc6763c-4eaf-4330-914d-56486ebfd68e"`
}

func (o Origin) Validate() (err error) {
	if o.ID == "" {
		return fmt.Errorf("Required field 'id' cannot be empty")
	}
	if o.Type == "" {
		return fmt.Errorf("Required field 'type' cannot be empty")
	}
	if o.Provider == "" {
		return fmt.Errorf("Required field 'provider' cannot be empty")
	}
	return
}

func (o Origin) ToGRPC() *common.Origin {
	return &common.Origin{
		Id:       o.ID,
		Type:     o.Type,
		Provider: o.Provider,
	}
}

func (o *Origin) FromGRPC(origin common.Origin) {
	o.ID = origin.Id
	o.Type = origin.Type
	o.Provider = origin.Provider
}
