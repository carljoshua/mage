package mage

import (
    "image"
    "image/color"
    "math"
)

// Edges() finds and returns the edges in the image. (Improve noise reduction in this)
func (i *Image) Edges() *Image {
    bw_img := i.GaussianBlur(1.0).Grayscale()
    b := bw_img.img.Bounds()

    kernelX := []float32{-1, 0, 1,
                         -2, 0, 2,
                         -1, 0, 1 }
    kernelY := []float32{-1,-2,-1,
                          0, 0, 0,
                          1, 2, 1 }

    var edgeX, edgeY [][]float32
    edgeX = applyFilter(i.img, kernelX, 3)[0]
    edgeY = applyFilter(i.img, kernelY, 3)[0]

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
func (i *Image) GaussianBlur(sigma float32) *Image {
    kernel := []float32{ 4,3,2,3,4,
                         3,2,1,2,3,
                         2,1,0,1,2,
                         3,2,1,2,3,
                         4,3,2,3,4 }
    // for i, x := range(kernel) {
    //     tmp := 1/math.Sqrt(float64(2*math.Pi*(sigma*sigma)))
    //     kernel[i] = float32(tmp * math.Exp(-float64((x*x)/2*(sigma*sigma))))
    // }
    return i.ApplyFilter(kernel, 5)
}

// MeanBlur() applies returns an *Image with Mean Blur applied to it.
func (i *Image) MeanBlur() *Image {
    kernel := []float32{ 1,1,1,
                         1,1,1,
                         1,1,1 }
    return i.ApplyFilter(kernel, 3)
}
