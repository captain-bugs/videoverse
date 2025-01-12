package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	id, err := uuid.NewV7()
	if err != nil {
		return ""
	}
	return id.String()
}
