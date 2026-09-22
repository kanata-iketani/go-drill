// lesson10 ch02: else if と else（elif の代わり）
package main

import "fmt"

func main() {
	score := 72

	// elif は else if。直前の } と同じ行に書きます
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 60 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	temp := 31
	if temp >= 30 {
		fmt.Println("暑い")
	} else if temp >= 15 {
		fmt.Println("快適")
	} else {
		fmt.Println("寒い")
	}
}
