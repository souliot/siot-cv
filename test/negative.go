package test

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"

	"github.com/souliot/cv/core/point"
)

func negativeGray(src, dst string) (err error) {
	file, err := os.Open(src)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}
	n := new(point.NegativeGray)
	grayAll := n.Negative(img)
	fAll, err := os.Create(dst + "_negative_gray.png")
	if err != nil {
		return
	}
	defer fAll.Close()
	png.Encode(fAll, grayAll)
	return
}

func negativeRGBA(src, dst string) (err error) {
	file, err := os.Open(src)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}
	n := new(point.NegativeRGBA)
	All := n.Negative(img)
	fAll, err := os.Create(dst + "_negative_RGBA.png")
	if err != nil {
		return
	}
	defer fAll.Close()
	png.Encode(fAll, All)
	return
}
