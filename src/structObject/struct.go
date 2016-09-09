// struct project struct.go
package structObject

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func EchoPerson() {

	a := Person{"张三", 20}
	fmt.Println(a)

	a.age = 30
	fmt.Println(a)

	p := &a
	p.name = "李四"
	fmt.Println(a)
}

func StructGrammar() {

	var (
		a = Person{"张三", 10}
		b = Person{name: "李四"}
		c = Person{}
		p = &Person{"王五", 20}
	)
	fmt.Println("-----")
	fmt.Println(a.name, b, c, p)
	//{张三 10} {李四 0} { 0} &{王五 20} p返回一个指向结构体的指针

}
