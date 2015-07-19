package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/yuya-takeyama/argf"
)

const AppName = "base64dec"

type Options struct {
	ShowVersion bool `short:"v" long:"version" description:"Show version"`
}

var opts Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = AppName
	parser.Usage = "[OPTIONS] FILES..."

	args, err := parser.Parse()
	if err != nil {
		fmt.Print(err)
		return
	}

	r, err := argf.From(args)
	if err != nil {
		panic(err)
	}

	err = base64dec(r, os.Stdout, os.Stderr, opts)
	if err != nil {
		panic(err)
	}
}

func base64dec(r io.Reader, stdout io.Writer, stderr io.Writer, opts Options) error {
	if opts.ShowVersion {
		io.WriteString(stdout, fmt.Sprintf("%s v%s, build %s\n", AppName, Version, GitCommit))
		return nil
	}

	decoder := base64.NewDecoder(base64.StdEncoding, r)
	_, err := io.Copy(stdout, decoder)
	if err != nil {
		return err
	}

	return nil
}
