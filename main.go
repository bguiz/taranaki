package main

import (
	"github.com/example/taranaki/cmd"

	// register protocols and resources
	_ "github.com/example/taranaki/protocols"
	_ "github.com/example/taranaki/resources"
)

func main() {
	cmd.RootCmd.Execute()
}
