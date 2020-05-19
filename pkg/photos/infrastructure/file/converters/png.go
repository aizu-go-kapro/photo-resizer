package converters

import (
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
	"image"
	"image/png"
	"os"
)

type PngConverter struct{}

func NewPngConverter() PngConverter {
	return PngConverter{}
}

func (c PngConverter) Open(imageFile *os.File) (photos.Photo, error) {
	var newPhoto []image.Image
	convertedImage, err := png.Decode(imageFile)
	newPhoto = append(newPhoto, convertedImage)
	if err != nil {
		return photos.Photo{}, err
	}
	return photos.NewPhoto(newPhoto), nil
}

func (c PngConverter) Save(out *os.File, photo photos.Photo) error {
	encodingImage := photo.Images()[0] // とりあえず1枚目以外は無視
	err := png.Encode(out, encodingImage)
	if err != nil {
		return err
	}
	return nil
}
