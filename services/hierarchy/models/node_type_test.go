package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateType_Success(t *testing.T) {
	assert.Nil(t, NodeTypeCompany.Validate())
	assert.Nil(t, NodeType("company").Validate())
	assert.Nil(t, NodeType("asset").Validate())
}

func Test_ValidateType_Failure(t *testing.T) {
	assert.NotNil(t, NodeType("Company").Validate())
	assert.NotNil(t, NodeType("AsSeT").Validate())
	assert.NotNil(t, NodeType("").Validate())
	assert.NotNil(t, NodeType("Batman").Validate())
}

func Test_ChildOfType_Success(t *testing.T) {
	assert.True(t, NodeTypeCompany.IsChildOf(NodeTypeCompany))
	assert.True(t, NodeTypeAsset.IsChildOf(NodeTypeFunctionalLocation))
	assert.True(t, NodeTypeSystem.IsChildOf(NodeTypePlant))
}

func Test_ChildOfType_Failure(t *testing.T) {
	assert.False(t, NodeTypeCompany.IsChildOf(NodeTypeAsset))
	assert.False(t, NodeTypeFunctionalLocation.IsChildOf(NodeTypeAsset))
	assert.False(t, NodeType("Batman").IsChildOf(NodeTypePlant))
}
