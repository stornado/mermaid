package mermaid

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"
	"net/http/httptest"
	"text/template"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/go-chi/chi"
)

//go:embed mermaid.min.js
var mermaidjs []byte

//go:embed mermaid.min.js.map
var mermaidjsmap []byte

//go:embed index.html
var index string

func RenderWithTheme(mermaidSource, mermaidTheme string) (string, []byte, error) {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("mermaid").Parse(index)
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		err = t.Execute(&buf, map[string]string{
			"MermaidSource": mermaidSource,
			"MermaidTheme":  mermaidTheme,
		})
		if err != nil {
			panic(err)
		}
		_, err = w.Write(buf.Bytes())
		if err != nil {
			panic(err)
		}
	})
	r.Get("/mermaid.min.js", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(mermaidjs)
		if err != nil {
			panic(err)
		}
	})
	r.Get("/mermaid.min.js.map", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(mermaidjsmap)
		if err != nil {
			panic(err)
		}
	})
	ts := httptest.NewServer(r)
	defer ts.Close()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var svg string
	var png []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.EmulateViewport(1920, 5780),
		chromedp.Reload(),
		chromedp.ScrollIntoView("svg", chromedp.ByQuery),
		chromedp.WaitReady("svg", chromedp.ByQuery),
		chromedp.OuterHTML("svg", &svg, chromedp.NodeReady, chromedp.ByQuery),
		chromedp.Screenshot("svg", &png, chromedp.NodeReady, chromedp.ByQuery),
	)
	if err != nil {
		return "", nil, err
	}

	return svg, png, nil
}

func Render(mermaidSource string) (string, []byte, error) {
	return RenderWithTheme(mermaidSource, "default")
}

func RenderDark(mermaidSource string) (string, []byte, error) {
	return RenderWithTheme(mermaidSource, "dark")
}

func RenderForest(mermaidSource string) (string, []byte, error) {
	return RenderWithTheme(mermaidSource, "forest")
}

func RenderNeutral(mermaidSource string) (string, []byte, error) {
	return RenderWithTheme(mermaidSource, "neutral")
}
