# bml.pk

Pickle for [GNU poke][poke] 2.0 or later to read [BMP files][bmp-format].

## Installation

Copy the `bmp.pk` file to `~/.local/share/poke`.

```console
$ mkdir -p ~/.local/share/poke

$ cp bmp.pk ~/.local/share/poke/bmp.pk
```

## Usage

To load this pickle use `load bmp` expression:

    $ poke

    (poke) load bmp

If the module can't be loaded, ensure that  `~/.local/share/poke` is in the
value of [`load_path`][poke-modules].

To inspect functions and types defined in the pickle, you can use the
[`.info`][poke-info] command:

    (poke) .info types BMP


## Testing

The implementation is tested against the examples in [BMP Suite][bmpsuite].


[bmpsuite]: https://entropymine.com/jason/bmpsuite/
[bmp-format]: https://en.wikipedia.org/wiki/BMP_(file_format)
[poke-info]: https://www.jemarch.net/poke-2.4-manual/html_node/info-command.html
[poke-modules]: https://www.jemarch.net/poke-2.4-manual/html_node/Loading-pickles-as-Modules.html
[poke]: https://www.gnu.org/software/poke/
