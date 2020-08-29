package godocgen

import (
	"fmt"
	"strings"
)

// Markdown is a collection of blocks
type Markdown struct {
	Blocks []Blockable
}

// NewMarkdown creates an empty markdown object
func NewMarkdown() *Markdown {
	return &Markdown{
		Blocks: []Blockable{},
	}
}

// Prepend prepends a block to Markdown
func (d *Markdown) Prepend(b ...Blockable) {
	d.Blocks = append(b, d.Blocks...)
}

// Append adds a block to Markdown
func (d *Markdown) Append(b Blockable) {
	d.Blocks = append(d.Blocks, b)
}

// Content returns a string representation of the markdown object
func (d *Markdown) Content() string {
	output := ""

	for _, b := range d.Blocks {
		output = fmt.Sprintf("%s\n%s", output, b.ToString())
	}

	return output
}

// Blockable is implement by content blocks
type Blockable interface {
	ToString() string
}

// Heading1 is H1
type Heading1 string

// ToString returns markdown string
func (h Heading1) ToString() string {
	return "# " + string(h) + "\n"
}

// Heading2 is H2
type Heading2 string

// ToString returns markdown string
func (h Heading2) ToString() string {
	return "## " + string(h) + "\n"
}

// Heading3 is H3
type Heading3 string

// ToString returns markdown string
func (h Heading3) ToString() string {
	return "### " + string(h) + "\n"
}

// Heading4 is H4
type Heading4 string

// ToString returns markdown string
func (h Heading4) ToString() string {
	return "#### " + string(h) + "\n"
}

// Content is plain text
type Content string

// ToString returns markdown string
func (c Content) ToString() string {
	return string(c)
}

// GoCode is a Go code block
type GoCode string

// ToString returns code block wrapped in `pre` quotes
func (c GoCode) ToString() string {
	return fmt.Sprintf("```go\n%s\n```", string(c))
}

// Index is a list of all labels in the document
type Index []IndexEntry

// ToString returns index serialized as a list
func (i Index) ToString() string {
	l := ""

	for _, e := range i {
		l += fmt.Sprintf("%s- %s\n", strings.Repeat("\t", e.Level), e.Label)
	}

	return l
}

// IndexEntry is an entry that is rendered as a list item
type IndexEntry struct {
	Label string
	Level int
}

// NewIndex creates a new index from all headings in the markdown doc
func NewIndex(md *Markdown) Index {
	idx := make(Index, 0)

	for _, b := range md.Blocks {
		switch b.(type) {
		case Heading2:
			idx = append(idx, IndexEntry{Label: string(b.(Heading2)), Level: 0})
		case Heading3:
			idx = append(idx, IndexEntry{Label: string(b.(Heading3)), Level: 1})
		case Heading4:
			idx = append(idx, IndexEntry{Label: string(b.(Heading4)), Level: 2})
		}
	}

	return idx
}
