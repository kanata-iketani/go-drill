// lesson22 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + ": ワン"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + ": ニャー"
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	kind := fields[0]
	name := fields[1]

	// Speaker 型の変数には「契約を満たす型」ならどれでも入れられます。
	// Dog も Cat も implements と書いていませんが、Speak() string を
	// 持っているだけで自動的に Speaker として扱えます（暗黙実装）。
	var s Speaker
	if kind == "dog" {
		s = Dog{Name: name}
	} else {
		s = Cat{Name: name}
	}
	fmt.Println(s.Speak())
}
