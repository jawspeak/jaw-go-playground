package main

import (
	"fmt"
	"jaw-playground-go/initexample/main/mainsub"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("main/main_test.go (1st)")
	mainsub.MainsubFunc()
}

func TestMain2(t *testing.T) {
	fmt.Println("main/main_test.go (2nd)")
	mainsub.MainsubFunc()
}
