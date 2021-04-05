package point

import (
	"image"
	"image/color"

	"github.com/souliot/cv/core/util"
)

type NegativeRGBA struct {
	bounds image.Rectangle
}

var _ Negative = new(NegativeRGBA)

func (m *NegativeRGBA) SetBounds(x0, y0, x1, y1 int) {
	m.bounds = image.Rect(x0, y0, x1, y1)
}

func (m *NegativeRGBA) Bounds() image.Rectangle {
	return m.bounds
}

func (m *NegativeRGBA) Negative(src image.Image) image.Image {
	if src == nil {
		return src
	}
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
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
			clr := util.GetColorRGBA(colorRgb, util.ImgChanAll)
			dst.SetRGBA(x, y, color.RGBA{255 - clr.R, 255 - clr.G, 255 - clr.B, clr.A})
		}
	}
	return dst
}
