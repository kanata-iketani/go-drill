// lesson14 ch03 解答
package main

import "fmt"

// double は各要素を 2 倍した新しいスライスを返す。
// Python の [x * 2 for x in nums] を関数に切り出した形で、
// 名前が付くことで呼び出す側は 1 行になり、意図も伝わりやすくなります。
func double(nums []int) []int {
	result := []int{}
	for _, x := range nums {
		result = append(result, x*2)
	}
	return result
}

func main() {
	fmt.Println(double([]int{1, 2, 3, 4, 5}))
}
