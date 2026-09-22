// lesson07 ch01: スライスの生成と len
package main

import "fmt"

func main() {
	// 要素の型を付けて作ります。int のスライスには int しか入りません
	nums := []int{10, 20, 30}
	fmt.Println(nums)
	fmt.Println(len(nums))

	// 文字列のスライスも同じ形です
	fruits := []string{"りんご", "みかん"}
	fmt.Println(fruits)

	// 空のスライスは要素 0 個から始まります
	var empty []int
	fmt.Println(len(empty))
}
