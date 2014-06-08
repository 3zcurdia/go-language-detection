go-language-detection
=====================

A language detection library for Go.

## Installation instructions
First install the package :
```go
go get github.com/AntoineFinkelstein/go-language/
```

The first time the library is called, it will download all the wordlists, create bloom filters and save time to save time later. Therefore, make sure you the package can write at `~/tmp/`.

## Usage

```go
package main

import (
  "fmt"
  "github.com/AntoineFinkelstein/go-language-detection"
)

func main() {
  result, validity := goLanguageDetection.Find("Dude where's my car ?")
  fmt.Println(result) // => English
  fmt.Println(validity) // Percentage of words found in the returned language
}
```

## Todo list

Help welcomed :-)

- [x] Use bloom filters for better performances
- [x] Allow the language files to be included in the binaries
- [x] Find a way not to read the language files everytime
- [ ] Write a few tests
