#!/bin/sh

cd cmd/coffee2go

go build -o ../../coffee2go
cd ../../
./coffee2go