Introduction
============
Factbeat is a [Beat](https://www.elastic.co/products/beats) that ships
[Facter](https://puppetlabs.com/facter) facts to
[Elasticsearch](https://www.elastic.co/products/elasticsearch), where
they can be stored, analyzed and compared over time.

Facter Version
--------------
Facter 3 (cFacter) is recommended, and is the default. The output from
Facter 3 is more structured and more stable than that of Facter 2,
which makes for a much better experience with Factbeat and
Elasticsearch.

If you really want to, you can configure the path to Facter, so you
could point Factbeat at Facter 2 instead.

Installing
===========
Pre-built binaries are availble for Windows and Linux (x86_64) on the
[releases page](https://github.com/jarpy/factbeat/releases).

The downloads for both platforms contain:
* The `factbeat` or `factbeat.exe` binary
* An example `factbeat.yml` config file
* The Elasticsearch mapping template: `factbeat.template.json`

The Windows archive also contains two Powershell scripts, for
registering and removing the Factbeat Windows service.

Building
========
[![Build Status](https://travis-ci.org/jarpy/factbeat.svg?branch=master)](https://travis-ci.org/jarpy/factbeat)

Factbeat was created in accordance with the [Beats Developer Guide][guide] and
thus uses the common build system. With the appropriate pre-requisites in
place, you should be able to simply:
```
make
```
to get a `./factbeat` binary.

[guide]: https://www.elastic.co/guide/en/beats/libbeat/5.2/new-beat.html

Acceptance Tests
================
Factbeat ships with a containerized test suite.

The containerized test system requires that you have:
* `docker`
* `docker-compose`
* `make`

However you don't need to install Elasticsearch, Python etc. They
are all packaged for you in Docker containers.

To run the suite, invoke:
```
make acceptance-test
```

TODO
====
* Convert percents to beats style ie. "83.3%" -> 0.833
* Improved mapping template.
* _Your suggestions_.
