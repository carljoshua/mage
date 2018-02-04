package filter

import (
    "image"
)

func MeanBlur(img image.Image) image.Image {
    kernel := []uint32{1,1,1,1,1,1,1,1,1}
    return ApplyFilter(img, kernel, 3)
}

func GaussianBlur(img image.Image) image.Image {
    kernel := []uint32{1,2,1,2,4,2,1,2,1}
    return ApplyFilter(img, kernel, 3)
}

func Crop(img image.Image, xi, yi, xf, yf int) image.Image {
    area := image.Rect(xi, yi, xf, yf)
    cropped := image.NewRGBA(area)
    for y := area.Min.Y; y < area.Max.Y; y++ {
        for x := area.Min.X; x < area.Max.X; x++ {
            cropped.Set(x, y, img.At(x, y))
        }
    }
    return cropped
}
