# lightbank
[![Go Report Card](https://goreportcard.com/badge/github.com/tonnytg/lightbank)](https://goreportcard.com/report/github.com/tonnytg/lightbank)

### Start Database

```
docker run -it --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:latest
```

Create a database `lightbank` and run app `go run main.go` using


