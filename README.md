Install [Mongodb](https://www.mongodb.com/download-center) and [Redis](https://redis.io/download), and start those services.
1. Get Steam game list `go run crawlers/steam_gamelist/main.go`
2. Get Steam games `go run crawlers/steam_games/main.go`
3. Host App `go run main.go`

Example pages:

[Game detail page](http://localhost:8080/games/424780/Imperia-Online)

[Game list page](http://localhost:8080/games)

[Catalog page](http://localhost:8080/catalog/1/Action)

[Multiple Catalog page](http://localhost:8080/catalog/1-2/Action-Strategy)
