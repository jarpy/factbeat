package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/jarpy/factbeat/config"

	"encoding/json"
	"os/exec"
	"strings"
)

type Factbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

func deDot(m map[string]interface{}) map[string]interface{} {
	for key, value := range m {
		// Recurse into sub-trees.
		switch value := value.(type) {
		case map[string]interface{}:
			m[key] = deDot(value)
		}

		if strings.ContainsRune(key, '.') {
			delete(m, key)
			key := strings.Replace(key, ".", "_", -1)
			m[key] = value
		}
	}
	return m
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Factbeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Factbeat) Run(b *beat.Beat) error {
	logp.Info("factbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		// We do not know the shape of the JSON we'll get from Facter,
		// so we can't define a proper type for it. Instead, we'll
		// define the most generic thing that can store deserialized JSON.
		// REF: https://michaelheap.com/golang-encodedecode-arbitrary-json/
		var facts map[string]interface{}

		// Run Facter, and feed STDOUT a JSON decoder.
		// REF: http://www.darrencoxall.com/golang/executing-commands-in-go/
		cmd := exec.Command(*bt.config.Facter, "--json")
		facterOutput, err := cmd.StdoutPipe()

		if err != nil {
			logp.Err("Error opening Facter: %v", err)
		}
		if err := cmd.Start(); err != nil {
			logp.Err("Error running Facter: %v", err)
		}

		decoder := json.NewDecoder(facterOutput)
		decoder.UseNumber() // REF: http://stackoverflow.com/a/22346593
		if err := decoder.Decode(&facts); err != nil {
			logp.Err("Error making JSON parser: %v", err)
		}
		if err := cmd.Wait(); err != nil {
			logp.Err("Error parsing JSON from Facter: %v", err)
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
		}
		for key, value := range deDot(facts) {
			event[key] = value
		}

		bt.client.PublishEvent(event)
		logp.Info("Event sent")
	}
}

func (bt *Factbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
