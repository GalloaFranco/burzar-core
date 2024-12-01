# Commands

Execute app
```shell
go run ./...
```

Execute unit tests
```shell
go test -v ./...
```

Create & update mocks
```shell
mockery --all
# After that you can update mocks with: 
go generate
```