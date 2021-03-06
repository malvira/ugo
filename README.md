Getting Started
===============

build runtime.a with:

    cd runtime
    GOARCH=arm go tool compile -pack runtime.go

and do the same for math.a

Compile main.o with:

    GOARCH=arm go tool compile main.go

Link to a.out with:

    GOARCH=arm go tool link -w -c -n -v -v -L runtime -L math main.o

can omit -c -n -v and -v (just debugging)

dwarf output doesn't work (hence the -w).

Should result in an a.out of about 18kB. Then to make binary:

    arm-none-eabi-objcopy -o BINARY a.out

Disassembly for verification:

    arm-none-eabi-objdump --disassembler-options=force-thumb -S main.elf

The binary will be constructed properly, but need to force-thumb when disassembling with objdump.


TODO
====

 * starting and interrupt vector
 
     - looks like will have to extend the link command to place
       external sections in places. Might be command line options that
       take an .o and an address and the .o gets inserted at those
       locations. Or something like that

  * For now, can get by without an interrupt table and just put everything at the starting address (just to blink an LED)

ARM / Thumb

https://www.altera.com/content/dam/altera-www/global/en_US/pdfs/literature/third-party/archives/ddi0100e_arm_arm.pdf

f44f 6280       mov.w   r2, #1024       ; 0x400
condition f will do thumb2. so need to implement this.
seems to be f44f 6280       mov.w   r2, #1024       ; 0x400
load/store immediate offset
see page 68

6280 f44f is op code according to
https://www.onlinedisassembler.com/odaweb/

this pdf has the T32 encoding
http://vision.gel.ulaval.ca/~jflalonde/cours/1001/h15/docs/ARM_v7.pdf
