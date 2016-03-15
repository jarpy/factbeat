package factbeat

import (
	"encoding/json"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"log"
	"os/exec"
	"strings"
	"time"
)

type Factbeat struct {
	period     time.Duration
	facterPath string
	done       chan struct{}
	FbConfig   ConfigSettings
	events     publisher.Client
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

func New() *Factbeat {
	return &Factbeat{}
}

func (fb *Factbeat) Config(b *beat.Beat) error {
	err := cfgfile.Read(&fb.FbConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	if fb.FbConfig.Input.Period != nil {
		fb.period = time.Duration(*fb.FbConfig.Input.Period) * time.Second
	} else {
		fb.period = 60 * time.Second
	}

	if fb.FbConfig.Input.Facter != nil {
		fb.facterPath = *fb.FbConfig.Input.Facter
	} else {
		fb.facterPath = FACTER_DEFAULT_PATH
	}

	return nil
}

func (fb *Factbeat) Setup(b *beat.Beat) error {
	fb.events = b.Events
	return nil
}

func (fb *Factbeat) Run(b *beat.Beat) error {
	for {
		// We do not know the shape of the JSON we'll get from Facter,
		// so we can't define a proper type for it. Instead, we'll
		// define the most generic thing that can store deserialized JSON.
		// REF: https://michaelheap.com/golang-encodedecode-arbitrary-json/
		var facts map[string]interface{}

		// Run Facter, and feed STDOUT a JSON decoder.
		// REF: http://www.darrencoxall.com/golang/executing-commands-in-go/
		cmd := exec.Command(fb.facterPath, "--json")
		facterOutput, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		decoder := json.NewDecoder(facterOutput)
		decoder.UseNumber() // REF: http://stackoverflow.com/a/22346593
		if err := decoder.Decode(&facts); err != nil {
			log.Fatal(err)
		}
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       "facts",
		}
		for key, value := range deDot(facts) {
			event[key] = value
		}
		fb.events.PublishEvent(event)

		// FIXME: Does not include processing time.
		time.Sleep(fb.period)
	}
}

func (fb *Factbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (fb *Factbeat) Stop() {
	close(fb.done)
}

