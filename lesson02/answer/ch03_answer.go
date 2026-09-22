// lesson02 ch03 解答
package main

import "fmt"

func main() {
	// %s に文字列、%d に整数、%.1f に小数が順番に対応します。
	// %.1f を使うと 169.500000 ではなく 169.5 と桁を制御でき、
	// 期待出力に正確に合わせられます（Printf は改行しないので \n も忘れずに）。
	fmt.Printf("%s さん (%d) の身長は %.1f cm です\n", "Tanaka", 28, 169.5)
}
