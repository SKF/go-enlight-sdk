package models

import (
	"fmt"

	grpcapi "github.com/SKF/proto/hierarchy"
)

type AssetNode struct {
	Criticality Criticality `json:"criticality"`
	Type        string      `json:"type,omitempty"`
	Class       string      `json:"class,omitempty"`
	Sequence    string      `json:"sequence,omitempty"`
}

type Criticality string

const (
	CriticalityA       Criticality = "criticality_a"
	CriticalityB       Criticality = "criticality_b"
	CriticalityC       Criticality = "criticality_c"
	CriticalityUnknown Criticality = "criticality_unknown"
)

var ciriticalities = []Criticality{
	CriticalityA, CriticalityB, CriticalityC,
}

func ParseCriticality(dbCriticality string) Criticality {
	switch dbCriticality {
	case "criticality_a":
		return CriticalityA
	case "criticality_b":
		return CriticalityB
	case "criticality_c":
		return CriticalityC
	default:
		return CriticalityUnknown
	}
}

func (asset AssetNode) Validate() error {
	if asset.Criticality == "" {
		return fmt.Errorf("criticality field on asset cannot be empty")
	}
	if asset.Type != "" && asset.Class == "" {
		return fmt.Errorf("cannot set asset type without specifying class")
	} else if asset.Sequence != "" && (asset.Type == "" || asset.Class == "") {
		return fmt.Errorf("cannot set asset sequnce without specifying class and type")
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

func (asset AssetNode) ToGRPC() *grpcapi.AssetNode {
	return &grpcapi.AssetNode{
		Criticality: string(asset.Criticality),
		Class:       asset.Class,
		Type:        asset.Type,
		Sequence:    asset.Sequence,
	}
}

func (asset *AssetNode) FromGRPC(assetNode grpcapi.AssetNode) {
	asset.Criticality = Criticality(assetNode.Criticality)
	asset.Class = assetNode.Class
	asset.Type = assetNode.Type
	asset.Sequence = assetNode.Sequence
}
