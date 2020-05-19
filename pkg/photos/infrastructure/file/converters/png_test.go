package converters

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/browser"
)

func TestPngConverter_Open(t *testing.T) {
	pngConverter := NewPngConverter()
	pngFile, err := getPngResource()
	if err != nil {
		t.Fatal(err)
	}

	pngPhotos, err := pngConverter.Open(pngFile)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range pngPhotos.Images() {
		file, err := writeImage(m)
		if err != nil {
			t.Fatal(err)
		}
		filePath := file.Name()
		t.Log("The file path : ", filePath)
		err = browser.OpenURL(filePath)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestPngConverter_Save(t *testing.T) {
	pngConverter := NewPngConverter()
	photosData, err := getAnimatedPhoto()
	if err != nil {
		t.Fatal(err)
	}

	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.png")
	if err != nil {
		panic(err)
	}

	err = pngConverter.Save(file, photosData)
	if err != nil {
		t.Fatal(err)
	}

	filePath := file.Name()
	t.Log("The file path : ", filePath)
	err = browser.OpenURL(filePath)
	if err != nil {
		t.Fatal(err)
	}
}
