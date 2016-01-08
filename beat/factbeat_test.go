package factbeat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var deDotFixture = map[string]interface{}{
	"key.name": true,
	"parent": map[string]interface{}{
		"dotted.child": true,
	},
}

func TestDeDotTranslatesKeys(t *testing.T) {
	assert.True(t, deDot(deDotFixture)["key_name"].(bool))
}

func TestDeDotRemovesDottedKeys(t *testing.T) {
	assert.Zero(t, deDot(deDotFixture)["key.name"])
}

func TestDeDotTranlsatesNestedKeys(t *testing.T) {
	parent := deDot(deDotFixture)["parent"]
	child := parent.(map[string]interface{})["dotted_child"]
	assert.True(t, child.(bool))
}

func TestDeDotRemovesDottedNestedKeys(t *testing.T) {
	parent := deDot(deDotFixture)["parent"]
	assert.Zero(t, parent.(map[string]interface{})["dotted.child"])
}
