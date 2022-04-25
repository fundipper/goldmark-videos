package videos

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// VideoAttributeFilter defines attribute names which video elements can have.
var VideoAttributeFilter = html.ImageAttributeFilter.Extend(
	[]byte("scrolling"),
	[]byte("frameborder"),
	[]byte("framespacing"),
	[]byte("allowfullscreen"),
)

// Renderer struct is a renderer.NodeRenderer implementation for the extension.
type Renderer struct {
	html.Config
}

// NewRenderer builds a new Renderer with given options and returns it.
func NewRenderer() renderer.NodeRenderer {
	return &Renderer{}
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs.
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindVideo, r.renderVideo)
}

func (r *Renderer) renderVideo(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*Video)
	for k, v := range extender.Attribute {
		n.SetAttributeString(k, util.StringToReadOnlyBytes(v))
	}

	if entering {
		_, _ = w.WriteString("<iframe src=\"")
		if r.Unsafe || !html.IsDangerousURL(n.Destination) {
			_, _ = w.Write(util.EscapeHTML(util.URLEscape(n.Destination, true)))
		}
		_ = w.WriteByte('"')
		if n.Attributes() != nil {
			html.RenderAttributes(w, n, VideoAttributeFilter)
		}
		_ = w.WriteByte('>')
	} else {
		_, _ = w.WriteString("</iframe>")
	}
	return ast.WalkContinue, nil
}
