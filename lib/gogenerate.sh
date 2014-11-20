#!/bin/bash -e
#
# re-builds the lexer and parser
#
source goinit.sh
go install github.com/blynn/nex
go generate xoba/...
