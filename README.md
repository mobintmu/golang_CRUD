# golang_CRUD
## golang simple project



## Run server
go run cmd/main.go


## test health

```
curl http://localhost:3000/api/health
curl -X POST http://localhost:3000/api/users      -H "Content-Type: application/json"      -d '{"id":"1","username":"john","email":"john@example.com"}'
```

## test docker file

```
curl http://localhost:8080/api/health
curl -X POST http://localhost:8080/api/users      -H "Content-Type: application/json"      -d '{"id":"1","username":"john","email":"john@example.com"}'
```


## build docker

```
docker build -t golang-crud .
```

## docker compose
docker compose up

Install Docker Compose on Ubuntu

https://phoenixnap.com/kb/install-docker-compose-on-ubuntu