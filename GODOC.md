
# godocgen

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package godocgen parses Go packages and indexes constants, functions and types.

`godocgen` keeps your private Go repositories documented with the same source code & repository.



## <a name="pkg-index">Index</a>


* [func GetSourcePackages(path string) ([]string, error)](#getsourcepackages)
* [func ToMD(w io.Writer, text string)](#tomd)
* [func WriteMDTemplate(idx *PackageIndex, outPath string) error](#writemdtemplate)
* [func WriteMarkdown(idx *PackageIndex, outPath string) error](#writemarkdown)
* [type Blockable](#blockable)
* [type Const](#const)
* [type Content](#content)
	* [func (c Content) ToString() string](#content-tostring)
* [type Example](#example)
* [type Func](#func)
* [type GoCode](#gocode)
	* [func (c GoCode) ToString() string](#gocode-tostring)
* [type Heading1](#heading1)
	* [func (h Heading1) ToString() string](#heading1-tostring)
* [type Heading2](#heading2)
	* [func (h Heading2) ToString() string](#heading2-tostring)
* [type Heading3](#heading3)
	* [func (h Heading3) ToString() string](#heading3-tostring)
* [type Heading4](#heading4)
	* [func (h Heading4) ToString() string](#heading4-tostring)
* [type Index](#index)
	* [func NewIndex(md *Markdown) Index](#newindex)
	* [func (i Index) ToString() string](#index-tostring)
* [type IndexEntry](#indexentry)
* [type Markdown](#markdown)
	* [func NewMarkdown() *Markdown](#newmarkdown)
	* [func (d *Markdown) Append(b Blockable)](#markdown-append)
	* [func (d *Markdown) Content() string](#markdown-content)
	* [func (d *Markdown) Prepend(b ...Blockable)](#markdown-prepend)
* [type PackageIndex](#packageindex)
	* [func GetPackageIndex(src string) (*PackageIndex, error)](#getpackageindex)
* [type Type](#type)
* [type Var](#var)





## <a name="getsourcepackages">func GetSourcePackages</a>
``` go
func GetSourcePackages(path string) ([]string, error)
```
GetSourcePackages returns all package directories in path
excludes vendor directories

## <a name="tomd">func ToMD</a>
``` go
func ToMD(w io.Writer, text string)
```
ToMD converts comment text to formatted Markdown.
The comment was prepared by DocReader,
so it is known not to have leading, trailing blank lines
nor to have trailing spaces at the end of lines.
The comment markers have already been removed.

Each span of unindented non-blank lines is converted into
a single paragraph. There is one exception to the rule: a span that
consists of a single line, is followed by another paragraph span,
begins with a capital letter, and contains no punctuation
is formatted as a heading.

A span of indented lines is converted into a pre block,
with the common indent prefix removed.

URLs in the comment text are converted into links.

## <a name="writemdtemplate">func WriteMDTemplate</a>
``` go
func WriteMDTemplate(idx *PackageIndex, outPath string) error
```
WriteMDTemplate creates a markdown output using an template

## <a name="writemarkdown">func WriteMarkdown</a>
``` go
func WriteMarkdown(idx *PackageIndex, outPath string) error
```
WriteMarkdown writes indexed documents as markdown content



## <a name="blockable">type Blockable</a>
``` go
type Blockable interface {
	ToString() string
}
```
Blockable is implement by content blocks





## <a name="const">type Const</a>
``` go
type Const struct {
	Doc	string
	Snippet	string
}
```
Const is a single or a block of constants in a package





## <a name="content">type Content</a>
``` go
type Content string
```
Content is plain text



### <a name="content-tostring">func (Content) ToString</a>
``` go
func (c Content) ToString() string
```
ToString returns markdown string



## <a name="example">type Example</a>
``` go
type Example struct {
	Name	string
	ForFunc	string
	Suffix	string
	Doc	string
	Snippet	string
	Output	string
}
```
Example demonstrates use of functions & types in a package





## <a name="func">type Func</a>
``` go
type Func struct {
	Name		string
	Doc		string
	Snippet		string
	Recv		string
	Examples	[]Example
}
```
Func is a package level function





## <a name="gocode">type GoCode</a>
``` go
type GoCode string
```
GoCode is a Go code block



### <a name="gocode-tostring">func (GoCode) ToString</a>
``` go
func (c GoCode) ToString() string
```
ToString returns code block wrapped in `pre` quotes



## <a name="heading1">type Heading1</a>
``` go
type Heading1 string
```
Heading1 is H1



### <a name="heading1-tostring">func (Heading1) ToString</a>
``` go
func (h Heading1) ToString() string
```
ToString returns markdown string



## <a name="heading2">type Heading2</a>
``` go
type Heading2 string
```
Heading2 is H2



### <a name="heading2-tostring">func (Heading2) ToString</a>
``` go
func (h Heading2) ToString() string
```
ToString returns markdown string



## <a name="heading3">type Heading3</a>
``` go
type Heading3 string
```
Heading3 is H3



### <a name="heading3-tostring">func (Heading3) ToString</a>
``` go
func (h Heading3) ToString() string
```
ToString returns markdown string



## <a name="heading4">type Heading4</a>
``` go
type Heading4 string
```
Heading4 is H4



### <a name="heading4-tostring">func (Heading4) ToString</a>
``` go
func (h Heading4) ToString() string
```
ToString returns markdown string



## <a name="index">type Index</a>
``` go
type Index []IndexEntry
```
Index is a list of all labels in the document


### <a name="newindex">func NewIndex</a>
``` go
func NewIndex(md *Markdown) Index
```
NewIndex creates a new index from all headings in the markdown doc


### <a name="index-tostring">func (Index) ToString</a>
``` go
func (i Index) ToString() string
```
ToString returns index serialized as a list



## <a name="indexentry">type IndexEntry</a>
``` go
type IndexEntry struct {
	Label	string
	Level	int
}
```
IndexEntry is an entry that is rendered as a list item





## <a name="markdown">type Markdown</a>
``` go
type Markdown struct {
	Blocks []Blockable
}
```
Markdown is a collection of blocks


### <a name="newmarkdown">func NewMarkdown</a>
``` go
func NewMarkdown() *Markdown
```
NewMarkdown creates an empty markdown object


### <a name="markdown-append">func (*Markdown) Append</a>
``` go
func (d *Markdown) Append(b Blockable)
```
Append adds a block to Markdown

### <a name="markdown-content">func (*Markdown) Content</a>
``` go
func (d *Markdown) Content() string
```
Content returns a string representation of the markdown object

### <a name="markdown-prepend">func (*Markdown) Prepend</a>
``` go
func (d *Markdown) Prepend(b ...Blockable)
```
Prepend prepends a block to Markdown



## <a name="packageindex">type PackageIndex</a>
``` go
type PackageIndex struct {
	Name	string
	Dir	string
	Doc	string

	Vars		[]Var
	Consts		[]Const
	Funcs		[]Func
	Types		[]Type
	Examples	[]Example
}
```
PackageIndex contains all parsed types & functions information


### <a name="getpackageindex">func GetPackageIndex</a>
``` go
func GetPackageIndex(src string) (*PackageIndex, error)
```
GetPackageIndex parses and returns all the documentation content in a package




## <a name="type">type Type</a>
``` go
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
Type captures public types available at package level





## <a name="var">type Var</a>
``` go
type Var struct {
	Doc	string
	Snippet	string
}
```
Var is a single or a block of variables in a package






