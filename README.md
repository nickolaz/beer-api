# beer-api

Api for a Falabella Challenge.
localhost:8080/swagger/index.html
## Usage

You can use docker and docker-compose to run the project.

```bash
docker-compose up --build
```
Using Make

```bash
Make up
```

## Usage

In [SWAGGER](localhost:8080/swagger/index.html) you can see all endpoints of the project

- GET /beers: Get all beers
- GET /beers/day: the beer of the day is a random beer
- POST /beer/{beerID}: List a detail of one beer
- GET /beers/{beerID}/boxprice: List price of a box by params
- POST /beer: Create a beer