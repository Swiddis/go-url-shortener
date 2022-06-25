# Go URL Shortener
This is a url shortener based on [Eddy Wm's implementation](https://www.eddywm.com/lets-build-a-url-shortener-in-go/). Written in Go, deployed with Docker.

## Description
This is mostly an experiment in me both learning Go and getting better at Docker. To do so, I've started with the base provided by [Eddy Wm](https://www.eddywm.com/), and added to it as I've seen fit. There's still a bit to be done.

## Getting Started

### Dependencies
* [Docker](https://www.docker.com/)

### Running
The project is fully tested and deployed in Docker through `docker-compose`.

Optionally, create environment variables:
```sh
REDIS_PASSWORD=[a secure password] # defaults to 'admin'
```

deploy the application with:
```sh
docker-compose up --build
```

Run unit/integration tests with:
```sh
docker-compose -f docker-compose.test.yml up --build --exit-code-from api-test
```

## Usage
The project defines 3 endpoints for running a simple URL shortener:
```sh
# Ping endpoint
GET http://localhost:9808/
> {
  "message": "Hello, World! Looks like the API is running."
}

# Shorten a url
POST http://localhost:9808/create-short-url
{
  "long_url": "https://github.com/Swiddis/go-url-shortener",
  "user_id": "0"
}
> {
  "message": "short url created successfully",
  "short_url": "http://localhost:9808/ZT5syNhZ"
}

# Retrieve a URL
GET http://localhost:9808/ZT5syNhZ
> [Redirect to https://github.com/Swiddis/go-url-shortener]
```

## To-do
* Web front-end
* Load balancer
* More sophisticated storage scheme
  * Use Redis as a cache
  * Use a SQL database as persistent storage
* Switch from deprecated `docker-compose` to `docker compose`
* General functionality additions
