# Tests

This code here demonstrates writing tests with Go.

## Running Tests

Go has a [testing module](https://pkg.go.dev/testing)

Run tests with `go test` or `go test -v`

Run a single test case:

```bash
go test -run TestDoSomething -v
```

Many IDEs support running individual test cases directly.

## Benchmarks

- Use `testing.B` for setting up benchmarks
- Run with `go test -v -bench .` the `.` says run all of the benchmarks
- Specify a non-existant test to run and it will skip running tests at the same time
  `go test -v bench . -run SOMEMADEUPTESTNAME`

## Profiling

Where is the program spending its time?

```bash
go test -v -bench . -run ZZZ -cpuprofile=prof.out
go tool pprof prof.out
list doSomething
```
