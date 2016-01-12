# Build Factbeat in a dedicated Golang container.
build: goget
	docker-compose run builder go fmt
	docker-compose run builder go clean
	docker-compose run builder go build

# Fetch Go dependencies.
goget:
	docker-compose run builder go get -d -v
	docker-compose run builder go get github.com/stretchr/testify/assert

# Run the "flake8" checker over the acceptance tests, which are in Python.
flake8:
	docker-compose run linter flake8 --exclude='.*' test

# Build all our Docker containers. See docker-compose.yml for the list.
docker: build
	docker-compose stop
	docker-compose rm --force
	docker-compose build

# Run Factbeat (and Elasticsearch) in the foreground of your terminal.
run: docker
	docker-compose up elasticsearch factbeat

# Run the unit tests. These are true _unit_ tests, exercising individual Go functions.
unit-test: goget
	docker-compose run builder go test -v . ./beat

# Run the "black box" acceptance tests, injecting data into an Elasticsearch
# instance, and making assertions about what's in there afterwards.
acceptance-test: build docker
	docker-compose up -d elasticsearch factbeat
	docker-compose run tester py.test test
	docker-compose stop
	docker-compose rm --force

test: unit-test acceptance-test
