package models

import (
	"encoding/json"
	"fmt"

	"github.com/SKF/go-utility/uuid"
	grpcapi "github.com/SKF/proto/hierarchy"
)

type Component struct {
	Id            uuid.UUID       `json:"id"`
	Type          string          `json:"type"`
	Props         json.RawMessage `json:"props,omitempty"`
	SubComponents []Component     `json:"subComponents,omitempty"`
}

type AssetNode struct {
	Criticality Criticality `json:"criticality"`
	Type        string      `json:"type,omitempty"`
	Class       string      `json:"class,omitempty"`
	Sequence    string      `json:"sequence,omitempty"`
	Components  []Component `json:"components,omitempty"`
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
	components := make([]*grpcapi.Component, 0)
	for _, c := range asset.Components {
		components = append(components, c.ToGRPC())
	}

	return &grpcapi.AssetNode{
		Criticality: string(asset.Criticality),
		Class:       asset.Class,
		Type:        asset.Type,
		Sequence:    asset.Sequence,
		Components:  components,
	}
}

func (asset *AssetNode) FromGRPC(assetNode grpcapi.AssetNode) {
	asset.Criticality = Criticality(assetNode.Criticality)
	asset.Class = assetNode.Class
	asset.Type = assetNode.Type
	asset.Sequence = assetNode.Sequence
	asset.Components = make([]Component, 0)
	if assetNode.Components != nil {
		for _, gc := range assetNode.Components {
			c := &Component{}
			c.FromGRPC(gc)
			asset.Components = append(asset.Components, *c)
		}
	}
}

func (component Component) ToGRPC() *grpcapi.Component {
	subComponents := make([]*grpcapi.Component, 0)
	if component.SubComponents != nil {
		for _, sc := range component.SubComponents {
			subComponents = append(subComponents, sc.ToGRPC())
		}
	}

	ret := &grpcapi.Component{
		Id:            component.Id.String(),
		Type:          component.Type,
		SubComponents: subComponents,
	}

	if component.Props != nil {
		if buf, err := component.Props.MarshalJSON(); err == nil {
			ret.Props = string(buf)
		}
	}

	return ret
}

func (component *Component) FromGRPC(c *grpcapi.Component) {
	component.Id = uuid.UUID(c.Id)
	component.Type = c.Type
	component.SubComponents = make([]Component, 0)
	component.Props.UnmarshalJSON([]byte(c.Props)) // nolint
	if c.SubComponents != nil {
		for _, gsc := range c.SubComponents {
			sc := &Component{}
			sc.FromGRPC(gsc)
			component.SubComponents = append(component.SubComponents, *sc)
		}
	}
}
