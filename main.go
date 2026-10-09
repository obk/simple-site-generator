package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type Post struct {
	Title, Description, Slug string
	Time                     time.Time
	Body                     template.HTML
}

func (p Post) Date() string { return p.Time.Format("January 2, 2006") }

type Page struct {
	Title, Description string
	Posts, Nav         []Post
	Post
}

var nav []Post

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var notes = regexp.MustCompile(`\{(sn|mn)\|([^}]*)\}`)

// notes turns {sn|text} and {mn|text} into Tufte-style side and margin notes.
func addNotes(md string) string {
	n := 0
	return notes.ReplaceAllStringFunc(md, func(m string) string {
		n++
		g := notes.FindStringSubmatch(m)
		id := fmt.Sprintf("%s-%d", g[1], n)
		label, class := "&#8853;", "marginnote"
		toggle := "margin-toggle"
		if g[1] == "sn" {
			label, class = "", "sidenote"
			toggle = "margin-toggle sidenote-number"
		}
		return fmt.Sprintf(`<label for="%s" class="%s">%s</label><input type="checkbox" id="%s" class="margin-toggle"><span class="%s">%s</span>`,
			id, toggle, label, id, class, g[2])
	})
}

func parse(path string) Post {
	raw, err := os.ReadFile(path)
	must(err)
	p := Post{Slug: strings.TrimSuffix(filepath.Base(path), ".md")}
	parts := strings.SplitN(string(raw), "---\n", 3)
	if len(parts) != 3 || parts[0] != "" {
		log.Fatalf("%s: missing front matter", path)
	}
	for _, line := range strings.Split(parts[1], "\n") {
		k, v, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		v = strings.TrimSpace(v)
		switch strings.TrimSpace(k) {
		case "title":
			p.Title = v
		case "description":
			p.Description = v
		case "date":
			p.Time, err = time.Parse("2006-01-02", v)
			must(err)
		}
	}
	var buf bytes.Buffer
	md := goldmark.New(goldmark.WithExtensions(extension.GFM), goldmark.WithRendererOptions(html.WithUnsafe()))
	must(md.Convert([]byte(addNotes(parts[2])), &buf))
	p.Body = template.HTML(buf.String())
	return p
}

func render(tpl, out string, data Page) {
	t := template.Must(template.ParseFiles("templates/base.html", "templates/"+tpl))
	data.Nav = nav
	must(os.MkdirAll(filepath.Dir(out), 0o755))
	f, err := os.Create(out)
	must(err)
	defer f.Close()
	must(t.ExecuteTemplate(f, "base", data))
}

func copyDir(src, dst string) {
	must(filepath.WalkDir(src, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, p)
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		return os.WriteFile(target, b, 0o644)
	}))
}

func load(dir string) []Post {
	files, err := filepath.Glob(dir + "/*.md")
	must(err)
	var out []Post
	for _, f := range files {
		out = append(out, parse(f))
	}
	return out
}

func main() {
	must(os.RemoveAll("public"))
	copyDir("static", "public")

	posts := load("content/posts")
	nav = load("content/pages")
	sort.Slice(nav, func(i, j int) bool { return nav[i].Title < nav[j].Title })
	sort.Slice(posts, func(i, j int) bool { return posts[i].Time.After(posts[j].Time) })

	for _, p := range posts {
		render("post.html", filepath.Join("public", "posts", p.Slug, "index.html"),
			Page{Title: p.Title, Description: p.Description, Post: p})
	}
	for _, p := range nav {
		render("page.html", filepath.Join("public", p.Slug, "index.html"),
			Page{Title: p.Title, Description: p.Description, Post: p})
	}
	render("index.html", "public/index.html", Page{Title: "My Site", Posts: posts})
	log.Printf("built %d posts, %d pages -> public/", len(posts), len(nav))
}
