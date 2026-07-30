# Verified Derivatives for Fast Filtering and Schema Validation of Semi-Structured Data (TestSuite)

[![Build Status](https://github.com/katydid/validator-testsuite/actions/workflows/build.yml/badge.svg)](https://github.com/katydid/validator-testsuite/actions)

Results of running the benchmarks from 2016 to 2026 can be found in the [benches](./benches/) folder.

<!-- Instructions for running the testsuite can be found at [katydid/paper-2026-verified-filter-go](https://github.com/katydid/paper-2026-verified-filter-go). -->

The test suite is a language agnostic test and benchmark suite, so that it can be used by multiple implementation languages.
The test suite contains some Go code to help to generate tests for multiple serialization formats.
The output is just a bunch of files and folders that can be read by any programming language.

<!-- This repo was originally based [katydid/validator-testsuite](https://github.com/katydid/validator-testsuite/commit/565d7259b4a086251c11862a68b5619f731156bb), which explains how to add more tests and benchmarks. -->

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
`./src/github.com/katydid/validator-testsuite`.

Inside each codec folder is a list of benchcase folders, each with a name.
Files in the codec folder are schemas that might be required by a benchcase.
For example, marshaled protocol buffer file descriptor sets.

