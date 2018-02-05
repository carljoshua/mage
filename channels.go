package mage

import (
    "image"
    "image/color"
)

// Channels() splits the image in Red, Green and Blue channel.
func (i *Image) Channels() (*Image, *Image, *Image) {
    b := i.img.Bounds()
    r_img := image.NewRGBA(b)
    g_img := image.NewRGBA(b)
    b_img := image.NewRGBA(b)

    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            r, g, b, a := i.img.At(x, y).RGBA()

            r_img.Set(x, y, color.RGBA{ R: uint8(r>>8), G: 0, B: 0, A: uint8(a>>8) })
            g_img.Set(x, y, color.RGBA{ R: 0, G: uint8(g>>8), B: 0, A: uint8(a>>8) })
            b_img.Set(x, y, color.RGBA{ R: 0, G: 0, B: uint8(b>>8), A: uint8(a>>8) })
        }
    }
    return &Image{ img: r_img }, &Image{ img: g_img }, &Image{ img: b_img }
}

// R() returns the Red channel of the image.
func (i *Image) R() *Image {
    r, _, _ := i.Channels()
    return r
}

// G() returns the Green channel of the image.
func (i *Image) G() *Image {
    _, g, _ := i.Channels()
    return g
}

// B() returns the Blue channel of the image.
func (i *Image) B() *Image {
    _, _, b := i.Channels()
    return b
}

// Grayscale() converts the image into grayscale.
func (i *Image) Grayscale() *Image {
    b := i.img.Bounds()
    tmp := image.NewGray(b)

    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            tmp.Set(x, y, i.img.At(x, y))
        }
    }
    return &Image{ img: tmp }
}
