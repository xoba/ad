#!/bin/bash
export GO=`go run lib/files.go src/xoba | xargs`
echo *.md *.txt *.r lib/*.sh *.sh $GO


