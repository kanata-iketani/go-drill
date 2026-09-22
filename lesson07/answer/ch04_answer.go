// lesson07 ch04 解答
package main

import "fmt"

func main() {
	// var で宣言した空スライスに追加していきます。
	var langs []string

	// append はその場で変更せず「追加後のスライス」を返すため、
	// s = append(s, x) と再代入するのが Python の a.append(x) との違いです。
	langs = append(langs, "Go")
	langs = append(langs, "Python")
	langs = append(langs, "Ruby")

	fmt.Println(langs)
	fmt.Println(len(langs))
}
