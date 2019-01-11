// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	esDockerImage = "elasticsearch:6.5.4"
)

func NewElasticsearch() (*elasticsearch, error) {
	cmd := exec.Command("docker", "run", "-d", "-p", "9200:9200", "-p", "9300:9300", esDockerImage)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	go func() {
		if err := cmd.Run(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}()

	es := &elasticsearch{}
	if err := checkES(es); err != nil {
		return es, err
	}
	es.dockerId = strings.TrimSpace(stdout.String())
	return es, nil
}

type elasticsearch struct {
	dockerId string
}

func (es *elasticsearch) ID() string {
	if utf8.RuneCountInString(es.dockerId) > 12 {
		return es.dockerId[:12]
	}
	return ""
}

func (es *elasticsearch) ping() chan error {
	out := make(chan error, 1)

	// HTTP check (port 9200)
	resp, err := http.DefaultClient.Get("http://localhost:9200/_cluster/health")
	if err != nil {
		out <- err
		return out
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		out <- fmt.Errorf("elasticsearch: bogus ping status code: %v", resp.Status)
		return out
	}
	out <- nil
	return out
}

func checkES(es *elasticsearch) error {
	fmt.Println("Waiting for Elasticsearch to be healthy...")
	ticker := time.After(60 * time.Second)
	for {
		select {
		case <-ticker:
			return fmt.Errorf("NewES: timeout waiting for startup")
		case err := <-es.ping():
			if err != nil {
				time.Sleep(1 * time.Second)
			} else {
				return nil
			}
		}
	}
}

func (es *elasticsearch) Stop() error {
	if es.dockerId != "" {
		return exec.Command("docker", "kill", es.dockerId).Run()
	}
	return nil
}
