package service

import (
	"fmt"
	"strings"
)

func ValidateBlock(hash string, difficulty int, characterToFind string) bool {
	prefix := strings.Repeat(characterToFind, difficulty)
	val := strings.HasPrefix(hash, prefix)
	fmt.Print(val)
	return val
}
