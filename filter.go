// Package filter provides methods for blurs and image processing.
package filter

import (
    "image"
    "image/color"
)

func ApplyFilter(img image.Image, kernel []uint32, ksize int) image.Image {
    b := img.Bounds()
    border := int(ksize/2)

    tmp := image.NewRGBA(b)

    for y := b.Min.Y + border; y < b.Max.Y - border; y++ {
        for x := b.Min.X + border; x < b.Max.X - border; x++ {
            c := convolve(get_patch(img, x, y, border), kernel)
            tmp.Set(x, y, c)
        }
    }
    return tmp
}

func convolve(input [][]uint32, kernel []uint32) color.Color {
    var sum uint32
    tmp := make([]uint32, 4)

    for _, v := range(kernel) {
        sum = sum + v
    }

    for i, channel := range(input) {
        var total uint32
        total = 0
        for i, _ := range(kernel) {
            total = total + (channel[i] * kernel[i])
        }
        tmp[i] = uint32(total / sum)
    }

    return color.RGBA{ R: uint8(tmp[0]>>8),
        G: uint8(tmp[1]>>8),
        B: uint8(tmp[2]>>8),
        A: uint8(tmp[3]>>8) }
}

func get_patch(img image.Image, x int, y int, border int) [][]uint32 {
    patch := make([][]uint32, 4)
    for i := -1*border; i <= border; i++ {
        for j := -1*border; j <= border; j++ {
            r, g, b, a := img.At(x + j, y + i).RGBA()
            patch[0] = append(patch[0], r)
            patch[1] = append(patch[1], g)
            patch[2] = append(patch[2], b)
            patch[3] = append(patch[3], a)
        }
    }
    return patch
}
