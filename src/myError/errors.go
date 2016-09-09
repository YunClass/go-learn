// errors project errors.go
package myError

import (
	"bytes"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	Msg  string
}

func (myError *MyError) Error() string {

	var buffer bytes.Buffer
	buffer.WriteString("时间：")
	buffer.WriteString(myError.When.String())
	buffer.WriteString(",说明：")
	buffer.WriteString(myError.Msg)

	return buffer.String()
}

func run() error {
	return &MyError{time.Now(), "出错啦"}
}

func PrintError() {

	if error := run(); error != nil {
		fmt.Println(error.Error())
	}

}
