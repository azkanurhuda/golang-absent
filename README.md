# Product Transaction
Golang Absent is service for absent and can detect IP Address, Latitude and Longitude
## Libraries
- [gorilla/mux](https://github.com/gorilla/mux)

## Getting Started
### Setup ENV
- Create account from [APIStack](https://ipstack.com/)
- Copy your Aour API Access Key into docker-compose.yaml -> API_STACK_KEY

### Run Server
1. Run containers
    ```shell
    docker compose up -d
    ```
2. Stop containers
   ```shell
   docker compose stop
   ```

## API
### Sign Up
`POST localhost:8888/signup`
#### Request
```shell
curl --location --request POST 'localhost:8888/signup' \
--form 'email=azka@email.com' \
--form 'password=azka123' \
--form 'username=azkapass'
```
#### Response
```json
{
    "user_id": "ea82e099-2add-4be9-a524-b36a17cd25d6",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZWE4MmUwOTktMmFkZC00YmU5LWE1MjQtYjM2YTE3Y2QyNWQ2IiwiZXhwIjoxNzExMTUxMzM0fQ.rTPRpUKe3QnTSFVTEjKrZ7cYnReFk2ofhFgt0qwbGSs",
    "expires_at": "2024-03-22T23:48:54Z"
}
```

### Login
`POST localhost:8888/login`
#### Request
```shell
curl --location --request POST 'localhost:8888/login' \
--form 'email=azka@email.com' \
--form 'password=azkapass'
```
#### Response 
```json
{
    "user_id": "ea82e099-2add-4be9-a524-b36a17cd25d6",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZWE4MmUwOTktMmFkZC00YmU5LWE1MjQtYjM2YTE3Y2QyNWQ2IiwiZXhwIjoxNzExMTUxMzk3fQ.zXHFiFUK-sC86V0vYjpD3ysBfl7fC-PKRKKhbSJ-8vI",
    "expires_at": "2024-03-22T23:49:57Z"
}
```

### Me
`GET localhost:8888/v1/me`
#### Request
```shell
curl --location --request GET 'localhost:8888/v1/me' \
--header "Authorization: Bearer $JWT_TOKEN"
```
#### Response
```json
{
    "user_id": "ea82e099-2add-4be9-a524-b36a17cd25d6",
    "username": "azkanurhuda",
    "email": "nurhudaazka@gmail.com"
}
```

### Check In
`POST localhost:8888/v1/checkin`
#### Request
```shell
curl --location --request POST 'http://localhost:8888/v1/checkin' \
--header 'Authorization: Bearer $JWT_TOKEN"
```
#### Response
```json
{
    "id": "2e7ea3bb-8ea1-42d5-a6ac-cbc4995636a4",
    "user_id": "ea82e099-2add-4be9-a524-b36a17cd25d6",
    "ip_address": "203.190.113.254",
    "latitude": -7.812300205230713,
    "longitude": 110.29810333251953,
    "status": "Check In"
}
```

### Check Out
`POST localhost:8888/v1/checkout`
#### Request
```shell
curl --location --request POST 'http://localhost:8888/v1/checkout' \
--header 'Authorization: Bearer $JWT_TOKEN"
```
#### Response
```json
{
    "id": "2e7ea3bb-8ea1-42d5-a6ac-cbc4995636a4",
    "user_id": "ea82e099-2add-4be9-a524-b36a17cd25d6",
    "ip_address": "203.190.113.254",
    "latitude": -7.812300205230713,
    "longitude": 110.29810333251953,
    "status": "Check Out"
}
```
