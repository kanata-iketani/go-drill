// lesson07 ch01 解答
package main

import "fmt"

func main() {
	// [1, 2, 3] に当たる書き方。[]int が「int のスライス」という型名です。
	nums := []int{10, 20, 30}

	// fmt.Println にスライスを渡すと [10 20 30] の形式で表示されます。
	fmt.Println(nums)
	fmt.Println(len(nums))
}
