# coding-rule-linter
coding-rule-linter is linter for my coding rules.
My coding rules are:
- variable identifier should not contain symbols or numbers, except for blank identifier
- constant identifier should not contain lower case letter, symbols(except for underscore) or numbers.
- test function identifier should separate "Test" and "Identifier of Target" with underscore.
## Badges
[![CI](https://github.com/KazuwoKiwame12/coding-rule-linter/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/KazuwoKiwame12/coding-rule-linter/actions/workflows/ci.yml)
 [![Coverage Status](https://coveralls.io/repos/github/KazuwoKiwame12/coding-rule-linter/badge.svg?branch=main)](https://coveralls.io/github/KazuwoKiwame12/coding-rule-linter?branch=main)
[![License](https://img.shields.io/github/license/KazuwoKiwame12/coding-rule-linter)](/LICENSE)

## Installation
```bash
$ go install github.com/KazuwoKiwame12/coding-rule-linter/cmd/coding-rule-linter@latest
```

## Usage
### main.go
```golang
package main

// variable identifier should not contain symbols or numbers, except for blank identifier
var sample int = 1
var sample1 int = 2

// constant identifier should not contain lower case letter, symbols(except for underscore) or numbers.
const SAMPLE_SAMPLE int = 1
const SamPle int = 2
```
### main_test.go
```golang
package main_test

import (
	"fmt"
	"testing"
)

// test function identifier should separate "Test" and "Identifier of Target" with underscore.
func Test_Sample(t *testing.T) {
	fmt.Println("no problem")
}

func TestSample(t *testing.T) {
	fmt.Println("problem")
}

func Sample(t *string) {
	fmt.Println("not target")
}
```
### Run
```bash
$ go vet -vettool=$(which coding-rule-linter) ./...
# coding-rule-linter
./main.go:7:1: var identifier should follow the coding rules, you should change from "sample1" to correct format
./main.go:11:1: const identifier should follow the coding rules, you should change from "SamPle" to correct format
# coding-rule-linter_test
./main_test.go:12:1: test function identifier should follow the coding rules, you should change from "TestSample" to correct format
```
