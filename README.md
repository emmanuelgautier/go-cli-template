# Go CLI Template

<p align="left">
    <a href="https://github.com/emmanuelgautier/go-cli-template/actions/workflows/ci.yml"><img src="https://github.com/emmanuelgautier/go-cli-template/actions/workflows/ci.yml/badge.svg?branch=main&event=push" alt="CI Tasks for Go Cli template"></a>
    <a href="https://goreportcard.com/report/github.com/emmanuelgautier/go-cli-template"><img src="https://goreportcard.com/badge/github.com/emmanuelgautier/go-cli-template" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/emmanuelgautier/go-cli-template"><img src="https://pkg.go.dev/badge/www.github.com/emmanuelgautier/go-cli-template" alt="PkgGoDev"></a>
</p>

This is a simple Go CLI (Command Line Interface) template that you can use as a starting point for building your own command-line applications in Go. It provides a basic structure and some common features to help you get started quickly.

## Features

- Command-line argument parsing using the [cobra](https://pkg.go.dev/github.com/spf13/cobra) package.
- Simple subcommand support.
- Github Actions workflow
- GoReleaser preconfigured for Docker build, Github Release binaries (multi-archi) builds and snapcraft publishing.

## Usage

1. Clone or download this repository:

```bash
git clone https://github.com/emmanuelgautier/go-cli-template.git
cd go-cli-template
```

2. Build the CLI tool:

```bash
go build -o go-cli-template
```

3. Run the CLI tool with the `--help` flag to see the available commands:

```bash
./go-cli-template --help
```

You should see output similar to the following:

```
A simple Go CLI template.

Usage:
  go-cli-template [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  hello       Prints a friendly greeting
  help        Help about any command

Flags:
  -h, --help   help for go-cli-template

Use "go-cli-template [command] --help" for more information about a command.
```

4. Run the `hello` subcommand:

```bash
./go-cli-template hello --name YourName
```

Replace `YourName` with your actual name. This command will print a greeting.

## License

This Go CLI template is open-source and available under the MIT License. Feel free to use it as a starting point for your own CLI applications. Contributions and improvements are welcome!

## Author

[Emmanuel Gautier](https://www.emmanuelgautier.com/)
