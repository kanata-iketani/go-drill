// lesson07 ch03: スライス式 s[i:j]
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// i 番目から j-1 番目まで。Python の a[1:4] と同じです
	fmt.Println(nums[1:4])

	// 始点・終点は省略できます
	fmt.Println(nums[:2])
	fmt.Println(nums[3:])

	// a[:-1]（末尾を除く）は len を使って書きます
	fmt.Println(nums[:len(nums)-1])
}
