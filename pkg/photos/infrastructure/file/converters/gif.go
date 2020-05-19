package converters

import (
	"image"
	"os"

	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
)

type GifConverter struct {
	// アニメーションのタイミング等を保存時に使いたい場合は、
	// ここで最後に読んだファイルを覚えておくといいかも？
}

func NewGifConverter() GifConverter {
	return GifConverter{}
}

func (c GifConverter) Open(imageFile *os.File) (photos.Photo, error) {
	var newPhoto []image.Image
	// TODO: gif画像をimage.Imageの配列に変換する処理
	return photos.NewPhoto(newPhoto), nil
}

func (c GifConverter) Save(out *os.File, photo photos.Photo) error {
	// TODO: image.Imageの配列をgif画像に変換する処理。変換したファイルはoutに書き込む。
	return nil
}
