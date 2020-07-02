package models

import (
	"fmt"

	"github.com/SKF/go-utility/v2/uuid"
	"github.com/pkg/errors"
)

type Component struct {
	BaseComponent
	*ShaftComponent
	*BearingComponent
}

type ComponentType string

const (
	ShaftComponentType   ComponentType = "shaft"
	BearingComponentType ComponentType = "bearing"
)

var AllComponentTypes = []ComponentType{
	ShaftComponentType, BearingComponentType,
}

func (ct ComponentType) String() string {
	return string(ct)
}

type BaseComponent struct {
	ID         uuid.UUID     `json:"id" swaggertype:"string" format:"uuid" example:"7bcd1711-21bd-4eb7-8349-b053d6d5226f"`
	Type       ComponentType `json:"type" swaggertype:"string" example:"bearing" enums:"shaft,bearing"`
	AttachedTo *uuid.UUID    `json:"attachedTo,omitempty" swaggertype:"string" format:"uuid" example:"7bcd1711-21bd-4eb7-8349-b053d6d5226f"`
	Position   uint32        `json:"position"`
}

type ShaftComponent struct {
	FixedSpeed *float64 `json:"fixedSpeed,omitempty"`
}

type RotatingRing string

const (
	InnerRotatingRing RotatingRing = "inner"
	OuterRotatingRing RotatingRing = "outer"
	BothRotatingRing  RotatingRing = "both"
)

type BearingComponent struct {
	PositionDescription string       `json:"positionDescription"`
	Manufacturer        string       `json:"manufacturer"`
	Designation         string       `json:"designation"`
	SerialNumber        string       `json:"serialNumber"`
	ShaftSide           ShaftSide    `json:"shaftSide" swaggertype:"string" example:"nde" enums:"de,nde,unknown"`
	RotatingRing        RotatingRing `json:"rotatingRing" swaggertype:"string" example:"inner" enums:"inner,outer,both"`
}

func (component Component) Validate() (err error) {
	if err := component.BaseComponent.Validate(); err != nil {
		return err
	}
	switch component.Type {
	case ShaftComponentType:
		if component.AttachedTo != nil && *component.AttachedTo != "" {
			return errors.New("a shaft component can not be attached to another component")
		}
		if component.ShaftComponent != nil {
			if err := component.ShaftComponent.Validate(); err != nil {
				return err
			}
		}
	case BearingComponentType:
		if component.AttachedTo == nil || *component.AttachedTo == "" {
			return errors.New("a bearing component must be attached to a shaft")
		}
		if component.BearingComponent != nil {
			if err := component.BearingComponent.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (component BaseComponent) Validate() (err error) {
	if err := component.ID.Validate(); err != nil {
		return errors.Wrap(err, "required field 'id' is invalid")
	}
	if err := component.Type.Validate(); err != nil {
		return errors.Wrap(err, "required field 'type' is invalid")
	}
	if component.AttachedTo != nil && *component.AttachedTo != "" {
		if err := component.AttachedTo.Validate(); err != nil {
			return errors.Wrap(err, "optional field 'attachedTo' is invalid")
		}
	}
	if component.Position == 0 {
		return errors.New("required field 'position' is invalid")
	}
	return nil
}

func (ct ComponentType) Validate() (err error) {
	for _, componentType := range AllComponentTypes {
		if ct == componentType {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid component type", ct)
}

func (shaft ShaftComponent) Validate() (err error) {
	if shaft.FixedSpeed != nil && *shaft.FixedSpeed < 0.0 {
		return errors.New("fixedSpeed should be a positive number")
	}
	return nil
}

func (bearing BearingComponent) Validate() (err error) {
	if err := bearing.ShaftSide.Validate(); bearing.ShaftSide != "" && err != nil {
		return errors.Wrap(err, "optional field 'shaftSide' is invalid")
	}
	if err := bearing.RotatingRing.Validate(); bearing.RotatingRing != "" && err != nil {
		return errors.Wrap(err, "optional field 'rotatingRing' is invalid")
	}
	return nil
}

func (rr RotatingRing) Validate() (err error) {
	if rr != InnerRotatingRing && rr != OuterRotatingRing && rr != BothRotatingRing {
		return fmt.Errorf("'%s' is not a valid rotating ring", rr)
	}
	return nil
}
