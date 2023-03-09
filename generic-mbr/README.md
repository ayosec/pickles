# `mbr.pk`

Pickle for [GNU poke][poke] 2.0 or later to read a [classical generic MBR][mbr].

## Installation

Copy the `mbr.pk` file to `~/.local/share/poke`.

```console
$ mkdir -p ~/.local/share/poke

$ cp mbr.pk ~/.local/share/poke/mbr.pk
```

## Usage

To load this pickle use `load mbr` expression:

    $ poke

    (poke) load mbr

If the module can't be loaded, ensure that  `~/.local/share/poke` is in the
value of [`load_path`][poke-modules].

To inspect functions and types defined in the pickle, you can use the
[`.info`][poke-info] command:

    (poke) .info types MBR


[mbr]: https://en.wikipedia.org/wiki/Master_boot_record
[poke-info]: https://www.jemarch.net/poke-2.4-manual/html_node/info-command.html
[poke-modules]: https://www.jemarch.net/poke-2.4-manual/html_node/Loading-pickles-as-Modules.html
[poke]: https://www.gnu.org/software/poke/
