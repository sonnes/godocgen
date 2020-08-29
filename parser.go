package godocgen

import (
	"bytes"
	"go/ast"
	"go/build"
	"go/doc"
	"go/printer"
	"go/token"
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

	files, err := parseFiles(fset, src, pkgfiles)

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	// ignore any errors - they are due to unresolved identifiers
	pkg, _ := ast.NewPackage(fset, files, poorMansImporter, nil)

	// get documentation
	pkgdoc := doc.New(pkg, src, doc.AllMethods)

	idx := &PackageIndex{
		Dir:      src,
		Name:     pkgdoc.Name,
		Doc:      pkgdoc.Doc,
		Consts:   []Const{},
		Funcs:    []Func{},
		Types:    []Type{},
		Examples: []Example{},
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
	idx.Funcs = collectFuncs(fset, pkgdoc.Funcs)

	// package level types
	for _, ptype := range pkgdoc.Types {
		s := bytes.NewBufferString("")
		printer.Fprint(s, fset, ptype.Decl)

		t := Type{
			Name:    ptype.Name,
			Doc:     ptype.Doc,
			Snippet: s.String(),
			Methods: collectFuncs(fset, ptype.Methods),
		}

		idx.Types = append(idx.Types, t)
	}

	// get examples
	testfiles := append(pkginfo.TestGoFiles, pkginfo.XTestGoFiles...)
	files, _ = parseFiles(fset, src, testfiles)

	idx.Examples = collectExamples(fset, pkg, files)

	return idx, nil
}

func collectFuncs(fset *token.FileSet, docFuncs []*doc.Func) []Func {
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

		funcs = append(funcs, f)
	}

	return funcs
}

func collectExamples(fset *token.FileSet, pkg *ast.Package, testfiles map[string]*ast.File) []Example {
	var files []*ast.File
	for _, f := range testfiles {
		files = append(files, f)
	}

	var examples []Example
	globals := globalNames(pkg)
	for _, e := range doc.Examples(files...) {
		name := stripExampleSuffix(e.Name)
		if name == "" || globals[name] {

			s := bytes.NewBufferString("")
			printer.Fprint(s, fset, e.Code)

			example := Example{
				Name:    e.Name,
				Suffix:  e.Suffix,
				Doc:     e.Doc,
				Snippet: s.String(),
			}

			examples = append(examples, example)
		}
	}

	return examples
}
