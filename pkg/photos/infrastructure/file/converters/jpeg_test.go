package converters

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/browser"
)

func TestJpegConverter_Open(t *testing.T) {
	jpegConverter := NewJpegConverter()
	jpegFile := getJpegResource()

	jpegPhotos, err := jpegConverter.Open(jpegFile)
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}

	err = jpegConverter.Save(file, photosData)
	if err != nil {
		t.Fatal(err)
	}

	filePath := file.Name()
	fmt.Println("The file path : ", filePath)
	_ = browser.OpenURL(filePath)
}
