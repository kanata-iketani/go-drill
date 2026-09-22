// lesson12 ch02: for {} 無限ループと break
package main

import "fmt"

func main() {
	// for {} は while True: に相当する無限ループです
	count := 0
	for {
		count++
		fmt.Println("count =", count)
		if count == 3 {
			break // 条件を満たしたらループを抜けます
		}
	}
	fmt.Println("ループを抜けました")
}
