package json

import "encoding/json"

func DeserializeStatus(jsonString string) (Status, error) {
	var jsonParsed Status

	err := json.Unmarshal([]byte(jsonString), &jsonParsed)

	if err != nil {
		return Status{}, err
	}

	return jsonParsed, nil
}
