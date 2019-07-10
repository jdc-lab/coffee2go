@echo off

cd cmd/lorca

go generate
go build -ldflags "-H windowsgui" -o ../../coffee2go.exe