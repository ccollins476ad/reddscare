package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ccollins476ad/bdfrscrape/fileutil"
	log "github.com/sirupsen/logrus"
)

func main() {
	env, err := parseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	if env.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	if !env.SkipBDFR {
		err := runBDFR(*env)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error running bdfr: %v\n", err)
			os.Exit(2)
		}
	}

	err = runBDFRscrape(*env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running bdfrscrape: %v\n", err)
		os.Exit(3)
	}

	err = runBDFRToHTML(*env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running bdfrtohtml: %v\n", err)
		os.Exit(4)
	}

	// Copy media files that bdfrscrape downloaded to the html media directory.
	err = fileutil.RecursiveCopyIf(env.WorkSubDir, env.HTMLMediaSubDir, func(info os.FileInfo) bool {
		return info.IsDir() || strings.HasPrefix(info.Name(), "_bdfrscrape_http")
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error copying bdfrscrape media: %v\n", err)
		os.Exit(5)
	}
}

func runBDFR(env Env) error {
	args := []string{"-m", "bdfr", "clone", env.Config.BDFRDir, "--subreddit", env.Subreddit}
	args = append(args, env.Config.BDFROpts...)

	return run("python", args...)
}

func runBDFRscrape(env Env) error {
	var args []string

	// flags
	args = append(args, env.Config.BDFRscrapeOpts...)
	if env.Verbose {
		args = append(args, "-v")
	}

	// arguments
	args = append(args, env.BDFRSubDir)
	args = append(args, env.WorkSubDir)

	return run(env.Config.BDFRscrapePath, args...)
}

func runBDFRToHTML(env Env) error {
	args := []string{"-m", "bdfrtohtml", "--input_folder", env.WorkSubDir, "--output_folder", env.HTMLSubDir}
	args = append(args, env.Config.BDFRToHTMLOpts...)

	return run("python", args...)
}
