# Clean & CQRS Architecture with Go

## Running in local
### Run each service in separates (without using docker compose)

```
go get -u -t -d -v ./...
```

```
go mod download
```

```
go mod tidy
```

To run `command-service`
```
go run command-service/cmd/main.go
```

To run `query-service` 
```
go run query-service/cmd/main.go
```

### Runing all in one using docker compoes (merged port with nginx reverse proxy)

```
docker compose up
```

