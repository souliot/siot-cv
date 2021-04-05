package point

import "image"

type Negative interface {
	SetBounds(x0, y0, x1, y1 int)
	Bounds() image.Rectangle
	Negative(image.Image) image.Image
}
