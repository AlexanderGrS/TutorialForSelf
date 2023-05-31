package kata

func FindOdd(seq []int) int {
	fcount, count := 0, 0
	for i, value := range seq {
		count = value
		for j, value1 := range seq {
			if value == value1 && count != value && value != 0 && j != i {
				count = value
				seq[j] = 0
			} else if value == value1 && count == value && value != 0 && j != i {
				count = 0
				seq[j] = 0
			}
		}
		if count != 0 {
			fcount = count
		}
	}
	return fcount
}
