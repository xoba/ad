#!/bin/bash -e
#
# re-builds the lexer and parser
#
source goinit.sh
./clean.sh
go install github.com/blynn/nex
./install.sh
go generate xoba/...
./install.sh
