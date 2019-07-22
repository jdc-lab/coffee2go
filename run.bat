@echo off

cd cmd/lorca

go build -ldflags "-H windowsgui" -o ../../coffee2go.exe
cd ../../
coffee2go.exe