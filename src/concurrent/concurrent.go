// concurrent project concurrent.go
package concurrent

import (
	"fmt"
	"time"
)

func Say(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

//同步
func SynByChannel(){
	
	a := []int{1,3,4,5,6}
	channel := make(chan int)
	sum := 0
	for _,v : range a {
		sum += v
	}
	
}