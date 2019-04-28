package commands

import (
	"github.com/abiosoft/ishell"
)

const mvUsage string = `
mv usage: mv src dst
Move parameter from src to dst.
`

func mv(c *ishell.Context) {
	err := ps.Move(parsePath(c.Args[0]), parsePath(c.Args[1]))
	if err != nil {
		shell.Println(err)
	}
}
