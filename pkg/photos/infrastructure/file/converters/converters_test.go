package converters

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
)

// edited: https://github.com/golang/tour/blob/master/pic/pic.go#L35
func writeImage(m image.Image) (*os.File, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.png")
	if err != nil {
		return nil, err
	}

	err = png.Encode(file, m)
	if err != nil {
		return nil, err
	}
	return file, nil
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
	file, _ := os.Open(filepath.Join("test-resources", "test-img.gif"))
	return file
}

func getJpegResource() *os.File {
	file, _ := os.Open(filepath.Join("test-resources", "smptebars.jpg"))
	return file
}

func getPngResource() *os.File {
	file, _ := os.Open(filepath.Join("test-resources", "rgbtest.png"))
	return file
}
