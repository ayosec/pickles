# ico.pk

Pickle for [GNU poke][poke] 2.0 or later to read [ICO files][ico-format].

## Installation

Copy the `ico.pk` file to `~/.local/share/poke`.

```console
$ mkdir -p ~/.local/share/poke

$ cp ico.pk ~/.local/share/poke/ico.pk
```

## Usage

To load this pickle use `load ico` expression:

    $ poke

    (poke) load ico

If the module can't be loaded, ensure that  `~/.local/share/poke` is in the
value of [`load_path`][poke-modules].

To inspect functions and types defined in the pickle, you can use the
[`.info`][poke-info] command:

    (poke) .info types ICO


[ico-format]: https://en.wikipedia.org/wiki/ICO_(file_format)
[poke-info]: https://www.jemarch.net/poke-2.4-manual/html_node/info-command.html
[poke-modules]: https://www.jemarch.net/poke-2.4-manual/html_node/Loading-pickles-as-Modules.html
[poke]: https://www.gnu.org/software/poke/
