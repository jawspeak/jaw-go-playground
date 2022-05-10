package tests_parallel_in_pkg

import "time"

func DoIt() string {
	time.Sleep(time.Duration(2) * time.Second)
	return "done"
}
