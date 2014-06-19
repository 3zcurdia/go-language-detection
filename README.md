go-language-detection
=====================

A language detection library for Go.

## Installation instructions
First install the package :
```go
go get github.com/AntoineFinkelstein/go-language-detection/
```

The first time the library is called, it will download all the wordlists, create bloom filters and save time to save time later. Therefore, make sure you the package can write at `~/tmp/`.

## Usage

```go
func main() {
  var detect goLanguageDetection.Detect = *goLanguageDetection.New()
  
  result, validity := detect.Text("Wikipedia is a collaboratively edited, multilingual, free-access, free content Internet encyclopedia that is supported and hosted by the non-profit Wikimedia Foundation. Volunteers worldwide collaboratively write Wikipedia's 30 million articles in 287 languages, including in the English Wikipedia. Anyone who can access the site can edit almost any of its articles, which on the Internet comprise the largest and most popular general reference work.")

  fmt.Println(result) // English
  fmt.Println(validity) // Ratio of words found in the returned language : 0.92
}
```

## Todo list

Help welcomed :-)

- [x] Use bloom filters for better performances
- [x] Allow the language files to be included in the binaries
- [x] Find a way not to read the language files everytime
- [ ] Write a few tests
