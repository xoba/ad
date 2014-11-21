#!/bin/bash
./clean.sh
source install.sh
go test -cover xoba/...