package utils

import (
	"strconv"
	"strings"
)

func ParseID(path, prefix string) (int, error) {
	idStr := strings.TrimPrefix(path, prefix)
	return strconv.Atoi(idStr)
}
