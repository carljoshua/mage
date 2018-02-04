package filter

import (
    "errors"
    "image"
    "image/png"
    "image/jpeg"
    "os"
)

type Image struct {
    image.Image

}

func Open(path string) (image.Image, error) {
    reader, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    img, _, err := image.Decode(reader)
    if err != nil {
        return nil, err
    }
    return img, nil
}

func Save(img image.Image, name string, ext string) error {
    file, err := os.Create(name + "." + ext)
    if err != nil {
        return err
    }
    switch ext {
    case "jpeg", "jpg":
        if err := jpeg.Encode(file, img, nil); err != nil {
            return err
        }
    case "png":
        if err := png.Encode(file, img); err != nil {
            return err
        }
    default:
        return errors.New("Invalid file extension")
    }
    return nil
}
