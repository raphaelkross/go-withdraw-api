package helpers

import (
	"reflect"
	"testing"
)

func TestSplitAmountIntoBills(t *testing.T) {
	amount := 187

	output := map[int]int{
		50: 3,
		10: 3,
		5:  1,
		1:  2,
	}

	result := SplitAmountIntoBills(amount)

	if !reflect.DeepEqual(output, result) {
		t.Errorf("Amount 1,000 failed. Expected %v, received %v", output, result)
	}
}
