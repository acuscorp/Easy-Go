### notes

**generate a coverage profile from tests**
```
go test -v -cover -coverprofile=cover.out ./...
go tool cover -html=cover.out
```
