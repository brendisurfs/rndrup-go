package main

import (
	"rndr-go/cmd"
	"rndr-go/env"
)

func main() {
	env.Env()
	cmd.Execute()
}
