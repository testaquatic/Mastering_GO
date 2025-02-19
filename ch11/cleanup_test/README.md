# 주의점

다음과 같이 실행해야 한다.

```bash
go test -v *.go
go clean -testcache
go test -v *.go
```
