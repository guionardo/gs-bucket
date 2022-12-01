package handlers

import (
	"io"
	"net/http"
	"strings"

	_ "embed"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/guionardo/gs-bucket/pkg/config"
)

//go:embed markdown.css
var markdownCss []byte

var _home []byte

func getHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write(_home)
	w.WriteHeader(http.StatusOK)
}

func getMarkdownCssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(markdownCss)
	w.WriteHeader(http.StatusOK)
}

func renderHookDropCodeBlock(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	// skip all nodes that are not CodeBlock nodes
	if _, ok := node.(*ast.CodeBlock); !ok {
		return ast.GoToNext, false
	}
	// custom rendering logic for ast.CodeBlock. By doing nothing it won't be
	// present in the output
	return ast.GoToNext, true
}
func SetupHome(home []byte) {
	extensions := parser.CommonExtensions
	parser := parser.NewWithExtensions(extensions)
	opts := html.RendererOptions{
		Flags:          html.CommonFlags | html.CompletePage,
		RenderNodeHook: renderHookDropCodeBlock,
		Title:          "GS Bucket",
		CSS:            "/markdown.css",
	}
	renderer := html.NewRenderer(opts)
	body := strings.ReplaceAll(string(markdown.ToHTML(home, parser, renderer)), "$THISVERSION$", config.Version)

	_home = []byte(body)
}
