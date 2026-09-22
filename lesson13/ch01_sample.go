// lesson13 ch01: for i, v := range スライス（enumerate 相当）
package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}

	// range はインデックスと値の 2 つを返します（enumerate 相当）
	for i, v := range fruits {
		fmt.Println(i, v)
	}

	// 番号を 1 から振りたいときは i+1 を使います
	for i, v := range fruits {
		fmt.Printf("%d番目: %s\n", i+1, v)
	}
}
