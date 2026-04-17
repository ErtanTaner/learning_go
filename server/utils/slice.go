package utils

import (
	"slices"
	"example.com/net-test/types"
)

func FindIdx(t []types.Todo, id string) int {
	return slices.IndexFunc(t, func(el types.Todo) bool {
		return el.ID == id
	})
}
