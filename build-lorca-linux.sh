#!/bin/sh

cd cmd/lorca

go generate
go build -o ../../coffee2go