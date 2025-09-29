package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"go.senan.xyz/taglib"
)

const imgpathKey = "img"

type tagMap map[string][]string

func (tm tagMap) add(name string, value *string) {
	if *value != "" {
		tm[name] = []string{*value}
	}
}

func tfd(tag string) string {
	return "The value for the '" + tag + "' tag."
}

func subWrite(file string, args []string) error {
	fset := flag.NewFlagSet("tagline write", flag.ContinueOnError)

	tagFile := fset.String("f", "", "The file to read the tags from instead of the flags.")
	tagFileType := fset.String("ft", "toml", "The type of file to treat the file given with 'f' as; valid values are 'toml' or 'json'.")
	imgpath := fset.String("img", "", "A JPEG, PNG, or GIF file to use as the art for the audio file.")
	tTitle := fset.String("title", "", tfd(taglib.Title))
	tGenre := fset.String("genre", "", tfd(taglib.Genre))
	tAlbum := fset.String("album", "", tfd(taglib.Album))
	tArtist := fset.String("artist", "", tfd(taglib.Arranger))
	tComposer := fset.String("composer", "", tfd(taglib.Composer))
	tDate := fset.String("date", "", tfd(taglib.Date))

	err := fset.Parse(args)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}

		return err
	}

	tags := tagMap{}

	if *tagFile != "" {
		var decode func(io.Reader, *tagMap) error

		tagFileTypeLow := strings.ToLower(*tagFileType)

		switch tagFileTypeLow {
		case "json":
			decode = decoderJson
		case "toml":
			decode = decoderToml
		default:
			return fmt.Errorf("Invalid file type '%s'", *tagFileType)
		}

		f, err := os.Open(*tagFile)
		if err != nil {
			return err
		}

		defer f.Close()

		err = decode(f, &tags)
		if err != nil {
			return err
		}
	} else {
		if *imgpath != "" {
			tags[imgpathKey] = []string{*imgpath}
		}

		tags.add(taglib.Title, tTitle)
		tags.add(taglib.Genre, tGenre)
		tags.add(taglib.Album, tAlbum)
		tags.add(taglib.Artist, tArtist)
		tags.add(taglib.Composer, tComposer)
		tags.add(taglib.Date, tDate)
	}

	if img, ok := tags[imgpathKey]; ok {
		delete(tags, imgpathKey)

		if len(img) > 0 {
			data, err := os.ReadFile(img[0])
			if err != nil {
				return err
			}

			err = taglib.WriteImage(file, data)
			if err != nil {
				return err
			}
		}
	}

	return taglib.WriteTags(file, tags, 0)
}

func decoderJson(r io.Reader, tags *tagMap) error {
	enc := json.NewDecoder(r)

	return enc.Decode(tags)
}

func decoderToml(r io.Reader, tags *tagMap) error {
	enc := toml.NewDecoder(r)

	_, err := enc.Decode(tags)
	return err
}

/*
func decodeImage(path string) (image.Image, error) {
	ext := filepath.Ext(path)

	var decode func(io.Reader) (image.Image, error)

	switch ext {
	case ".png":
		decode = png.Decode
	case ".jpg", ".jpeg":
		decode = jpeg.Decode
	case ".gif":
		decode = gif.Decode
	default:
		return nil, fmt.Errorf("Invalid image file extension '%s'", ext)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return decode(f)
}
*/
