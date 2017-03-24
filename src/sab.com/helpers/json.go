package helpers

import (
	"encoding/json"
	"io"
)

func JsonToObject(jsonSource io.Reader, targetObject interface{}) {
	decoder := json.NewDecoder(jsonSource)
	decoder.Decode(targetObject)
}

func ObjecToJsonByteBuffer(sourceObject interface{}) ([]byte, error) {
	payl, err := json.Marshal(sourceObject)

	return payl, err
}
