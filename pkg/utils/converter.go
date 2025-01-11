package utils

import (
	"encoding/json"
	"videoverse/pkg/logbox"
)

func ToMap(s interface{}) (newMap map[string]interface{}) {
	data, err := json.Marshal(s) // Convert to a json string
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap) // Convert to a map
	logbox.NewLogBox().Debug().Err(err).Msg("UNMARSHAL_ERROR")
	return
}
