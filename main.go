package main

import (
	"github.com/elastic/beats/libbeat/beat"
	factbeat "github.com/jarpy/factbeat/beat"
)

var Name = "factbeat"
var Version = "0.1.3"

func main() {
	beat.Run(Name, Version, factbeat.New())
}
