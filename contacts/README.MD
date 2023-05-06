### To test the whole application

```go test ./... -v```

### To test a specific file

```go test models/contact_test.go -v```

### To test a specific functionality

```go test -v -run TestContactToJson models/contact_test.go```

### To check which code is covered and which is not

```go test -timeout 30s -coverprofile=coverage.out contacts/models```

```go tool cover -func=coverage.out```

```go tool cover -html=coverage.out```

### to more coverage 

```go install github.com/axw/gocov@latest```

```go test -coverprofile=c.out```
```gocov convert c.out | gocov annotate -```

### to get report
```gocov test | gocov report```

### use tools to display the data 
https://github.com/matm/gocov-html


