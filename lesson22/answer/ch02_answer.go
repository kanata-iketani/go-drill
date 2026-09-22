// lesson22 ch02 解答
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

func (p Person) Introduce() string {
	return fmt.Sprintf("私は%s、%d歳です", p.Name, p.Age)
}

type Employee struct {
	Person
	Company string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	name := fields[0]
	age, _ := strconv.Atoi(fields[1])
	sc.Scan()
	company := sc.Text()

	// Person を「埋め込み」にしたので、Introduce() は Employee に昇格し、
	// e.Introduce() と直接呼べます。継承のような見た目ですが、
	// 実体は e.Person が持つメソッドへの近道（has-a）です。
	e := Employee{Person: Person{Name: name, Age: age}, Company: company}
	fmt.Println(e.Introduce())
	fmt.Println("勤務先は" + e.Company + "です")
}
