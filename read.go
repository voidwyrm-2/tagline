package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"go.senan.xyz/taglib"
)

var ignoreTags = map[string]struct{}{}

func subRead(file string, args []string) error {
	fset := flag.NewFlagSet("tagline read", flag.ContinueOnError)
	output := fset.String("o", "", "The file to write the tags to instead of stdout.")
	outputType := fset.String("ot", "toml", "The type of file to output as; valid values are 'toml' or 'json'.")

	err := fset.Parse(args)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}

		return err
	}

	tags, err := taglib.ReadTags(file)
	if err != nil {
		return fmt.Errorf("Error while reading tags from '%s': %w", filepath.Base(file), err)
	}

	if *output != "" {
		output := *output
		outputTypeLow := strings.ToLower(*outputType)

		var encode func(io.Writer, map[string][]string) error

		switch outputTypeLow {
		case "json":
			encode = encoderJson
		case "toml":
			encode = encoderToml
		default:
			return fmt.Errorf("Invalid output type '%s'", *outputType)
		}

		ext := "." + outputTypeLow

		if filepath.Ext(output) != ext {
			output += ext
		}

		f, err := os.Create(output)
		if err != nil {
			return err
		}

		defer f.Close()

		return encode(f, tags)
	} else {
		fmt.Printf("Tags for '%s':\n", filepath.Base(file))

		for k, v := range tags {
			if _, ok := ignoreTags[k]; !ok {
				fmt.Printf(" '%s': %s\n", k, v)
			}
		}
	}

	return nil
}

func encoderJson(w io.Writer, tags map[string][]string) error {
	enc := json.NewEncoder(w)

	return enc.Encode(tags)
}

func encoderToml(w io.Writer, tags map[string][]string) error {
	enc := toml.NewEncoder(w)

	return enc.Encode(tags)
}
