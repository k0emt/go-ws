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

## Resource Links

- [Go Lang tutorial](https://golang.org/doc/tutorial/getting-started)
- [Gin web services](https://golang.org/doc/tutorial/web-service-gin)
- <https://rmoff.net/2020/07/16/learning-golang-some-rough-notes-s02e07-splitting-go-code-into-separate-source-files-and-building-a-binary-executable/>
