package models

import (
	"encoding/json"
	"fmt"

	"github.com/SKF/go-utility/v2/uuid"
	grpcapi "github.com/SKF/proto/v2/hierarchy"
)

type Component struct {
	Id            uuid.UUID       `json:"id"`
	Type          string          `json:"type"`
	Props         json.RawMessage `json:"props"`
	SubComponents []*Component    `json:"subComponents"`
}

type AssetNode struct {
	Criticality   Criticality  `json:"criticality" example:"criticality_a"`
	AssetType     string       `json:"assetType,omitempty" example:"CO"`
	AssetClass    string       `json:"assetClass,omitempty" example:"AX"`
	AssetSequence string       `json:"assetSequence,omitempty" example:"02"`
	Manufacturer  string       `json:"manufacturer,omitempty" example:"Atlas Copco"`
	Model         string       `json:"model,omitempty" example:"SF"`
	SerialNumber  string       `json:"serialNumber,omitempty"`
	Components    []*Component `json:"components,omitempty"`
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
	for _, c := range asset.Components {
		if err := c.Validate(); err != nil {
			return err
		}
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

func (c *Component) Validate() error {
	if err := c.Id.Validate(); err != nil {
		return err
	}
	if c.Type == "" {
		return fmt.Errorf("Component type cannot be empty")
	}
	for _, sc := range c.SubComponents {
		if err := sc.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	for _, cmp := range asset.Components {
		ret.Components = append(ret.Components, cmp.ToGRPC())
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
	for _, cmp := range assetNode.Components {
		c := &Component{}
		c.FromGRPC(cmp)
		asset.Components = append(asset.Components, c)
	}
}

func (cmp *Component) ToGRPC() *grpcapi.Component {
	ret := &grpcapi.Component{
		Id:   cmp.Id.String(),
		Type: cmp.Type,
	}

	if cmp.Props != nil {
		if buf, err := cmp.Props.MarshalJSON(); err == nil {
			ret.Props = string(buf)
		}
	}

	for _, c := range cmp.SubComponents {
		ret.SubComponents = append(ret.SubComponents, c.ToGRPC())
	}

	return ret
}

func (cmp *Component) FromGRPC(c *grpcapi.Component) {
	cmp.Id = uuid.UUID(c.Id)
	cmp.Type = c.Type
	cmp.Props.UnmarshalJSON([]byte(c.Props)) // nolint: errcheck

	for _, sc := range c.SubComponents {
		x := &Component{}
		x.FromGRPC(sc)
		cmp.SubComponents = append(cmp.SubComponents, x)
	}
}
