#!/bin/bash

export DYLD_LIBRARY_PATH=.
go build
./go-worker-api