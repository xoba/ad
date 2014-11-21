#!/bin/bash
./clean.sh
source install.sh
go test -cover xoba/...
go test -coverprofile=coverage.out xoba/ad/parser/templates
go tool cover -html=coverage.out
