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

	"github.com/go-kit/kit/log"
)

var (
	esDockerImage = "elasticsearch:6.5.4"
)

// NewElasticsearch will launch and monitor a Docker container of Elasticsearch
func NewElasticsearch(logger log.Logger) (es *Elasticsearch, err error) {
	cmd := exec.Command("docker", "run", "-d", "-p", "9200:9200", "-p", "9300:9300", esDockerImage)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	go func() {
		if err = cmd.Run(); err != nil && logger != nil {
			logger.Log("elasticsearch", fmt.Sprintf("ERROR: %v\n", err))
		}
	}()

	es = &Elasticsearch{}
	if err := checkES(logger, es); err != nil {
		return es, err
	}
	es.dockerId = strings.TrimSpace(stdout.String())
	return es, nil
}

// Elasticsearch wraps a docker container of Elasticsearch
type Elasticsearch struct { // TODO(adam): move into internal/ package?
	dockerId string
}

// ID returns the docker ID for the running container
func (es *Elasticsearch) ID() string {
	if utf8.RuneCountInString(es.dockerId) > 12 {
		return es.dockerId[:12]
	}
	return ""
}

// Ping checks if Elasticsearch is up and running.
func (es *Elasticsearch) Ping() error {
	if es != nil {
		return <-es.ping()
	}
	return nil
}

func (es *Elasticsearch) ping() chan error {
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

func checkES(logger log.Logger, es *Elasticsearch) error {
	if logger != nil {
		logger.Log("elasticsearch", "Waiting for Elasticsearch to be healthy...")
	}
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

// Stop will shutdown the Elasticsearch docker container
func (es *Elasticsearch) Stop(logger log.Logger) error {
	if es == nil {
		return nil
	}
	if es.dockerId != "" {
		if logger != nil {
			logger.Log("elasticsearch", fmt.Sprintf("shutting down Elasticsearch (Docker container ID: %s)", es.ID()))
		}
		return exec.Command("docker", "kill", es.dockerId).Run()
	}
	return nil
}
