package array

func Sum(array []int) int{
	var sum int

	for _, number := range array{
		sum += number
	}
	
	return sum
}

