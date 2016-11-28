#!/usr/bin/env bash

cd $GOPATH/src/github.com/isindir/goproxmoxapi
go build
go install

#go test -test.v -parallel 20 -run "TestRecentTasksAPI"
#LONG_RUN_TEST=1 go test -test.v -parallel 20
go test -test.v -parallel 20
