// lesson09 ch01 解答
package main

import "fmt"

// Python なら return w*h, 2*(w+h) でタプルを返す場面です。
// Go にタプルはないため、(int, int) と宣言して 2 つの値を直接返します。
func rect(w, h int) (int, int) {
	return w * h, 2 * (w + h)
}

func main() {
	area, perimeter := rect(3, 4)
	fmt.Println(area)
	fmt.Println(perimeter)
}
