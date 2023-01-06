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
	t.Run("should pass for case 371", func(t *testing.T) {
		testCase := Testcase{
			value:    371,
			expected: true,
		}
		testCase.actual = golangunittesting.CalculateIsArmstrong(testCase.value)
		if testCase.expected != testCase.actual {
			t.Fail()
		}

	})
	t.Run("should pass for case 370", func(t *testing.T) {
		testCase := Testcase{
			value:    370,
			expected: true,
		}
		testCase.actual = golangunittesting.CalculateIsArmstrong(testCase.value)
		if testCase.expected != testCase.actual {
			t.Fail()
		}

	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("should fail for case 350", func(t *testing.T) {
		testCase := Testcase{
			value:    350,
			expected: false,
		}
		testCase.actual = golangunittesting.CalculateIsArmstrong(testCase.value)
		if testCase.expected != testCase.actual {
			t.Fail()
		}

	})
	t.Run("should fail for case 300", func(t *testing.T) {
		testCase := Testcase{
			value:    300,
			expected: false,
		}
		testCase.actual = golangunittesting.CalculateIsArmstrong(testCase.value)
		if testCase.expected != testCase.actual {
			t.Fail()
		}

	})
}
