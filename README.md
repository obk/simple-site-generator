# simple-site-generator

A tiny static site generator written in Go. It turns a folder of Markdown files into a plain HTML website that you can host anywhere. There is no server, database or JavaScript build step.


## What it does

1. Reads Markdown files with front matter (title, description, date) from `content/`.
2. Renders them into the HTML templates in `templates/`.
3. Writes the finished site to `public/`, together with everything in `static/`.

```
content/posts/*.md  ─┐
content/pages/*.md  ─┼─►  main.go  ─►  public/
templates/*.html    ─┤
static/**           ─┘
```

Features:

- **Posts**, listed on the home page, newest first.
- **Pages** such as About and Contact, linked in the footer menu.
- **Markdown** with tables, task lists, strikethrough and raw HTML.
- **Sidenotes and margin notes**, written as `{sn|text}` and `{mn|text}`. They sit in the margin on wide screens and open inline on narrow ones.

## Use

It is a single Go file, with one dependency ([goldmark](https://github.com/yuin/goldmark)).

Requires Go 1.24+.

```sh
go run .                              # build into public/
cd public && python3 -m http.server   # preview
```

## Deploy

Upload `public/` to any static host eg: Sourcehut Pages or Codeberg Pages.

## Content

**Posts** go in `content/posts/`, are listed on the home page, and are published at `/posts/<file-name>/`:

```markdown
---
title: My Post
description: Shown in link previews   (optional)
date: 2026-01-01                      (optional, YYYY-MM-DD)
---

Markdown here.
```

**Pages** (About, Contact...) go in `content/pages/`, use the same front matter without a date, are published at `/<file-name>/`, and are linked in the menu (sorted by title).

Files in `static/` are copied to `public/` as-is.

## Theme

The look is Swiss minimal: Inter, black on white, and a single red accent that is kept for interaction (link hover, selection, note numbers). The stylesheet lowercases all text with one `text-transform` rule, so you can write content with standard grammar. Code stays as written.

- **Colours** are four CSS variables at the top of `static/css/style.css`: `--fg`, `--bg`, `--muted` and `--accent`. Change them there and the whole site follows.
- **Dark mode** follows the visitor's system setting. It swaps the same variables, and SVG images loaded with `<img>` are inverted so they stay readable.
- **Font** is Inter, loaded from `static/fonts/InterVariable.woff2`. That file is not included, so add it yourself. Without it the site falls back to Helvetica or Arial.
- **Layout** is one centred column, 44rem wide. On screens wider than 76rem, sidenotes and margin notes move into the right margin. On narrow screens they open inline when you tap the number.
- **Elements** covered: headings, links, lists, blockquotes, code, tables, figures with captions, and margin figures. The style guide post shows every one of them.

Edit `static/css/style.css` (plain CSS, with comments) and `templates/*.html`. The site name ("My Site") is set in `templates/base.html` and `main.go`.

## License

[MIT](LICENSE)
