package converters

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/browser"
)

func TestPngConverter_Open(t *testing.T) {
	pngConverter := NewPngConverter()
	pngFile := getPngResource()

	pngPhotos, err := pngConverter.Open(pngFile)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range pngPhotos.Images() {
		filePath := writeImage(m).Name()
		fmt.Println("The file path : ", filePath)
		_ = browser.OpenURL(filePath)
	}
}

func TestPngConverter_Save(t *testing.T) {
	pngConverter := NewPngConverter()
	photosData := getAnimatedPhoto()

	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.png")
	if err != nil {
		panic(err)
	}

	err = pngConverter.Save(file, photosData)
	if err != nil {
		t.Fatal(err)
	}

	filePath := file.Name()
	fmt.Println("The file path : ", filePath)
	_ = browser.OpenURL(filePath)
}
