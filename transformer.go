package videos

import (
	"fmt"
	"net/url"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (a *Transformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if n.Kind() != ast.KindImage {
			return ast.WalkContinue, nil
		}

		image := n.(*ast.Image)
		u, err := url.Parse(string(image.Destination))
		if err != nil {
			msg := ast.NewString([]byte(fmt.Sprintf("<!-- %s -->", err)))
			msg.SetCode(true)
			n.Parent().InsertAfter(n.Parent(), n, msg)
			return ast.WalkContinue, nil
		}

		path, ok := extender.Source[u.Host]
		if !ok || u.Path != path {
			return ast.WalkContinue, nil
		}

		n.Parent().ReplaceChild(n.Parent(), n, NewVideo(image))

		return ast.WalkContinue, nil
	})
}
