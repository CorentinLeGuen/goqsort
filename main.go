package main

import "fmt"

func main() {
	toSort := []int{1, 5, 4, 7, 98, 69, 3, 8, 5, 48, 45, 4, 8, 45, 4, 84, 7}
	qsort(toSort, 0, len(toSort)-1)
	fmt.Printf("%#v\n", toSort)
}

func qsort(fn []int, l int, h int) {
	if l < h {
		p := partition(fn, l, h)
		qsort(fn, l, p-1)
		qsort(fn, p+1, h)
	}
}

func partition(fn []int, l int, h int) int {
	t := fn[h]
	i := l
	for j := l; j < h; j++ {
		if fn[j] < t {
			fn[i], fn[j] = fn[j], fn[i]
			i++
		}
	}
	fn[i], fn[h] = fn[h], fn[i]
	return i
}
