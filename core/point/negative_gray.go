package point

import (
	"image"
	"image/color"

	"github.com/souliot/cv/core/util"
)

type NegativeGray struct {
	bounds image.Rectangle
}

var _ Negative = new(NegativeGray)

func (m *NegativeGray) SetBounds(x0, y0, x1, y1 int) {
	m.bounds = image.Rect(x0, y0, x1, y1)
}

func (m *NegativeGray) Bounds() image.Rectangle {
	return m.bounds
}

func (m *NegativeGray) Negative(src image.Image) image.Image {
	if src == nil {
		return src
	}
	bounds := src.Bounds()
	dst := image.NewGray(bounds)
	if !m.bounds.Empty() {
		bounds = m.bounds
	}
	x0, y0, x1, y1 := bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			colorRgb := src.At(x, y)
			clr8 := util.GetColorGray(colorRgb, util.ImgChanAll)
			dst.SetGray(x, y, color.Gray{255 - clr8})
		}
	}
	return dst
}
