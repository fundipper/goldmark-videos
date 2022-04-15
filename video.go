package videos

import (
	"github.com/yuin/goldmark/ast"
)

// KindVideo is a NodeKind of the video node.
var KindVideo = ast.NewNodeKind("Video")

// Video struct represents a video of the Markdown text.
type Video struct {
	ast.Image
}

// NewVideo returns a new video node.
func NewVideo(img *ast.Image) *Video {
	c := &Video{
		Image: *img,
	}

	c.Destination = img.Destination
	c.Title = img.Title
	return c
}

// Kind implements Node.Kind.
func (v *Video) Kind() ast.NodeKind {
	return KindVideo
}
