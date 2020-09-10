# godocgen

`godocgen` is a CLI tool to generate godocs in markdown format in your repository.

This tool helps you in creating & maintaining documentation for your private Go codebases. The generated markdown files can be kept in the same repository, removing the need for running a separate `godoc` server.

## Usage

```
$ godocgen
NAME:
   godocgen - Document & maintain your Go documentation in markdown files

USAGE:
   godocgen [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --source value, --src value  path to packages that have to be documented
   --name value                 name of markdown file to output to (default: "GODOC.md")
   --help, -h                   show help (default: false)
godocgen version
```
