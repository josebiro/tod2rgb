/* Daytime color calculations */
package color

type Color struct {
	Red        uint8
	Green      uint8
	Blue       uint8
	Brightness uint8
}

func (c *Color) SetRed(v uint8) {
	c.Red = v
}

func (c *Color) SetGreen(v uint8) {
	c.Green = v
}

func (c *Color) SetBlue(v uint8) {
	c.Blue = v
}

func (c *Color) SetBrightness(v uint8) {
	c.Brightness = v
}

func (c *Color) SetNight() {
	// Hardcode night color and brightness
	c.Red = 0
	c.Green = 0
	c.Blue = 255
	c.Brightness = 127
}

func NewColor() *Color {
	c := Color{Red: 0, Green: 0, Blue: 0, Brightness: 0}
	return &c
}
