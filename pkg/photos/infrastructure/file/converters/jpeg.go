package converters

import (
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
	"image"
	"os"
)

type JpegConverter struct{}

func NewJpegConverter() JpegConverter {
	return JpegConverter{}
}

func (c JpegConverter) Open(imageFile *os.File) (photos.Photo, error) {
	var newPhoto []image.Image
	// TODO: Jpeg画像をimage.Imageの配列に変換する処理
	return photos.NewPhoto(newPhoto), nil
}

func (c JpegConverter) Save(out *os.File, photo photos.Photo) error {
	// TODO: image.Imageの配列をJpeg画像に変換する処理。変換したファイルはoutに書き込む。
	return nil
}
