# Clean CQRS Architecture with Go

![image](https://user-images.githubusercontent.com/7758970/209476984-88fed3c1-31cc-4ed0-bfed-a401a5226c3f.png)

#### Built With

* Go (Mux)
* Gorm
* Docker
* Docker Compose
* Nginx Reverse Proxy

#### Stacks
* Redis
* PostgreDB
* Nats Streaming
* ElasticSearch

<br>

## Run from deployed

Temporarily, i've deployed this project to my personal VPS. You can access from the host of `http://103.134.154.18`.

For example, if you want to get list of article, access `http://103.134.154.18:8080/articles`

<br>

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
- `POST` localhost:8080/articles

<br>

To run `query-service` 
```
go run query-service/cmd/main.go
```

available endpoints :
- `GET` localhost:8080/articles
- `GET` localhost:8080/articles/search?query=your-query-here

<br>
<br>


### Runing all in one using docker compose (merged port with nginx reverse proxy)

```
docker compose up
```

available endpoints :
- `POST` localhost:8080/articles
- `GET` localhost:8080/articles
- `GET` localhost:8080/articles/search?query=your-query-here

<br>

### Example Request Body Payload for `POST` /articles :
```
{
    "title": "this is title 123",
    "content": "this is content 123"
}
```
