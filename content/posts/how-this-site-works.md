---
title: How This Site Works
description: From Markdown file to finished page, and how to add your own.
date: 2026-01-05
---

A single Go program reads the Markdown files in `content/`, puts them into the templates in `templates/`, and writes plain HTML to `public/`. There is no server or database.

## Add a post

Create `content/posts/my-post.md`:

```markdown
---
title: My Post
description: Shown in link previews
date: 2026-02-01
---

Write anything here.
```

The file name is the URL: this one is published at `/posts/my-post/`. The `date` is optional, and posts are listed newest first.

## Add a page

Pages such as About and Contact go in `content/pages/`, use the same front matter (without a date), and are linked from the menu, sorted by title. `about.md` is published at `/about/`.

## Build

```sh
go run .
```

## Styling

`static/css/style.css` is a plain, hand-written stylesheet with comments. Edit it directly and reload. Colours are variables at the top, and dark mode follows the visitor's system setting.
