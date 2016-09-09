// interface project interface.go
package interfaces

import (
	"fmt"
)

type Person interface {
	Say()
}

type Children string

func (child Children) Say() {
	fmt.Println("Children say")
}

type YoungMan struct {
	Name string
}

func (man *YoungMan) Say() {
	fmt.Println("YoungMan say")
}

func Interfaces() {

	var person Person

	child := Children("我是小孩")
	youngman := YoungMan{"我是年轻人"}

	person = child //child实现了Person接口
	person.Say()

	person = &youngman //youngman实现了Person的接口
	person.Say()

}
