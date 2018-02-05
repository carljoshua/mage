package mage

import (
    "errors"
    "image"
    "image/png"
    "image/jpeg"
    "os"
)

type Image struct {
    img     image.Image
}

// Open() opens the file specified by the path.
// It returns an *Image if successful.
func Open(path string) (*Image, error) {
    if _, err := os.Stat(path); err != nil {
        return nil, err
    }
    reader, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    tmp, _, err := image.Decode(reader)
    if err != nil {
        return nil, err
    }
    return &Image{ img: tmp }, nil
}

// Save() create and store the *Image in a file.
// There are only two file extension that this function accepts:
// jpeg/jpg and png.
func (i *Image) Save(name string, ext string) error {
    file, err := os.Create(name + "." + ext)
    if err != nil {
        return err
    }
    switch ext {
    case "jpeg", "jpg":
        if err := jpeg.Encode(file, i.img, nil); err != nil {
            return err
        }
    case "png":
        if err := png.Encode(file, i.img); err != nil {
            return err
        }
    default:
        return errors.New("Invalid file extension")
    }
    return nil
}
