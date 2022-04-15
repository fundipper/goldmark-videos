package videos

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Renderer struct is a renderer.NodeRenderer implementation for the extension.
type Renderer struct{}

// NewRenderer builds a new Renderer with given options and returns it.
func NewRenderer() renderer.NodeRenderer {
	r := &Renderer{}
	return r
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs.
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindVideo, r.renderVideo)
}

func (r *Renderer) renderVideo(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		return ast.WalkContinue, nil
	}

	v := node.(*Video)
	w.WriteString(`<iframe width="560" height="315" src="`)
	w.Write(v.Destination)
	w.WriteString(`" scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>`)
	return ast.WalkContinue, nil
}
