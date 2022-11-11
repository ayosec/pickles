# go-heapdump for GNU Poke

Pickle for [GNU poke][poke] 2.0 or later to read [heap-dumps for Go programs][heapdump].


## Installation

Copy the `go-heapdump.pk` file to `~/.local/share/poke`.

```console
$ mkdir -p ~/.local/share/poke

$ wget -O ~/.local/share/poke/go-heapdump.pk https://github.com/ayosec/go-heapdump.pk/raw/main/go-heapdump.pk
```


## Usage

To load this pickle use `load "go-heapdump.pk"` expression:

    $ poke

    (poke) load "go-heapdump.pk"

If the module can't be loaded, ensure that  `~/.local/share/poke` is in the
value of [`load_path`][poke-modules].

All types and variables have the `GHD_` prefix. To see available types you can
use the [`.info`][poke-info] command:

    (poke) .info types GHD


## Example

First, you need to create a heap-dump in the Go program:

```go
package main

import (
	"os"
	"runtime/debug"
)

func main() {
	if f, err := os.Create("heap-dump"); err == nil {
		debug.WriteHeapDump(f.Fd())
		f.Close()
	}
}
```

Then, the file `heap-dump` can be parsed with the pickle in this repository:

    $ poke ./heap-dump

    (poke) load "go-heapdump.pk"

    (poke) GHD_File @ 0#B





[heapdump]: https://github.com/golang/go/wiki/heapdump15-through-heapdump17
[poke-info]: https://www.jemarch.net/poke-2.3-manual/html_node/info-command.html
[poke-modules]: https://www.jemarch.net/poke-2.3-manual/html_node/Loading-pickles-as-Modules.html
[poke]: https://www.gnu.org/software/poke/
