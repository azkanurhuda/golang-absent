# Product Transaction
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
### POST /signup
```shell
curl --location --request POST 'localhost:8888/signup' \
--form 'email=azka@email.com' \
--form 'password=azka123' \
--form 'username=azkapass'
```

### POST /login
```shell
curl --location --request POST 'localhost:8888/login' \
--form 'email=azka@email.com' \
--form 'password=azkapass'

```

### GET /v1/me
```shell
curl --location --request GET 'localhost:8888/v1/me' \
--header "Authorization: Bearer $JWT_TOKEN"
```

### POST /v1/checkin
```shell
curl --location --request POST 'http://localhost:8888/v1/checkin' \
--header 'Authorization: Bearer $JWT_TOKEN"
```

### POST /v1/checkout
```shell
curl --location --request POST 'http://localhost:8888/v1/checkout' \
--header 'Authorization: Bearer $JWT_TOKEN"
```