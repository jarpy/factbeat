BEATNAME=factbeat
BEAT_DIR=github.com/jarpy/factbeat
SYSTEM_TESTS=false
TEST_ENVIRONMENT=false
ES_BEATS?=./vendor/github.com/elastic/beats
GOPACKAGES=$(shell glide novendor)
PREFIX?=.
NOTICE_FILE=NOTICE

# Path to the libbeat Makefile
-include $(ES_BEATS)/libbeat/scripts/Makefile

# Initial beat setup
.PHONY: setup
setup: copy-vendor
	make update

# Copy beats into vendor directory
.PHONY: copy-vendor
copy-vendor:
	mkdir -p vendor/github.com/elastic/
	cp -R ${GOPATH}/src/github.com/elastic/beats vendor/github.com/elastic/
	rm -rf vendor/github.com/elastic/beats/.git

.PHONY: git-init
git-init:
	git init
	git add README.md CONTRIBUTING.md
	git commit -m "Initial commit"
	git add LICENSE
	git commit -m "Add the LICENSE"
	git add .gitignore
	git commit -m "Add git settings"
	git add .
	git reset -- .travis.yml
	git commit -m "Add factbeat"
	git add .travis.yml
	git commit -m "Add Travis CI"


# Build all our Docker images. See docker-compose.yml for the list.
docker:
	docker-compose stop
	docker-compose rm --force
	docker-compose build

docker-build:
	docker-compose run builder rm -f factbeat
	docker-compose run -e GOOS=linux builder go build

# Run the "black box" acceptance tests, injecting data into an Elasticsearch
# instance, and making assertions about what's in there afterwards.
acceptance-test: factbeat docker
	docker-compose up -d elasticsearch factbeat
	docker-compose run tester py.test tests-acceptance
	docker-compose stop
	docker-compose rm --force

# This is called by the beats packer before building starts
.PHONY: before-build
before-build:

# Collects all dependencies and then calls update
.PHONY: collect
collect:
