package converters

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/browser"
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
