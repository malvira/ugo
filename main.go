package main

// blink from
// https://github.com/libopencm3/libopencm3-examples/blob/master/examples/stm32/f4/nucleo-f411re/blink/blink.c

// unsafe pointers to setup registers
import "unsafe"

func MemLocation(addr uint32) (mem *uint32) {
	ptr := uintptr(addr)
	mem = (*uint32)(unsafe.Pointer(ptr))
	return mem
}

// constants for the stm32f4
// someday these will be wrapped up in a HAL 

const PERIPH_BASE = 0x40000000
const PERIPH_BASE_AHB1 = PERIPH_BASE + 0x20000
const RCC_BASE = PERIPH_BASE_AHB1 + 0x3800

const GPIO_PORT_A_BASE = PERIPH_BASE_AHB1 + 0x0000
const GPIOA = GPIO_PORT_A_BASE

const GPIO_CNF_OUTPUT_PUSHPULL = 0
const GPIO_MODE_OUTPUT_2_MHZ = 0x02

const GPIO5 = (uint32)(1 << 5)

func gpio_init() {
	RCC_AHB1ENR := MemLocation(uint32(RCC_BASE + 0x30))
	*RCC_AHB1ENR = (1 << 0)

	GPIOA_MODER := MemLocation(uint32(GPIOA + 0x0))
	*GPIOA_MODER = 1024

}

func main() {

	gpio_init()

	// hack for alignment
	// add an extra instruction to get the alignment at the word section correct
	j := uint8(0x55)
	j = j + j 
	
	GPIOA_BSRR := MemLocation(GPIOA + 0x18)
	
	for {
		*GPIOA_BSRR = GPIO5 /* bit set */

		for i := 0; i < 254; i ++ {
			for i := 0; i < 254; i ++ {
			}
		}

		// could work except don't have this instruction
		// (which moves a 32-bit literal :(
		// so need to come up with a way to do this.
		*GPIOA_BSRR = (GPIO5 << 16) /* bit reset */

		for i := 0; i < 254; i ++ {
			for i := 0; i < 254; i ++ {
			}
		}

		
	}
}
