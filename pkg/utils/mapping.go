package utils

import (
	"bytes"
	"encoding/json"
)

// Mapping from Map[string]interface{} to interface{}/struct{}
func Mapping(in, out interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		return err
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		return err
	}
	return nil
}

func CompareMap(m1 any, m2 any) (bool, string, error) {
	n2, err := json.Marshal(m2)
	if err != nil {
		return false, string(n2), err
	}

	n1, err := json.Marshal(m1)
	if err != nil {
		return false, string(n2), err
	}

	// fmt.Println("n1", string(n1))
	// fmt.Println("n2", string(n2))

	return string(n1) != string(n2), string(n2), nil
}
