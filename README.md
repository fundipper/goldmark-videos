# goldmark-videos

goldmark-embed is an extension for the [goldmark][goldmark] library that extends
the Markdown `![]()` image embed syntax to support additional media formats.

~~YouTube only at first.~~

Base on [goldmark-embed](https://github.com/fundipper/goldmark-embed), Support customize option to embed video use `ifmarme`.

## Demo

This markdown:

```md
# Hello goldmark-videos

![](https://www.youtube.com/embed/LDflrf85h9Y)

![](https://v.qq.com/txp/iframe/player.html?vid=i0042v2fm34)

![](//player.bilibili.com/player.html?aid=634140852&bvid=BV1sb4y1t7xV&cid=442265383&page=1)
```

Becomes this HTML:

```html
<h1>Hello goldmark-videos</h1>
<p>
	<iframe
		width="560"
		height="315"
		src="https://www.youtube.com/embed/dQw4w9WgXcQ"
		frameborder="0"
		allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
		allowfullscreen
	></iframe>
</p>

<p>
	<iframe
		width="560"
		height="315"
		src="https://v.qq.com/txp/iframe/player.html?vid=i0042v2fm34"
		frameborder="0"
		allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
		allowfullscreen
	></iframe>
</p>

<p>
	<iframe
		width="560"
		height="315"
		src="//player.bilibili.com/player.html?aid=634140852&bvid=BV1sb4y1t7xV&cid=442265383&page=1"
		frameborder="0"
		allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
		allowfullscreen
	></iframe>
</p>
```

### Installation

```bash
go get github.com/fundipper/goldmark-videos
```

## Usage

```go
  markdown := goldmark.New(
	goldmark.WithExtensions(
		videos.NewExtender(
			map[string]string{
				"www.youtube.com":     "/embed",
				"v.qq.com":            "/txp/iframe/player.html",
				"player.bilibili.com": "/player.html",
			},
			map[string]string{
				"width":           "560",
				"height":          "315",
				"border":          "0",
				"scrolling":       "no",
				"frameborder":     "no",
				"framespacing":    "0",
				"allowfullscreen": "true",	
			},
		),
	),
)
  var buf bytes.Buffer
  if err := markdown.Convert([]byte(source), &buf); err != nil {
    panic(err)
  }
  fmt.Print(buf)
}
```

## thanks

[Goldmark](https://github.com/yuin/goldmark)

[13rac1](https://github.com/13rac1/goldmark-embed)
