package utils

import (
	"fmt"
	"strconv"
)

func ToInt(value any) (i int) {
	var err error

	switch v := value.(type) {
	case string:
		if v == "" {
			return 0
		}
		i, err = strconv.Atoi(v)
	case int:
		i = v
	default:
		panic(fmt.Errorf("unhandled type in integer conversion: %w", err))
	}
	if err != nil {
		panic(fmt.Errorf("error converting %v to integer: %w", value, err))
	}

	return i
}
