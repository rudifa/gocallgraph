# gocallgraph

A prototype tool to generate a call graph for a Go program.

Some manual assembly required.

## Usage - preliminary

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

Launch `gocallgraph` commands in the directory of the go module you want to analyze.

Examples below use the `gocallgraph` module itself as the module tovisualize.

1. Run `gocallgraph raw` to get the raw call graph (in the default format for go tool callgraph).

```
gocallgraph % gocallgraph raw
2023/12/01 14:30:46 Directory created: .tmp
2023/12/01 14:30:46 File created: .tmp/callers.txt
2023/12/01 14:30:46 File created: .tmp/targets.txt
2023/12/01 14:30:47 Wrote the raw callgraph file .tmp/callgraph.raw
```

2. Run `gocallgraph dot` to convert the raw call graph to a dot file and to a svg file.

```
gocallgraph % gocallgraph dot -s
2023/12/01 14:31:39 callers:
2023/12/01 14:31:39 targets:
2023/12/01 14:31:39 DOT file generated successfully: .tmp/callgraph.dot
2023/12/01 14:31:39 Wrote SVG file: .tmp/callgraph.svg
```

This openns the svg file in your browser... and it is empty.

Note that Callers and Targets are initially empty.

3. Use the `find` command to find the  signature of a function of interest to you in the raw call graph.

```
gocallgraph % gocallgraph find -f 'main'
found: 1
github.com/rudifa/gocallgraph.main
```

4. Use the `add` command to add the function (as returned by `find`) to the list of callers or targets.

```
gocallgraph % gocallgraph add -c 'github.com/rudifa/gocallgraph.main'
add caller: github.com/rudifa/gocallgraph.main
add target:
```

Note that you should  put simple qoutes around the full signature of the function to prevent the interpretation of the special characters by your shell.

5. Run `gocallgraph dot` again

```
gocallgraph % gocallgraph dot -s
2023/12/01 14:43:16 Callers:
2023/12/01 14:43:16 github.com/rudifa/gocallgraph.main
2023/12/01 14:43:16 targets:
2023/12/01 14:43:16 DOT file generated successfully: .tmp/callgraph.dot
2023/12/01 14:43:16 Wrote SVG file: .tmp/callgraph.svg
```

This time the svg file shows the function `main` and the target(s) of its calls, just one in this case:

```
github.com/rudifa/gocallgraph/cmd.Execute
```

6. From here on, you can add more callers and targets, and run `gocallgraph dot` again to update the svg file.

7. Use the `show` command to re-open the svg file in your browser.

8. Use the `sub` command to remove a function from the list of callers or targets, then use the `dot -s` to update the svg file and show it in your browser.

```
gocallgraph % gocallgraph sub -c 'github.com/rudifa/gocallgraph.main'
sub caller: github.com/rudifa/gocallgraph.main
sub target:
```

9. When you modify yout go module's code you need to run `gocallgraph raw` again to get the latest raw call graph, and then `gocallgraph dot -s` to update the svg file and  display it.
