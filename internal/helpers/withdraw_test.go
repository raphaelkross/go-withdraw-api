package helpers

import (
	"reflect"
	"testing"
)

func TestSplitAmountIntoBills(t *testing.T) {
	amount := 187

	output := map[int]int{
		100: 1,
		50:  1,
		20:  1,
		10:  1,
		5:   1,
		1:   2,
	}

	result := SplitAmountIntoBills(amount)

	if !reflect.DeepEqual(output, result) {
		t.Errorf("Amount 1,000 failed. Expected %v, received %v", output, result)
	}
}
