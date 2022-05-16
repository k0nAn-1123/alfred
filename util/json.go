package util

import "encoding/json"

func JsonMarshal(item interface{}) (string, error) {
	if bytes, err := json.Marshal(item); err != nil {
		return "", err
	} else {
		str := string(bytes)
		return str, nil
	}
}

func JsonUnmarshal(str string, item *interface{}) error {
	bytes := []byte(str)
	if err := json.Unmarshal(bytes, item); err != nil {
		return err
	} else {
		return nil
	}
}
