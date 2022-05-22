package main

import "fmt"

func recoveryHelper(fromWhere string) {
	if err := recover(); err != nil {
		fmt.Println("Recovered in", fromWhere, "- I recovered in the helper from panic, err was:", err)
	}
}
func funcA() {
	defer func() { fmt.Println("  defer in funcA") }()
	//defer recoveryHelper("A") // try commenting out diff combo's of these in each func{A,B,C}
	fmt.Println("funcA start")
	funcB()
	fmt.Println("funcA end, (only accessed if recovered in B, or C)")
}
func funcB() {
	defer func() { fmt.Println("  defer in funcB") }()
	//defer recoveryHelper("B") // try commenting out diff combo's of these in each func{A,B,C}
	fmt.Println("funcB start")
	funcC()
	fmt.Println("funcB end, (only accessed if recovered in C)")
}
func funcC() {
	defer func() { fmt.Println("  defer in funcC") }()
	defer recoveryHelper("C") // try commenting out diff combo's of these in each func{A,B,C}
	fmt.Println("funcC start")
	panic("I paniced in C")
	// Below is unreachable code
	fmt.Println("funcC end, (never accessed b/c of panic (even if recovered in C this isn't executed b/c it is too late)")
}
func main() {
	funcA()
}
