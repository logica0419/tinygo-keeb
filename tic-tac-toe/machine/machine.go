package machine

import (
	"image/color"
	"machine"
	"time"

	pio "github.com/tinygo-org/pio/rp2-pio"
	"github.com/tinygo-org/pio/rp2-pio/piolib"
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

type WS2812B struct {
	ws *piolib.WS2812B

	colPins [4]machine.Pin
	rowPins [3]machine.Pin

	display ssd1306.Device
}

func NewWS2812B() *WS2812B {
	s, _ := pio.PIO0.ClaimStateMachine()
	ws, _ := piolib.NewWS2812B(s, machine.GPIO1)
	ws.EnableDMA(true)

	colPins := [4]machine.Pin{
		machine.GPIO5,
		machine.GPIO6,
		machine.GPIO7,
		machine.GPIO8,
	}

	rowPins := [3]machine.Pin{
		machine.GPIO9,
		machine.GPIO10,
		machine.GPIO11,
	}

	for _, c := range colPins {
		c.Configure(machine.PinConfig{Mode: machine.PinOutput})
		c.Low()
	}

	for _, c := range rowPins {
		c.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 2.8 * machine.MHz,
		SDA:       machine.GPIO12,
		SCL:       machine.GPIO13,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address:  0x3C,
		Width:    128,
		Height:   64,
		Rotation: drivers.Rotation180,
	})

	return &WS2812B{
		ws: ws,

		colPins: colPins,
		rowPins: rowPins,
		display: display,
	}
}

func (ws *WS2812B) PutColor(c color.Color) {
	ws.ws.PutColor(c)
}

func (ws *WS2812B) WriteRaw(rawGRB []uint32) error {
	return ws.ws.WriteRaw(rawGRB)
}

func (ws *WS2812B) DisplayString(s string) {
	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}

	ws.display.ClearDisplay()
	time.Sleep(50 * time.Millisecond)

	tinyfont.WriteLine(&ws.display, &freemono.Bold9pt7b, 5, 10, s, white)
	ws.display.Display()
}

func (ws *WS2812B) Scan() int {
	defer time.Sleep(500 * time.Millisecond)

	for {
		// COL1
		ws.colPins[0].High()
		ws.colPins[1].Low()
		ws.colPins[2].Low()
		ws.colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if ws.rowPins[0].Get() {
			return 0
		}
		if ws.rowPins[1].Get() {
			return 3
		}
		if ws.rowPins[2].Get() {
			return 6
		}

		// COL2
		ws.colPins[0].Low()
		ws.colPins[1].High()
		ws.colPins[2].Low()
		ws.colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if ws.rowPins[0].Get() {
			return 1
		}
		if ws.rowPins[1].Get() {
			return 4
		}
		if ws.rowPins[2].Get() {
			return 7
		}

		// COL3
		ws.colPins[0].Low()
		ws.colPins[1].Low()
		ws.colPins[2].High()
		ws.colPins[3].Low()
		time.Sleep(1 * time.Millisecond)

		if ws.rowPins[0].Get() {
			return 2
		}
		if ws.rowPins[1].Get() {
			return 5
		}
		if ws.rowPins[2].Get() {
			return 8
		}

		// COL4
		// ws.colPins[0].Low()
		// ws.colPins[1].Low()
		// ws.colPins[2].Low()
		// ws.colPins[3].High()
		// time.Sleep(1 * time.Millisecond)

		// if ws.rowPins[0].Get() {
		// 	fmt.Printf("sw4 pressed\n")
		// 	time.Sleep(100 * time.Millisecond)
		// }
		// if ws.rowPins[1].Get() {
		// 	fmt.Printf("sw8 pressed\n")
		// 	time.Sleep(100 * time.Millisecond)
		// }
		// if ws.rowPins[2].Get() {
		// 	fmt.Printf("sw12 pressed\n")
		// 	time.Sleep(100 * time.Millisecond)
		// }
	}
}
