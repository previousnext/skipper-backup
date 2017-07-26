#!/bin/bash

export GOPATH=$(pwd)/vendor:$(pwd)

go test $1
