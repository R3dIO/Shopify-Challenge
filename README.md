# Shopify-Challenge
# Shopify Production Engineer Challenge Fall 2022

## How to Run

1. To run the program make sure you have latest version of Go
2. Install dependency
```sh
    go get
```
3. Export environment variable in shell for your openwether API key
```sh
    export KEY=adfhnasdjklasdjo123
```
4. Go into main directory and execute below command to start server
```sh
    go run main.go
```
5. After running server check localhost:8080 in your browser for application status or
```sh
    curl http://localhost:8080
```

# Crud on Items
- Create a item
```sh
    curl --location --request POST 'http://localhost:8080/items/' \
    --header 'Content-Type: text/plain' \
    --data-raw '{
        "name": "Monitor",
        "city": "Delhi",
        "quantity": 3
    }'
```

- Get a item
```sh
    curl --location --request GET 'http://localhost:8080/items/1'
```

- Update a item
```sh
    curl --location --request PATCH 'http://localhost:8080/items/1' \
    --header 'Content-Type: text/plain' \
    --data-raw '{
        "city": "London",
        "quantity": 4
    }'
```

- Delete a item
```sh
   curl --location --request DELETE 'http://localhost:8080/items/1'
```