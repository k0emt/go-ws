# Notes

- <https://golang.org/doc/tutorial/getting-started>
- <https://golang.org/doc/tutorial/web-service-gin>
- <https://rmoff.net/2020/07/16/learning-golang-some-rough-notes-s02e07-splitting-go-code-into-separate-source-files-and-building-a-binary-executable/>

```bash
curl http://localhost:8080/articles \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id":"4","title":"Jaws","author":"someone","wordCount":7331}'
```

## Say what?

- In order for a go method to be exported from its package, it must start with a capital letter. :-/
- you import packages and not files
