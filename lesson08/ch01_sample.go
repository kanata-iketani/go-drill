// lesson08 ch01: strings.Join（スライス→文字列）
package main

import (
	"fmt"
	"strings"
)

func main() {
	week := []string{"月", "火", "水"}

	// ",".join(week) に当たります。スライスが先、区切りが後です
	fmt.Println(strings.Join(week, ","))

	// 区切りには自由な文字列を使えます
	fmt.Println(strings.Join(week, " と "))

	// 空文字列ならそのまま連結です
	fmt.Println(strings.Join(week, ""))
}
