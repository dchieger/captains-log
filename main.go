package main

import (
	"github.com/dchieger/captains-log/cmd"
)

func init() {
}

func main() {

	cmd.Execute()

	select {}
}
