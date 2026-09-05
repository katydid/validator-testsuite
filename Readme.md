# Verified Derivatives for Fast Filtering and Schema Validation of Semi-Structured Data (TestSuite)

[![Build Status](https://git.katydid.org.za/validator-testsuite/actions/workflows/build.yml/badge.svg)](https://git.katydid.org.za/validator-testsuite/actions)

This is a language agnostic test and benchmark suite, which can be used by multiple implementation languages.

Results of running the benchmarks from 2016 to 2026 can be found in the [benches](./benches/) folder.

Even though the project contains some Go code, it is only there to help to generate tests and benchmarks.
The output is a set of files and folders that can be read by any programming language.

## Setup

1. Install [Go](https://golang.org) version 1.26.3, which can be downloaded [here](https://go.dev/dl/) for [Mac](https://go.dev/dl/go1.26.3.darwin-arm64.pkg) or [Linux](https://go.dev/dl/go1.26.3.linux-amd64.tar.gz).

2. Install protoc version 29.3 by [downloading](https://github.com/protocolbuffers/protobuf/releases/download/v29.3/protoc-29.3-linux-x86_64.zip) it from [Protobuf releases](https://github.com/protocolbuffers/protobuf/releases) and placing the `protoc` binary in your `PATH`.

3. Clone this repo to `./src/katydid.org.za/go/validator-testsuite`:

```
mkdir -p ./src/katydid.org.za/go/
(cd ./src/katydid.org.za/go/ && git clone https://git.katydid.org.za/validator-testsuite)
```

4. Generate benchmarks by going to your cloned testsuite `./src/katydid.org.za/go/validator-testsuite` and chose how much you want to generate by either running:

* `(cd ./src/katydid.org.za/go/validator-testsuite && make regenerate-paper-benchmarks)` to only generate the benchmarks for the paper or
* `(cd ./src/katydid.org.za/go/validator-testsuite && make regenerate-all)` to generate a bigger variety of benchmarks.

This can take a few minutes.

5. Choose an implementation to benchmark and clone it:

* The Go protobuf implementation: `(cd ./src/katydid.org.za/go/ && git clone https://git.katydid.org.za/validator-go-proto)` 
* The Go JSON/XML/Reflect implementation: `(cd ./src/katydid.org.za/go/ && git clone https://git.katydid.org.za/validator-go)`

6. Go to your implementation and choose how much benchmarks or tests you want to run by either running:

* `(cd ./src/katydid.org.za/go/validator-go-proto && make paper_benchmarks)` to only run the benchmarks for the paper or
* `(cd ./src/katydid.org.za/go/validator-go-proto && make bench)` to run a bigger variety of benchmarks on Protocol Buffers and JSON.
* `(cd ./src/katydid.org.za/go/validator-go-proto && make test)` to tun the Go tests for the Protocol Buffers and JSON.
* `(cd ./src/katydid.org.za/go/validator-go && make bench)` to run a bigger variety of benchmarks for the XML, JSON and Reflect.
* `(cd ./src/katydid.org.za/go/validator-go && make test)` to tun the Go tests for the XML, JSON and Reflect.

## Tests

The validator tests are located in the `./validator/tests` folder.
Tests are grouped by codec: json, xml, etc. for the various serialization formats.
Inside each codec folder is a list of testcase folders, each with a name.
Also found in the testcase folder is a filename starting with `valid` or `invalid` depending on whether the contents of the file is valid with respect to the provide grammar of invalid.

## Benchmarks

The validator benchmarks are located in the `./validator/benches` folder.

This folder is not checked in, because of its size.
Instead this folder can be generated, by running `make regenerate-all` or `make regenerate-paper-benchmarks` depending on which benchmarks you want to run.
This will require `go` to be installed and this folder to checked out to
`./src/katydid.org.za/go/validator-testsuite`.

Inside each codec folder is a list of benchcase folders, each with a name.
Files in the codec folder are schemas that might be required by a benchcase.
For example, marshaled protocol buffer file descriptor sets.

