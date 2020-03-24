package main

import (
	photosApplication "github.com/aizu-go-kapro/photo-resizer/pkg/photos/application"
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/infrastructure/file"
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/infrastructure/file/converters"
	photosInfra "github.com/aizu-go-kapro/photo-resizer/pkg/photos/infrastructure/photos"
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/interfaces/public/cli"
	"log"
	"os"
)

func main() {
	photosRepo := photosInfra.NewPhotoRepository()
	fileInfra := file.NewBroker(converters.NewJpegConverter(), converters.NewPngConverter(), converters.NewGifConverter())
	photosService := photosApplication.NewPhotoService(photosRepo, fileInfra)

	app := cli.GenerateClientApp(photosService)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
