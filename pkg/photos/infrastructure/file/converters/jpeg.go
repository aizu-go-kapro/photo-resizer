package converters

import (
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
	"image"
	"image/jpeg"
	"os"
)

type JpegConverter struct{}

func NewJpegConverter() JpegConverter {
	return JpegConverter{}
}

func (c JpegConverter) Open(imageFile *os.File) (photos.Photo, error) {
	var newPhoto []image.Image
	convertedImage, err := jpeg.Decode(imageFile)
	newPhoto = append(newPhoto, convertedImage)
	if err != nil {
		return photos.Photo{}, err
	}
	return photos.NewPhoto(newPhoto), nil
}

func (c JpegConverter) Save(out *os.File, photo photos.Photo) error {
	encodingImage := photo.Images()[0] // とりあえず1枚目以外は無視
	err := jpeg.Encode(out, encodingImage, &jpeg.Options{})
	if err != nil {
		return err
	}
	return nil
}
