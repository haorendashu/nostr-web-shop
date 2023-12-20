package utils

import (
	"github.com/google/uuid"
	"strings"
)

func RandomId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
