package helpers

import (
	"encoding/json"
	"io"
)

func JsonToObject(jsonSource io.Reader, targetObject interface{}) error {
	decoder := json.NewDecoder(jsonSource)
	return decoder.Decode(targetObject)
}
