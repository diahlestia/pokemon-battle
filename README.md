# 10xers
> **The backend API for recording pokemon battle data.**


## Running in Your Local Environment

Do these steps firstly:
1. Clone this repository to your local environment
2. Copy `.env.example` to `.env` and set the environment variables in `.env` adjust the value with your local env
3. Run `go run migrate/migrate.go` on the terminal. This will create the table to your local database.
4. Run `go run main.go` Open http://localhost:8080. If it shows `{"success": true}`, congratulations! Your setup is successful.

This API contains some method:
| METHOD | ENDPOINT                                           | REQUEST EXAMPLE                | DESCRIPTION                                                               |
|--------|----------------------------------------------------|--------------------------------|---------------------------------------------------------------------------|
| GET    | /pokemon/count                                     |                                | Get total of pokemons from API                                            |
| POST   | /match                                             | {"name": "match 1"}            | Register the match                                                        |
| GET    | /matches                                           |                                | Get list of matches                                                       |
| POST   | /pokemon/match/create                              | {"matchId": 1}                 | Get random 5 pokemons for the match and register the pokemon to the match |
| PUT    | /pokemon/match/start                               | {"matchId": 1}                 | Start the match and get the winner                                        |
| GET    | /pokemon/match?startDate=20200101&endDate=20221223 |                                | Get list of pokemon matches with date filter                              |
| GET    | /pokemon/matches                                   |                                | Get list of pokemon matches                                               |
| PUT    | /pokemon/discualification                          | {"matchId": 1, "pokemonId": 1} | Discualified a pokemon from a match                                       |