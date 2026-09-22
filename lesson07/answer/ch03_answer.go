// lesson07 ch03 解答
package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}

	// s[i:j] は「i 番目から j-1 番目まで」。Python と同じ半開区間です。
	fmt.Println(nums[1:4])

	// 始点を省略すると先頭から、終点を省略すると末尾までになります。
	fmt.Println(nums[:2])
	fmt.Println(nums[3:])
}
