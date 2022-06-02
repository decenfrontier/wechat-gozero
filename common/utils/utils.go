package utils

import "github.com/google/uuid"

func GenUuid() string {
	uuid := uuid.New()
	key := uuid.String()
	return key
}