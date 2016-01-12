Introduction
============
Factbeat is a [Beat](https://www.elastic.co/products/beats) that ships
[Facter](https://puppetlabs.com/facter) facts to
[Elasticsearch](https://www.elastic.co/products/elasticsearch), where
they can be stored, analyzed and compared over time.

Building
========
[![Build Status](https://travis-ci.org/jarpy/factbeat.svg?branch=master)](https://travis-ci.org/jarpy/factbeat)

Like all Beats, Factbeat is written in Go. If you are familiar with
Go, and have a development environment set up, feel free to build
Factbeat like any other Go program.

However...

Containers, containers, containers.
==================================
Factbeat ships with a fully containerized build and test pipeline. It
provides containers that can build the Go source code and run
its unit tests. There are also containers that manage acceptance
testing using Python and a live Elasticsearch instance.

The containerized build/test sytem requires that you have:
* `docker`
* `docker-compose`
* `make`

However you don't need to install Go, Elasticsearch, Python etc. They
are all packaged for you in Docker containers.

Given the above dependencies, you should be able to simply:
```
make
```
to get a `./factbeat` binary.

Though let's not forget:
```
make unit-test
```
and
```
make acceptance-test
```

TODO
====
* Allow blacklist/whitelist of facts.
* Convert percents to beats style ie. "83.3%" -> 0.833
* Consider removing disk stats all together, since Topbeat has them
  covered.
* Improved mapping template.
* Automate installing the mapping template.
* _Your suggestions_.
