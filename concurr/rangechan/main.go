package main

func concurrentFib(n int) []int {
	res := make([]int, 0)
	fibChan := make(chan int)
	go fibonacci(n, fibChan)
	for num := range fibChan {
		res = append(res, num)
	}
	return res
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
