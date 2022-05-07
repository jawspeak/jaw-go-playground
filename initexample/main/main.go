package main

import (
	"fmt"
	"jaw-playground-go/initexample/main/mainsub"
	"jaw-playground-go/initexample/understandinit"
	"time"
)

func init() {
	fmt.Println("main.main init()")
}

var _ = func() int {
	fmt.Println("in main.main var _ assignment trick which runs before main's init()")
	return 0
}()

func main() {
	fmt.Println("main.main hello world")
	fmt.Println("main.main understandinit.X is", understandinit.X) // import pkg above runs init
	// mainsub.MainsubFunc()
	fmt.Println("main.main calling mainsub.MainsubFunc()=", mainsub.MainsubFunc())
	go mainsub.MainsubFunc() // a go routine will not result in init() getting called again
	time.Sleep(time.Duration(2) * time.Second)
}
