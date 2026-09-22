// lesson07 ch04: append で追加
package main

import "fmt"

func main() {
	var nums []int // 空のスライスから始めます

	// append は「追加後のスライス」を返すので、必ず再代入します
	nums = append(nums, 10)
	nums = append(nums, 20)
	fmt.Println(nums)

	// 一度に複数追加することもできます
	nums = append(nums, 30, 40)
	fmt.Println(nums)
	fmt.Println(len(nums))
}
