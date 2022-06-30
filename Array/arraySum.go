package array

func Sum(array []int) int{
	var sum int

	for _, number := range array{
		sum += number
	}
	
	return sum
}

func SumAll(numbersSum ...[]int) []int{
	var sums []int

	for _, numbers := range numbersSum{
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTail(numbersSumTails ...[]int) []int{
	var sums []int

	for _, numbers := range numbersSumTails{
		if len(numbers) == 0 || len(numbers) == 1{
			sums = append(sums, 0)
		}else{
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
		
		
		
	}

	return sums
}
