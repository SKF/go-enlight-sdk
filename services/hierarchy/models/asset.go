package models

import (
	"fmt"
)

type AssetNode struct {
	Criticality Criticality `json:"criticality"`
}

type Criticality string

const (
	CriticalityA Criticality = "criticality_a"
	CriticalityB Criticality = "criticality_b"
	CriticalityC Criticality = "criticality_c"
)

var ciriticalities = []Criticality{
	CriticalityA, CriticalityB, CriticalityC,
}

func (asset AssetNode) Validate() error {
	if asset.Criticality == "" {
		return fmt.Errorf("Criticality field on asset cannot be empty")
	}
	return nil
}

func (cr Criticality) ValidateCriticaltiy() error {
	for _, ctrl := range ciriticalities {
		if cr == ctrl {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid asset criticality", cr)
}
