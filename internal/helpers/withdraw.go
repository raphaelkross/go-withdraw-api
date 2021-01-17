package helpers

// SplitAmountIntoBills - Split amount into the minimum number of bills.
func SplitAmountIntoBills(amount int) map[int]int {

	// Default output where the number of bills are stored.
	output := map[int]int{
		50: 0,
		10: 0,
		5:  0,
		1:  0,
	}

	// Available bills.
	bills := [...]int{
		50,
		10,
		5,
		1,
	}

	// For each bill...
	for _, bill := range bills {

		// While amount is bigger than current bill value...
		for amount >= bill {
			// Reduce bill val from amount...
			amount -= bill

			// And... add one to bill's count.
			output[bill] = output[bill] + 1
		}
	}

	return output
}
