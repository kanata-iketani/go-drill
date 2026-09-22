// lesson08 ch02: strings.Split（文字列→スライス）
package main

import (
	"fmt"
	"strings"
)

func main() {
	csv := "赤,青,黄"

	// "赤,青,黄".split(",") に当たります。結果は []string です
	colors := strings.Split(csv, ",")
	fmt.Println(colors)
	fmt.Println(len(colors))

	// 分割後はふつうのスライスとして使えます
	fmt.Println(colors[1])
}
