#!/bin/bash -e
source goinit.sh
go install github.com/blynn/nex
go generate xoba/...