package identicon

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"math/rand"
)

const (
	length = 50
	col    = 5
	row    = 5
)

// A Data provides to construct identicon.
type Data struct {
	id    string
	image *image.RGBA
	color color.Color
}

// NewData returns a Data object.
func NewData(id string) Data {
	d := Data{
		id:    id,
		image: image.NewRGBA(image.Rect(0, 0, col*length, row*length)),
		color: color.Black,
	}

	// XXX:

	draw.Draw(d.image, d.image.Bounds(),
		&image.Uniform{color.White}, image.ZP, draw.Src)

	for x := 0; x < col*length; x += length {
		for y := 0; y < row*length; y += length {
			fill := &image.Uniform{d.color}
			if rand.Intn(10)%2 == 0 {
				fill = &image.Uniform{color.White}
			}
			draw.Draw(d.image, image.Rect(x, y, x+length, y+length), fill, image.ZP, draw.Src)
		}
	}

	return d
}

// Encode writes the Image d.img to w in PNG format. Any Image may be encoded,
func (d *Data) Encode(w io.Writer) error {
	return png.Encode(w, d.image)
}
