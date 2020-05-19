package converters

import (
	"fmt"
	"github.com/pkg/browser"
	"io/ioutil"
	"os"
	"testing"
)

func TestJpegConverter_Open(t *testing.T) {
	jpegConverter := NewJpegConverter()
	jpegFile := getJpegResource()

	jpegPhotos, err := jpegConverter.Open(jpegFile)
	if err != nil {
		t.Error(err)
	}

	for _, m := range jpegPhotos.Images() {
		filePath := writeImage(m).Name()
		fmt.Println("The file path : ", filePath)
		_ = browser.OpenURL(filePath)
	}
}

func TestJpegConverter_Save(t *testing.T) {
	jpegConverter := NewJpegConverter()
	photosData := getAnimatedPhoto()

	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.jpg")
	if err != nil {
		panic(err)
	}

	err = jpegConverter.Save(file, photosData)
	if err != nil {
		t.Error(err)
	}

	filePath := file.Name()
	fmt.Println("The file path : ", filePath)
	_ = browser.OpenURL(filePath)
}
