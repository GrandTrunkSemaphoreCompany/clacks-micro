//build cortexm baremetal linux arm atsamd51j19 atsamd51 sam sam atsamd51j19a feather_m4 tinygo gc.conservative scheduler.tasks
package main

import (
	"fmt"
	"image/color"
	"machine"
	"math/bits"
	"time"
	"tinygo.org/x/drivers/ssd1306"
)

// change these to test a different UART or pins if available
var (
	uart = machine.UART0
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
)

func main() {
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})
	uart.Write([]byte("Echo console enabled. Type something then press enter:\r\n"))

		machine.I2C0.Configure(machine.I2CConfig{
			Frequency: machine.TWI_FREQ_400KHZ,
		})

		display := ssd1306.NewI2C(machine.I2C0)
		display.Configure(ssd1306.Config{
			Address: ssd1306.Address_128_32,
			Width:   128,
			Height:  32,
		})

		display.ClearDisplay()


	input := make([]byte, 64)
	n := 0
	for {
		var uByte uint8
		var bitHigh bool

		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()

			switch data {
			case 13:
				// return key

				for c := 0; c<n; c++ {
					uByte = input[c]
					for i := 0; i < 2; i++ {
						for j := 0; j < 4; j++ {
							c := color.RGBA{0, 0, 0, 255}

							bitHigh = bits.LeadingZeros8(uByte) != 0
							//fmt.Printf("%q = high: %t - %d\n\r", input[c], bitHigh, bits.LeadingZeros8(uByte))


							for w := 1; w < 15; w++ {
								for h := 1; h < 31; h++ {

									if bitHigh {
										c = color.RGBA{255, 255, 255, 255}
									}

									var x int16 = int16(h + (j * 32))
									var y int16 = int16(w + (i * 16))

									display.SetPixel(x, y, c)
								}
							}
							uByte = bits.RotateLeft8(uByte, 1)
						}
					}
					display.Display()
					// Flash
					time.Sleep(2 * time.Second)
					display.ClearDisplay()
					time.Sleep(2 * time.Second)
					fmt.Printf("\n\r")
				}
				n = 0
			default:
				// just echo the character
				uart.WriteByte(data)
				input[n] = data
				n++
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

//
//
//
//
//func main() {
//
//	machine.I2C0.Configure(machine.I2CConfig{
//		Frequency: machine.TWI_FREQ_400KHZ,
//	})
//
//	display := ssd1306.NewI2C(machine.I2C0)
//	display.Configure(ssd1306.Config{
//		Address: ssd1306.Address_128_32,
//		Width:   128,
//		Height:  32,
//	})
//
//	display.ClearDisplay()
//
//	for {
//		var message = []byte("Here is a string....")
//		var uByte uint8
//
//		for c := range message {
//			uByte = uint8(c)
//			for i := 0; i < 2; i++ {
//				for j := 0; j < 4; j++ {
//
//					var bitHigh bool = bits.LeadingZeros8(uByte) != 0
//					fmt.Printf("%s - %q = high: %t - %d\n", c, uByte, bitHigh, bits.LeadingZeros8(uByte))
//
//					for w := 0; w < 16; w++ {
//						for h := 0; h < 32; h++ {
//
//							c := color.RGBA{0, 0, 0, 255}
//
//							if !bitHigh {
//								c = color.RGBA{255, 255, 255, 255}
//							}
//
//							var x int16 = int16(h + (j * 32))
//							var y int16 = int16(w + (i * 16))
//
//							display.SetPixel(x, y, c)
//						}
//					}
//					uByte = bits.RotateLeft8(uByte, 1)
//				}
//			}
//
//			display.Display()
//			// Flash
//			time.Sleep(2 * time.Second)
//			display.ClearDisplay()
//			time.Sleep(2 * time.Second)
//		}
//	}
//}
