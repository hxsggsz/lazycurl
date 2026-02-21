package main

import (
	"lazycurl/cmd"
	_ "lazycurl/cmd/collections" // calls init()
)

func main() {
	cmd.Execute()
}
