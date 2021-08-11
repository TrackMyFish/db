
.PHONY: run
run:
	go run main.go

.PHONY: docker-build
docker-build:
	docker build -f ./Dockerfile -t trackmyfish_migrations .

.PHONY: docker-tag
docker-tag:
	docker tag trackmyfish_migrations simondrake/trackmyfish_migrations:v1alpha1

.PHONY: docker-push
docker-push:
	docker push simondrake/trackmyfish_migrations:v1alpha1

.PHONY: docker-all
docker-all: docker-build docker-tag docker-push
	echo "=================================="
	echo "Docker tagged, built and pushed"
	echo "=================================="
