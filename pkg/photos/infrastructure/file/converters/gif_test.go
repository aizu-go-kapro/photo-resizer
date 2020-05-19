package converters

import (
	"fmt"
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
	"github.com/pkg/browser"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestGifConverter_Open(t *testing.T) {
	gifConverter := NewGifConverter()
	gifFile := getOptimizedGifResource()

	gifPhotos, err := gifConverter.Open(gifFile)
	if err != nil {
		t.Error(err)
	}

	for _, m := range gifPhotos.Images() {
		filePath := writeImage(m).Name()
		fmt.Println("The file path : ", filePath)
		_ = browser.OpenURL(filePath)
	}
}

func TestGifConverter_Save(t *testing.T) {
	gifConverter := NewGifConverter()
	photosData := getAnimatedPhoto()

	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.gif")
	if err != nil {
		panic(err)
	}

	err = gifConverter.Save(file, photosData)
	if err != nil {
		t.Error(err)
	}

	filePath := file.Name()
	fmt.Println("The file path : ", filePath)
	_ = browser.OpenURL(filePath)
}

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

/*
About test cases:
All frame; test-resources/rgb.gif: http://tech.nitoyon.com/ja/blog/2016/01/07/go-animated-gif-gen/
Optimized; test-resources/nyancat.gif: https://blog.zhaytam.com/2018/08/21/creating-gifs-using-python-pillow/
*/
