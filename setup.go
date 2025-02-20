package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Env struct {
	// command line
	Verbose    bool
	SkipBDFR   bool
	ConfigPath string
	Subreddit  string

	// config file
	Config Config

	// derived settings
	BDFRSubDir      string
	WorkSubDir      string
	HTMLSubDir      string
	HTMLMediaSubDir string
}

func parseArgs() (*Env, error) {
	verbose := flag.Bool("v", false, "verbose output")
	skipBDFR := flag.Bool("skip-bdfr", false, "skip the bdfr-clone step")
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) < 1 {
		return nil, fmt.Errorf("missing required argument: config_yaml")
	}

	if len(flag.Args()) < 2 {
		return nil, fmt.Errorf("missing required argument: subreddit")
	}

	configPath := flag.Args()[0]
	subreddit := flag.Args()[1]

	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}
	if !cfg.IsComplete() {
		return nil, fmt.Errorf("invalid config: %+v", cfg)
	}

	htmlSubDir := filepath.Join(cfg.HTMLDir, subreddit)

	env := &Env{
		Verbose:         *verbose,
		SkipBDFR:        *skipBDFR,
		ConfigPath:      configPath,
		Subreddit:       subreddit,
		Config:          cfg,
		BDFRSubDir:      filepath.Join(cfg.BDFRDir, subreddit),
		WorkSubDir:      filepath.Join(cfg.WorkDir, subreddit),
		HTMLSubDir:      htmlSubDir,
		HTMLMediaSubDir: filepath.Join(htmlSubDir, "media"),
	}

	j, err := json.MarshalIndent(env, "    ", "    ")
	if err != nil {
		return nil, err
	}
	log.Infof("[config]\n%s", string(j))

	return env, nil
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags...] <config_yaml> <subreddit>\n", filepath.Base(os.Args[0]))
}
