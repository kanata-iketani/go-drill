// lesson11 ch02: && と || と !（and or not の代わり）
package main

import "fmt"

func main() {
	x := 7

	// and は && になります
	fmt.Println(x > 0 && x < 10)

	// or は || になります
	fmt.Println(x < 0 || x > 5)

	// not は ! になります
	fmt.Println(!(x == 7))

	// Python の 0 < x < 10 のような連結はできず、&& で分けます
	if x > 0 && x < 10 {
		fmt.Println("1桁の正の数です")
	}
}
