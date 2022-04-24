package config_test

import (
	"fmt"
	"patricpuola/gocart/config"
	"strconv"
	"testing"
)

type VerbosityTestCase struct {
	verbosityLevel int
	isVerbose      bool
	isVeryVerbose  bool
}

var verbosityHappyCases = []VerbosityTestCase{
	{config.NOT_VERBOSE, false, false},
	{config.VERBOSE, true, false},
	{config.VERY_VERBOSE, true, true},
}

func TestVerbosity(t *testing.T) {
	for _, test := range verbosityHappyCases {
		config.SetVerbosity(test.verbosityLevel)
		if config.IsVerbose() != test.isVerbose {
			t.Error(fmt.Sprintf("With config.SetVerbosity(%d) config.IsVerbose() returns %s, expected %s", test.verbosityLevel, strconv.FormatBool(test.isVerbose), strconv.FormatBool(test.isVeryVerbose)))
		}
		if config.IsVeryVerbose() != test.isVeryVerbose {
			t.Error(fmt.Sprintf("With config.SetVerbosity(%d) config.IsVeryVerbose() returns %s, expected %s", test.verbosityLevel, strconv.FormatBool(test.isVerbose), strconv.FormatBool(test.isVeryVerbose)))
		}
	}
}
