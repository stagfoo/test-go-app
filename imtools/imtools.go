package imtools

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

type circle struct {
	centerPoint image.Point
	radius      int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(
		c.centerPoint.X-c.radius,
		c.centerPoint.Y-c.radius,
		c.centerPoint.X+c.radius,
		c.centerPoint.Y+c.radius,
	)
}

func (c *circle) At(x, y int) color.Color {
	xpos := float64(x-c.centerPoint.X) + 0.5
	ypos := float64(y-c.centerPoint.Y) + 0.5
	radiusSquared := float64(c.radius * c.radius)
	if xpos*xpos+ypos*ypos < radiusSquared {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// Circle a rectangle source image.
func Circle(src image.Image) image.Image {
	dst := image.NewRGBA(src.Bounds())
	r := int(math.Min(
		float64(src.Bounds().Dx()),
		float64(src.Bounds().Dy()),
	) / 2)
	p := image.Point{
		X: src.Bounds().Dx() / 2,
		Y: src.Bounds().Dy() / 2,
	}
	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{p, r}, image.ZP, draw.Over)
	return dst
}
