package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/yusufcanb/tlm/app"
)

//go:embed VERSION
var version string
var sha1ver string

func main() {
	tlm := app.New(version, sha1ver)
	if err := tlm.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
