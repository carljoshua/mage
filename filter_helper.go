package mage

import (
    "image"
    "image/color"
)

// ApplyFilter() is used in applying custom filters into the image.
// The kernel should be square(3x3, 5x5) and must have odd number of element
// on each side(3x3 is valid while 4x4 is not). This is to ensure that the
// kernel have a center which will where the output of the kernel is applied
// in the image.
func applyFilter(img image.Image, kernel []float32, ksize int) [][][]float32 {
    b := img.Bounds()
    border := int(ksize/2)

    tmp := make([][][]float32, 4)

    for y := b.Min.Y; y < b.Max.Y; y++ {
        var tmpR, tmpG, tmpB, tmpA []float32
        for x := b.Min.X; x < b.Max.X; x++ {
            r, g, b, a := convolve(get_patch(img, x, y, border), kernel)
            tmpR = append(tmpR, r)
            tmpG = append(tmpG, g)
            tmpB = append(tmpB, b)
            tmpA = append(tmpA, a)
        }
        tmp[0] = append(tmp[0], tmpR)
        tmp[1] = append(tmp[1], tmpG)
        tmp[2] = append(tmp[2], tmpB)
        tmp[3] = append(tmp[3], tmpA)
    }

    return tmp
}

func (i *Image) ApplyFilter(kernel []float32, ksize int) *Image {
    tmp := tensor_to_image(applyFilter(i.img, kernel, ksize))
    return &Image{ img: tmp }
}

func convolve(input [][]uint32, kernel []float32) (float32, float32, float32, float32) {
    var sum float32
    tmp := make([]float32, 4)

    for _, v := range(kernel) {
        sum = sum + v
    }

    for i, channel := range(input) {
        total := float32(0)
        for i, _ := range(kernel) {
            total = total + (float32(channel[i]) * kernel[i])
        }
        if sum == 0 {
            tmp[i] = total
        }else {
            tmp[i] = total / sum
        }
    }

    return tmp[0], tmp[1], tmp[2], tmp[3]
}

func get_patch(img image.Image, x int, y int, border int) [][]uint32 {
    patch := make([][]uint32, 4)
    for i := -1*border; i <= border; i++ {
        for j := -1*border; j <= border; j++ {
            r, g, b, a := img.At(x + j, y + i).RGBA()
            patch[0] = append(patch[0], r>>8)
            patch[1] = append(patch[1], g>>8)
            patch[2] = append(patch[2], b>>8)
            patch[3] = append(patch[3], a>>8)
        }
    }
    return patch
}

func tensor_to_image(img [][][]float32) image.Image {
    b := image.Rect(0, 0, len(img[0][0]), len(img[0]))
    tmp := image.NewRGBA(b)
    for y, _ := range(img[0]) {
        for x, _ := range(img[0][y]) {
            c := color.RGBA{
                R: uint8(img[0][y][x]),
                G: uint8(img[1][y][x]),
                B: uint8(img[2][y][x]),
                A: uint8(img[3][y][x]) }
            tmp.Set(x, y, c)
        }
    }
    return tmp
}
