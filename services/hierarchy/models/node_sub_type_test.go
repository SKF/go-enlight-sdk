package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateSubType_Success(t *testing.T) {
	assert.Nil(t, NodeSubTypeCompany.Validate())
	assert.Nil(t, NodeSubType("company").Validate())
	assert.Nil(t, NodeSubType("ship").Validate())
	assert.Nil(t, NodeSubType("asset").Validate())
}

func Test_ValidateSubType_Failure(t *testing.T) {
	assert.NotNil(t, NodeSubType("Company").Validate())
	assert.NotNil(t, NodeSubType("AsSeT").Validate())
	assert.NotNil(t, NodeSubType("").Validate())
	assert.NotNil(t, NodeSubType("Batman").Validate())
}

func Test_ChildOfSubType_Success(t *testing.T) {
	assert.True(t, NodeSubTypeShip.IsTypeOf(NodeTypePlant))
	assert.True(t, NodeSubTypePlant.IsTypeOf(NodeTypePlant))
	assert.True(t, NodeSubTypeCompany.IsTypeOf(NodeTypeCompany))
	assert.True(t, NodeSubTypeFunctionalLocation.IsTypeOf(NodeTypeFunctionalLocation))
	assert.True(t, NodeSubTypePlant.IsTypeOf(NodeTypePlant))
}

func Test_ChildOfSubType_Failure(t *testing.T) {
	assert.False(t, NodeSubTypeShip.IsTypeOf(NodeTypeCompany))
	assert.False(t, NodeSubTypeAsset.IsTypeOf(NodeTypeCompany))
	assert.False(t, NodeSubTypeAsset.IsTypeOf(NodeType("")))
	assert.False(t, NodeSubType("Batman").IsTypeOf(NodeTypeCompany))
}
