package cli

import (
	"errors"
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/application"
	"github.com/urfave/cli/v2"
)

var MissingArgumentError = errors.New("not enough argument error")

// 入れさせたいコマンド photo-converter --rate 50 --grayscale --pixelate 3 [input file] [output file]

func GenerateClientApp(service application.PhotoService) *cli.App {
	return &cli.App{
		// TODO: CLIの作成
		// マニュアル https://github.com/urfave/cli/blob/master/docs/v2/manual.md
		Name:      "photo-converter",
		Usage:     "resize photo with some effects",
		ArgsUsage: "INPUT_FILE OUTPUT_FILE",
		HideHelp:  true,

		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:     "rate",
				Aliases:  []string{"r"},
				Value:    100.0,
				Usage:    "resizing `PERCENTAGE`",
				Required: true,
			},
			&cli.IntFlag{
				Name:    "pixelate",
				Aliases: []string{"p"},
				Value:   0,
				Usage:   "pixelize one `SIDE`",
			},
			&cli.BoolFlag{
				Name:    "grayscale",
				Aliases: []string{"g"},
				Value:   false,
				Usage:   "convert color image to black and white",
			},
		},

		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				_ = cli.ShowAppHelp(c)
				return MissingArgumentError
			}

			INPUT_FILE_PATH := c.Args().Get(0)
			OUTPUT_FILE_PATH := c.Args().Get(1)
			PIXELATE_SIDE := c.Int("pixelate")
			IS_GRAYSCALE := c.Bool("grayscale")

			err := service.ImportPhotoFromPath(INPUT_FILE_PATH)
			if err != nil {
				return err
			}

			err = service.ImportPhotoFromPath(INPUT_FILE_PATH)
			if err != nil {
				return err
			}

			if IS_GRAYSCALE {
				err = service.GrayScalePhoto()
				if err != nil {
					return err
				}
			}

			if PIXELATE_SIDE > 0 {
				err = service.PixelatePhoto(PIXELATE_SIDE)
				if err != nil {
					return err
				}
			}

			err = service.ExportPhotoToPath(OUTPUT_FILE_PATH)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
