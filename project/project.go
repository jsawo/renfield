package project

import (
	"github.com/google/uuid"
)

func GetNewId() string {
	return uuid.NewString()
}
