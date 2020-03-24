package converters

import (
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
	"image"
	"os"
)

type PngConverter struct{}

func NewPngConverter() PngConverter {
	return PngConverter{}
}

func (c PngConverter) Open(imageFile *os.File) (photos.Photo, error) {
	var newPhoto []image.Image
	// TODO: png画像をimage.Imageの配列に変換する処理
	return photos.NewPhoto(newPhoto), nil
}

func (c PngConverter) Save(out *os.File, photo photos.Photo) error {
	// TODO: image.Imageの配列をpng画像に変換する処理。変換したファイルはoutに書き込む。
	return nil
}
