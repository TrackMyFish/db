# db

## Purpose

This repository hold the database design and migration scripts of TrackMyFish

## Running DB migrations

* Add database connection details to `local.env`
* Run migrations with `go run main.go`

### [migrate-cli](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

**NOTE:** This should only be used for local development purposes.

* Get current version - `migrate -path migrations -database "postgres://trackmyfish:supersecretpassword@localhost:5432/trackmyfish?sslmode=disable" version`
* Apply one down migration (e.g. from version 8 -> 7): `migrate -path migrations -database "postgres://trackmyfish:supersecretpassword@localhost:5432/trackmyfish?sslmode=disable" down 1`

## Logging into the postgres container

Sometimes it's useful to log into the Postgres container, for testing purposes. This can be done with the `docker exec -it trackmyfish_db psql -U trackmyfish -W trackmyfish` command, followed by the specified password (`supersecretpassword` by default).

It can also be done in one command: `docker exec -it trackmyfish_db psql postgresql://trackmyfish:supersecretpassword@localhost:5432/trackmyfish`

**NOTE:** This MUST only be done in the development/testing environment only.

# Running the Dockerfile

## Build the image

```
docker build -f ./Dockerfile -t trackmyfish_migrations .
```

## Run the image

```
docker run -e DATABASE_HOST=localhost -e DATABASE_PORT=5432 -e DATABASE_USERNAME=trackmyfish -e DATABASE_PASSWORD=supersecretpassword -e DATABASE_NAME=trackmyfish trackmyfish_migrations
```

## Publish the docker image

```
docker login

docker tag trackmyfish_migrations simondrake/trackmyfish_migrations:v1alpha1

docker push simondrake/trackmyfish_migrations:v1alpha1
```



## ToDo
