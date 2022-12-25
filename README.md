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
Then open `localhost:8080` on your browser

To run `query-service` 
```
go run query-service/cmd/main.go
```
Then open `localhost:8080` on your browser

### Runing all in one using docker compoes (merged port with nginx reverse proxy)

```
docker compose up
```

