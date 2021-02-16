package main

func Sum(numbers []int) (total int) {
	for _, value := range numbers {
		total += value
	}
	return
}

func SumAll(allSlices ...[]int) (sliceOfSums []int) {
	numberOfSlices := len(allSlices)
	sliceOfSums = make([]int, numberOfSlices)
	for i, oneSlice := range allSlices{
		sliceOfSums[i] = Sum(oneSlice)
	}
	return
}
