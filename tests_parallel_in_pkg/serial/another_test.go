package serial

import (
	"fmt"
	"jaw-playground-go/tests_parallel_in_pkg"
	"testing"
)

func TestOnotherDo1(t *testing.T) {
	fmt.Println("TestAnotherDo1")
	if tests_parallel_in_pkg.DoIt() != "done" {
		t.Fail()
	}
}

func TestAnotherDo2(t *testing.T) {
	fmt.Println("TestAnotherDo2")
	if tests_parallel_in_pkg.DoIt() != "done" {
		t.Fail()
	}
}
