Tests are run in parallel per package, as separate processes.

But within a package, they are run with `-p` `-parallel` flags defaulting to `GOMAXPROCS`.

Go into each subfolder and run `time go test` and see speed.


Sample output in 4 ways of running with `-p` and `-parallel`:

```
~/Development/playground-go/tests_parallel_in_pkg master time go test ./... -count 1 -p 1 -parallel 1 | ruby -pe 'print Time.now.strftime("[%Y-%m-%d %H:%M:%S] ")'
[2022-05-09 18:04:31] ?   	jaw-playground-go/tests_parallel_in_pkg	[no test files]
[2022-05-09 18:04:39] ok  	jaw-playground-go/tests_parallel_in_pkg/serial	8.181s
[2022-05-09 18:04:47] ok  	jaw-playground-go/tests_parallel_in_pkg/simple_parallel	8.095s
[2022-05-09 18:04:47] ok  	jaw-playground-go/tests_parallel_in_pkg/zanother	0.090s
go test ./... -count 1 -p 1 -parallel 1  0.27s user 0.18s system 2% cpu 16.721 total
ruby -pe 'print Time.now.strftime("[%Y-%m-%d %H:%M:%S] ")'  0.02s user 0.01s system 0% cpu 16.723 total


~/Development/playground-go/tests_parallel_in_pkg master echo 'max parallel, packages and within packages to the extent t.Parallel() is called'
max parallel, packages and within packages to the extent t.Parallel() is called
~/Development/playground-go/tests_parallel_in_pkg master time go test ./... -count 1 -p 8 -parallel 8 | ruby -pe 'print Time.now.strftime("[%Y-%m-%d %H:%M:%S] ")'
[2022-05-09 18:05:19] ?   	jaw-playground-go/tests_parallel_in_pkg	[no test files]
[2022-05-09 18:05:27] ok  	jaw-playground-go/tests_parallel_in_pkg/serial	8.084s
[2022-05-09 18:05:27] ok  	jaw-playground-go/tests_parallel_in_pkg/simple_parallel	2.131s
[2022-05-09 18:05:27] ok  	jaw-playground-go/tests_parallel_in_pkg/zanother	0.176s
go test ./... -count 1 -p 8 -parallel 8  0.31s user 0.29s system 7% cpu 8.279 total
ruby -pe 'print Time.now.strftime("[%Y-%m-%d %H:%M:%S] ")'  0.02s user 0.01s system 0% cpu 8.281 total


~/Development/playground-go/tests_parallel_in_pkg master echo 'packages test serially, but parallel within packages to the extent t.Parallel() is called'
packages test serially, but parallel within packages to the extent t.Parallel() is called
~/Development/playground-go/tests_parallel_in_pkg master time go test ./... -count 1 -p 1 -parallel 8 | ruby -pe 'print Time.now.strftime("[%Y-%m-%d %H:%M:%S] ")'
[2022-05-09 18:06:05] ?   	jaw-playground-go/tests_parallel_in_pkg	[no test files]
[2022-05-09 18:06:14] ok  	jaw-playground-go/tests_parallel_in_pkg/serial	8.225s
[2022-05-09 18:06:16] ok  	jaw-playground-go/tests_parallel_in_pkg/simple_parallel	2.096s
[2022-05-09 18:06:16] ok  	jaw-playground-go/tests_parallel_in_pkg/zanother	0.088s
go test ./... -count 1 -p 1 -parallel 8  0.27s user 0.18s system 4% cpu 10.762 total
ruby -pe 'print Time.now.strftime("[%Y-%m-%d %H:%M:%S] ")'  0.02s user 0.01s system 0% cpu 10.765 total


~/Development/playground-go/tests_parallel_in_pkg master echo 'packages test in parallel, but within packages serially - aka t.Parallel() does not do anything'
packages test in parallel, but within packages serially - aka t.Parallel() does not do anything
~/Development/playground-go/tests_parallel_in_pkg master time go test ./... -count 1 -p 8 -parallel 1 | ruby -pe 'print Time.now.strftime("[%Y-%m-%d %H:%M:%S] ")'
[2022-05-09 18:06:51] ?   	jaw-playground-go/tests_parallel_in_pkg	[no test files]
[2022-05-09 18:06:59] ok  	jaw-playground-go/tests_parallel_in_pkg/serial	8.083s
[2022-05-09 18:06:59] ok  	jaw-playground-go/tests_parallel_in_pkg/simple_parallel	8.137s
[2022-05-09 18:06:59] ok  	jaw-playground-go/tests_parallel_in_pkg/zanother	0.182s
go test ./... -count 1 -p 8 -parallel 1  0.31s user 0.28s system 6% cpu 8.325 total
ruby -pe 'print Time.now.strftime("[%Y-%m-%d %H:%M:%S] ")'  0.02s user 0.01s system 0% cpu 8.327 total
```
