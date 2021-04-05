package util

import "image/color"

type ImgChan int

const (
	ImgChanAll ImgChan = iota
	ImgChanR
	ImgChanG
	ImgChanB
)

func GetColorGray(clr color.Color, ichan ImgChan) (clr8 uint8) {
	r, g, b, _ := clr.RGBA()
	switch ichan {
	case ImgChanAll:
		clr8 = uint8((r*19595 + g*38469 + b*7472) >> 16)
		return
	case ImgChanR:
		clr8 = uint8(r >> 8)
		return
	case ImgChanG:
		clr8 = uint8(g >> 8)
		return
	case ImgChanB:
		clr8 = uint8(b >> 8)
		return
	}
	return
}

func GetColorRGBA(clr color.Color, ichan ImgChan) (clrRGBA color.RGBA) {
	r, g, b, a := clr.RGBA()
	switch ichan {
	case ImgChanAll:
		r_u8, g_u8, b_u8, a_u8 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)
		clrRGBA = color.RGBA{r_u8, g_u8, b_u8, a_u8}
		return
	case ImgChanR:
		clr8, a_u8 := uint8(r>>8), uint8(a>>8)
		clrRGBA = color.RGBA{clr8, 0, 0, a_u8}
		return
	case ImgChanG:
		clr8, a_u8 := uint8(g>>8), uint8(a>>8)
		clrRGBA = color.RGBA{0, clr8, 0, a_u8}
		return
	case ImgChanB:
		clr8, a_u8 := uint8(b>>8), uint8(a>>8)
		clrRGBA = color.RGBA{0, 0, clr8, a_u8}
		return
	}
	return
}
