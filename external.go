package godocgen

// functions that are copied from go/tools/godoc

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	pathpkg "path"
	"strings"
	"unicode"
	"unicode/utf8"
)

func parseFiles(fs *token.FileSet, path string, localnames []string) (map[string]*ast.File, error) {
	files := make(map[string]*ast.File)
	for _, f := range localnames {
		absname := pathpkg.Join(path, f)

		src, err := ioutil.ReadFile(absname)
		if err != nil {
			return nil, err
		}

		file, err := parser.ParseFile(fs, absname, src, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		files[pathpkg.Join(path, f)] = file
	}

	return files, nil
}

// globalNames returns a set of the names declared by all package-level
// declarations. Method names are returned in the form Receiver_Method.
func globalNames(pkg *ast.Package) map[string]bool {
	names := make(map[string]bool)
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			addNames(names, decl)
		}
	}
	return names
}

// addNames adds the names declared by decl to the names set.
// Method names are added in the form ReceiverTypeName_Method.
func addNames(names map[string]bool, decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		name := d.Name.Name
		if d.Recv != nil {
			var typeName string
			switch r := d.Recv.List[0].Type.(type) {
			case *ast.StarExpr:
				typeName = r.X.(*ast.Ident).Name
			case *ast.Ident:
				typeName = r.Name
			}
			name = typeName + "_" + name
		}
		names[name] = true
	case *ast.GenDecl:
		for _, spec := range d.Specs {
			switch s := spec.(type) {
			case *ast.TypeSpec:
				names[s.Name.Name] = true
			case *ast.ValueSpec:
				for _, id := range s.Names {
					names[id.Name] = true
				}
			}
		}
	}
}

func startsWithUppercase(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

// stripExampleSuffix strips lowercase braz in Foo_braz or Foo_Bar_braz from name
// while keeping uppercase Braz in Foo_Braz.
func stripExampleSuffix(name string) string {
	if i := strings.LastIndex(name, "_"); i != -1 {
		if i < len(name)-1 && !startsWithUppercase(name[i+1:]) {
			name = name[:i]
		}
	}
	return name
}

func splitExampleName(s string) (name, suffix string) {
	i := strings.LastIndex(s, "_")
	if 0 <= i && i < len(s)-1 && !startsWithUppercase(s[i+1:]) {
		name = s[:i]
		suffix = " (" + strings.Title(s[i+1:]) + ")"
		return
	}
	name = s
	return
}

// poorMansImporter returns a (dummy) package object named
// by the last path component of the provided package path
// (as is the convention for packages). This is sufficient
// to resolve package identifiers without doing an actual
// import. It never returns an error.
//
func poorMansImporter(imports map[string]*ast.Object, path string) (*ast.Object, error) {
	pkg := imports[path]
	if pkg == nil {
		// note that strings.LastIndex returns -1 if there is no "/"
		pkg = ast.NewObj(ast.Pkg, path[strings.LastIndex(path, "/")+1:])
		pkg.Data = ast.NewScope(nil) // required by ast.NewPackage for dot-import
		imports[path] = pkg
	}
	return pkg, nil
}
