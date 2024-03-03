package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		numDBs int
	}
	tests := []testCase{
		{1},
		{3},
		{4},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{0},
			{13},
		}...)
	}

	for _, test := range tests {
		dbChan := getDBsChannel(test.numDBs)
		waitForDBs(test.numDBs, dbChan)
		if len(dbChan) != 0 {
			t.Errorf(`Test Failed: (%v) ->
expected length: %v
actual length: %v
=======================
`,
				test.numDBs,
				0,
				len(dbChan),
			)
		} else {
			fmt.Printf(`Test Passed: (%v) ->
expected length: %v
actual length: %v
=======================
`,
				test.numDBs,
				0,
				len(dbChan),
			)
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
