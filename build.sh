#!/bin/sh

set -x

export GOARCH=arm

cd runtime
go tool compile -pack runtime.go
cd ../math
go tool compile -pack math.go

cd ..
go tool compile main.go
go tool link -w -c -n -v -v -L runtime -L math -o main.elf main.o 
arm-none-eabi-objcopy -O binary main.elf main.bin
