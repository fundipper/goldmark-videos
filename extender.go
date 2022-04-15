package videos

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var extender *Extender

// Option is a functional option type for this extension.
type Option struct {
	Host string
	Path string
}

type Extender struct {
	Options []Option
}

// New returns a new Embed extension.
func NewExtender(opts ...Option) goldmark.Extender {
	extender = &Extender{
		Options: opts,
	}
	return extender
}

func (e *Extender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(NewTransformer(), 500),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewRenderer(), 500),
		),
	)
}
