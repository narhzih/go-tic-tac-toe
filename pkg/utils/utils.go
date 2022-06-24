package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitStringToInt(stringToSplit string) (int, int, error) {
	values := strings.Split(stringToSplit, ":")
	if len(values) > 0 {
		value1, err := strconv.Atoi(values[0])
		value2, err := strconv.Atoi(values[1])
		if err != nil {
			return 0, 0, err
		}
		return value1, value2, nil
	}
	return 0, 0, fmt.Errorf("invalid values")
}
