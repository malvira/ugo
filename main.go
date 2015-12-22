package main

// blink from
// https://github.com/libopencm3/libopencm3-examples/blob/master/examples/stm32/f4/nucleo-f411re/blink/blink.c

// unsafe pointers to setup registers
import "unsafe"

func Register(addr uint32) (reg *uint32) {
	ptr := uintptr(addr)
	reg = (*uint32)(unsafe.Pointer(ptr))
	return reg
}

const PERIPH_BASE = 0x40000000
const PERIPH_BASE_AHB1 = PERIPH_BASE + 0x20000
const RCC_BASE = PERIPH_BASE_AHB1 + 0x3800

func gpio_init() {
	RCC_AHB1ENR := Register(uint32(RCC_BASE + 0x30))
	
	*RCC_AHB1ENR = (1 << 0)
}

func main() {

	gpio_init()
	
	for {}
}
