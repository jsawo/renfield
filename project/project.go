package project

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

func GetNewId() string {
	return strings.ToLower(ulid.Make().String())
}
