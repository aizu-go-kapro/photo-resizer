package cli

import (
	"fmt"
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/application"
	"gopkg.in/urfave/cli.v2"
)

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
			&cli.Int64Flag{
				Name:    "pixelate",
				Aliases: []string{"p"},
				Value:   30,
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
			fmt.Println("boom! I say!")
			return nil
		},
	}
}
