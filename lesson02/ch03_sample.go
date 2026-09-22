// lesson02 ch03: fmt.Printf（書式指定つき出力）
package main

import "fmt"

func main() {
	name := "gopher"
	age := 13
	pi := 3.14159

	fmt.Printf("name は %s です\n", name)
	fmt.Printf("age は %d 歳です\n", age)
	fmt.Printf("pi はおよそ %f です\n", pi)
	fmt.Printf("小数第1位までなら %.1f です\n", pi)

	// Python の print(2026, 9, 20, sep="-") はこう書きます
	fmt.Printf("%d-%d-%d\n", 2026, 9, 20)
}
