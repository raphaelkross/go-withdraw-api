package helpers

import (
	"reflect"
	"testing"
)

func TestSplitAmountIntoBills(t *testing.T) {
	amount := 1050

	output := map[int]int{
		100: 10,
		50:  0,
		20:  0,
		10:  0,
		5:   0,
		1:   0,
	}

	result := SplitAmountIntoBills(amount)

	if !reflect.DeepEqual(output, result) {
		t.Errorf("Amount 1,000 failed. Expected %v, received %v", output, result)
	}
}
