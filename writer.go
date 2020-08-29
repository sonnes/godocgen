package godocgen

import (
	"fmt"
	"io/ioutil"
)

// test constats
const (
	A = iota
	B
)

// WriteMarkdown writes indexed documents as markdown content
func WriteMarkdown(idx *PackageIndex, outPath string) error {

	md := NewMarkdown()

	// constants
	if len(idx.Consts) > 0 {
		md.Append(Heading2("Constants"))
		for _, c := range idx.Consts {
			md.Append(Content(c.Doc))
			md.Append(GoCode(c.Snippet))
		}
	}

	// functions
	if len(idx.Funcs) > 0 {
		md.Append(Heading2("Functions"))
		for _, f := range idx.Funcs {
			md.Append(Heading3(fmt.Sprintf("func %s", f.Name)))
			md.Append(GoCode(f.Snippet))
			md.Append(Content(f.Doc))
		}
	}

	// types
	if len(idx.Types) > 0 {
		md.Append(Heading2("Types"))
		for _, t := range idx.Types {
			md.Append(Heading3(fmt.Sprintf("type %s", t.Name)))
			md.Append(Content(t.Doc))
			md.Append(GoCode(t.Snippet))

			if t.Methods != nil && len(t.Methods) > 0 {
				for _, f := range t.Methods {
					md.Append(Heading4(fmt.Sprintf("func (%s) %s", f.Recv, f.Name)))
					md.Append(GoCode(f.Snippet))
					md.Append(Content(f.Doc))
				}
			}
		}
	}

	// examples
	if len(idx.Examples) > 0 {
		md.Append(Heading2("Examples"))
		for _, e := range idx.Examples {
			md.Append(Heading3(e.Name))
			md.Append(Content(e.Doc))
			md.Append(GoCode(e.Snippet))
		}
	}

	md.Prepend(
		// package details
		Heading1(idx.Name),
		Content(idx.Doc),
		// index
		Heading4("Index"),
		NewIndex(md),
	)

	body := []byte(md.Content())
	err := ioutil.WriteFile(outPath, body, 0600)

	if err != nil {
		return err
	}

	return nil
}
