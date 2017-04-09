package helpers

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConvertToObjectFromJsonSource(t *testing.T) {
	type NamedObject struct {
		Name string
	}

	var object NamedObject

	jsonSource := `{"name": "Patate"}`

	JsonToObject(strings.NewReader(jsonSource), &object)

	assert.Equal(t, "Patate", object.Name)
}
