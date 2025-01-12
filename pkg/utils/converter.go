package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
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

func ConvertSecondsToDuration(seconds float64) string {

	intPart := int(seconds)
	//milliseconds := int((seconds - float64(intPart)) * 1000)

	duration := time.Duration(intPart) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60

	// Format the string as HH:MM:SS.mmm
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs)
}

func StringToMap(s string) map[string]any {
	var newMap map[string]any
	err := json.Unmarshal([]byte(s), &newMap)
	if err != nil {
		logbox.NewLogBox().Error().Err(err).Msg("error in json unmarshal")
	}
	return newMap
}

func MapToStruct[M any, S any](m M) S {
	str, err := json.Marshal(m)
	if err != nil {
		logbox.NewLogBox().Error().Err(err).Msg("error in json marshal")
	}
	var structure S
	if err := json.Unmarshal(str, &structure); err != nil {
		logbox.NewLogBox().Error().Err(err).Msg("error in json unmarshal")
	}
	return structure
}

func ToFlatCase(s string) string {
	x := strings.ToLower(s)
	return strings.ReplaceAll(x, " ", "_")
}
