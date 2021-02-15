package main

func Sum(numbers [5]int) (total int) {
	for _, value := range numbers {
		total += value
	}
	return
}
