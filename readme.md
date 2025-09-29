# Tagline

A utilty for reading and writing audio file metadata.

## Example

```bash
go run . write Towns.mp3 -f Towns.toml
```

## Installation

Installation requires the Go compiler, which can be downloaded at [go.dev](https://go.dev)

```bash
go install github.com/voidwyrm-2/tagline@latest
tagline help
```

## Building

```bash
git clone https://github.com/voidwyrm-2/tagline
cd tagline
go build -o tagline .
./tagline help
```

## Licensing

`WellTraveledCompanions.png` and `Towns.mp3` are licensed under:

`Copyright (c) 2025 Nuclear Pasta, all rights reserved.`

Everything else in this repo is licensed under:
```
MIT License

Copyright (c) 2025 Nuclear Pasta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
