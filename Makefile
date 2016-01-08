build: goget
	docker-compose run builder go fmt
	docker-compose run builder go clean
	docker-compose run builder go build

goget:
	docker-compose run builder go get -d -v
	docker-compose run builder go get github.com/stretchr/testify/assert

flake8:
	docker-compose run linter flake8 --exclude='.*' test

docker: build
	docker-compose stop
	docker-compose rm --force
	docker-compose build

run: docker
	docker-compose up elasticsearch factbeat

unit-test: goget
	docker-compose run builder go test -v . ./beat

acceptance-test: build docker
	docker-compose up -d elasticsearch factbeat
	docker-compose run tester py.test
	docker-compose stop
	docker-compose rm --force

test: unit-test acceptance-test
