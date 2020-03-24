package cli

import (
	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/application"
	"gopkg.in/urfave/cli.v2"
)

func GenerateClientApp(service application.PhotoService) *cli.App {
	return &cli.App{
		// TODO: CLIの作成
		// マニュアル https://github.com/urfave/cli/blob/master/docs/v2/manual.md
	}
}
