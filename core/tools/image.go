package tools

import (
	"bytes"
	"image"

	"golang.org/x/image/font/gofont/goregular"

	"github.com/gonutz/gofont"
)

func NewChaptaImage(code string) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, 240, 64))
	font, err := gofont.Read(bytes.NewReader(goregular.TTF))
	if err != nil {
		return nil, err
	}
	font.HeightInPixels = 25
	font.R, font.G, font.B, font.A = 255, 255, 255, 255 // black
	font.Write(img, code, (240-25*4)/2, (64-25)/2)
	return img, nil
}
