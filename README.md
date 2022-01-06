# Go WS

This research project was created to explore Go by creating a simple web service.

## make it Go

Build and run:

```bash
cd src
go build .
go run .
```

Use your browser or favorite tool to interact with the web service on `http://localhost:8080/`

## Spike code

See `src/ws.go` for the various endpoints.

Add an article using the service end point

```bash
curl http://localhost:8080/articles \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id":"4","title":"Jaws","author":"someone","wordCount":7331}'
```

## Say what?

Things learned:

- In order for a Go method to be exported from its package, it must start with a capital letter. :-/
- you import packages (whole directories or external references) and not files
- `make()` vs `new()`
- slices vs arrays
  Kind of like lists vs fixed size array.  Manipulate the slice with the `append` method.  With `make()` can cap the size of the slice.  Slices are ordered.  Can use the `sort` package to sort the list.
  
  Sorting a slice ascending and descending:

  ```go
  var numbers = []int{5, 4, 8, 3, 2, 1}
  sort.Ints(numbers)
  fmt.Println(numbers)

  sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
  fmt.Println(numbers)
  ```

## Some GO commands

```bash
go mod init path.to/module
go run .
go build .
go mod tidy
go install
go test
go get
```

## Next Steps

Some more things I'm going to tackle with this project:

- parse out the result in the time service
  - add another end point that uses/does the parsing
- gRPC / HTTP2 endpoints
- GraphQL
  - another one to utilize the time service, specify what you want back
- HTTPS
- how to create and run unit tests
- how to use with swagger.io
- Authentication / Authorization: API key?  JWT
- read from a configuration file
- log to an application log file
- write process ID to file
- `recover` error handling
- [fuzzing test](https://go.dev/doc/fuzz/) example when it is released into production

## Learning Resources

### Books

- [The Go Programming Language](https://www.amazon.com/gp/product/B0184N7WWS)
- [Learning Go: An Idiomatic Approach to Real-World Go Programming](https://www.amazon.com/gp/product/B08XYGCM71)

### LinkedIN Learning

- [Learning Go](https://www.linkedin.com/learning/learning-go-8399317)
- [Go Essential Training](https://www.linkedin.com/learning/go-essential-training)
- [Learning the Go Standard Library](https://www.linkedin.com/learning/learning-the-go-standard-library)
- [GO Design Patterns](https://www.linkedin.com/learning/go-design-patterns)

### Udemy Courses

- [Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/)
- [REST based microservices API development in Golang](https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang/)
- [gRPC GoLang Master Class: Build Modern API & Microservices](https://www.udemy.com/course/grpc-golang/)
- [Go Bootcamp: Master Golang with 1000+ Exercises and Projects](https://www.udemy.com/course/learn-go-the-complete-bootcamp-course-golang/)

### Web

- [Robin Moffat - Learning Go](https://rmoff.net/2020/06/25/learning-golang-some-rough-notes-s01e00/)
  - [separate source files](https://rmoff.net/2020/07/16/learning-golang-some-rough-notes-s02e07-splitting-go-code-into-separate-source-files-and-building-a-binary-executable/)

## Resource Links

- [Go Lang tutorial](https://golang.org/doc/tutorial/getting-started)
- [A Tour of Go](https://tour.golang.org/welcome/1)
- [Effective Go](https://golang.org/doc/effective_go#web_server)
- [Go Standard Library](https://pkg.go.dev/std)
- [Gin web services](https://golang.org/doc/tutorial/web-service-gin)
- [JSON-to-Go](https://mholt.github.io/json-to-go/)
- [Concurrency is _not_ Parallelism video](https://www.youtube.com/watch?v=oV9rvDllKEg) [(Slides)](https://talks.golang.org/2012/waza.slide#1)

## Libraries to check out

- [Fyne - cross platform GUI](https://github.com/fyne-io/fyne)
