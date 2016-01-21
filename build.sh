#!/bin/sh

set -x

export GOARCH=arm

cd runtime
go tool compile runtime.go
#go tool asm start.s
go tool pack cv runtime.a *.o
cd ../math
go tool compile -pack math.go

cd ..
go tool compile main.go
go tool link -T 0x80001ac -w -c -n -v -v -L runtime -L math -o main.elf main.o

#dd if=/home/malvira/repos/libopencm3-examples/examples/stm32/f4/nucleo-f401re/blink/blink.bin of=header.bin bs=1 count=428

#echo -ne "\x00\x80\x01\x20\x08\x00\x00\x08" > header.bin

cat header-bin main.bin > full.bin

# use change addresses to offset to target arch memory (stm32 starts execution at 0x4
# stm32 top of stack value is at 0x0
# stm32 first instruction is at 0x4
#arm-none-eabi-objcopy --change-addresses -0x11000 main.elf main.elf
#arm-none-eabi-objcopy --change-addresses 4 main.elf main.elf 
arm-none-eabi-objcopy -O binary main.elf main.bin
