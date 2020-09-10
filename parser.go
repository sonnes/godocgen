package godocgen

import (
	"bytes"
	"go/build"
	"go/doc"
	"go/printer"
	"go/token"
	"strings"
)

// GetPackageIndex parses and returns all the documentation content in a package
func GetPackageIndex(src string) (*PackageIndex, error) {

	if src == "" {
		return nil, nil
	}

	ctxt := build.Default

	pkginfo, err := ctxt.ImportDir(src, 0)

	if err != nil {
		return nil, err
	}

	// build package AST
	fset := token.NewFileSet()

	// collect source files
	pkgfiles := append(pkginfo.GoFiles, pkginfo.CgoFiles...)
	pkgfiles = append(pkgfiles, pkginfo.TestGoFiles...)
	pkgfiles = append(pkgfiles, pkginfo.XTestGoFiles...)

	files, err := parseFiles(fset, src, pkgfiles)

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	// ignore any errors - they are due to unresolved identifiers
	pkgdoc, err := doc.NewFromFiles(fset, files, "test.com/", doc.AllMethods)

	if err != nil {
		return nil, err
	}

	idx := &PackageIndex{
		Dir:      src,
		Name:     pkgdoc.Name,
		Doc:      pkgdoc.Doc,
		Consts:   []Const{},
		Vars:     []Var{},
		Funcs:    []Func{},
		Types:    []Type{},
		Examples: []Example{},
	}

	// extract example first
	for _, e := range doc.Examples(files...) {

		s := bytes.NewBufferString("")
		printer.Fprint(s, fset, e.Code)

		forFunc, suffix := splitExampleName(e.Name)

		example := Example{
			Name:    formatExampleName(e.Name),
			ForFunc: forFunc,
			Suffix:  strings.TrimSpace(suffix),
			Doc:     e.Doc,
			Snippet: s.String(),
			Output:  e.Output,
		}

		idx.Examples = append(idx.Examples, example)
	}

	// package level constants
	for _, pvar := range pkgdoc.Vars {
		s := bytes.NewBufferString("")
		printer.Fprint(s, fset, pvar.Decl)

		v := Var{
			Doc:     pvar.Doc,
			Snippet: s.String(),
		}

		idx.Vars = append(idx.Vars, v)
	}

	// package level constants
	for _, pconst := range pkgdoc.Consts {
		s := bytes.NewBufferString("")
		printer.Fprint(s, fset, pconst.Decl)

		c := Const{
			Doc:     pconst.Doc,
			Snippet: s.String(),
		}

		idx.Consts = append(idx.Consts, c)
	}

	// package level functions
	idx.Funcs = collectFuncs(fset, pkgdoc.Funcs, idx.Examples)

	// package level types
	for _, ptype := range pkgdoc.Types {
		s := bytes.NewBufferString("")
		printer.Fprint(s, fset, ptype.Decl)

		t := Type{
			Name:    ptype.Name,
			Doc:     ptype.Doc,
			Snippet: s.String(),
			Methods: collectFuncs(fset, ptype.Methods, idx.Examples),
			Funcs:   collectFuncs(fset, ptype.Funcs, idx.Examples),
		}

		idx.Types = append(idx.Types, t)
	}

	return idx, nil
}

func collectFuncs(fset *token.FileSet, docFuncs []*doc.Func, examples []Example) []Func {
	funcs := []Func{}

	for _, pfunc := range docFuncs {
		s := bytes.NewBufferString("")
		printer.Fprint(s, fset, pfunc.Decl)

		f := Func{
			Name:    pfunc.Name,
			Recv:    pfunc.Recv,
			Doc:     pfunc.Doc,
			Snippet: s.String(),
		}

		globalName := getGlobalName(pfunc.Decl)

		for _, e := range examples {
			if e.ForFunc != globalName {
				continue
			}

			f.Examples = append(f.Examples, e)
		}

		funcs = append(funcs, f)
	}

	return funcs
}
