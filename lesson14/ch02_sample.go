// lesson14 ch02: 条件付き内包表記の代替
package main

import "fmt"

func main() {
	nums := []int{7, 2, 9, 4, 11, 6}

	// Python: evens = [x for x in nums if x % 2 == 0]
	evens := []int{}
	for _, x := range nums {
		if x%2 == 0 { // 条件を満たしたときだけ append
			evens = append(evens, x)
		}
	}
	fmt.Println(evens) // [2 4 6]

	// フィルタしながら変換もできます: 奇数だけ 10 倍
	oddsTen := []int{}
	for _, x := range nums {
		if x%2 == 1 {
			oddsTen = append(oddsTen, x*10)
		}
	}
	fmt.Println(oddsTen) // [70 90 110]
}
