package util

import (
	"bytes"
	"encoding/json"
)

func CreateRequestBody(requestBody interface{}) (*bytes.Buffer, error) {
	rawRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(rawRequestBody), nil
}
