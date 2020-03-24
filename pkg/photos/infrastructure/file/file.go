package file

import (
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
	"github.com/pkg/errors"
	"os"
	"strings"
)

var InvalidFileExtension = errors.New("invalid file extension")

type ImageFileConverter interface {
	Open(imageFile *os.File) (photos.Photo, error)
	Save(out *os.File, photo photos.Photo) error
}

type Broker struct {
	jpegConverter ImageFileConverter
	pngConverter  ImageFileConverter
	gifConverter  ImageFileConverter
}

func NewBroker(jpegConverter ImageFileConverter, pngConverter ImageFileConverter, gifConverter ImageFileConverter) *Broker {
	return &Broker{jpegConverter, pngConverter, gifConverter}
}

func (fb *Broker) ReadPhoto(path string) (photos.Photo, error) {
	file, err := os.Open(path)
	if err != nil {
		return photos.Photo{}, errors.Wrap(err, "invalid path")
	}

	extension := path[strings.LastIndex(path, "."):]
	if extension == ".jpeg" || extension == ".jpg" {
		photo, err := fb.jpegConverter.Open(file)
		if err != nil {
			return photos.Photo{}, err
		}
		return photo, nil
	} else if extension == ".png" {
		photo, err := fb.pngConverter.Open(file)
		if err != nil {
			return photos.Photo{}, err
		}
		return photo, nil
	} else if extension == ".gif" {
		photo, err := fb.gifConverter.Open(file)
		if err != nil {
			return photos.Photo{}, err
		}
		return photo, nil
	}
	return photos.Photo{}, InvalidFileExtension
}

func (fb *Broker) WritePhoto(path string, photo photos.Photo) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}

	extension := path[strings.LastIndex(path, "."):]
	if extension == ".jpeg" || extension == ".jpg" {
		err = fb.jpegConverter.Save(out, photo)
		if err != nil {
			return err
		}
		return nil
	} else if extension == ".png" {
		err = fb.pngConverter.Save(out, photo)
		if err != nil {
			return err
		}
		return nil
	} else if extension == ".gif" {
		err = fb.gifConverter.Save(out, photo)
		if err != nil {
			return err
		}
		return nil
	}
	return InvalidFileExtension
}
