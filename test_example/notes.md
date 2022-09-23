### notes

**generate a coverage profile from tests**
```cli
go test -v -cover -coverprofile=cover.out ./... --> to generate a.out coverage profile
go tool cover -html=cover.out -->to see it in the web browser
go tool cover -html=cover.out -o coverage.html --> to save to the file
-race // to check race problems

```
