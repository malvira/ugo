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

	GPIOA_CRH := MemLocation(uint32(GPIOA + 0x4))

	*GPIOA_CRH = GPIO_CNF_OUTPUT_PUSHPULL << (((8 - 8) * 4) + 2)
	*GPIOA_CRH |= (GPIO_MODE_OUTPUT_2_MHZ << ((8 - 8) * 4))
}

func main() {

	gpio_init()

	GPIOA_BSRR := MemLocation(GPIOA + 0x18)
	GPIOA_BRR := MemLocation(GPIOA + 0x14)
	
	for {
		*GPIOA_BRR = GPIO5 /* LED on */

		for i := 0; i < 254; i ++ {
			for i := 0; i < 254; i ++ {
			}
		}
		
		*GPIOA_BSRR = GPIO5 /* LED off */

		for i := 0; i < 254; i ++ {
			for i := 0; i < 254; i ++ {
			}
		}
		
	}
}
