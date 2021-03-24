package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, item := range numbersToSum {
	sums = append(sums, Sum(item))
	}
	return
}


func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, item := range numbersToSum {
		t := []int{0}
		if len(item) > 1 {
			t = item[1:]
		}
		sums = append(sums, Sum(t))
	}
	return
}
