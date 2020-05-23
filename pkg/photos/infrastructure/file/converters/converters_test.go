package converters

import (
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

func getAnimatedPhoto() (photos.Photo, error) {
	var animatedImages []image.Image
	fileInfos, err := ioutil.ReadDir(filepath.Join("test-resources", "motions"))
	if err != nil {
		return photos.Photo{}, err
	}

	for _, fileInfo := range fileInfos {
		file, err := os.Open(filepath.Join("test-resources", "motions", fileInfo.Name()))
		if err != nil {
			return photos.Photo{}, err
		}

		pngImage, err := png.Decode(file)
		if err != nil {
			return photos.Photo{}, err
		}
		animatedImages = append(animatedImages, pngImage)
	}
	return photos.NewPhoto(animatedImages), nil
}

func getOptimizedGifResource() (*os.File, error) {
	file, err := os.Open(filepath.Join("test-resources", "test-img.gif"))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getJpegResource() (*os.File, error) {
	file, err := os.Open(filepath.Join("test-resources", "smptebars.jpg"))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getPngResource() (*os.File, error) {
	file, err := os.Open(filepath.Join("test-resources", "rgbtest.png"))
	if err != nil {
		return nil, err
	}
	return file, nil
}
