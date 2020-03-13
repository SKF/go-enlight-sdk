package models

import (
	"fmt"

	grpcapi "github.com/SKF/proto/hierarchy"
)

type AssetNode struct {
	Criticality   Criticality `json:"criticality" example:"criticality_a"`
	AssetType     string      `json:"assetType,omitempty" example:"CO"`
	AssetClass    string      `json:"assetClass,omitempty" example:"AX"`
	AssetSequence string      `json:"assetSequence,omitempty" example:"02"`
	Manufacturer  string      `json:"manufacturer,omitempty" example:"Atlas Copco"`
	Model         string      `json:"model,omitempty" example:"SF"`
	SerialNumber  string      `json:"serialNumber,omitempty"`
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
	if err := asset.Criticality.ValidateCriticaltiy(); err != nil {
		return err
	}
	if asset.AssetType != "" && asset.AssetClass == "" {
		return fmt.Errorf("cannot set asset type without specifying class")
	} else if asset.AssetSequence != "" && (asset.AssetType == "" || asset.AssetClass == "") {
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
	ret := grpcapi.AssetNode{
		Criticality:  string(asset.Criticality),
		Class:        asset.AssetClass,
		Type:         asset.AssetType,
		Sequence:     asset.AssetSequence,
		Manufacturer: asset.Manufacturer,
		Model:        asset.Model,
		SerialNumber: asset.SerialNumber,
	}
	return &ret
}

func (asset *AssetNode) FromGRPC(assetNode grpcapi.AssetNode) {
	asset.Criticality = Criticality(assetNode.Criticality)
	asset.AssetClass = assetNode.Class
	asset.AssetType = assetNode.Type
	asset.AssetSequence = assetNode.Sequence
	asset.Manufacturer = assetNode.Manufacturer
	asset.Model = assetNode.Model
	asset.SerialNumber = assetNode.SerialNumber
}
