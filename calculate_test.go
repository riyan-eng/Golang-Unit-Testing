package golangunittesting_test

import (
	"testing"

	golangunittesting "github.com/riyan-eng/Golang-Unit-Testing"
)

type Testcase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	testCase := Testcase{
		value:    371,
		expected: true,
	}
	testCase.actual = golangunittesting.CalculateIsArmstrong(testCase.value)
	if testCase.expected != testCase.actual {
		t.Fail()
	}
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	testCase := Testcase{
		value:    350,
		expected: false,
	}
	testCase.actual = golangunittesting.CalculateIsArmstrong(testCase.value)
	if testCase.expected != testCase.actual {
		t.Fail()
	}
}
