package converters

import (
	"fmt"
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
	"image"
	"image/png"
	"io/ioutil"
	"os"
)

// edited: https://github.com/golang/tour/blob/master/pic/pic.go#L35
func writeImage(m image.Image) *os.File {
	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(file, m)
	if err != nil {
		panic(err)
	}
	return file
}

func getAnimatedPhoto() photos.Photo {
	var animatedImages []image.Image
	for i := 1; i <= 10; i++ {
		file, _ := os.Open(fmt.Sprintf("test-resources/test-img%02d.png", i))
		pngImage, _ := png.Decode(file)
		animatedImages = append(animatedImages, pngImage)
	}
	return photos.NewPhoto(animatedImages)
}

func getOptimizedGifResource() *os.File {
	file, _ := os.Open("test-resources/nyancat.gif")
	return file
}
