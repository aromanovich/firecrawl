package main

import (
	"C"
	// "log"
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
)

//export ConvertHTMLToMarkdown
func ConvertHTMLToMarkdown(html *C.char) *C.char {
	converter := converter.NewConverter()

	markdown, err := converter.ConvertString(C.GoString(html))
	if err != nil {
		// log.Fatal(err)
	}
	return C.CString(markdown)
}

func main() {
	// This function is required for the main package
}
