# godocgen

`godocgen` is a CLI tool to generate godocs in markdown format in your repository.

This tool helps you in creating & maintaining documentation for your private Go codebases. The generated markdown files can be kept in the same repository, removing the need for running a separate `godoc` server.

## Usage

```
$ godocgen markdown
NAME:
   commands markdown - generates markdown documentation for all packages in `src` folder

USAGE:
   commands markdown [command options] [arguments...]

OPTIONS:
   --source value, --src value   path to packages that have to be documented
   --markdown value, --md value  name of markdown file to output to (default: "GODOC.md")
   --help, -h                    show help (default: false)
```
