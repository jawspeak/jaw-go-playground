package mainsub

import (
	"fmt"
	"jaw-playground-go/initexample/understandinit"
	"time"
)

func init() {
	fmt.Println("mainsub.init() start")
	fmt.Println("mainsub.init() end")
}

func MainsubFunc() string {
	fmt.Println("mainsub.MainsubFunc()")
	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println("mainsub.MainsubFunc() understandinit.X is", understandinit.X)
	return "<retval>"
}
