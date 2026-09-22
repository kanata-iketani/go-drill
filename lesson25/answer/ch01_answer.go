// lesson25 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	name := fields[0]
	age, _ := strconv.Atoi(fields[1])

	original := Person{Name: name, Age: age}
	// 構造体は値型なので、この代入で中身が丸ごと複製されます。
	// Python のように「同じオブジェクトを 2 つの名前で指す」のではないため、
	// copied を書き換えても original には影響しません。
	copied := original
	copied.Age += 10

	fmt.Printf("original: %s %d\n", original.Name, original.Age)
	fmt.Printf("copied: %s %d\n", copied.Name, copied.Age)
}
