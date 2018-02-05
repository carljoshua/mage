// Package mage provides methods for blurs and image processing.
package mage

import (
    "image"
    "fmt"
)

const (
    FLIP_X          =   0
    FLIP_Y          =   1
    ROTATE_90       =   2
    ROTATE_270      =   3
)

// Crop() returns the portion of the image selected by the argument values.
func (i *Image) Crop(xi, yi, xf, yf int) *Image {
    area := image.Rect(xi, yi, xf, yf)
    cropped := image.NewRGBA(area)
    for y := area.Min.Y; y < area.Max.Y; y++ {
        for x := area.Min.X; x < area.Max.X; x++ {
            cropped.Set(x, y, i.img.At(x, y))
        }
    }
    return &Image{ img: cropped }
}

// Paste() pastes the image into another image in the selected coordinates.
// For example, to paste the image2 the top-left most part of image1,
// do "image1.Paste(image2, 0, 0)" without the quotes.
func (i *Image) Paste(m *Image, x_coor, y_coor int) *Image {
    b := i.img.Bounds()
    b2 := m.img.Bounds()
    tmp := image.NewRGBA(b)

    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            if x >= x_coor && y >= y_coor && x-x_coor < b2.Max.X && y-y_coor < b2.Max.Y {
                tmp.Set(x, y, m.img.At(x-x_coor, y-y_coor))
            }else {
                tmp.Set(x, y, i.img.At(x, y))
            }
        }
    }
    return &Image{ img: tmp }
}


// Resize() resizes the image into the given dimension.
// Due to the nature of the process, the some information are lost.
// This is noticable in the edges of the resized image.
func (i *Image) Resize(width, height int) *Image {
    b := i.img.Bounds()
    skipX := float32(b.Max.X)/float32(width)
    skipY := float32(b.Max.Y)/float32(height)

    tmp := image.NewRGBA(image.Rect(0, 0, width, height))
    b = tmp.Bounds()
    j, k := float32(0), float32(0)

    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            tmp.Set(x, y, i.img.At(int(j), int(k)))
            j += skipX

            fmt.Printf("%d-%d, %f-%f\n", x, y, j, k)
        }
        k += skipY
        j = 0
    }
    return &Image{ img: tmp }
}

// Transpose() applies the rotates or flips that it is passed. The available
// rotations of flips are: FLIP_X (flip the image horizontally), FLIP_Y (flip the
// image vertically), ROTATE_90 (rotate the image 90 degrees counter-clockwise)
// and ROTATE_270 (rotate the image 270 degrees counter-clockwise)
func (i *Image) Transpose(d int) *Image {
    b := i.img.Bounds()
    tmp := image.NewRGBA(b)
    if d == 2 || d == 3 {
        rect := image.Rect(b.Min.Y, b.Min.X, b.Max.Y, b.Max.X)
        tmp = image.NewRGBA(rect)
    }

    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            if d == 0 {
                tmp.Set(b.Max.X - x, y, i.img.At(x, y))
            }else if d == 1 {
                tmp.Set(x, b.Max.Y - y, i.img.At(x, y))
            }else if d == 2 {
                tmp.Set(y, b.Max.X - x, i.img.At(x, y))
            }else if d == 3 {
                tmp.Set(b.Max.Y - y, x, i.img.At(x, y))
            }else {
                return nil
            }
        }
    }
    return &Image{ img: tmp }
}
