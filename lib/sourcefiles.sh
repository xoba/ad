#!/bin/bash
export GO=`go run lib/files.go src/xoba | xargs`
echo *.md *.asm *.txt *.r lib/*.sh *.sh $GO


