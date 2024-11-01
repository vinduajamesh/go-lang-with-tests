package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, v := range numbers {
		sums = append(sums, Sum(v))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, v := range numbers {
		if len(v) == 0 {
			sums = append(sums, Sum(v))
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
