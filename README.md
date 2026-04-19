# TextTiling

A Go implementation of the TextTiling algorithm for automatic text segmentation.

## Installation

```bash
go install github.com/davicbtoliveira/texttiling/cmd/textiling@latest
```

## Usage

```bash
textiling <file.txt>
```

### Options

- `-w`: Pseudosentence size (default: 20)
- `-k`: Block size (default: 10)
- `-method`: Similarity method - `block` or `vocab` (default: block)
- `-policy`: Cutoff policy - `hc` or `lc` (default: hc)
- `-debug`: Show debug information

## Algorithm

TextTiling identifies topic boundaries in documents by computing lexical similarity between blocks of pseudosentences. It detects gaps in lexical co-occurrence that indicate topic shifts.
