# Go URL Shortener

This is a url shortener based on [Eddy Wm's implementation](https://www.eddywm.com/lets-build-a-url-shortener-in-go/). Written in Go, deployed with Docker.

## Description

This is mostly an experiment in me both learning Go and getting better at Docker. To do so, I've started with the base provided by [Eddy Wm](https://www.eddywm.com/), and added to it as I've seen fit. There's still a bit to be done. Some differences:

- Fully containerized in Docker
- Testing environment also in Docker
- Load balancing through Nginx
  - Pass `--scale api=N` to balance between `N` instances.
- More detailed healthcheck endpoint

## Getting Started

### Dependencies

- [Docker](https://www.docker.com/)

### Running

The project is fully tested and deployed in Docker through `docker-compose`.

Optionally, create environment variables:

```sh
API_REPLICAS=[number] # defaults to 2
REDIS_PASSWORD=[a secure password] # defaults to 'admin'
```

Deploy the application with:

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
# Check if running
GET http://localhost:8080/health/ready
> {
  "status": "READY"
  "redis": "READY"
}

# Shorten a url
POST http://localhost:8080/
{
  "long_url": "https://github.com/Swiddis/go-url-shortener",
  "user_id": "0"
}
> {
  "message": "short url created successfully",
  "short_url": "http://localhost:8080/ZT5syNhZ"
}

# Retrieve a URL
GET http://localhost:8080/ZT5syNhZ
> [Redirect to https://github.com/Swiddis/go-url-shortener]
```

## To-do

- Web front-end
- More sophisticated storage scheme
  - Use Redis as a cache
  - Use a SQL database as persistent storage
- Switch from deprecated `docker-compose` to `docker compose`
- General functionality additions
