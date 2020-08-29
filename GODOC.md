# godocgen

Package godocgen parses Go packages and indexes constants, functions and types.

`godocgen` keeps your private Go repositories documented with the same source code & repository.

#### Index

- Constants
- Functions
  - func GetSourcePackages
  - func WriteMarkdown
- Types
  - type Blockable
  - type Const
  - type Content
    - func (Content) ToString
  - type Example
  - type Func
  - type GoCode
    - func (GoCode) ToString
  - type Heading1
    - func (Heading1) ToString
  - type Heading2
    - func (Heading2) ToString
  - type Heading3
    - func (Heading3) ToString
  - type Heading4
    - func (Heading4) ToString
  - type Index
    - func (Index) ToString
  - type IndexEntry
  - type Markdown
    - func (\*Markdown) Append
    - func (\*Markdown) Content
    - func (\*Markdown) Prepend
  - type PackageIndex
  - type Type

## Constants

test constats

```go
const (
	A	= iota
	B
)
```

## Functions

### func GetSourcePackages

```go
func GetSourcePackages(path string) ([]string, error)
```

GetSourcePackages returns all package directories in path
excludes vendor directories

### func WriteMarkdown

```go
func WriteMarkdown(idx *PackageIndex, outPath string) error
```

WriteMarkdown writes indexed documents as markdown content

## Types

### type Blockable

Blockable is implement by content blocks

```go
type Blockable interface {
	ToString() string
}
```

### type Const

Const is a single or a block of constants in a package

```go
type Const struct {
	Doc	string
	Snippet	string
}
```

### type Content

Content is plain text

```go
type Content string
```

#### func (Content) ToString

```go
func (c Content) ToString() string
```

ToString returns markdown string

### type Example

Example demonstrates use of functions & types in a package

```go
type Example struct {
	Name	string
	Suffix	string
	Doc	string
	Snippet	string
}
```

### type Func

Func is a package level function

```go
type Func struct {
	Name		string
	Doc		string
	Snippet		string
	Recv		string
	Examples	[]Example
}
```

### type GoCode

GoCode is a Go code block

```go
type GoCode string
```

#### func (GoCode) ToString

```go
func (c GoCode) ToString() string
```

ToString returns code block wrapped in `pre` quotes

### type Heading1

Heading1 is H1

```go
type Heading1 string
```

#### func (Heading1) ToString

```go
func (h Heading1) ToString() string
```

ToString returns markdown string

### type Heading2

Heading2 is H2

```go
type Heading2 string
```

#### func (Heading2) ToString

```go
func (h Heading2) ToString() string
```

ToString returns markdown string

### type Heading3

Heading3 is H3

```go
type Heading3 string
```

#### func (Heading3) ToString

```go
func (h Heading3) ToString() string
```

ToString returns markdown string

### type Heading4

Heading4 is H4

```go
type Heading4 string
```

#### func (Heading4) ToString

```go
func (h Heading4) ToString() string
```

ToString returns markdown string

### type Index

Index is a list of all labels in the document

```go
type Index []IndexEntry
```

#### func (Index) ToString

```go
func (i Index) ToString() string
```

ToString returns index serialized as a list

### type IndexEntry

IndexEntry is an entry that is rendered as a list item

```go
type IndexEntry struct {
	Label	string
	Level	int
}
```

### type Markdown

Markdown is a collection of blocks

```go
type Markdown struct {
	Blocks []Blockable
}
```

#### func (\*Markdown) Append

```go
func (d *Markdown) Append(b Blockable)
```

Append adds a block to Markdown

#### func (\*Markdown) Content

```go
func (d *Markdown) Content() string
```

Content returns a string representation of the markdown object

#### func (\*Markdown) Prepend

```go
func (d *Markdown) Prepend(b ...Blockable)
```

Prepend prepends a block to Markdown

### type PackageIndex

PackageIndex contains all parsed types & functions information

```go
type PackageIndex struct {
	Name	string
	Dir	string
	Doc	string

	Consts		[]Const
	Funcs		[]Func
	Types		[]Type
	Examples	[]Example
}
```

### type Type

Type captures public types available at package level

```go
type Type struct {
	Name		string
	Doc		string
	Snippet		string
	Consts		[]Const
	Funcs		[]Func
	Methods		[]Func
	Examples	[]Example
}
```
