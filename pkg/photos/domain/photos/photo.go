package photos

import (
	"image"
)

type Photo struct {
	images []image.Image
}

func NewPhoto(images []image.Image) Photo {
	return Photo{images}
}

func (ph Photo) Images() []image.Image {
	return ph.images
}

// リサイズメソッド。rateに拡大倍率を入れて実行すると新しいPhotoが返る。
func (ph *Photo) Resize(rate float64) Photo {
	processingPhoto := *ph
	// TODO: リサイズ処理
	return processingPhoto
}

// グレールケール化メソッド。この関数を呼ぶとカラーの写真がグレースケールになる。
func (ph *Photo) ConvertToGrayscale() Photo {
	processingImages := make([]image.Image, len(ph.Images()))
	for i, img := range ph.Images() {
		processingImage := image.NewGray(img.Bounds())
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y += 1 {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x += 1 {
				processingImage.Set(x, y, img.At(x, y))
			}
		}
		processingImages[i] = processingImage
	}
	processingPhoto := NewPhoto(processingImages)
	return processingPhoto
}

// モザイク処理メソッド。sideのピクセル数を一辺の長さとしてモザイクをかける。
func (ph *Photo) Pixelate(side int) Photo {
	processingPhoto := *ph
	// TODO: ピクセル化処理
	return processingPhoto
}
