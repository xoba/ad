#!/bin/bash
export GO=`go run lib/files.go src/xoba | xargs`
echo *.md *.txt lib/*.sh *.sh $GO


