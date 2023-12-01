# gocallgraph

A prototype tool to generate a call graph for a Go program.

Some manual assembly required.

## Usage - preliminary, not up to date

```
Gets call graph of a go module, converts it to dot and displays it...

Usage:
  gocallgraph [command]

Available Commands:
  add         Add a new function to visualize, as a caller or as a target.
  dot         Convert the raw call graph to dot file and to svg file
  find        Find a function in the raw callgraph
  help        Help about any command
  raw         Get the raw call for the currrent module
  show        Display the svg file in browser
  sub         Remove a function to visualize, as a caller or as a target.

Flags:
  -h, --help      help for gocallgraph
  -v, --verbose   Enable verbose output

Use "gocallgraph [command] --help" for more information about a command.

```

## Install

Clone this repo and run `go install` in the repo directory.

## How to use

Launch `gocallgraph` command in the directory of the go module you want to analyze.

1. Run `gocallgraph raw` to get the raw call graph (in the default format for go tool callgraph).

```
% gocallgraph raw                                                                             [sidekick-rudifa|…2]
2023/07/23 16:36:36 Wrote callgraph to .tmp/callgraph.raw
```

2. Run `gocallgraph svg` to convert the raw call graph to a dot file and to a svg file.

```
% gocallgraph svg                                                                                                      [dev L|✚10…1]
2023/07/23 16:39:39 open .tmp/callers.txt: no such file or directory
```

What?!

Well, this is where the manual assembly comes in.

Open the file `.tmp/callgraph.raw` in your favorite editor and look for symbols that you want to include in the call graph.

For example, you want to include `main` as a caller.

In `.tmp/callgraph.raw` you find the line

```
github.com/rudifa/gocallgraph.main  --static-9:13-->    github.com/rudifa/gocallgraph/cmd.Execute
```

Add the full symbol `github.com/rudifa/gocallgraph.main` to the file `.tmp/callers.txt` and ...

3. Run `gocallgraph svg` again to convert the raw call graph to dot and to svg.

4. Run `gocallgraph show` to display the svg file in browser, or just open it from your terminal: `open .tmp/callgraph.svg`.

...

TODO getraw: touch .tmp/callers.txt and .tmp/calees.txt.
