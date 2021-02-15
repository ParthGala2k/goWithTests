package main

func Sum(numbers []int) (total int) {
	for _, value := range numbers {
		total += value
	}
	return
}
