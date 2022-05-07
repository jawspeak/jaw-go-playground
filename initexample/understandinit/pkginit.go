package understandinit

import "fmt"

var X int // set in init() here

func init() {
	fmt.Println("package understandinit, init() start")
	fmt.Println("X before", X)
	X = 33
	fmt.Println("X after", X)
	fmt.Println("package understandinit, init() end")
}
