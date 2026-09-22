// lesson14 ch02 解答
package main

import "fmt"

func main() {
	nums := []int{12, 5, 8, 20, 3, 16, 7}

	// Python の [x for x in nums if x >= 10] に相当します。
	// for で全要素を回し、if で条件を満たすものだけ append します。
	// フィルタは「空スライス + for + if + append」の 4 点セットです。
	big := []int{}
	for _, x := range nums {
		if x >= 10 {
			big = append(big, x)
		}
	}
	fmt.Println(big)
}
