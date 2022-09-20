package cmd

import (
	"fmt"
	"github.com/jsiebens/ionscale/internal/key"
	"github.com/muesli/coral"
)

func keyCommand() *coral.Command {
	command := &coral.Command{
		Use:          "genkey",
		SilenceUsage: true,
	}

	var disableNewLine bool

	command.Flags().BoolVarP(&disableNewLine, "no-newline", "n", false, "do not output a trailing newline")

	command.RunE = func(command *coral.Command, args []string) error {
		serverKey := key.NewServerKey()
		if disableNewLine {
			fmt.Print(serverKey.String())
		} else {
			fmt.Println(serverKey.String())
		}
		return nil
	}

	return command
}
