package videos

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (a *Transformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	replaceImages := func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if n.Kind() != ast.KindImage {
			return ast.WalkContinue, nil
		}

		img := n.(*ast.Image)
		u, err := url.Parse(string(img.Destination))
		if err != nil {
			msg := ast.NewString([]byte(fmt.Sprintf("<!-- %s -->", err)))
			msg.SetCode(true)
			n.Parent().InsertAfter(n.Parent(), n, msg)
			return ast.WalkContinue, nil
		}

		var ok bool
		for _, v := range extender.Options {
			if u.Host == v.Host && strings.HasPrefix(u.Path, v.Path) {

				ok = true
				break
			}
		}

		if !ok {
			return ast.WalkContinue, nil
		}

		v := NewVideo(img)
		n.Parent().ReplaceChild(n.Parent(), n, v)

		return ast.WalkContinue, nil
	}

	ast.Walk(node, replaceImages)
}
