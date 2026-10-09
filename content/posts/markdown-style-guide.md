---
title: Style Guide
description: Every element this site supports.
date: 2026-09-30
---

Body text is set in Inter at a comfortable measure, in black on white, with red kept for interaction.{sn|A sidenote sits in the right margin on wide screens. On narrow ones it opens inline when you tap the number.} Links are [underlined](/about/) and turn red on hover.

## Notes

A sidenote is numbered and tied to a word.{sn|Numbers are assigned automatically, in the order they appear.} A margin note is unnumbered.{mn|Like this one. It carries context without asking for attention.} Both are written in Markdown as `{sn|text}` and `{mn|text}`.

## Quotes

<blockquote><p>Simplicity is the ultimate sophistication.</p><footer>Attributed to Leonardo da Vinci</footer></blockquote>

## Lists

- Unordered lists stay plain
- With no decoration
- Beyond the bullet

1. Ordered lists
2. Work the same way

## Code

Inline code like `go run .` and blocks:

```go
func main() {
	fmt.Println("hello, world")
}
```

## Tables

| Element  | Weight | Size    |
|----------|-------:|--------:|
| Title    |    700 | 2.5rem  |
| Heading  |    700 | 1.25rem |
| Body     |    400 | 1rem    |

## Figures

<figure>
<img src="/images/bell-curve.svg" alt="A normal distribution with the mean and one standard deviation marked">
<figcaption>A figure with a caption: the normal distribution, with the mean and one standard deviation marked.</figcaption>
</figure>

A small figure can sit in the margin instead.{mn|<img src="/images/sparks.svg" alt="Four small multiples"><br>A margin figure: four small multiples.}
