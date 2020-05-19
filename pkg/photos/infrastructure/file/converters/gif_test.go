package converters

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/browser"
)

func TestGifConverter_Open(t *testing.T) {
	gifConverter := NewGifConverter()
	gifFile, err := getOptimizedGifResource()
	if err != nil {
		t.Fatal(err)
	}

	gifPhotos, err := gifConverter.Open(gifFile)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range gifPhotos.Images() {
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

func TestGifConverter_Save(t *testing.T) {
	gifConverter := NewGifConverter()
	photosData, err := getAnimatedPhoto()
	if err != nil {
		t.Fatal(err)
	}

	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.gif")
	if err != nil {
		t.Fatal(err)
	}

	err = gifConverter.Save(file, photosData)
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
