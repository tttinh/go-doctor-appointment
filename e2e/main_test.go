package e2e

import (
	"testing"
)

func TestABC(t *testing.T) {
	t.Run("TestStepOne", testStepOne)
	t.Run("TestStepTwo", testStepTwo)
	t.Run("TestStepThree", testStepThree)
}

func testStepOne(t *testing.T) {
	t.Log("Running Test Step One")
	// Your test logic for step one
}

func testStepTwo(t *testing.T) {
	t.Log("Running Test Step Two")
	// Your test logic for step two
}

func testStepThree(t *testing.T) {
	t.Log("Running Test Step Three")
	// Your test logic for step three
	t.Error("hic")
}
