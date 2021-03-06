package main

import (
	"os"

	"github.com/spf13/pflag"
)

func parseCLIFlags() cliOptions {
	flag := pflag.NewFlagSet("schmp", pflag.ExitOnError)
	mode := flag.StringP("mode", "m", "json", "input file format.  Allowed values: json, yaml, toml")
	outType := flag.StringP("out-type", "o", "stdout", "Output format. Allowed values: stdout, json")
	outFile := flag.String("out-file", "", "Output file. Only used if `--out-type` is not stdout")
	inFiles := flag.StringArrayP("file", "f", []string{}, "Files to compare. Use this flag multiple times, once for each file.")
	flag.Parse(os.Args[1:])

	options := cliOptions{
		mode:    *mode,
		outType: *outType,
		outFile: *outFile,
		inFiles: *inFiles,
	}
	return options
}
