# Billy

Billy is an app made to keep track of all your recurring subscriptions

## Requirements

Section currently still a work in progress

## Endpoints

Base API URL: http://localhost:4000

| Method | URL Pattern         | Handler            | Remarks                      |
| ------ | ------------------- | ------------------ | ---------------------------- |
| GET    | /api/v1/healthcheck | healthcheckHandler | Show application information |

## Tech Stack

1. `go-chi` – Routing
2. `pq` – PostgreSQL database driver
3. `golang-migrate` – SQL migrations
4. `godotenv` – Load environment variables

## Contact

Khairul Akmal – [@mofodox](https://twitter.com/mofodox)
