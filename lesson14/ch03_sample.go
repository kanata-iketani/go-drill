// lesson14 ch03: 関数に切り出すのが Go 流
package main

import "fmt"

// squares は各要素を 2 乗した新しいスライスを返す
func squares(nums []int) []int {
	result := []int{}
	for _, x := range nums {
		result = append(result, x*x)
	}
	return result
}

// bigOnly は 10 以上の値だけを集めた新しいスライスを返す
func bigOnly(nums []int) []int {
	result := []int{}
	for _, x := range nums {
		if x >= 10 {
			result = append(result, x)
		}
	}
	return result
}

func main() {
	nums := []int{3, 12, 5, 20}
	// 呼び出す側は内包表記と同じ 1 行で読めます
	fmt.Println(squares(nums)) // [9 144 25 400]
	fmt.Println(bigOnly(nums)) // [12 20]
}
