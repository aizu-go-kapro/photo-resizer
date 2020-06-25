package photos

import (
	"image"
	"image/color"
	"math"
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
	if rate <= 0 {
		rate = 1
	}
	processingImages := make([]image.Image, len(ph.Images()))
	for i, img := range ph.Images() {
		// rateからキャンバスサイズを求め、Rectangleを生成する
		yDotsOriginal := img.Bounds().Max.Y - img.Bounds().Min.Y
		xDotsOriginal := img.Bounds().Max.X - img.Bounds().Min.X
		yDots := int(float64(yDotsOriginal) * (rate / 100.0))
		xDots := int(float64(xDotsOriginal) * (rate / 100.0))
		processingImage := image.NewRGBA(image.Rectangle{
			Min: image.Point{X: img.Bounds().Min.X, Y: img.Bounds().Min.Y},
			Max: image.Point{X: img.Bounds().Min.X + xDots, Y: img.Bounds().Min.Y + yDots},
		})

		for y := 0; y < processingImage.Bounds().Max.Y; y += 1 {
			for x := 0; x < processingImage.Bounds().Max.X; x += 1 {
				// 変換の画像と変換後の画像で最も近い距離の点を求め、そこのドットをセットする (Nearest Neighbor)
				originalY := math.Max(0, math.Min(float64(y)/(rate/100.0), float64(yDotsOriginal)))
				originalX := math.Max(0, math.Min(float64(x)/(rate/100.0), float64(xDotsOriginal)))
				processingImage.Set(x, y, img.At(img.Bounds().Min.X+int(originalX), img.Bounds().Min.Y+int(originalY)))
			}
		}
		processingImages[i] = processingImage
	}
	processingPhoto := NewPhoto(processingImages)
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
	if side <= 0 {
		side = 1
	}
	processingImages := make([]image.Image, len(ph.Images()))
	for i, img := range ph.Images() {
		processingImage := image.NewRGBA(img.Bounds())
		yDots := img.Bounds().Max.Y - img.Bounds().Min.Y
		xDots := img.Bounds().Max.X - img.Bounds().Min.X
		yBlocks := yDots / side
		xBlocks := xDots / side
		for yBlock := 0; yBlock < yBlocks; yBlock += 1 {
			for xBlock := 0; xBlock < xBlocks; xBlock += 1 {
				// ブロックの部分内の色の平均を取る
				var sumR, sumG, sumB, sumA, count uint64
				for y := img.Bounds().Min.Y + yBlock*side; y < int(math.Min(float64(img.Bounds().Min.Y+(yBlock+1)*side), float64(img.Bounds().Max.Y))); y += 1 {
					for x := img.Bounds().Min.X + xBlock*side; x < int(math.Min(float64(img.Bounds().Min.X+(xBlock+1)*side), float64(img.Bounds().Max.X))); x += 1 {
						r, g, b, a := img.At(x, y).RGBA()
						sumR += uint64(r)
						sumG += uint64(g)
						sumB += uint64(b)
						sumA += uint64(a)
						count += 1
					}
				}
				var aveR, aveG, aveB, aveA uint16
				aveR = uint16(sumR / count)
				aveG = uint16(sumG / count)
				aveB = uint16(sumB / count)
				aveA = uint16(sumA / count)
				aveColor := color.RGBA64{R: aveR, G: aveG, B: aveB, A: aveA}

				// ブロック内を求めた色の平均で埋める
				for y := img.Bounds().Min.Y + yBlock*side; y < int(math.Min(float64(img.Bounds().Min.Y+(yBlock+1)*side), float64(img.Bounds().Max.Y))); y += 1 {
					for x := img.Bounds().Min.X + xBlock*side; x < int(math.Min(float64(img.Bounds().Min.X+(xBlock+1)*side), float64(img.Bounds().Max.X))); x += 1 {
						processingImage.Set(x, y, aveColor)
					}
				}
			}
		}
		processingImages[i] = processingImage
	}
	processingPhoto := NewPhoto(processingImages)
	return processingPhoto
}
