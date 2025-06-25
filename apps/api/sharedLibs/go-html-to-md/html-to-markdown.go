package main

import (
	"C"
	"log"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
)
import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/table"
)

//export ConvertHTMLToMarkdown
func ConvertHTMLToMarkdown(html *C.char) *C.char {
	conv := converter.NewConverter(
		converter.WithPlugins(
			base.NewBasePlugin(),
			commonmark.NewCommonmarkPlugin(
				commonmark.WithStrongDelimiter("**"),
			),
		),
	)
	// https://github.com/JohannesKaufmann/html-to-markdown/blob/9c53576b59f527f91e1328981a1f454de605b9ed/cli/html2markdown/cmd/cmd_convert.go#L95
	conv.Register.Plugin(
		table.NewTablePlugin(
			table.WithSkipEmptyRows(false),
			table.WithHeaderPromotion(false),
			table.WithSpanCellBehavior(table.SpanCellBehavior(table.SpanBehaviorEmpty)),
			table.WithPresentationTables(false),
			table.WithNewlineBehavior(table.NewlineBehavior(table.NewlineBehaviorPreserve)),
		),
	)

	markdown, err := htmltomarkdown.ConvertString(C.GoString(html))
	if err != nil {
		log.Fatal(err)
	}
	return C.CString(markdown)
}

func main() {
	// This function is required for the main package
}
