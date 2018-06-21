package models

import (
	"fmt"
)

type Origin struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Provider string `json:"provider"`
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
