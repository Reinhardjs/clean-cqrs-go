# Clean & CQRS Architecture with Go

## Running in local
If you want to run this project on your local machine, do the followings

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

<br>
<br>

To run `command-service`
```
go run command-service/cmd/main.go
```

available endpoints :
- `POST` /articles

<br>

To run `query-service` 
```
go run query-service/cmd/main.go
```

available endpoints :
- `GET` /articles
- `GET` /articles/search?query=your-query-here

<br>
<br>


### Runing all in one using docker compose (merged port with nginx reverse proxy)

```
docker compose up
```

available endpoints :
- `POST` /articles
- `GET` /articles
- `GET` /articles/search?query=your-query-here

