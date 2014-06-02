go-language-detection
=====================

A language detection library for Go.

**This library is currently in development**, although it can be used in it's current form.

## Installation instructions
You can't import (at least for now) this library has you would normally. The files containing the languages aren't .go files. Therefore it is required to download the repository, and import it from where you need it.

## Usage

```go
package main

import (
  "fmt"
  "github.com/AntoineFinkelstein/go-language-detection"
)

func main() {
  result, validity := goLanguageDetection.Find("Dude where's my car ? towel change awesome")
  fmt.Println(result) // => English
  fmt.Println(validity) // Percentage of words found in the returned language
}
```

## Todo list

Help welcomed :-)

- [ ] Use bloom filters for better performances
- [ ] Allow the language files to be built (this package will the work like any other one)
- [ ] Write a few tests
