package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/jarpy/factbeat/beater"
)

func main() {
	err := beat.Run("factbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
