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
    