# Clean & CQRS Architecture with Go

This is a simple Rest API of articles, built with CQRS and Clean Architecture implemented.

![image](https://user-images.githubusercontent.com/7758970/209476768-17f067ba-200c-416e-a016-ff48622b8345.png)

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

