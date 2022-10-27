package project

import "github.com/oklog/ulid/v2"

func GetNewId() string {
	return ulid.Make().String()
}
