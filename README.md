# stx
A simple command line tool to analyze file stats similar to the Unix `wc` command

## Installation
```bash
go build -o stx ./cmd/main.go
```

## Usage
```bash
stx [FLAGS] [file...]
```
### Flags
| Flag | Description      |
| ---- | ---------------  |
| `-l` | Count lines only |
| `-w` | Count words only |
| `-c` | Count chars only |
| `-u` | Count unique words only |
| `-h` | Show help        |

> Flags can be combined (`-lwc`)
> If no flags are specified, all stats are shown

## Examples
```bash
# Count all stats for a single file
stx file.txt

# Count only lines
stx -l file.txt

# Count lines and words for multiple files
stx -lw file1.txt file2.txt

# Count all stats for multiple files (shows totals)
stx *.txt

# Show help
stx -h
```

## Output Format
```bash
File                     Lines     Words     Chars
----------------------------------------------------
example.txt                 42       156       892
another.txt                 18        67       403
----------------------------------------------------
TOTAL                       60       223      1295
```
