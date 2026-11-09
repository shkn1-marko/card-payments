// Command card-payments serves the research chapters as a small website.
//
// Chapters are Markdown files in content/chapters, named "<n>-<slug>.md".
// They are rendered once at startup and served from memory.
package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

//go:embed content templates static
var files embed.FS

type chapter struct {
	Number   int
	Slug     string
	Title    string
	Sections []section // the chapter's "## " headings, for the contents sidebar
	HTML     template.HTML
}

type section struct {
	ID    string
	Title string
}

type site struct {
	base     string // URL prefix the site is served under, e.g. "/card-payments"
	chapters []chapter
	bySlug   map[string]*chapter
	tmpl     *template.Template
}

var chapterName = regexp.MustCompile(`^(\d+)-([a-z0-9-]+)\.md$`)

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "address to listen on")
	base := flag.String("base", "", `URL prefix used in links, e.g. "/card-payments" (empty for root)`)
	flag.Parse()

	s, err := newSite(strings.TrimRight(*base, "/"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("serving %d chapter(s) on http://%s%s/", len(s.chapters), *addr, s.base)
	log.Fatal(http.ListenAndServe(*addr, s.routes()))
}

func newSite(base string) (*site, error) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	)

	entries, err := fs.ReadDir(files, "content/chapters")
	if err != nil {
		return nil, err
	}

	s := &site{base: base, bySlug: map[string]*chapter{}}
	for _, e := range entries {
		m := chapterName.FindStringSubmatch(e.Name())
		if e.IsDir() || m == nil {
			continue
		}
		src, err := files.ReadFile(path.Join("content/chapters", e.Name()))
		if err != nil {
			return nil, err
		}
		doc := md.Parser().Parse(text.NewReader(src))
		var buf bytes.Buffer
		if err := md.Renderer().Render(&buf, src, doc); err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		n, _ := strconv.Atoi(m[1])
		s.chapters = append(s.chapters, chapter{
			Number:   n,
			Slug:     m[2],
			Title:    titleOf(src, m[2]),
			Sections: sectionsOf(doc, src),
			HTML:     template.HTML(buf.String()),
		})
	}
	sort.Slice(s.chapters, func(i, j int) bool { return s.chapters[i].Number < s.chapters[j].Number })
	for i := range s.chapters {
		s.bySlug[s.chapters[i].Slug] = &s.chapters[i]
	}

	s.tmpl, err = template.ParseFS(files, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// titleOf returns the text of the first "# " heading, without a leading "1. " number.
func titleOf(src []byte, fallback string) string {
	for _, line := range strings.Split(string(src), "\n") {
		if t, ok := strings.CutPrefix(strings.TrimSpace(line), "# "); ok {
			if i := strings.Index(t, ". "); i > 0 {
				if _, err := strconv.Atoi(t[:i]); err == nil {
					t = t[i+2:]
				}
			}
			return t
		}
	}
	return fallback
}

// sectionsOf lists the level-2 headings of a parsed chapter with their anchor IDs.
func sectionsOf(doc ast.Node, src []byte) []section {
	var out []section
	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		h, ok := n.(*ast.Heading)
		if !entering || !ok || h.Level != 2 {
			return ast.WalkContinue, nil
		}
		id, _ := h.AttributeString("id")
		idb, _ := id.([]byte)
		out = append(out, section{ID: string(idb), Title: plainText(h, src)})
		return ast.WalkSkipChildren, nil
	})
	return out
}

// plainText returns the text inside a node, without formatting.
func plainText(n ast.Node, src []byte) string {
	var b strings.Builder
	ast.Walk(n, func(c ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch t := c.(type) {
		case *ast.Text:
			b.Write(t.Segment.Value(src))
			if t.SoftLineBreak() {
				b.WriteByte(' ')
			}
		case *ast.String:
			b.Write(t.Value)
		}
		return ast.WalkContinue, nil
	})
	return b.String()
}

func (s *site) routes() http.Handler {
	mux := http.NewServeMux()
	static, _ := fs.Sub(files, "static")
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(static)))
	mux.HandleFunc("GET /{$}", s.index)
	mux.HandleFunc("GET /{slug}", s.chapter)

	if s.base == "" {
		return mux
	}
	// Serve under the prefix too, so links work when run without nginx.
	outer := http.NewServeMux()
	outer.Handle(s.base+"/", http.StripPrefix(s.base, mux))
	outer.Handle("/", mux)
	return outer
}

type pageData struct {
	Base     string
	Title    string
	Chapters []chapter
	Chapter  *chapter
}

func (s *site) index(w http.ResponseWriter, r *http.Request) {
	s.render(w, "index.html", pageData{Base: s.base, Title: "Online Card Payments", Chapters: s.chapters})
}

func (s *site) chapter(w http.ResponseWriter, r *http.Request) {
	c, ok := s.bySlug[r.PathValue("slug")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	s.render(w, "chapter.html", pageData{Base: s.base, Title: c.Title, Chapters: s.chapters, Chapter: c})
}

func (s *site) render(w http.ResponseWriter, name string, d pageData) {
	var buf bytes.Buffer
	if err := s.tmpl.ExecuteTemplate(&buf, name, d); err != nil {
		log.Printf("render %s: %v", name, err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
