package photos

import (
	"errors"
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
)

var NoImagesFound = errors.New("stored images is empty")

type MemoryRepository struct {
	photo photos.Photo
}

func NewPhotoRepository() *MemoryRepository {
	return &MemoryRepository{photos.Photo{}}
}

func (m *MemoryRepository) Get() (*photos.Photo, error) {
	if len(m.photo.Images()) == 0 {
		return &photos.Photo{}, NoImagesFound
	}
	return &m.photo, nil
}

func (m *MemoryRepository) Save(photo *photos.Photo) error {
	m.photo = *photo
	return nil
}
