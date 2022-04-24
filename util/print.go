package util

import (
	"fmt"
	"patricpuola/gocart/config"
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

func PrintVerbose(print string) {
	if config.IsVerbose() {
		fmt.Print(print)
	}
}

func PrintVeryVerbose(print string) {
	if config.IsVeryVerbose() {
		fmt.Print(print)
	}
}
