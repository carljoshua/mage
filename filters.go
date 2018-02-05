package mage

import (
    "image"
    "image/color"
    "math"
)

// Edges() finds and returns the edges in the image. (Improve noise reduction in this)
func (i *Image) Edges() *Image {
    bw_img := i.GaussianBlur().Grayscale()
    b := bw_img.img.Bounds()
    var edgeX, edgeY [][]int

    kernelX := []int{   -1, 0, 1,
                        -2, 0, 2,
                        -1, 0, 1 }
    kernelY := []int{   -1,-2,-1,
                         0, 0, 0,
                         1, 2, 1 }

    for y := b.Min.Y; y < b.Max.Y; y++ {
        var tmpX, tmpY []int
        for x := b.Min.X; x < b.Max.X; x++ {
            tmp := make([]int, 9)
            i := 0
            for j := -1; j <= 1; j++ {
                for k := -1; k <= 1; k++ {
                    v, _, _, _ := bw_img.img.At(x+k, y+j).RGBA()
                    tmp[i] = int(v >> 8)
                    i++
                }
            }

            totalX := 0
            totalY := 0
            for i, _ := range(tmp) {
                totalX = totalX + (tmp[i] * kernelX[i])
                totalY = totalY + (tmp[i] * kernelY[i])
            }

            tmpX = append(tmpX, totalX)
            tmpY = append(tmpY, totalY)
        }
        edgeX = append(edgeX, tmpX)
        edgeY = append(edgeY, tmpY)
    }

    tmp := image.NewGray(b)
    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            a := edgeX[y][x]
            b := edgeY[y][x]
            c := ((a*a) + (b*b))
            if c <= 0 {
                c = 0
            }
            tmp.Set(x, y, color.Gray{ Y: uint8(math.Sqrt(float64(c))) })
        }
    }

    return &Image{ img: tmp }
}

// GaussianBlur() applies returns an *Image with Gaussian Blur applied to it.
func (i *Image) GaussianBlur() *Image {
    kernel := []uint32{ 1,2,1,
                        2,4,2,
                        1,2,1 }
    return &Image{ img: i.ApplyFilter(kernel, 3) }
}

// MeanBlur() applies returns an *Image with Mean Blur applied to it.
func (i *Image) MeanBlur() *Image {
    kernel := []uint32{ 1,1,1,
                        1,1,1,
                        1,1,1 }
    return &Image{ img: i.ApplyFilter(kernel, 3) }
}
