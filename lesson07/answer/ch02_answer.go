// lesson07 ch02 解答
package main

import "fmt"

func main() {
	days := []string{"月", "火", "水"}

	// a[1] = "金" と同じ書き方で、その場で要素が書き換わります。
	days[1] = "金"
	fmt.Println(days)

	// Python の a[-1] は使えないため、末尾は len(s)-1 番目と指定します。
	fmt.Println(days[len(days)-1])
}
