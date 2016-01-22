package commands

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/wsxiaoys/terminal"

	"github.com/mondough/orchestra/services"
)

var UpdateCommand = &cli.Command{
	Name:         "update",
	Usage:        "Update service(s)",
	Action:       BeforeAfterWrapper(UpdateAction),
	BashComplete: ServicesBashComplete,
}

func UpdateAction(c *cli.Context) {
	args := []string{
		"get",
		"-d",
		"-t",
	}

	for _, service := range FilterServices(c) {
		args = append(args, service.Path)
	}

	if len(args) > 3 {
		cmd := exec.Command("go", args...)
		output := new(bytes.Buffer)
		cmd.Stdout = output
		cmd.Stderr = output

		spacing := strings.Repeat(" ", services.MaxServiceNameLength+2-len("tmp"))

		if cmd.Run() != nil {
			outputStr := output.String()
			appendError(fmt.Errorf("Failed to update services"))
			terminal.Stdout.Colorf("%s%s| @{r} error: @{|}Failed to update: %s\n", "tmp", spacing, outputStr)
		} else {
			terminal.Stdout.Colorf("%s%s| @{g} (re)built\n", "tmp", spacing)
		}
	}
}
