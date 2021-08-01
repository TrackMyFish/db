# db

## Purpose

This repository hold the database design and migration scripts of TrackMyFish

## Running DB migrations

* Run postgres using `docker compose` - `docker-compose up -d` (**Note**: unless you're just testing TrackMyFish, you should change the secrets to something much more secure)
* Run migrations with `go run main.go`

## Logging into the postgres container

* Sometimes it's useful to log into the Postgres container, for testing purposes. This can be done with the `docker exec -it trackmyfish_db psql -U trackmyfish -W trackmyfish` command, followed by the specified password (`supersecretpassword` by default). **NOTE:** This MUST only be done in the development/testing environment only.


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

docker tag trackmyfish simondrake/trackmyfish:v1alpha1

docker push simondrake/trackmyfish:v1alpha1
```



## ToDo
