package plural

import (
	"github.com/pluralsh/pluralctl/pkg/console"
	"github.com/urfave/cli/v3"
)

const ApplicationName = "pluralctl"

type Plural struct {
	ConsoleClient console.ConsoleClient
}

func CreateNewApp(plural *Plural) *cli.App {
	app := cli.NewApp()

	return app
}

func RunPlural(arguments []string) error {
	return CreateNewApp(&Plural{}).Run(arguments)
}
