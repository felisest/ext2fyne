package widgets

import "image/color"

func invertColor(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	return color.NRGBA{
		R: 255 - uint8(r>>8),
		G: 255 - uint8(g>>8),
		B: 255 - uint8(b>>8),
		A: uint8(a >> 8),
	}
}
