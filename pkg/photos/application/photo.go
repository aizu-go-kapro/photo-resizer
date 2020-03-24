package application

import (
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
	"github.com/pkg/errors"
)

type fileBroker interface {
	ReadPhoto(path string) (photos.Photo, error)
	WritePhoto(path string, photo photos.Photo) error
}

// このサービスの使い方
// 1. ImportPhotoFromPathで画像を読み込む
// 2. ResizePhoto / GrayScalePhoto / PixelatePhoto を必要に応じて実行し画像を編集
// 3. ExportPhotoToPathで指定したパスに画像を保存
type PhotoService struct {
	photoRepository photos.Repository
	fileBroker      fileBroker
}

func NewPhotoService(photoRepository photos.Repository, broker fileBroker) PhotoService {
	return PhotoService{photoRepository, broker}
}

// リサイズしたい拡大率を入れると、内部で保存されている画像がリサイズされるユースケース。
func (ps PhotoService) ResizePhoto(rate float64) error {
	currentPhoto, err := ps.photoRepository.Get()
	if err != nil {
		return errors.Wrap(err, "no photo found. please import photo first")
	}
	newPhoto := currentPhoto.Resize(rate)
	err = ps.photoRepository.Save(&newPhoto)
	if err != nil {
		return err
	}
	return nil
}

// グレイスケールにしたいときに実行すると、内部で保存されている画像がグレイスケールにされるユースケース。
func (ps PhotoService) GrayScalePhoto() error {
	currentPhoto, err := ps.photoRepository.Get()
	if err != nil {
		return errors.Wrap(err, "no photo found. please import photo first")
	}
	newPhoto := currentPhoto.ConvertToGrayscale()
	err = ps.photoRepository.Save(&newPhoto)
	if err != nil {
		return err
	}
	return nil
}

// モザイク化するタイルの一片の長さを入力すると、内部で保存されている画像がモザイク化されるユースケース。
func (ps PhotoService) PixelatePhoto(side int) error {
	currentPhoto, err := ps.photoRepository.Get()
	if err != nil {
		return errors.Wrap(err, "no photo found. please import photo first")
	}
	newPhoto := currentPhoto.Pixelate(side)
	err = ps.photoRepository.Save(&newPhoto)
	if err != nil {
		return err
	}
	return nil
}

// 編集したい画像のパスを入れると、その画像を読み込んで編集可能状態にするユースケース。
func (ps PhotoService) ImportPhotoFromPath(path string) error {
	photo, err := ps.fileBroker.ReadPhoto(path)
	if err != nil {
		return errors.Wrap(err, "cannot get image")
	}
	err = ps.photoRepository.Save(&photo)
	if err != nil {
		return errors.Wrap(err, "cannot import photo")
	}
	return nil
}

// 編集した画像を保存したいときに実行すると、画像を元のファイル形式に変換して保存するユースケース。
func (ps PhotoService) ExportPhotoToPath(path string) error {
	photo, err := ps.photoRepository.Get()
	if err != nil {
		return errors.Wrap(err, "cannot export photo")
	}
	err = ps.fileBroker.WritePhoto(path, *photo)
	if err != nil {
		return errors.Wrap(err, "cannot convert photo to the file")
	}
	return nil
}
