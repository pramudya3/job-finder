package utils

import "encoding/json"

func Copy(dest, src interface{}) error {
	dataByte, err := json.Marshal(dest)
	if err != nil {
		return err
	}

	return json.Unmarshal(dataByte, &src)
}
