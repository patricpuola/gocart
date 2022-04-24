package util

import (
	"fmt"
	"strings"
)

func PrettyPrintMap(stringMap map[string]string, indent string) {
	var maxLenKey int
	for k := range stringMap {
		if len(k) > maxLenKey {
			maxLenKey = len(k)
		}
	}

	for k, v := range stringMap {
		fmt.Println(indent + k + ": " + strings.Repeat(" ", maxLenKey-len(k)) + v)
	}
}
