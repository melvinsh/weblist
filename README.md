# weblist

A Go command-line tool that analyzes web pages and lists all words sorted by frequency of occurrence.

## Description

`weblist` fetches a web page using a headless browser (to render JavaScript content), extracts all text content, and outputs a sorted list of words by their frequency of appearance. Words are case-insensitive and sorted from most frequent to least frequent.

## Features

- Renders JavaScript-heavy websites using headless Chrome via [go-rod](https://github.com/go-rod/rod)
- Extracts text content from HTML
- Counts word frequency
- Outputs words sorted by frequency (descending order)

## Requirements

- Go 1.23.2 or later
- Chrome/Chromium browser (used by rod for headless browsing)

## Installation

```bash
git clone https://github.com/melvinsh/weblist.git
cd weblist
go mod download
```

## Usage

```bash
go run main.go <URL>
```

### Example

```bash
go run main.go https://example.com
```

This will output all words found on the page, sorted by frequency:

```
the
example
domain
this
...
```

## Building

To build a standalone binary:

```bash
go build -o weblist
./weblist https://example.com
```

## Dependencies

- [go-rod/rod](https://github.com/go-rod/rod) - High-level browser automation library
- [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html) - HTML parsing

## License

See LICENSE file for details.
