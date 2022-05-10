package simple_parallel

import (
	"fmt"
	"jaw-playground-go/tests_parallel_in_pkg"
	"testing"
)

func TestDo1(t *testing.T) {
	t.Parallel()
	fmt.Println("TestDo1")
	if tests_parallel_in_pkg.DoIt() != "done" {
		t.Fail()
	}
}

func TestDo2(t *testing.T) {
	t.Parallel()
	fmt.Println("TestDo2")
	if tests_parallel_in_pkg.DoIt() != "done" {
		t.Fail()
	}
}
