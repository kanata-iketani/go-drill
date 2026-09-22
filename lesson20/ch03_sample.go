// lesson20 ch03: 可変長引数 ...int（*args の代わり）
package main

import "fmt"

// nums ...int と書くと、関数の中では nums は []int として扱えます
func total(nums ...int) int {
	t := 0
	for _, n := range nums {
		t += n
	}
	return t
}

func main() {
	// 好きな個数で呼び出せます
	fmt.Println(total(1, 2, 3))
	fmt.Println(total(10, 20))
	fmt.Println(total()) // 0 個でもよい

	// スライスを渡すときは後ろに ... を付けてばらします
	nums := []int{5, 6, 7}
	fmt.Println(total(nums...))
}
