package converters

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestJpegConverter_Open(t *testing.T) {
	jpegConverter := NewJpegConverter()
	jpegFile, err := getJpegResource()
	if err != nil {
		t.Fatal(err)
	}

	jpegPhotos, err := jpegConverter.Open(jpegFile)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range jpegPhotos.Images() {
		file, err := writeImage(m)
		if err != nil {
			t.Fatal(err)
		}
		filePath := file.Name()
		t.Log("The file path : ", filePath)
		imageUrl, err := uploadImageToImgur(file)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("Url: ", imageUrl)
	}
}

func TestJpegConverter_Save(t *testing.T) {
	jpegConverter := NewJpegConverter()
	photosData, err := getAnimatedPhoto()
	if err != nil {
		t.Fatal(err)
	}

	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.jpg")
	if err != nil {
		t.Fatal(err)
	}

	err = jpegConverter.Save(file, photosData)
	if err != nil {
		t.Fatal(err)
	}

	filePath := file.Name()
	t.Log("The file path : ", filePath)
	imageUrl, err := uploadImageToImgur(file)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Url: ", imageUrl)
}
