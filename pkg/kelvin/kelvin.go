/* Kelvin - convert localtime to RGB and brightness color values

This is where the magic happens.package daylight. Algorithm adapted from
https://gist.github.com/paulkaplan/5184275 which itseelf was derived from
http://www.tannerhelland.com/4435/convert-temperature-rgb-algorithm-code/

*/
package kelvin

import (
	"math"

	color "github.com/josebiro/tod2rgb/pkg/color"
)

// Night is solid blue, but using the calculator is like a gazillion kelvin
const Night float64 = 100000000000000000000000000000000000

func KelvinToRGB(kelvin float64) *color.Color {
	c := color.NewColor()
	var temp float64

	var red float64
	var green float64
	var blue float64

	temp = kelvin / 100

	if temp <= 66 {
		red = 255

		green = float64(temp)
		green = 99.4708025861*math.Log(green) - 161.1195681661

		if temp <= 19 {
			blue = 0
		} else {
			blue = float64(temp - 10)
			blue = 138.5177312231*math.Log(blue) - 305.0447927307
		}
	} else {
		red = float64(temp - 60)
		red = 329.698727446 * math.Pow(red, -0.1332047592)

		green = float64(temp - 60)
		green = 288.1221695283 * math.Pow(green, -0.0755148492)

		blue = 255
	}

	c.SetRed(uint8(clamp(red, 0, 255)))
	c.SetGreen(uint8(clamp(green, 0, 255)))
	c.SetBlue(uint8(clamp(blue, 0, 255)))
	c.SetBrightness(250)
	return c
}

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
