# Build Factbeat in a dedicated Golang container.
build: build-linux build-windows

build-linux: go-fmt go-get
	docker-compose run -e GOOS=linux builder go build

build-windows: go-fmt go-get
	docker-compose run -e GOOS=windows builder go build

go-fmt:
	docker-compose run builder go fmt

# Fetch Go dependencies.
go-get:
	docker-compose run -e GOOS=linux builder go get -d -v
	docker-compose run -e GOOS=windows builder go get -d -v
	docker-compose run builder go get github.com/stretchr/testify/assert

# Run the "flake8" checker over the acceptance tests, which are in Python.
flake8:
	docker-compose run linter flake8 --exclude='.*' test

# Build all our Docker images. See docker-compose.yml for the list.
docker:
	docker-compose stop
	docker-compose rm --force
	docker-compose build

# Run Factbeat (and Elasticsearch) in the foreground of your terminal.
# Also provide Kibana, for browsing the results.
run: docker build-linux
	docker-compose up elasticsearch factbeat kibana

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
